package domain

import "time"

const (
	LicenseCurrentVersion = "v1"
)

type License struct {
	
	ID string `json:"id"`

	Features string `json:"features"`
	Status string `json:"status"` // archived, suspended, expired, active
	Version string `json:"version"`
	
	Note string `json:"note"`
	Contact string `json:"contact"`
	Mail string `json:"mail"`
	
	ProductIDs []string `json:"product_ids"`
	OrganizationID string `json:"organization_id"`

	Secret string `json:"-"`

	ExpirationDate time.Time `json:"expiration_date"`
	ActivationDate time.Time `json:"activation_date"`
	DateCreated time.Time `json:"date_created"`
	LastUpdated time.Time `json:"last_updated"`
	LastAccessed time.Time `json:"last_accessed"`

	AccessCount int `json:"access_count"`
	LastAccessIP string `json:"-"`
}