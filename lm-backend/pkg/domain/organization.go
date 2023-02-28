package domain

import "time"

type Organization struct {

	ID string
	Name string
	Contact string
	Mail string
	Address string
	ZipCode string
	Country string
	//LicenseCount int
	Licenses []string // license IDs

	DateCreated time.Time
	LastUpdated time.Time

	
}
