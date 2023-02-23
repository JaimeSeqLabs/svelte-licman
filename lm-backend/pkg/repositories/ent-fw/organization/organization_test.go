package organization_repo_test

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/organization"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestOrganizationBasic(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	repo := organization_repo.NewOrganizationEntRepo(
		organization_repo.WithEntClient(client),
		organization_repo.WithAutoMigration(true),
	)

	target := domain.Organization{
		Name: "Org1",
		Location: "Barcelona, Spain",
	}

	err := repo.Save(target)
	if err != nil {
		t.Fatal(err)
	}

	org, err := repo.FindByName("Org1")
	if err != nil {
		t.Fatal(err)
	}

	// patch obj with ID coming from ent
	target.ID = org.ID

	if !reflect.DeepEqual(org, target) {
		t.Fatalf("want %+v but got %+v", target, org)
	}




}