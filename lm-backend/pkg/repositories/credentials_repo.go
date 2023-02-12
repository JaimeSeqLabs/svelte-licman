package repositories

import "license-manager/pkg/domain"

type CredentialsRepository interface {
	Save(domain.Credentials) error
	FindByUserNameAndPasswordHash(name, psswdHash string) (domain.Credentials, error)
	Update(domain.Credentials) error
	MergeClaimsFor(name, hash string, claims domain.Claims) (domain.Claims, error)
	DeleteByUserNameAndPasswordHash(name, hash string) error
}