package service

import (
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
)

type licenseService struct {
	licenseRepo repositories.LicenseRepository
	orgRepo     repositories.OrganizationRepository
	prodRepo    repositories.ProductRepository
}

func NewLicenseService(
	licenseRepo repositories.LicenseRepository,
	orgRepo repositories.OrganizationRepository,
	prodRepo repositories.ProductRepository,
) LicenseService {
	return &licenseService{licenseRepo, orgRepo, prodRepo}
}

func (ls *licenseService) ListLicenses() ([]domain.License, error) {
	return ls.licenseRepo.FindAll(), nil
}

func (ls *licenseService) DescribeLicense(id string) (domain.License, error) {
	return ls.licenseRepo.FindByID(id)
}

func (ls *licenseService) CreateLicense(req exchange.CreateLicenseRequest) (domain.License, error) {
	
	org, err := ls.orgRepo.FindByName(req.OrganizationName)
	if err != nil {
		return domain.License{}, err
	}

	prodIDs := []string{}

	for _, sku := range req.ProductSKUs {

		prod, err := ls.prodRepo.FindBySKU(sku)
		if err != nil {
			return domain.License{}, fmt.Errorf("cannot create license, product sku %s not found: %w", sku, err)
		}
	
		prodIDs = append(prodIDs, prod.ID)
	}

	license, err := ls.licenseRepo.Save(domain.License{
		Features: req.Features,
		Status: req.Status,
		Version: req.Version,
		Note: req.Note,
		Contact: req.Contact,
		Mail: req.Mail,
		ProductIDs: prodIDs,
		OrganizationID: org.ID,
		Secret: req.Secret,
		ExpirationDate: req.ExpirationDate,
		ActivationDate: req.ActivationDate,
	})
	if err != nil {
		return domain.License{}, err
	}

	return license, nil
}

func (ls *licenseService) UpdateLicense(id string, req exchange.UpdateLicenseRequest) (domain.License, error) {

	newLicense := req.License
	newLicense.Quotas = req.Quotas
	newLicense.ProductIDs = req.Products

	updated, err := ls.licenseRepo.UpdateByID(id, newLicense)
	if err != nil {
		return domain.License{}, err
	}

	return updated, nil
}

func (ls *licenseService) DeleteLicense(id string) error {
	return ls.licenseRepo.DeleteByID(id)
}

func (ls *licenseService) SuspendLicense(id string) error {

	license, err := ls.licenseRepo.FindByID(id)
	if err != nil {
		return err
	}

	license.Status = "suspended"

	_, err = ls.licenseRepo.UpdateByID(license.ID, license)

	return err
}

func (ls *licenseService) FindQuotasByID(id string) (map[string]string, error) {

	license, err := ls.licenseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return license.Quotas, nil
}

func (ls *licenseService) SetQuotasByID(id string, q map[string]string) error {

	license, err := ls.licenseRepo.FindByID(id)
	if err != nil {
		return err
	}

	license.Quotas = q

	_, err = ls.licenseRepo.UpdateByID(id, license)
	if err != nil {
		return err
	}

	return nil
}
