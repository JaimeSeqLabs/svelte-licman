package auth

import (
	"context"

	"github.com/go-chi/jwtauth/v5"
)

type JWTService interface {
	GenTokenFor(claims map[string]any) (string, error)
	GetClaimsFromCtx(context.Context) (map[string]any, error)
	GetJWTAuth() *jwtauth.JWTAuth
}