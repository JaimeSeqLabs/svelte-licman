package license_repo_test

import (
	"context"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/license"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateRepo(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	_ = license_repo.NewLicenseEntRepo(client)

}

func TestCreateData(t *testing.T) {
	
}

func TestReadData(t *testing.T) {
	
}

func TestUpdateData(t *testing.T) {
	
}

func TestDeleteData(t *testing.T) {
	
}