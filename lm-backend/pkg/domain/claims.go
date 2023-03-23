package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UserKind string `json:"user_kind"`
}