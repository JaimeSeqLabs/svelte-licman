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

	Quotas map[string]string `json:"quotas"`

	Secret string `json:"secret"`

	ExpirationDate time.Time `json:"expiration_date"`
	ActivationDate time.Time `json:"activation_date"`
}

type UpdateLicenseRequest struct {

	License domain.License `json:"license"`
	Products []string `json:"product_ids"`// product IDs
	Quotas map[string]string`json:"quotas"`
	
}

type ListAllLicensesResponse struct {
	Licenses []ListAllLicensesItem `json:"licenses"`
}

type ListAllLicensesItem domain.License

type DescribeLicenseResponse struct {
	License domain.License `json:"license"`
	Quotas map[string]string `json:"quotas"`
}

type DecodeLicenseRequest struct {
	Encoded string `json:"encoded"`
}

type DescribeLicenseStatusResponse struct {
	
	ID string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Mail string `json:"mail"`
	ActivationDate time.Time `json:"activation_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	Status string `json:"status"` // archived, suspended, expired, active
	Contact string `json:"contact"`
	ProductIDs []string `json:"product_ids"`
	Quotas map[string]string `json:"quotas"`

}
