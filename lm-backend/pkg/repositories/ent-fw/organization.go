package entfw

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"log"
	"strconv"
)

type orgEntRepo struct {
	client *ent.Client
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
		SetLocation(org.Location).
		SetNillableContactID(getContactIdAsNillableInt(org)).
		Save(context.TODO())
	
	return err
}

func (repo *orgEntRepo) FindByName(name string) []domain.Organization {	
	
	res, err := repo.client.Organization.
		Query().
		Where(organization.NameEQ(name)).
		All(context.TODO())
	
	if err != nil {
		return nil
	}

	if len(res) == 0 {
		return []domain.Organization{}
	}

	orgs := make([]domain.Organization, len(res))

	for i, dto := range res {
		orgs[i] = ToEntity(dto)
	}

	return orgs
}

func (repo *orgEntRepo) Update(org domain.Organization) (updated bool, err error) {

	updates, err := repo.client.Organization.
		Update().
		SetName(org.Name).
		SetLocation(org.Location).
		SetNillableContactID(getContactIdAsNillableInt(org)).
		Where(
			organization.NameEQ(org.Name),
		).
		Save(context.TODO())

	if updates == 0 {
		return false, err
	}

	return true, err
}

func (repo *orgEntRepo) DeleteByName(name string) error {
	_, err := repo.client.Organization.
		Delete().
		Where(organization.NameEQ(name)).
		Exec(context.TODO())
	return err
}

func ToEntity(dto *ent.Organization) domain.Organization {
	cid := ""
	if dto.ContactID > 0 {
		cid = strconv.Itoa(dto.ContactID)
	}
	return domain.Organization{
		Name: dto.Name,
		Location: dto.Location,
		ContactID: cid,
	}
}

func getContactIdAsNillableInt(org domain.Organization) *int {
	var contactID *int
	if cid, err := strconv.Atoi(org.ContactID); err != nil {
		contactID = nil
	} else {
		contactID = &cid
	}
	return contactID
}


