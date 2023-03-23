package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	jwt.RegisteredClaims
	UserKind string `json:"user_kind"`
	RWPermission string `json:"rw_permission"`
}

func (claims Claims) Merge(updates Claims) Claims {
	
	// application claims
	if updates.UserKind != "" {
		claims.UserKind = updates.UserKind
	}
	if updates.RWPermission != "" {
		claims.RWPermission = updates.RWPermission
	}

	// JWT claims
	if updates.ID != "" {
		claims.ID = updates.ID
	}
	if updates.Issuer != "" {
		claims.Issuer = updates.Issuer
	}
	if updates.Subject != "" {
		claims.Subject = updates.Subject
	}
	if len(updates.Audience) == 0 {
		// TODO: merge audience
		claims.Audience = updates.Audience
	}
	if updates.ExpiresAt != nil {
		claims.ExpiresAt = updates.ExpiresAt
	}
	if updates.NotBefore != nil {
		claims.NotBefore = updates.NotBefore
	}
	if updates.IssuedAt != nil {
		claims.IssuedAt = updates.IssuedAt
	}

	return claims
}