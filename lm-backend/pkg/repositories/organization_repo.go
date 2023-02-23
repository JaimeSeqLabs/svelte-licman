package repositories

import "license-manager/pkg/domain"

type OrganizationRepository interface {
	Save(org domain.Organization) error
	FindByID(id string) (domain.Organization, error)
	FindByName(name string) (domain.Organization, error)
	FindAll() []domain.Organization
	UpdateByID(org domain.Organization) (updated bool, err error)
	UpdateByName(org domain.Organization) (updated bool, err error)
	DeleteByID(id string) error
	DeleteByName(name string) error
}
