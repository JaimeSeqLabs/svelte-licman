package controller

import (
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type licenseController struct {
	licenseService service.LicenseService
	certificateService service.CertificateService
}

func NewLicenseController(licenseService service.LicenseService, certificateService service.CertificateService) *licenseController {
	return &licenseController{ licenseService, certificateService}
}

func (lc *licenseController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", lc.ListAllLicenses)
    r.Post("/", lc.CreateLicense)
    r.Get("/{id}", lc.DescribeLicense)
    r.Put("/{id}", lc.UpdateLicense)
    r.Post("/decode", lc.DecodeLicense)
    r.Get("/status/{id}", lc.GetLicenseStatus)
    r.Patch("/expire/{id}", lc.ExpireLicense)
    r.Get("/download/{id}", lc.DownloadLicenseFile)
    r.Patch("/delete/{id}", lc.DeleteLicense)
    r.Patch("/restore/{id}", lc.RestoreLicense)

	return r
}

func (lc *licenseController) ListAllLicenses(w http.ResponseWriter, r *http.Request) {
	
	licenses, err := lc.licenseService.ListLicenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := exchange.ListAllLicensesResponse {
		Licenses: make([]exchange.ListAllLicensesItem, len(licenses)),
	}

	for i, lic := range licenses {
		resp.Licenses[i] = exchange.ListAllLicensesItem(lic)
	}

	sendJSON(w, resp)
}

func (lc *licenseController) CreateLicense(w http.ResponseWriter, r *http.Request) {

	var req exchange.CreateLicenseRequest
	
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	created, err := lc.licenseService.CreateLicense(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, created)
}

func (lc *licenseController) DescribeLicense(w http.ResponseWriter, r *http.Request) {
	licenseID := chi.URLParam(r, "id")

	license, err := lc.licenseService.DescribeLicense(licenseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSON(w, exchange.DescribeLicenseResponse {
		License: license,
		Quotas: license.Quotas,
	})
}

func (lc *licenseController) DecodeLicense(w http.ResponseWriter, r *http.Request) {
	
	var req exchange.DecodeLicenseRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := lc.certificateService.DescribeCertificateData(req.Encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	license, err := lc.licenseService.DescribeLicense(data.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSON(w, exchange.DescribeLicenseResponse {
		License: license,
		Quotas: license.Quotas,
	})
}

func (lc *licenseController) DeleteLicense(w http.ResponseWriter, r *http.Request) {

	licenseID := chi.URLParam(r, "id")

	if err := lc.licenseService.DeleteLicense(licenseID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK) // no content
}

func (lc *licenseController) RestoreLicense(w http.ResponseWriter, r *http.Request) {
	// TODO
	http.Error(w, "unimplemented method", http.StatusNotImplemented)
}

func (lc *licenseController) DownloadLicenseFile(w http.ResponseWriter, r *http.Request) {

	licenseID := chi.URLParam(r, "id")

	license, err := lc.licenseService.DescribeLicense(licenseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cert, err := lc.certificateService.CreateCertificate(license)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment; filename=cert.txt")
	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write([]byte(cert))
}

func (lc *licenseController) ExpireLicense(w http.ResponseWriter, r *http.Request) {
	
	licenseID := chi.URLParam(r, "id")

	if err := lc.licenseService.SuspendLicense(licenseID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func (lc *licenseController) GetLicenseStatus(w http.ResponseWriter, r *http.Request) {
	
	licenseID := chi.URLParam(r, "id")

	license, err := lc.licenseService.DescribeLicense(licenseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sendJSON(w, exchange.DescribeLicenseStatusResponse {
		ID: license.ID,
		OrganizationID: license.OrganizationID,
		Mail: license.Mail,
		ActivationDate: license.ActivationDate,
		ExpirationDate: license.ExpirationDate,
		Status: license.Status,
		Contact: license.Contact,
		ProductIDs: license.ProductIDs,
		Quotas: license.Quotas,
	})
}

func (lc *licenseController) UpdateLicense(w http.ResponseWriter, r *http.Request) {

	licenseID := chi.URLParam(r, "id")

	var req exchange.UpdateLicenseRequest
	if err := readJSON(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := lc.licenseService.UpdateLicense(licenseID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, exchange.DescribeLicenseResponse {
		License: updated,
		Quotas: updated.Quotas,
	})

}
