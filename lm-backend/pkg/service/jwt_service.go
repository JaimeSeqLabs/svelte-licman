package service

import (
	"context"
	"license-manager/pkg/domain"

	"github.com/go-chi/jwtauth/v5"
)

type JWTService interface {
	GenTokenFor(claims domain.Claims) (domain.Token, error)
	GetClaimsFromCtx(context.Context) (domain.Claims, error)
	GetJWTAuth() *jwtauth.JWTAuth
}