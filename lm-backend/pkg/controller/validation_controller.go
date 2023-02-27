package controller

import (
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/service"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type validationController struct {
	certService service.CertificateService
}

func NewValidationController(certService service.CertificateService) *validationController {
	return &validationController { certService }
}

func (vc *validationController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post("/", vc.Check)
	router.Post("/{product}", vc.CheckProduct)

	router.Post("/v2", vc.CheckV2)
	router.Post("/v2/{product}", vc.CheckProductV2)

	return router
}

func (vc *validationController) Check(w http.ResponseWriter, r *http.Request) {
	
	var req exchange.LicenseValidateRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := vc.certService.ValidateCertificate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK) // empty body
}

func (vc *validationController) CheckProduct(w http.ResponseWriter, r *http.Request) {
	
	prodSKU := chi.URLParam(r, "product")

	var req exchange.LicenseValidateRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ProductSKU = strings.ToUpper(prodSKU)

	_, err := vc.certService.ValidateCertificate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK) // empty body
}

func (vc *validationController) CheckV2(w http.ResponseWriter, r *http.Request) {
	
	var req exchange.LicenseValidateRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	license, err := vc.certService.ValidateCertificate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSON(w, exchange.LicenseValidateResponse {
		LicenseID: license.ID,
		LicenseVersion: license.Version,
		Quotas: license.Quotas,
	})
}

func (vc *validationController) CheckProductV2(w http.ResponseWriter, r *http.Request) {
	
	prodSKU := chi.URLParam(r, "product")

	var req exchange.LicenseValidateRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.ProductSKU = strings.ToUpper(prodSKU)
	
	license, err := vc.certService.ValidateCertificate(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSON(w, exchange.LicenseValidateResponse {
		LicenseID: license.ID,
		LicenseVersion: license.Version,
		Quotas: license.Quotas,
	})
}