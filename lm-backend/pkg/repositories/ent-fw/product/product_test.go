package product_repo_test

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/product"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestCreateRepo(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	_ = product_repo.NewProductEntRepo(client)

}

func TestCreateData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := product_repo.NewProductEntRepo(client)

	prod := domain.Product{
		SKU: "<sku_code>",
		Name: "TestProduct",
		InstallInstructions: "Press install button",
		LicenseCount: 3,
	}

	err := repo.Save(prod)
	if err != nil {
		t.Fatal(err)
	}

}

func TestReadData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := product_repo.NewProductEntRepo(client)

	prod := domain.Product{
		SKU: "<sku_code>",
		Name: "TestProduct",
		InstallInstructions: "Press install button",
		LicenseCount: 3,
	}

	err := repo.Save(prod)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.FindBySKU(prod.SKU)
	if err != nil {
		t.Fatal(err)
	}

	if prod.SKU != got.SKU {
		t.Fatal("different SKU field after find")
	}

	if prod.Name != got.Name {
		t.Fatal("different Name field after find")
	}

	if prod.InstallInstructions != got.InstallInstructions {
		t.Fatal("different InstallInstructions field after find")
	}

	if prod.LicenseCount != got.LicenseCount {
		t.Fatal("different LicenseCount field after find")
	}

	if got.ID == "" {
		t.Fatal("find returns an object without ID")
	}

	if got.DateCreated.Equal(time.Time{}) {
		t.Fatal("find returns an object without DateCreated")
	}

	if got.LastUpdated.Equal(time.Time{}) {
		t.Fatal("find returns an object without LastUpdated")
	}

}

func TestUpdateData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := product_repo.NewProductEntRepo(client)

	prod := domain.Product{
		SKU: "<sku_code>",
		Name: "TestProduct",
		InstallInstructions: "Press install button",
		LicenseCount: 3,
	}

	err := repo.Save(prod)
	if err != nil {
		t.Fatal(err)
	}

	// read to populate all extra fields
	prod, err = repo.FindBySKU(prod.SKU)
	if err != nil {
		t.Fatal(err)
	}

	prod.Name = "ProdName"
	prod.InstallInstructions = "No setup needed"

	// sleep to ensure "last_updated" field reflects changes
	time.Sleep(1 * time.Second)

	updated, err := repo.UpdateByID(prod)
	if err != nil {
		t.Fatal(err)
	}

	if !updated {
		t.Fatal("'updated' flag expected to be true")
	}

	got, err := repo.FindByID(prod.ID)
	if err != nil {
		t.Fatal(err)
	}

	// per-field comparisons

	if prod.SKU != got.SKU {
		t.Fatal("different SKU field after update")
	}

	if prod.Name != got.Name {
		t.Fatal("different Name field after update")
	}

	if prod.InstallInstructions != got.InstallInstructions {
		t.Fatal("different InstallInstructions field after update")
	}

	if prod.LicenseCount != got.LicenseCount {
		t.Fatal("different LicenseCount field after update")
	}

	if !prod.DateCreated.Equal(got.DateCreated) {
		t.Fatalf("different DateCreated field after update, want %v but got %v", prod.DateCreated, got.DateCreated)
	}

	if !prod.LastUpdated.Before(got.LastUpdated) {
		t.Fatal("updated product expected to have a LastUpdated date higher than original product")
	}

}

func TestDeleteData(t *testing.T) {
	

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := product_repo.NewProductEntRepo(client)

	prod := domain.Product{
		SKU: "<sku_code>",
		Name: "TestProduct",
		InstallInstructions: "Press install button",
		LicenseCount: 3,
	}

	err := repo.Save(prod)
	if err != nil {
		t.Fatal(err)
	}

	prod, _ = repo.FindBySKU(prod.SKU)

	err = repo.DeleteByID(prod)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.FindBySKU(prod.SKU)
	
	if !ent.IsNotFound(err) {
		t.Fatal("'not found' error expected")
	}

}






