package controller_test

import (
	"license-manager/pkg/controller"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories/ent-fw/product"
	"net/http"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestProductCRUD(t *testing.T) {

	db := makeClient(t)
	defer db.Close()

	prodController := controller.NewProductController(
		product_repo.NewProductEntRepo(db),
	)

	r := prodController.Routes()

	testProd := domain.Product {
		SKU: "<sku_code>",
		Name: "TestProduct",
		InstallInstructions: "Press install",
		LicenseCount: 0,
	}

	// create product

	res := callChiRouter(
		r, http.MethodPost, "/",
		exchange.CreateProductRequest {
			SKU: testProd.SKU,
			Name: testProd.Name,
			InstallInstructions: testProd.InstallInstructions,
		},
		nil,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatal("create request failed")
	}

	// list all created

	var listAllRes exchange.ListAllProductsResponse
	callChiRouter(
		r, http.MethodGet, "/",
		nil,
		&listAllRes,
	)

	if len(listAllRes.Products) != 1 {
		t.Fatal("unexpected length when listing all")
	}

	created := listAllRes.Products[0]

	if created.SKU != testProd.SKU {
		t.Fatal("unexpected SKU field when listing all")
	}

	if created.Name != testProd.Name {
		t.Fatal("unexpected Name field when listing all")
	}

	if created.InstallInstructions != testProd.InstallInstructions {
		t.Fatal("unexpected InstallInstructions field when listing all")
	}

	if created.LicenseCount != testProd.LicenseCount {
		t.Fatal("unexpected LicenseCount field when listing all")
	}

	testProd.ID = created.ID
	testProd.LicenseCount = created.LicenseCount
	testProd.DateCreated = created.DateCreated
	testProd.LastUpdated = created.LastUpdated

	// describe by ID

	var describeRes exchange.DescribeProductResponse
	callChiRouter(
		r, http.MethodGet, "/" + testProd.ID,
		nil,
		&describeRes,
	)

	if describeRes.ID != testProd.ID {
		t.Fatal("unexpected ID field when describing by ID")
	}

	if describeRes.SKU != testProd.SKU {
		t.Fatal("unexpected SKU field when describing by ID")
	}

	if describeRes.Name != testProd.Name {
		t.Fatal("unexpected Name field when describing by ID")
	}

	if describeRes.InstallInstructions != testProd.InstallInstructions {
		t.Fatal("unexpected InstallInstructions field when describing by ID")
	}

	if describeRes.LicenseCount != testProd.LicenseCount {
		t.Fatal("unexpected LicenseCount field when describing by ID")
	}

	if describeRes.DateCreated != testProd.DateCreated {
		t.Fatal("unexpected DateCreated field when describing by ID")
	}

	if describeRes.LastUpdated != testProd.LastUpdated {
		t.Fatal("unexpected LastUpdated field when describing by ID")
	}

	// update by ID

	res = callChiRouter(
		r, http.MethodPut, "/" + testProd.ID,
		exchange.UpdateProductRequest {
			SKU: "<sku_code_v2>",
			Name: testProd.Name,
			InstallInstructions: testProd.InstallInstructions,
		},
		nil,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatalf("update request failed")
	}

	var updatedRes exchange.DescribeProductResponse
	callChiRouter(
		r, http.MethodGet, "/" + testProd.ID,
		nil,
		&updatedRes,
	)

	if updatedRes.SKU != "<sku_code_v2>"{
		t.Fatal("update request didn't update SKU code")
	}

	// delete by ID

	res = callChiRouter(
		r, http.MethodPatch, "/delete/" + testProd.ID,
		nil,
		nil,
	)
	if res.StatusCode != http.StatusOK {
		t.Fatal("delete request failed")
	}

	res = callChiRouter(
		r, http.MethodGet, "/" + testProd.ID,
		nil,
		nil, // ignore response
	)

	if res.StatusCode == http.StatusOK {
		t.Fatal("find by ID request expected to fail after delete by ID")
	}

}
