// Code generated by ent, DO NOT EDIT.

package credentials

const (
	// Label holds the string label denoting the credentials type in the database.
	Label = "credentials"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldClaims holds the string denoting the claims field in the database.
	FieldClaims = "claims"
	// Table holds the table name of the credentials in the database.
	Table = "credentials"
)

// Columns holds all SQL columns for credentials fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPasswordHash,
	FieldClaims,
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
