package repositories

import "license-manager/pkg/domain"

type OrganizationRepository interface {
	Save(org domain.Organization) error
	FindByName(name string) []domain.Organization
	Update(org domain.Organization) (updated bool, err error)
	DeleteByName(name string) error
}
