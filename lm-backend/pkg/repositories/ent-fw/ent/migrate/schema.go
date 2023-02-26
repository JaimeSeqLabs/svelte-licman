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
		{Name: "issuer_id", Type: field.TypeString},
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
	// LicensesColumns holds the columns for the "licenses" table.
	LicensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "features", Type: field.TypeString},
		{Name: "status", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "note", Type: field.TypeString},
		{Name: "contact", Type: field.TypeString},
		{Name: "mail", Type: field.TypeString},
		{Name: "quotas", Type: field.TypeJSON},
		{Name: "secret", Type: field.TypeString},
		{Name: "expiration_date", Type: field.TypeTime},
		{Name: "activation_date", Type: field.TypeTime},
		{Name: "last_accessed", Type: field.TypeTime, Nullable: true},
		{Name: "last_access_ip", Type: field.TypeString, Nullable: true},
		{Name: "access_count", Type: field.TypeInt, Default: 0},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "last_updated", Type: field.TypeTime},
		{Name: "organization_licenses", Type: field.TypeString, Nullable: true},
	}
	// LicensesTable holds the schema information for the "licenses" table.
	LicensesTable = &schema.Table{
		Name:       "licenses",
		Columns:    LicensesColumns,
		PrimaryKey: []*schema.Column{LicensesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "licenses_organizations_licenses",
				Columns:    []*schema.Column{LicensesColumns[16]},
				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
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
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "sku", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "install_instr", Type: field.TypeString, Default: ""},
		{Name: "license_count", Type: field.TypeInt, Default: 0},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "last_updated", Type: field.TypeTime},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
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
	// LicenseLicenseProductsColumns holds the columns for the "license_license_products" table.
	LicenseLicenseProductsColumns = []*schema.Column{
		{Name: "license_id", Type: field.TypeString},
		{Name: "product_id", Type: field.TypeString},
	}
	// LicenseLicenseProductsTable holds the schema information for the "license_license_products" table.
	LicenseLicenseProductsTable = &schema.Table{
		Name:       "license_license_products",
		Columns:    LicenseLicenseProductsColumns,
		PrimaryKey: []*schema.Column{LicenseLicenseProductsColumns[0], LicenseLicenseProductsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "license_license_products_license_id",
				Columns:    []*schema.Column{LicenseLicenseProductsColumns[0]},
				RefColumns: []*schema.Column{LicensesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "license_license_products_product_id",
				Columns:    []*schema.Column{LicenseLicenseProductsColumns[1]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContactsTable,
		CredentialsTable,
		JwtTokensTable,
		LicensesTable,
		OrganizationsTable,
		ProductsTable,
		UsersTable,
		LicenseLicenseProductsTable,
	}
)

func init() {
	JwtTokensTable.ForeignKeys[0].RefTable = UsersTable
	LicensesTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationsTable.ForeignKeys[0].RefTable = ContactsTable
	LicenseLicenseProductsTable.ForeignKeys[0].RefTable = LicensesTable
	LicenseLicenseProductsTable.ForeignKeys[1].RefTable = ProductsTable
}
