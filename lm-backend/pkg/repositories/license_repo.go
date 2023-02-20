package repositories

import "license-manager/pkg/domain"

type LicenseRepository interface {
	
	Save(domain.License) (ID string, err error)
	
	FindAll() []domain.License
	FindByID(string) (domain.License, error)
	FindByOrgID(orgID string) []domain.License
	FindByProductID(prodID string) []domain.License

	UpdateByID(id string, license domain.License) (domain.License, error)
	DeleteByID(id string, license domain.License) error

}