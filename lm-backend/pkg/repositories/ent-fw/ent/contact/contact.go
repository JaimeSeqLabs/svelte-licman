// Code generated by ent, DO NOT EDIT.

package contact

const (
	// Label holds the string label denoting the contact type in the database.
	Label = "contact"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldMail holds the string denoting the mail field in the database.
	FieldMail = "mail"
	// Table holds the table name of the contact in the database.
	Table = "contacts"
)

// Columns holds all SQL columns for contact fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldMail,
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
	// MailValidator is a validator for the "mail" field. It is called by the builders before save.
	MailValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
