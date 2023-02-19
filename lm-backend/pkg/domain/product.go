package domain

import "time"

type Product struct {
	ID                  string
	SKU                 string
	Name                string
	InstallInstructions string
	LicenseCount        int

	DateCreated time.Time
	LastUpdated time.Time
}