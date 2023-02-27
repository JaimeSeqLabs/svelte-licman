package service

import (
	"context"
	"license-manager/pkg/domain"

	"github.com/go-chi/jwtauth/v5"
)

type JWTService interface {
	GenTokenFor(issuer domain.User, claims domain.Claims) (domain.Token, error)
	GetClaimsFromCtx(context.Context) (domain.Claims, error)
	GetIssuedBy(userID string) ([]domain.Token, error)
	GetJWTAuth() *jwtauth.JWTAuth
	RevokeTokensFor(issuer domain.User) (revoked int, err error)
}