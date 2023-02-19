package controller

import (
	"fmt"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type productController struct {
	productRepo repositories.ProductRepository
}

func NewProductController(productRepo repositories.ProductRepository) *productController {
	return &productController {
		productRepo: productRepo,
	}
}

func (pc *productController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", pc.ListAllProducts)
	router.Get("/{id}", pc.DescribeProduct)
	router.Put("/{id}", pc.UpdateProduct)
	router.Post("/", pc.CreateProduct)
	router.Patch("/delete/{id}", pc.DeleteProduct)
	router.Patch("/restore/{id}", pc.RestoreProduct)

	return router
}

func (pc *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req exchange.CreateProductRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := pc.productRepo.Save(domain.Product {
		SKU: req.SKU,
		Name: req.Name,
		InstallInstructions: req.InstallInstructions,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (pc *productController) ListAllProducts(w http.ResponseWriter, r *http.Request) {

	products := pc.productRepo.FindAll()

	response := exchange.ListAllProductsResponse{
		Products: make([]exchange.ListAllProductsItem, len(products)),
	}

	for i, prod := range products {
		response.Products[i] = exchange.ListAllProductsItem {
			ID: prod.ID,
			SKU: prod.SKU,
			Name: prod.Name,
			InstallInstructions: prod.InstallInstructions,
			LicenseCount: prod.LicenseCount,
			DateCreated: prod.DateCreated,
			LastUpdated: prod.LastUpdated,
		}
	}

	sendJSON(w, response)

}

func (pc *productController) DescribeProduct(w http.ResponseWriter, r *http.Request) {

	prodID := chi.URLParam(r, "id")

	product, err := pc.productRepo.FindByID(prodID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, exchange.DescribeProductResponse {
		ID: product.ID,
		SKU: product.SKU,
		Name: product.Name,
		InstallInstructions: product.InstallInstructions,
		LicenseCount: product.LicenseCount,
		DateCreated: product.DateCreated,
		LastUpdated: product.LastUpdated,
	})

}

func (pc *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	prodID := chi.URLParam(r, "id")

	var req exchange.UpdateProductRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := pc.productRepo.UpdateByID(domain.Product {
		ID: prodID,
		SKU: req.SKU,
		Name: req.Name,
		InstallInstructions: req.InstallInstructions,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, fmt.Sprintf("unable to update product %s, no changes made in DB", prodID), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (pc *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	prodID := chi.URLParam(r, "id")

	err := pc.productRepo.DeleteByID(domain.Product { ID: prodID })
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (pc *productController) RestoreProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logical org delete/restore
	http.Error(w, "unsupported operation", http.StatusNotImplemented)
}


