package entfw

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/credentials"
)

type credentialsEntRepo struct {
	client *ent.Client
}

// *credentialsEntRepo implements repo interface
var _ repositories.CredentialsRepository = (*credentialsEntRepo)(nil)

func NewCredentialsEntRepo() *credentialsEntRepo {
	return &credentialsEntRepo{

	}
}

func (repo *credentialsEntRepo) DeleteByUserNameAndPasswordHash(name string, hash string) error {
	repo.client.Credentials.Delete().Where()
}

// FindByUserName implements repositories.CredentialsRepository
func (repo *credentialsEntRepo) FindByUserName(name string) (domain.Credentials, error) {
	panic("unimplemented")
}

// FindByUserNameAndPasswordHash implements repositories.CredentialsRepository
func (repo *credentialsEntRepo) FindByUserNameAndPasswordHash(name string, psswdHash string) (domain.Credentials, error) {
	panic("unimplemented")
}

// MergeClaimsFor implements repositories.CredentialsRepository
func (repo *credentialsEntRepo) MergeClaimsFor(name string, hash string, claims domain.Claims) (domain.Claims, error) {
	panic("unimplemented")
}

// Save implements repositories.CredentialsRepository
func (repo *credentialsEntRepo) Save(domain.Credentials) error {
	panic("unimplemented")
}

// Update implements repositories.CredentialsRepository
func (repo *credentialsEntRepo) Update(domain.Credentials) error {
	panic("unimplemented")
}

