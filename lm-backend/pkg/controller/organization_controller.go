package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"license-manager/pkg/controller/exchange"
	"license-manager/pkg/domain"
	"license-manager/pkg/repositories"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type organizationController struct {
	orgRepo repositories.OrganizationRepository
}

func NewOrganizationController(orgRepo repositories.OrganizationRepository) *organizationController {
	return &organizationController{
		orgRepo: orgRepo,
	}
}

func (oc *organizationController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", oc.listAllOrgs)
	router.Get("/{id}", oc.describeOrg)
	router.Put("/{id}", oc.updateOrg)
	router.Post("/", oc.createOrg)
	router.Patch("/delete/{id}", oc.deleteOrg)
	router.Patch("/restore/{id}", oc.restoreOrg)

	return router
}

func (oc *organizationController) listAllOrgs(w http.ResponseWriter, r *http.Request) {

	orgs := oc.orgRepo.FindAll()

	response := exchange.ListAllOrgsResponse{
		Organizations: make([]exchange.ListAllOrgsItem, len(orgs)),
	}

	for i, org := range orgs {
		response.Organizations[i] = exchange.ListAllOrgsItem{
			ID:       org.ID,
			Name:     org.Name,
			Location: org.Location,
		}
	}

	sendJSON(w, response)
}

func (oc *organizationController) describeOrg(w http.ResponseWriter, r *http.Request) {

	orgID := chi.URLParam(r, "id")

	org, err := oc.orgRepo.FindByID(orgID)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot describe organization with ID %s, reason: %s", orgID, err.Error()), http.StatusInternalServerError)
		return
	}

	sendJSON(w, exchange.DescribeOrgResponse{
		ID:        org.ID,
		Name:      org.Name,
		Location:  org.Location,
		ContactID: org.ContactID,
	})
}

func (oc *organizationController) updateOrg(w http.ResponseWriter, r *http.Request) {

	orgID := chi.URLParam(r, "id")

	var updateReq exchange.UpdateOrgRequest
	if err := readJSON(r, &updateReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updated, err := oc.orgRepo.UpdateByID(domain.Organization{
		ID: orgID,
		Name: updateReq.Name,
		Location: updateReq.Location,
		ContactID: "", // does not update contact ID
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, fmt.Sprintf("unable to update organization %s, no changes made in DB", orgID), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (oc *organizationController) createOrg(w http.ResponseWriter, r *http.Request) {

	var createReq exchange.CreateOrgRequest
	if err := readJSON(r, &createReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := oc.orgRepo.Save(domain.Organization{
		Name: createReq.Name,
		Location: createReq.Location,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (oc *organizationController) deleteOrg(w http.ResponseWriter, r *http.Request) {

	orgID := chi.URLParam(r, "id")

	err := oc.orgRepo.DeleteByID(orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (oc *organizationController) restoreOrg(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logical org delete/restore
	http.Error(w, "unsupported operation", http.StatusNotImplemented)
}

func sendJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func readJSON(r *http.Request, ptr any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, ptr)
}
