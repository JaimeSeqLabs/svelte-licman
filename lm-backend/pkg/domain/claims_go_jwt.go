package domain

import "github.com/golang-jwt/jwt/v5"

type ClaimsJWT struct {
	jwt.RegisteredClaims
	UserKind string `json:"user_kind"`
}
