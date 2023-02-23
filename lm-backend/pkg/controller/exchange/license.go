package exchange

import "time"

type CreateLicenseRequest struct {
	
	Features string
	Status string // archived, suspended, expired, active
	Version string
	
	Note string
	Contact string
	Mail string
	
	// NOTE: skus not IDs
	ProductSKUs []string
	// NOTE: name not ID
	OrganizationName string

	Secret string

	ExpirationDate time.Time
	ActivationDate time.Time
}

type UpdateLicenseRequest struct {
	
	Features string

	Status string // archived, suspended, expired, active
	Version string
	
	Note string
	Contact string
	Mail string
	
	ExpirationDate time.Time
}