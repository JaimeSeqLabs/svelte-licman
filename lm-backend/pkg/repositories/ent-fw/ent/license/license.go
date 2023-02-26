// Code generated by ent, DO NOT EDIT.

package license

import (
	"time"
)

const (
	// Label holds the string label denoting the license type in the database.
	Label = "license"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFeatures holds the string denoting the features field in the database.
	FieldFeatures = "features"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldMail holds the string denoting the mail field in the database.
	FieldMail = "mail"
	// FieldQuotas holds the string denoting the quotas field in the database.
	FieldQuotas = "quotas"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// FieldActivationDate holds the string denoting the activation_date field in the database.
	FieldActivationDate = "activation_date"
	// FieldLastAccessed holds the string denoting the last_accessed field in the database.
	FieldLastAccessed = "last_accessed"
	// FieldLastAccessIP holds the string denoting the last_access_ip field in the database.
	FieldLastAccessIP = "last_access_ip"
	// FieldAccessCount holds the string denoting the access_count field in the database.
	FieldAccessCount = "access_count"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeLicenseProducts holds the string denoting the license_products edge name in mutations.
	EdgeLicenseProducts = "license_products"
	// EdgeOwnerOrg holds the string denoting the owner_org edge name in mutations.
	EdgeOwnerOrg = "owner_org"
	// Table holds the table name of the license in the database.
	Table = "licenses"
	// LicenseProductsTable is the table that holds the license_products relation/edge. The primary key declared below.
	LicenseProductsTable = "license_license_products"
	// LicenseProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	LicenseProductsInverseTable = "products"
	// OwnerOrgTable is the table that holds the owner_org relation/edge.
	OwnerOrgTable = "licenses"
	// OwnerOrgInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OwnerOrgInverseTable = "organizations"
	// OwnerOrgColumn is the table column denoting the owner_org relation/edge.
	OwnerOrgColumn = "organization_licenses"
)

// Columns holds all SQL columns for license fields.
var Columns = []string{
	FieldID,
	FieldFeatures,
	FieldStatus,
	FieldVersion,
	FieldNote,
	FieldContact,
	FieldMail,
	FieldQuotas,
	FieldSecret,
	FieldExpirationDate,
	FieldActivationDate,
	FieldLastAccessed,
	FieldLastAccessIP,
	FieldAccessCount,
	FieldDateCreated,
	FieldLastUpdated,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "licenses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"organization_licenses",
}

var (
	// LicenseProductsPrimaryKey and LicenseProductsColumn2 are the table columns denoting the
	// primary key for the license_products relation (M2M).
	LicenseProductsPrimaryKey = []string{"license_id", "product_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultQuotas holds the default value on creation for the "quotas" field.
	DefaultQuotas map[string]string
	// DefaultAccessCount holds the default value on creation for the "access_count" field.
	DefaultAccessCount int
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultLastUpdated holds the default value on creation for the "last_updated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "last_updated" field.
	UpdateDefaultLastUpdated func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
