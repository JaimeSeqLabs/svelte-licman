package controller_test

import (
	"context"
	"license-manager/pkg/controller"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	license_repo "license-manager/pkg/repositories/ent-fw/license"
	organization_repo "license-manager/pkg/repositories/ent-fw/organization"
	"license-manager/pkg/repositories/ent-fw/product"
	"license-manager/pkg/service"
	"net/http"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestLicenseCRUD(t *testing.T) {
	
	db := makeClient(t)
	defer db.Close()

	orgRepo := organization_repo.NewOrganizationEntRepo(organization_repo.WithEntClient(db))
	prodRepo := product_repo.NewProductEntRepo(db)
	licenseRepo := license_repo.NewLicenseEntRepo(db)

	licenseService := service.NewLicenseService(licenseRepo, orgRepo, prodRepo)
	certService := service.NewCertificateService(licenseService, prodRepo)

	licenseController := controller.NewLicenseController(licenseService, certService)
	r := licenseController.Routes()

	// test data

	org := db.Organization.Create().
		SetName("org").
		SetLocation("Barcelona").
		SaveX(context.TODO())
	prod1 := db.Product.Create().
		SetSku("sku1").
		SetName("prod1").
		SaveX(context.TODO())
	prod2 := db.Product.Create().
		SetSku("sku2").
		SetName("prod2").
		SaveX(context.TODO())

	now := time.Now()

	// create license

	createReq := exchange.CreateLicenseRequest {
		Features: "",
		Status: "active",
		Note: "",
		Contact: "jaime",
		Mail: "jaime@mail.com",
		ProductSKUs: []string{ prod1.Sku, prod2.Sku },
		OrganizationName: org.Name,
		Quotas: map[string]string{ "maxUsers": "5" },
		Secret: "<this_is_very_secret>",
		ExpirationDate:	now.Add(5 * 30 * 24 * time.Hour), // 5 months
		ActivationDate: now,
	}

	var createRes domain.License

	res := callChiRouter(
		r, http.MethodPost, "/",
		createReq,
		&createRes,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("create request failed with %s", res.Status)
	}

	// read license

	var describeRes exchange.DescribeLicenseResponse

	res = callChiRouter(
		r, http.MethodGet, "/" + createRes.ID,
		nil,
		&describeRes,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("read request failed with %s", res.Status)
	}

	if describeRes.License.ID != createRes.ID {
		t.Fatalf("unexpected license ID, want %s but got %s", describeRes.License.ID, createRes.ID)
	}

	if describeRes.License.Status != createRes.Status {
		t.Fatalf("unexpected license status, want %s but got %s", describeRes.License.ID, createRes.ID)
	}

	if !equal(describeRes.License.ProductIDs, []string{prod1.ID, prod2.ID}) {
		t.Fatalf("unexpected product ID list, want %+v but got %+v", describeRes.License.ProductIDs, []string{prod1.ID, prod2.ID})
	}

	if describeRes.License.OrganizationID != org.ID {
		t.Fatalf("unexpected license organization, want %s but got %s", describeRes.License.OrganizationID, org.ID)
	}

	if !equalMap(describeRes.License.Quotas, createRes.Quotas) {
		t.Fatalf("unexpected license quota values, want %s but got %s", describeRes.License.Quotas, createRes.Quotas)
	}

	// update license

	updateReq := exchange.UpdateLicenseRequest {
		License: describeRes.License,
		Quotas: map[string]string{ "maxWorkspaces": "3" },
		Products: []string{ prod2.ID },
	}
	var updateRes exchange.DescribeLicenseResponse

	res = callChiRouter(
		r, http.MethodPut, "/" + createRes.ID,
		updateReq,
		&updateRes,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("update request failed with %s", res.Status)
	}

	if !equalMap(updateReq.Quotas, updateRes.Quotas) {
		t.Fatalf("unexpected quotas after update, want %+v but got %+v", updateReq.Quotas, updateRes.Quotas)
	}

	updatedProds, _ := licenseService.DescribeLicense(updateRes.License.ID)

	if !equal(updateReq.Products, updatedProds.ProductIDs) {
		t.Fatalf("unexpected product IDs after update, want %+v but got %+v", updateReq.Products, updatedProds.ProductIDs)
	}

	// delete license

	res = callChiRouter(
		r, http.MethodPatch, "/delete/" + createRes.ID,
		nil,
		nil,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("delete request failed with %s", res.Status)
	}

	res = callChiRouter(
		r, http.MethodGet, "/" + createRes.ID, // describe
		nil,
		nil,
	)
	if res.StatusCode == http.StatusOK {
		t.Fatal("expected to fail describe after deleting the license")
	}

}

func equal[T comparable](s1, s2 []T) bool {
	
	for _, e := range s1 {

		found := false

		for _, e2 := range s2 {
			if e == e2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}

	}

	return true
}

func equalMap[K, V comparable](m1, m2 map[K]V) bool {

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