//go:build ignore

package service

import (
	"context"
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTClaimsCtxKey = "Claims"
	JWTCookieName   = "jwt"
)

type jwtService struct {
	secret    []byte
	tokenRepo repositories.JwtTokenRepository
}

func NewJWTService(secret string, tokenRepo repositories.JwtTokenRepository) JWTService {
	return &jwtService{
		secret:    []byte(secret),
		tokenRepo: tokenRepo,
	}
}

func (jwts *jwtService) GenTokenFor(issuer domain.User, claims domain.ClaimsJWT) (domain.Token, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if token == nil {
		return domain.Token{}, fmt.Errorf("unable to generate token for claims %+v", claims)
	}

	domainToken := domain.Token{
		Value:    token.Raw,
		Revoked:  false,
		Claims:   claims,
		IssuerID: issuer.ID,
	}

	return domainToken, jwts.tokenRepo.Save(domainToken)
}

func (jwts *jwtService) GetClaimsFromCtx(ctx context.Context) (domain.ClaimsJWT, error) {
	claims, ok := ctx.Value(JWTClaimsCtxKey).(domain.ClaimsJWT)
	if !ok {
		return domain.ClaimsJWT{}, fmt.Errorf("unable to retrieve claims from context")
	}
	return claims, nil
}

func (jwts *jwtService) GetIssuedBy(userID string) ([]domain.Token, error) {
	res, err := jwts.tokenRepo.FindByIssuer(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (jwts *jwtService) GetJWTAuth() *jwtauth.JWTAuth {
	return nil
}

func (jwts *jwtService) RevokeTokensFor(issuer domain.User) (int, error) {
	return jwts.tokenRepo.DeleteAllByIssuer(issuer.ID)
}

func (jwts *jwtService) NewJWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {

			tokenString := getTokenFromRequest(r)
			if tokenString == "" {
				http.Error(w, "jwt token not found", http.StatusUnauthorized)
				return
			}

			// claims embeds jwt.RegisteredClaims
			claims := domain.ClaimsJWT{}

			token, err := jwt.ParseWithClaims(
				tokenString,
				&claims,
				func(t *jwt.Token) (interface{}, error) { return jwts.secret, nil },
				jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if !token.Valid {
				http.Error(w, "invalid jwt token", http.StatusUnauthorized)
				return
			}

			// add claims to ctx
			ctx := context.WithValue(r.Context(), JWTClaimsCtxKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(f)
	}
}

func getTokenFromURL(r *http.Request) string {
	return chi.URLParam(r, JWTCookieName)
}

func getTokenFromCookies(r *http.Request) string {
	jwtCookie, err := r.Cookie(JWTCookieName)
	if err != nil {
		return ""
	}
	return jwtCookie.Value
}

func getTokenFromHeader(r *http.Request) string {

	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return ""
	}

	split := strings.Split(bearer, "Bearer ") // note space
	if len(split) != 2 {
		return ""
	}

	return split[1]
}

func getTokenFromRequest(r *http.Request) string {

	if token := getTokenFromHeader(r); token != "" {
		return token
	}

	if token := getTokenFromCookies(r); token != "" {
		return token
	}

	if token := getTokenFromURL(r); token != "" {
		return token
	}

	return ""
}
