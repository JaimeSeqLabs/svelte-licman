package controller

import (
	"fmt"
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

	router.Get("/", oc.ListAllOrgs)
	router.Get("/{id}", oc.DescribeOrg)
	router.Put("/{id}", oc.UpdateOrg)
	router.Post("/", oc.CreateOrg)
	router.Patch("/delete/{id}", oc.DeleteOrg)
	router.Patch("/restore/{id}", oc.RestoreOrg)

	return router
}

func (oc *organizationController) ListAllOrgs(w http.ResponseWriter, r *http.Request) {

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

func (oc *organizationController) DescribeOrg(w http.ResponseWriter, r *http.Request) {

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

func (oc *organizationController) UpdateOrg(w http.ResponseWriter, r *http.Request) {

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

func (oc *organizationController) CreateOrg(w http.ResponseWriter, r *http.Request) {

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

func (oc *organizationController) DeleteOrg(w http.ResponseWriter, r *http.Request) {

	orgID := chi.URLParam(r, "id")

	err := oc.orgRepo.DeleteByID(orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (oc *organizationController) RestoreOrg(w http.ResponseWriter, r *http.Request) {
	// TODO: implement logical org delete/restore
	http.Error(w, "unsupported operation", http.StatusNotImplemented)
}
