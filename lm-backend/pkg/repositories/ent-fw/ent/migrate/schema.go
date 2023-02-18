// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "mail", Type: field.TypeString},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
	}
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "claims", Type: field.TypeJSON},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
	}
	// JwtTokensColumns holds the columns for the "jwt_tokens" table.
	JwtTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "revoked", Type: field.TypeBool, Default: false},
		{Name: "claims", Type: field.TypeJSON},
		{Name: "user_issued", Type: field.TypeString},
	}
	// JwtTokensTable holds the schema information for the "jwt_tokens" table.
	JwtTokensTable = &schema.Table{
		Name:       "jwt_tokens",
		Columns:    JwtTokensColumns,
		PrimaryKey: []*schema.Column{JwtTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "jwt_tokens_users_issued",
				Columns:    []*schema.Column{JwtTokensColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "location", Type: field.TypeString},
		{Name: "contact_id", Type: field.TypeString, Nullable: true},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:       "organizations",
		Columns:    OrganizationsColumns,
		PrimaryKey: []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organizations_contacts_contact",
				Columns:    []*schema.Column{OrganizationsColumns[3]},
				RefColumns: []*schema.Column{ContactsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "mail", Type: field.TypeString, Unique: true},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "claims", Type: field.TypeJSON},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactsTable,
		CredentialsTable,
		JwtTokensTable,
		OrganizationsTable,
		UsersTable,
	}
)

func init() {
	JwtTokensTable.ForeignKeys[0].RefTable = UsersTable
	OrganizationsTable.ForeignKeys[0].RefTable = ContactsTable
}
