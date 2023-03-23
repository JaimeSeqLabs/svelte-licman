package ent_fw_common

import "license-manager/pkg/domain"

func MergeClaims(claims, updates domain.Claims) domain.Claims {

	// application claims
	if updates.UserKind != "" {
		claims.UserKind = updates.UserKind
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
