// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"
)

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldMail holds the string denoting the mail field in the database.
	FieldMail = "mail"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldZipcode holds the string denoting the zipcode field in the database.
	FieldZipcode = "zipcode"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeLicenses holds the string denoting the licenses edge name in mutations.
	EdgeLicenses = "licenses"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// LicensesTable is the table that holds the licenses relation/edge.
	LicensesTable = "licenses"
	// LicensesInverseTable is the table name for the License entity.
	// It exists in this package in order to avoid circular dependency with the "license" package.
	LicensesInverseTable = "licenses"
	// LicensesColumn is the table column denoting the licenses relation/edge.
	LicensesColumn = "organization_licenses"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldContact,
	FieldMail,
	FieldAddress,
	FieldZipcode,
	FieldCountry,
	FieldDateCreated,
	FieldLastUpdated,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultLastUpdated holds the default value on creation for the "last_updated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "last_updated" field.
	UpdateDefaultLastUpdated func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
