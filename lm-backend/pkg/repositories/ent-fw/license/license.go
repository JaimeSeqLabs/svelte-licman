package license_repo

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/license"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"license-manager/pkg/repositories/ent-fw/ent/product"
	"time"
)

type licenseEntRepo struct {
	client *ent.Client
}

func NewLicenseEntRepo(client *ent.Client) repositories.LicenseRepository {
	return &licenseEntRepo {
		client: client,
	}
}

func (repo *licenseEntRepo) Save(license domain.License) (ID string, err error) {
	
	res, err := repo.client.License.Create().
		SetFeatures(license.Features).
		SetStatus(license.Status).
		SetVersion(license.Version).
		SetNote(license.Note).
		SetContact(license.Contact).
		SetMail(license.Mail).
		AddLicenseProductIDs(license.ProductIDs...).
		SetOwnerOrgID(license.OrganizationID).
		SetSecret(license.Secret).
		SetExpirationDate(license.ExpirationDate).
		SetActivationDate(license.ActivationDate).
		SetLastAccessed(license.LastAccessed).
		SetAccessCount(license.AccessCount).
		SetLastAccessIP(license.LastAccessIP).
		Save(context.TODO())
	
	if err != nil {
		return "", err
	}
	return res.ID, nil
}

func (repo *licenseEntRepo) FindAll() []domain.License {
	
	all, err := repo.client.License.
		Query().
		WithLicenseProducts().
		WithOwnerOrg().
		All(context.TODO())
	
	if err != nil || len(all) == 0 {
		return []domain.License{}
	}

	return toEntitySlice(all)
}

func (repo *licenseEntRepo) FindByID(id string) (domain.License, error) {

	license, err := repo.client.License.
		Query().
		Where(
			license.IDEQ(id),
		).
		WithLicenseProducts().
		WithOwnerOrg().
		Only(context.TODO())
	
	if err != nil {
		return domain.License{}, err
	}

	return toEntity(license), nil
}

func (repo *licenseEntRepo) FindByOrgID(orgID string) []domain.License {
	
	licenses, err := repo.client.License.
		Query().
		WithOwnerOrg().
		WithLicenseProducts().
		Where(
			license.HasOwnerOrgWith(
				organization.IDEQ(orgID),
			),
		).
		All(context.TODO())
	
	if err != nil || licenses == nil {
		return []domain.License{}
	}
	return toEntitySlice(licenses)
}

func (repo *licenseEntRepo) FindByProductID(prodID string) []domain.License {
	
	licenses, err := repo.client.License.
		Query().
		WithLicenseProducts().
		WithOwnerOrg().
		Where(
			license.HasLicenseProductsWith(
				product.IDEQ(prodID),
			),
		).
		All(context.TODO())
	
	if err != nil || licenses == nil {
		return []domain.License{}
	}
	return toEntitySlice(licenses)
}

func (repo *licenseEntRepo) UpdateByID(id string, license domain.License) (domain.License, error) {
	
	res, err := repo.client.License.UpdateOneID(id).
		SetFeatures(license.Features).
		SetStatus(license.Status).
		SetVersion(license.Version).
		SetNote(license.Note).
		SetContact(license.Contact).
		SetMail(license.Mail).
		AddLicenseProductIDs(license.ProductIDs...).
		SetOwnerOrgID(license.OrganizationID).
		SetSecret(license.Secret).
		SetExpirationDate(license.ExpirationDate).
		SetActivationDate(license.ActivationDate).
		SetLastAccessed(time.Now()).
		AddAccessCount(1).
		Save(context.TODO())

	if err != nil {
		return domain.License{}, err
	}
	return toEntity(res), nil
}

func (repo *licenseEntRepo) DeleteByID(id string) error {
	return repo.client.License.DeleteOneID(id).Exec(context.TODO())
}

func toEntity(dto *ent.License) domain.License {
	return domain.License {
		ID: dto.ID,

		Features: dto.Features,
		Status: dto.Status,
		Version: dto.Version,
		
		Note: dto.Note,
		Contact: dto.Contact,
		Mail: dto.Mail,
		
		ProductIDs: getProductIDs(dto),
		OrganizationID: getOwnerOrgID(dto),

		Secret: dto.Secret,

		ExpirationDate: dto.ExpirationDate,
		ActivationDate: dto.ActivationDate,
		DateCreated: dto.DateCreated,
		LastUpdated: dto.LastUpdated,
		LastAccessed: dto.LastAccessed,

		AccessCount: dto.AccessCount,
		LastAccessIP: dto.LastAccessIP,
	}
}

func toEntitySlice(dtos []*ent.License) []domain.License {
	
	licenses := make([]domain.License, len(dtos))

	for i, dto := range dtos {
		licenses[i] = toEntity(dto)
	}

	return licenses
}

func getProductIDs(dto *ent.License) []string {

	prods := dto.Edges.LicenseProducts

	if len(prods) == 0 { // also if prods is nil
		return []string{}
	}

	ids := make([]string, len(prods))

	for i, p := range prods {
		ids[i] = p.ID
	}

	return ids
}

func getOwnerOrgID(dto *ent.License) string {
	
	org := dto.Edges.OwnerOrg

	if org == nil {
		return ""
	}

	return org.ID
}