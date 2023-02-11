package entfw_test

import (
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestOrganizationBasic(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	repo := entfw.NewOrganizaitionEntRepo(
		entfw.WithEntClient(client),
		entfw.WithAutoMigration(true),
	)

	target := domain.Organization{
		Name: "Org1",
		Location: "Barcelona, Spain",
	}

	err := repo.Save(target)
	if err != nil {
		t.Fatal(err)
	}

	orgs := repo.FindByName("Org1")
	if orgs == nil || len(orgs) != 1 {
		t.Fatal("unable to find inserted org")
	}

	org := orgs[0]

	if !reflect.DeepEqual(org, target) {
		t.Fatalf("want %+v but got %+v", target, org)
	}




}