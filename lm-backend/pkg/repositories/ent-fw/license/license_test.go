package license_repo_test

import (
	"context"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/ent"
	"license-manager/pkg/repositories/ent-fw/ent/enttest"
	"license-manager/pkg/repositories/ent-fw/license"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
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
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}
	
	repo := license_repo.NewLicenseEntRepo(client)

	now := time.Now()

	license := domain.License {

		Features: "licenseFeatures",
		Status: "active",
		Version: "1.0.0",
		
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		
		ProductIDs: []string{ uuid.NewString(), uuid.NewString() },
		OrganizationID: uuid.NewString(),

		Secret: "<this_is_very_secret>",

		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,

	}

	_, err := repo.Save(license)
	if err == nil {
		t.Fatal("expected error, should not be able to create license without associated entities")
	}

	org := client.Organization.Create().
		SetName("org").
		SetCountry("Spain").
		SetContact("jaime").
		SetMail("jaime@mail.com").
		SetAddress("Wallaby St.").
		SetZipcode("000000").
		SaveX(context.TODO())
	prod1 := client.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := client.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())

	license.OrganizationID = org.ID
	license.ProductIDs = []string{ prod1.ID, prod2.ID }

	res, err := repo.Save(license)
	if err != nil {
		t.Fatal(err)
	}

	if res.ID == "" {
		t.Fatal("expected to return a valid ID")
	}

}

func TestReadData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	org := client.Organization.Create().
		SetName("org").
		SetCountry("Spain").
		SetContact("jaime").
		SetMail("jaime@mail.com").
		SetAddress("Wallaby St.").
		SetZipcode("000000").
		SaveX(context.TODO())
	prod1 := client.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := client.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())
	
	repo := license_repo.NewLicenseEntRepo(client)

	now := time.Now().Local()

	license := domain.License {

		Features: "licenseFeatures",
		Status: "active",
		Version: "1.0.0",
		
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		
		ProductIDs: []string{ prod1.ID, prod2.ID },
		OrganizationID: org.ID,

		Secret: "<this_is_very_secret>",

		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,
		LastAccessed: now,

		AccessCount: 0,
		LastAccessIP: "192.168.1.1",
	}

	res, err := repo.Save(license)
	if err != nil {
		t.Fatal(err)
	}

	found, err := repo.FindByID(res.ID)
	if err != nil {
		t.Fatal(err)
	}

	// date comparisons

	dateTable := []struct {
		fieldName string
		want time.Time
		got time.Time
	}{
		{ "ExpirationDate", license.ExpirationDate, found.ExpirationDate},
		{ "ActivationDate", license.ActivationDate, found.ActivationDate},
		//{ "LastAccessed", license.LastAccessed, found.LastAccessed},
	}
	for _, e := range dateTable {
		if !e.want.Equal(e.got) {
			t.Fatalf("date %s check failed, want %v but got %v", e.fieldName, e.want, e.got)
		}
	}

	// check products independently, order may cause equal check to fail
	for _, p := range license.ProductIDs {
		contained := false
		for _, fp := range found.ProductIDs {
			if fp == p {
				contained = true
				break
			}
		}
		if !contained {
			t.Fatalf("product %s not contained in license retrieved from db", p)
		}
	}

	// restrict fields to compare

	want := domain.License {
		
		ID: res.ID,

		Features: license.Features,
		Status: license.Status,
		Version: license.Version,
		
		Note: license.Note,
		Contact: license.Contact,
		Mail: license.Mail,
		
		OrganizationID: license.OrganizationID,

		Secret: license.Secret,

	}

	got := domain.License {
		
		ID: res.ID,

		Features: found.Features,
		Status: found.Status,
		Version: found.Version,
		
		Note: found.Note,
		Contact: found.Contact,
		Mail: found.Mail,
		
		OrganizationID: found.OrganizationID,

		Secret: found.Secret,

	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf(`want %+v
		but got %+v`, want, got)
	}

}

