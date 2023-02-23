package service

import (
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
)

type LicenseService interface {
	
	ListLicenses() ([]domain.License, error)
	DescribeLicense(id string) (domain.License, error)
	CreateLicense(exchange.CreateLicenseRequest) (domain.License, error)
	UpdateLicense(id string, req exchange.UpdateLicenseRequest) (domain.License, error)
	DeleteLicense(id string) error

	SuspendLicense(id string) error
}