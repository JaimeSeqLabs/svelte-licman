package domain

import "time"

type Organization struct {

	ID string `json:"id"`
	Name string `json:"name"`
	Contact string `json:"contact"`
	Mail string `json:"mail"`
	Address string `json:"address"`
	ZipCode string `json:"zipcode"`
	Country string `json:"country"`
	//LicenseCount int
	Licenses []string `json:"licenses"`// license IDs

	DateCreated time.Time `json:"date_created"`
	LastUpdated time.Time `json:"last_updated"`

	
}
