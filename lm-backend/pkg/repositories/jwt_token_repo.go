package repositories

import "license-manager/pkg/domain"

type JwtTokenRepository interface {
	Save(token domain.Token) (domain.Token, error)
	FindByToken(token string) (domain.Token, error)
	FindByIssuer(userID string) ([]domain.Token, error)
	FindClaimsByToken(token string) (domain.Claims, error)
	IsRevoked(token string) (bool, error)
	Delete(token string) error
	DeleteAllByIssuer(userID string) (int, error)
}
