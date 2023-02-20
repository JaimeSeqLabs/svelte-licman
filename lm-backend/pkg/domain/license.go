package domain

import "time"

type License struct {
	
	ID string

	Features string
	Status string // archived, suspended, expired, active
	Version string
	
	Note string
	Contact string
	Mail string
	
	ProductIDs []string
	OrganizationID string

	Secret string // json ignore

	ExpirationDate time.Time
	ActivationDate time.Time
	DateCreated time.Time
	LastUpdated time.Time
	LastAccessed time.Time

	AccessCount int
	LastAccessIP string
}