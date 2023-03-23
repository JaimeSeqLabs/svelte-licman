package service

import (
	"context"
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTClaimsCtxKey = "claims"
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

func (jwts *jwtService) GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error) {
	
	if claims.IssuedAt == nil {
		claims.IssuedAt = jwt.NewNumericDate(time.Now())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	if token == nil {
		return domain.Token{}, fmt.Errorf("unable to generate token for claims %+v", claims)
	}

	domainToken, err := jwts.tokenRepo.Save(domain.Token{
		Value:    token.Raw,
		Revoked:  false,
		Claims:   claims,
		IssuerID: issuer.ID,
	})
	if err != nil {
		return domain.Token{}, err
	}

	return domainToken, nil
}

func (jwts *jwtService) GetIssuedBy(userID string) ([]domain.Token, error) {
	res, err := jwts.tokenRepo.FindByIssuer(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (jwts *jwtService) Validate(tokenStr string) error {
	_, err := jwts.ValidateWithClaims(tokenStr)
	return err
}

func (jwts *jwtService) ValidateWithClaims(tokenStr string) (domain.Claims, error) {
	
	// claims embeds jwt.RegisteredClaims
	claims := domain.Claims{}
	
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&claims,
		func(t *jwt.Token) (interface{}, error) { return jwts.secret, nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
	)
	if err != nil {
		return domain.Claims{}, err
	}
	if !token.Valid {
		return domain.Claims{}, fmt.Errorf("token %s is not valid", tokenStr)
	}
	
	return claims, nil
}

func (jwts *jwtService) RevokeTokensFor(issuer domain.User) (revoked int, err error) {
	return jwts.tokenRepo.DeleteAllByIssuer(issuer.ID)
}

func (jwts *jwtService) RevokeToken(tokenID string) (bool, error) {
	return true, jwts.tokenRepo.Delete(tokenID)
}

func (jwts *jwtService) GetClaimsFromCtx(ctx context.Context) (domain.Claims, error) {
	claims, ok := ctx.Value(JWTClaimsCtxKey).(domain.Claims)
	if !ok {
		return domain.Claims{}, fmt.Errorf("unable to retrieve claims from context")
	}
	return claims, nil
}
