package service

import (
	"context"
	"license-manager/pkg/domain"
)

type JWTService interface {

	GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error)
	GetIssuedBy(userID string) ([]domain.Token, error)
	
	Validate(tokenStr string) error	
	ValidateWithClaims(tokenStr string) (domain.Claims, error)
	
	RevokeTokensFor(issuer domain.User) (revoked int, err error)
	RevokeToken(tokenID string) (bool, error)
	
	GetClaimsFromCtx(context.Context) (domain.Claims, error)
}