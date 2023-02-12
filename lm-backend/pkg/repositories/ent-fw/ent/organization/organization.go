// Code generated by ent, DO NOT EDIT.

package organization

const (
	// Label holds the string label denoting the organization type in the database.
	Label = "organization"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldContactID holds the string denoting the contact_id field in the database.
	FieldContactID = "contact_id"
	// EdgeContact holds the string denoting the contact edge name in mutations.
	EdgeContact = "contact"
	// Table holds the table name of the organization in the database.
	Table = "organizations"
	// ContactTable is the table that holds the contact relation/edge.
	ContactTable = "organizations"
	// ContactInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactInverseTable = "contacts"
	// ContactColumn is the table column denoting the contact relation/edge.
	ContactColumn = "contact_id"
)

// Columns holds all SQL columns for organization fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLocation,
	FieldContactID,
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
	// LocationValidator is a validator for the "location" field. It is called by the builders before save.
	LocationValidator func(string) error
)