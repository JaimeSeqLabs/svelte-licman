package auth

import (
	"context"

	"github.com/go-chi/jwtauth/v5"
)
	
type jwtService struct {
	secret []byte
	tokenAuth *jwtauth.JWTAuth
}

func NewJWTServiceMOCK(secret string) JWTService {
	return &jwtService{
		secret: []byte(secret),
		tokenAuth: jwtauth.New("HS256", []byte(secret), nil),
	}
}

func (jwts *jwtService) GenTokenFor(claims map[string]any) (string, error) {
	_, tokenStr, err := jwts.tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (jwts *jwtService) GetClaimsFromCtx(ctx context.Context) (map[string]any, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (jwts *jwtService) GetJWTAuth() *jwtauth.JWTAuth {
	return jwts.tokenAuth
}