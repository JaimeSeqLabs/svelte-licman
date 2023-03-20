package service

import (
	"context"
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTClaimsCtxKey = "Claims"
	JWTCookieName   = "jwt"
)

type jwtService2 struct {
	secret    []byte
	tokenRepo repositories.JwtTokenRepository
}

func NewJWTService2(secret string, tokenRepo repositories.JwtTokenRepository) JWTService2 {
	return &jwtService2{
		secret:    []byte(secret),
		tokenRepo: tokenRepo,
	}
}

func (jwts *jwtService2) GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, fromDomainClaims(claims))
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

func (jwts *jwtService2) GetClaimsFromCtx(ctx context.Context) (domain.Claims, error) {
	claims, ok := ctx.Value(JWTClaimsCtxKey).(domain.ClaimsJWT)
	if !ok {
		return domain.Claims{}, fmt.Errorf("unable to retrieve claims from context")
	}
	return toDomainClaims(claims), nil
}

func (jwts *jwtService2) GetIssuedBy(userID string) ([]domain.Token, error) {
	res, err := jwts.tokenRepo.FindByIssuer(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (jwts *jwtService2) RevokeTokensFor(issuer domain.User) (int, error) {
	return jwts.tokenRepo.DeleteAllByIssuer(issuer.ID)
}

func (jwts *jwtService2) RevokeToken(tokenID string) (bool, error) {
	return true, jwts.tokenRepo.Delete(tokenID)
}

func (jwts *jwtService2) Validate(tokenStr string) error {
	_, err := jwts.ValidateWithClaims(tokenStr)
	return err
}

func (jwts *jwtService2) ValidateWithClaims(tokenStr string) (domain.Claims, error)  {

	// claims embeds jwt.RegisteredClaims
	claims := domain.ClaimsJWT{}
	
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
	
	return toDomainClaims(claims), nil
}

func toDomainClaims(claims domain.ClaimsJWT) domain.Claims {
	// TODO:
	return make(domain.Claims)
}

func fromDomainClaims(claims domain.Claims) domain.ClaimsJWT {
	return domain.ClaimsJWT{} // TODO:
}
