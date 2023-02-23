package sql

import (
	"fmt"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"strings"

	"github.com/jmoiron/sqlx"
)

const organizationSchema = `
	CREATE TABLE IF NOT EXISTS organization (
		org_id    		INTEGER PRIMARY KEY,
		name			VARCHAR(80) UNIQUE NOT NULL,
		location  		TEXT NOT NULL,
		contact_name	VARCHAR(80) NOT NULL,
		contact_mail	VARCHAR(80) NOT NULL
	);
`

type orgSqlRepo struct {
	db         *sqlx.DB
	migrations []string
}

// UpdateByID implements repositories.OrganizationRepository
func (*orgSqlRepo) UpdateByID(org domain.Organization) (updated bool, err error) {
	panic("unimplemented")
}

// UpdateByName implements repositories.OrganizationRepository
func (*orgSqlRepo) UpdateByName(org domain.Organization) (updated bool, err error) {
	panic("unimplemented")
}

// DeleteByID implements repositories.OrganizationRepository
func (*orgSqlRepo) DeleteByID(id string) error {
	panic("unimplemented")
}

// FindAll implements repositories.OrganizationRepository
func (*orgSqlRepo) FindAll() []domain.Organization {
	panic("unimplemented")
}

// FindByID implements repositories.OrganizationRepository
func (*orgSqlRepo) FindByID(id string) (domain.Organization, error) {
	panic("unimplemented")
}

// *orgSqlRepo implements repo interface
var _ repositories.OrganizationRepository = (*orgSqlRepo)(nil)

func NewOrganizationSQLRepo(opts ...func(*orgSqlRepo)) *orgSqlRepo {

	repo := &orgSqlRepo{
		migrations: []string{organizationSchema}, // default value, override with options
	}

	for _, opt := range opts {
		opt(repo)
	}

	// TODO: automate this from migration files
	// Migrations
	if repo.db != nil {
		for _, schema := range repo.migrations {
			repo.db.MustExec(schema) // panics if error
		}
	}

	return repo
}

func (repo *orgSqlRepo) Save(org domain.Organization) error {

	var dto OrganizationSQLDTO

	dto.FromEntity(org)

	_, err := repo.db.NamedExec(`
		INSERT INTO organization (
			name,
			location,
			contact_name,
			contact_mail
		)
		VALUES (
			:name,
			:location,
			:contact_name,
			:contact_mail
		)
	`, dto)

	return err
}

func (repo *orgSqlRepo) FindByName(name string) (domain.Organization, error) {

	rows, err := repo.db.Queryx(`
		SELECT * FROM organization
		WHERE organization.name = ?;
	`, name)
	if err != nil {
		return domain.Organization{}, err
	}

	if !rows.Next() {
		return domain.Organization{}, fmt.Errorf("organization %s not found", name)
	}

	var org OrganizationSQLDTO

	err = rows.StructScan(&org)
	if err != nil {
		rows.Close()
		panic(err)
	}

	return org.ToEntity(), nil
}

func (repo *orgSqlRepo) Update(org domain.Organization) (bool, error) {

	var dto OrganizationSQLDTO

	dto.FromEntity(org)

	res, err := repo.db.NamedExec(`
		UPDATE organization
		SET
			location=:location,
			contact_name=:contact_name,
			contact_mail=:contact_mail
		WHERE name = :name;
	`, dto)
	if err != nil {
		return false, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return (n != 0), nil
}

func (repo *orgSqlRepo) DeleteByName(name string) error {
	_, err := repo.db.Queryx(`
		DELETE FROM organization
		WHERE organization.name = ?;
	`, name)
	return err
}

type OrganizationSQLDTO struct {
	ID          int    `db:"org_id"`
	Name        string `db:"name"`
	Location    string `db:"location"`
	ContactName string `db:"contact_name"`
	ContactMail string `db:"contact_mail"`
}

func (dto *OrganizationSQLDTO) ToEntity() domain.Organization {
	return domain.Organization{
		Name:      dto.Name,
		Location:  dto.Location,
		ContactID: fmt.Sprintf("%s, %s", dto.ContactName, dto.ContactMail),
	}
}

func (dto *OrganizationSQLDTO) FromEntity(entity domain.Organization) {
	dto.ID = -1
	dto.Name = entity.Name
	dto.Location = entity.Location
	dto.ContactName, dto.ContactMail = parseContact(entity.ContactID)
}

func parseContact(s string) (name, mail string) {
	parts := strings.Split(s, ",")
	name = ""
	mail = ""

	for _, p := range parts {

		if strings.Contains(p, "@") {
			mail = strings.TrimSpace(p)
		} else {
			name = strings.TrimSpace(p)
		}

	}
	return
}
