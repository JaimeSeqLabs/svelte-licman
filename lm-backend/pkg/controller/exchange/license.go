package exchange

import (
	"license-manager/pkg/domain"
	"time"
)

type CreateLicenseRequest struct {
	
	Features string `json:"features"`
	Status string `json:"status"` // archived, suspended, expired, active
	Version string `json:"version"`
	
	Note string `json:"note"`
	Contact string `json:"contact"`
	Mail string `json:"mail"`
	
	// NOTE: skus not IDs
	ProductSKUs []string `json:"product_skus"`
	// NOTE: name not ID
	OrganizationName string `json:"organization_name"`

	Secret string `json:"secret"`

	ExpirationDate time.Time `json:"expiration_date"`
	ActivationDate time.Time `json:"activation_date"`
}

type UpdateLicenseRequest struct {
	
	Features string `json:"features"`

	Status string `json:"status"` // archived, suspended, expired, active
	Version string `json:"version"`
	
	Note string `json:"note"`
	Contact string `json:"contact"`
	Mail string `json:"mail"`
	
	ExpirationDate time.Time `json:"expiration_date"`
}

type ListAllLicensesResponse struct {
	Licenses []ListAllLicensesItem `json:"licenses"`
}

type ListAllLicensesItem domain.License