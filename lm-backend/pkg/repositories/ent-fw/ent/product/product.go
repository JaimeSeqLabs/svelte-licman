// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSku holds the string denoting the sku field in the database.
	FieldSku = "sku"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldInstallInstr holds the string denoting the install_instr field in the database.
	FieldInstallInstr = "install_instr"
	// FieldLicenseCount holds the string denoting the license_count field in the database.
	FieldLicenseCount = "license_count"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldLastUpdated holds the string denoting the last_updated field in the database.
	FieldLastUpdated = "last_updated"
	// EdgeLicense holds the string denoting the license edge name in mutations.
	EdgeLicense = "license"
	// Table holds the table name of the product in the database.
	Table = "products"
	// LicenseTable is the table that holds the license relation/edge. The primary key declared below.
	LicenseTable = "license_license_products"
	// LicenseInverseTable is the table name for the License entity.
	// It exists in this package in order to avoid circular dependency with the "license" package.
	LicenseInverseTable = "licenses"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldSku,
	FieldName,
	FieldInstallInstr,
	FieldLicenseCount,
	FieldDateCreated,
	FieldLastUpdated,
}

var (
	// LicensePrimaryKey and LicenseColumn2 are the table columns denoting the
	// primary key for the license relation (M2M).
	LicensePrimaryKey = []string{"license_id", "product_id"}
)

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
	// SkuValidator is a validator for the "sku" field. It is called by the builders before save.
	SkuValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultInstallInstr holds the default value on creation for the "install_instr" field.
	DefaultInstallInstr string
	// DefaultLicenseCount holds the default value on creation for the "license_count" field.
	DefaultLicenseCount int
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultLastUpdated holds the default value on creation for the "last_updated" field.
	DefaultLastUpdated func() time.Time
	// UpdateDefaultLastUpdated holds the default value on update for the "last_updated" field.
	UpdateDefaultLastUpdated func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
