package service

import (
	"context"
	"license-manager/pkg/domain"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Validator interface {
	ValidateWithClaims(tokenStr string) (domain.Claims, error)
}

func NewJWTMiddleware(validator Validator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		f := func(w http.ResponseWriter, r *http.Request) {

			// xtract token from url, cookies or header
			tokenString := getTokenFromRequest(r)
			if tokenString == "" {
				http.Error(w, "jwt token not found", http.StatusUnauthorized)
				return
			}

			// validate and xtract claims
			claims, err := validator.ValidateWithClaims(tokenString)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			
			// add claims to ctx
			ctx := context.WithValue(r.Context(), JWTClaimsCtxKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		
		return http.HandlerFunc(f)
	}
}

func getTokenFromRequest(r *http.Request) string {
	
	// 1. Try header
	if token := getTokenFromHeader(r); token != "" {
		return token
	}
	// 2. Try cookie
	if token := getTokenFromCookies(r); token != "" {
		return token
	}
	// 3. Try URL query param
	if token := getTokenFromURL(r); token != "" {
		return token
	}

	return ""
}

func getTokenFromHeader(r *http.Request) string {

	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		return ""
	}

	split := strings.Split(bearer, "Bearer ") // note space
	if len(split) != 2 {
		return ""
	}

	return split[1]
}

func getTokenFromCookies(r *http.Request) string {
	jwtCookie, err := r.Cookie(JWTCookieName)
	if err != nil {
		return ""
	}
	return jwtCookie.Value
}

func getTokenFromURL(r *http.Request) string {
	return chi.URLParam(r, JWTCookieName)
}