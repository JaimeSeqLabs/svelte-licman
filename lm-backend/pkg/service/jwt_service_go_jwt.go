package service

import (
	"license-manager/pkg/domain"
)

type JWTService2 interface {
	GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error)
	GetIssuedBy(userID string) ([]domain.Token, error)
	RevokeToken(tokenID string) (bool, error)
	RevokeTokensFor(issuer domain.User) (int, error)
	Validate(tokenStr string) error	
	ValidateWithClaims(tokenStr string) (domain.Claims, error)
}