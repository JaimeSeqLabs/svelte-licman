package service

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"

	"github.com/go-chi/jwtauth/v5"
)

type jwtService struct {
	secret    []byte
	tokenAuth *jwtauth.JWTAuth
	tokenRepo repositories.JwtTokenRepository
}

func NewJWTService(secret string, tokenRepo repositories.JwtTokenRepository) JWTService {
	return &jwtService{
		secret:    []byte(secret),
		tokenAuth: jwtauth.New("HS256", []byte(secret), nil),
		tokenRepo: tokenRepo,
	}
}

func (jwts *jwtService) GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error) {
	
	_, tokenStr, err := jwts.tokenAuth.Encode(claims)
	if err != nil {
		return domain.Token{}, err
	}
	
	token := domain.Token {
		Value: tokenStr,
		Revoked: false,
		Claims: claims,
		IssuerID: issuer.ID,
	}
	
	err = jwts.tokenRepo.Save(token)
	
	return token, err
}

func (jwts *jwtService) GetClaimsFromCtx(ctx context.Context) (domain.Claims, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	return domain.Claims(claims), nil
}

func (jwts *jwtService) GetIssuedBy(userID string) ([]domain.Token, error) {
	res, err := jwts.tokenRepo.FindByIssuer(userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (jwts *jwtService) GetJWTAuth() *jwtauth.JWTAuth {
	return jwts.tokenAuth
}

func (jwts *jwtService) RevokeTokensFor(issuer domain.User) (int, error) {
	return jwts.tokenRepo.DeleteAllByIssuer(issuer.ID)
}