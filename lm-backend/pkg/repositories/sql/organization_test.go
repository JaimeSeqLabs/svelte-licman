package sql_test

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/sql"
	"reflect"
	"testing"
)

func TestOrgSQLCreateOrg(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	err := repo.Save(domain.Organization{
		Name:     "BigCorpo",
		Location: "Barcelona, Spain",
		Contact:  "Jaime, jaime.munoz@mail.com",
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestOrgSQLFindOrg(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	tests := []domain.Organization{
		{
			Name:     "Org1",
			Location: "Madrid, Spain",
			Contact:  "Alice, alice@mail.com",
		},
		{
			Name:     "Org2",
			Location: "Barcelona, Spain",
			Contact:  "Bob, bob@mail.com",
		},
	}

	for _, tc := range tests {
		repo.Save(tc)
	}

	want := tests[0]
	got := repo.FindByName(want.Name)

	if got == nil {
		t.Fatal("db error, failed to query any result")
	}

	if len(got) != 1 {
		t.Fatal("failed to find the target entity")
	}

	if !reflect.DeepEqual(want, got[0]) {
		t.Fatalf("Got unexpected value, want %+v but got %+v\n", want, got[0])
	}
}

func TestOrgSQLUpdateOrg(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	org := domain.Organization{
		Name:     "BigCorpo",
		Location: "Barcelona, Spain",
		Contact:  "Jaime, jaime.munoz@mail.com",
	}

	repo.Save(org)

	org.Location = "Madrid, Spain"

	updated, err := repo.Update(org)
	if err != nil {
		t.Fatal(err)
	}

	if !updated {
		t.Fatal("should mark updated boolean flag")
	}

	got := repo.FindByName(org.Name)[0]

	if !reflect.DeepEqual(got, org) {
		t.Fatal("returned value from repository is not updated")
	}

}

func TestOrgSQLUpdateOrg_failure(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	org := domain.Organization{
		Name:     "BigCorpo",
		Location: "Barcelona, Spain",
		Contact:  "Jaime, jaime.munoz@mail.com",
	}

	updated, err := repo.Update(org)
	if err != nil {
		t.Fatal(err)
	}

	if updated {
		t.Fatal("updated boolean flag should not be set")
	}
}

func TestOrgSQLDeleteOrg(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	org := domain.Organization{
		Name:     "BigCorpo",
		Location: "Barcelona, Spain",
		Contact:  "Jaime, jaime.munoz@mail.com",
	}

	repo.Save(org)

	err := repo.DeleteByName(org.Name)
	if err != nil {
		t.Fatal(err)
	}

	res := repo.FindByName(org.Name)
	if len(res) != 0 {
		t.Fatal("repository found value even if it was deleted")
	}

}

func TestOrgSQLDeleteOrg_failure(t *testing.T) {

	repo := sql.NewOrganizationSQLRepo(
		sql.WithDriverAndURL("sqlite3", ":memory:"),
	)
	if repo == nil {
		t.Fatal("failed to create repo")
	}

	org := domain.Organization{
		Name:     "BigCorpo",
		Location: "Barcelona, Spain",
		Contact:  "Jaime, jaime.munoz@mail.com",
	}

	err := repo.DeleteByName(org.Name)
	if err != nil {
		t.Fatal(err)
	}

	// ok, nothing to delete
}