func TestUpdateData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	org := client.Organization.Create().
		SetName("org").
		SetCountry("Spain").
		SetContact("jaime").
		SetMail("jaime@mail.com").
		SetAddress("Wallaby St.").
		SetZipcode("000000").
		SaveX(context.TODO())
	prod1 := client.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := client.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())
	
	repo := license_repo.NewLicenseEntRepo(client)

	now := time.Now()

	license := domain.License {

		Features: "licenseFeatures",
		Status: "active",
		Version: "1.0.0",
		
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		
		ProductIDs: []string{ prod1.ID, prod2.ID },
		OrganizationID: org.ID,

		Secret: "<this_is_very_secret>",

		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,

		AccessCount: 0,
		LastAccessed: now,
		LastAccessIP: "192.168.1.1",
	}

	saved, err := repo.Save(license)
	if err != nil {
		t.Fatal(err)
	}

	newLicense := license
	newLicense.Status = "suspended"

	updated, err := repo.UpdateByID(saved.ID, newLicense)
	if err != nil {
		t.Fatal(err)
	}

	if updated.ID != saved.ID {
		t.Fatal("updated license has different ID")
	}

	if updated.Status != "suspended" {
		t.Fatal("updated license doesn't have updated value in field")
	}

	if !updated.LastUpdated.After(now) {
		t.Fatal("expected to have 'last_updated' field reflecting changes")
	}

	if updated.AccessCount != 1 {
		t.Fatal("expected to increase access counter")
	}

}

func TestDeleteData(t *testing.T) {
	
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	org := client.Organization.Create().
		SetName("org").
		SetCountry("Spain").
		SetContact("jaime").
		SetMail("jaime@mail.com").
		SetAddress("Wallaby St.").
		SetZipcode("000000").
		SaveX(context.TODO())
	prod1 := client.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := client.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())
	
	repo := license_repo.NewLicenseEntRepo(client)

	now := time.Now()

	license := domain.License {

		Features: "licenseFeatures",
		Status: "active",
		Version: "1.0.0",
		
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		
		ProductIDs: []string{ prod1.ID, prod2.ID },
		OrganizationID: org.ID,

		Secret: "<this_is_very_secret>",

		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,

		AccessCount: 0,
		LastAccessed: now,
		LastAccessIP: "192.168.1.1",
	}

	saved, err := repo.Save(license)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteByID(saved.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.FindByID(saved.ID)

	if !ent.IsNotFound(err) {
		t.Fatal("expected 'not found' error")
	}


}

func TestLicenseQuotas(t *testing.T) {

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		t.Fatal(err)
	}

	org := client.Organization.Create().
		SetName("org").
		SetCountry("Spain").
		SetContact("jaime").
		SetMail("jaime@mail.com").
		SetAddress("Wallaby St.").
		SetZipcode("000000").
		SaveX(context.TODO())
	prod1 := client.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := client.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())
	
	repo := license_repo.NewLicenseEntRepo(client)

	now := time.Now()

	quotas := map[string]string {
		"maxWorkspaces": "5",
		"maxUsers": "3",
	}

	license := domain.License {

		Features: "licenseFeatures",
		Status: "active",
		Version: "1.0.0",
		
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		
		ProductIDs: []string{ prod1.ID, prod2.ID },
		OrganizationID: org.ID,
		Quotas: quotas,

		Secret: "<this_is_very_secret>",

		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,

		AccessCount: 0,
		LastAccessed: now,
		LastAccessIP: "192.168.1.1",
	}

	saved, err := repo.Save(license)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.FindByID(saved.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !equal(quotas, got.Quotas) {
		t.Fatalf("the saved quota map does not match the original one, want %+v but got %+v", quotas, got.Quotas)
	}

	newLicense := license
	newLicense.Quotas = map[string]string {
		"maxWorkspaces": "15",
		"maxUsers": "10",
	}

	updated, err := repo.UpdateByID(saved.ID, newLicense)
	if err != nil {
		t.Fatal(err)
	}

	if !equal(newLicense.Quotas, updated.Quotas) {
		t.Fatalf("the saved quota map does not match the original one, want %+v but got %+v", quotas, updated.Quotas)
	}

}

func equal[K comparable, V any](m1, m2 map[K]V) bool {

	if m1 == nil {
		return m2 == nil
	}

	if m2 == nil {
		return m1 == nil
	}

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if v2, found := m2[k]; found {

			if !reflect.DeepEqual(v, v2) {
				return false
			}

		} else {
			return false // element not found
		}
	}

	return true
}
