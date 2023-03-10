package organization_repo

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"log"
)

type orgEntRepo struct {
	client  *ent.Client
	_runMig bool
}

// *orgSqlRepo implements repo interface
var _ repositories.OrganizationRepository = (*orgEntRepo)(nil)

func NewOrganizationEntRepo(opts ...func(*orgEntRepo)) *orgEntRepo {

	repo := &orgEntRepo{
		_runMig: false,
	}

	for _, opt := range opts {
		opt(repo)
	}

	if repo.client == nil {
		panic("must provide parameters to create an ent client")
	}

	if repo._runMig {
		if err := repo.client.Schema.Create(context.Background()); err != nil {
			repo.client.Close()
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return repo
}

func (repo *orgEntRepo) Save(org domain.Organization) error {

	_, err := repo.client.Organization.
		Create().
		SetName(org.Name).
		SetContact(org.Contact).
		SetMail(org.Mail).
		SetAddress(org.Address).
		SetZipcode(org.ZipCode).
		SetCountry(org.Country).
		AddLicenseIDs(org.Licenses...).
		Save(context.TODO())

	return err
}

func (repo *orgEntRepo) FindByID(id string) (domain.Organization, error) {

	res, err := repo.client.Organization.
		Query().
		WithLicenses().
		Where(organization.IDEQ(id)).
		Only(context.TODO())
	if err != nil {
		return domain.Organization{}, err
	}

	return toEntity(res), nil
}

func (repo *orgEntRepo) FindByName(name string) (domain.Organization, error) {

	res, err := repo.client.Organization.
		Query().
		WithLicenses().
		Where(organization.NameEQ(name)).
		Only(context.TODO())

	if err != nil {
		return domain.Organization{}, err
	}

	return toEntity(res), nil
}

func (repo *orgEntRepo) FindAll() []domain.Organization {
	
	res, err := repo.client.Organization.Query().All(context.TODO())
	
	if err != nil || len(res) == 0 {
		return []domain.Organization{}
	}

	orgs := make([]domain.Organization, len(res))

	for i, dto := range res {
		orgs[i] = toEntity(dto)
	}

	return orgs
}

func (repo *orgEntRepo) UpdateByName(org domain.Organization) (updated bool, err error) {

	updates, err := repo.client.Organization.
		Update().
		SetName(org.Name).
		SetContact(org.Contact).
		SetMail(org.Mail).
		SetAddress(org.Address).
		SetZipcode(org.ZipCode).
		SetCountry(org.Country).
		AddLicenseIDs(org.Licenses...).
		Where(
			organization.NameEQ(org.Name),
		).
		Save(context.TODO())

	if updates == 0 {
		return false, err
	}

	return true, err
}

func (repo *orgEntRepo) UpdateByID(org domain.Organization) (bool, error) {

	updated, err := repo.client.Organization.
		UpdateOneID(org.ID).
		SetName(org.Name).
		SetContact(org.Contact).
		SetMail(org.Mail).
		SetAddress(org.Address).
		SetZipcode(org.ZipCode).
		SetCountry(org.Country).
		AddLicenseIDs(org.Licenses...).
		Save(context.TODO())
	
	if err != nil {
		return false, err
	}

	if updated == nil {
		return false, err
	}

	return true, nil
}

func (repo *orgEntRepo) DeleteByID(id string) error {
	return repo.client.Organization.DeleteOneID(id).Exec(context.TODO())
}

func (repo *orgEntRepo) DeleteByName(name string) error {
	_, err := repo.client.Organization.
		Delete().
		Where(organization.NameEQ(name)).
		Exec(context.TODO())
	return err
}

func toEntity(dto *ent.Organization) domain.Organization {
	return domain.Organization{
		ID: dto.ID,
		Name:      dto.Name,
		Contact: dto.Contact,
		Mail: dto.Mail,
		Address: dto.Address,
		ZipCode: dto.Zipcode,
		Country: dto.Country,
		Licenses: getLicenseIDs(dto),
		DateCreated: dto.DateCreated,
		LastUpdated: dto.LastUpdated,
	}
}

func getLicenseIDs(dto *ent.Organization) []string {

	lics := dto.Edges.Licenses

	if len(lics) == 0 { // also if lics is nil
		return []string{}
	}

	ids := make([]string, len(lics))

	for i, e := range lics {
		ids[i] = e.ID
	}

	return ids
}
