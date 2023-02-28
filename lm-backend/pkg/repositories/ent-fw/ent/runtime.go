// Code generated by ent, DO NOT EDIT.

package ent

import (
	"license-manager/pkg/repositories/ent-fw/ent/contact"
	"license-manager/pkg/repositories/ent-fw/ent/credentials"
	"license-manager/pkg/repositories/ent-fw/ent/jwttoken"
	"license-manager/pkg/repositories/ent-fw/ent/license"
	"license-manager/pkg/repositories/ent-fw/ent/organization"
	"license-manager/pkg/repositories/ent-fw/ent/product"
	"license-manager/pkg/repositories/ent-fw/ent/schema"
	"license-manager/pkg/repositories/ent-fw/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescName is the schema descriptor for name field.
	contactDescName := contactFields[1].Descriptor()
	// contact.NameValidator is a validator for the "name" field. It is called by the builders before save.
	contact.NameValidator = contactDescName.Validators[0].(func(string) error)
	// contactDescMail is the schema descriptor for mail field.
	contactDescMail := contactFields[2].Descriptor()
	// contact.MailValidator is a validator for the "mail" field. It is called by the builders before save.
	contact.MailValidator = contactDescMail.Validators[0].(func(string) error)
	// contactDescID is the schema descriptor for id field.
	contactDescID := contactFields[0].Descriptor()
	// contact.DefaultID holds the default value on creation for the id field.
	contact.DefaultID = contactDescID.Default.(func() string)
	credentialsFields := schema.Credentials{}.Fields()
	_ = credentialsFields
	// credentialsDescUsername is the schema descriptor for username field.
	credentialsDescUsername := credentialsFields[1].Descriptor()
	// credentials.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	credentials.UsernameValidator = credentialsDescUsername.Validators[0].(func(string) error)
	// credentialsDescPasswordHash is the schema descriptor for password_hash field.
	credentialsDescPasswordHash := credentialsFields[2].Descriptor()
	// credentials.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	credentials.PasswordHashValidator = credentialsDescPasswordHash.Validators[0].(func(string) error)
	// credentialsDescID is the schema descriptor for id field.
	credentialsDescID := credentialsFields[0].Descriptor()
	// credentials.DefaultID holds the default value on creation for the id field.
	credentials.DefaultID = credentialsDescID.Default.(func() string)
	jwttokenFields := schema.JwtToken{}.Fields()
	_ = jwttokenFields
	// jwttokenDescToken is the schema descriptor for token field.
	jwttokenDescToken := jwttokenFields[1].Descriptor()
	// jwttoken.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	jwttoken.TokenValidator = jwttokenDescToken.Validators[0].(func(string) error)
	// jwttokenDescRevoked is the schema descriptor for revoked field.
	jwttokenDescRevoked := jwttokenFields[2].Descriptor()
	// jwttoken.DefaultRevoked holds the default value on creation for the revoked field.
	jwttoken.DefaultRevoked = jwttokenDescRevoked.Default.(bool)
	// jwttokenDescID is the schema descriptor for id field.
	jwttokenDescID := jwttokenFields[0].Descriptor()
	// jwttoken.DefaultID holds the default value on creation for the id field.
	jwttoken.DefaultID = jwttokenDescID.Default.(func() string)
	licenseFields := schema.License{}.Fields()
	_ = licenseFields
	// licenseDescQuotas is the schema descriptor for quotas field.
	licenseDescQuotas := licenseFields[7].Descriptor()
	// license.DefaultQuotas holds the default value on creation for the quotas field.
	license.DefaultQuotas = licenseDescQuotas.Default.(map[string]string)
	// licenseDescAccessCount is the schema descriptor for access_count field.
	licenseDescAccessCount := licenseFields[13].Descriptor()
	// license.DefaultAccessCount holds the default value on creation for the access_count field.
	license.DefaultAccessCount = licenseDescAccessCount.Default.(int)
	// licenseDescDateCreated is the schema descriptor for date_created field.
	licenseDescDateCreated := licenseFields[14].Descriptor()
	// license.DefaultDateCreated holds the default value on creation for the date_created field.
	license.DefaultDateCreated = licenseDescDateCreated.Default.(func() time.Time)
	// licenseDescLastUpdated is the schema descriptor for last_updated field.
	licenseDescLastUpdated := licenseFields[15].Descriptor()
	// license.DefaultLastUpdated holds the default value on creation for the last_updated field.
	license.DefaultLastUpdated = licenseDescLastUpdated.Default.(func() time.Time)
	// license.UpdateDefaultLastUpdated holds the default value on update for the last_updated field.
	license.UpdateDefaultLastUpdated = licenseDescLastUpdated.UpdateDefault.(func() time.Time)
	// licenseDescID is the schema descriptor for id field.
	licenseDescID := licenseFields[0].Descriptor()
	// license.DefaultID holds the default value on creation for the id field.
	license.DefaultID = licenseDescID.Default.(func() string)
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[1].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
	// organizationDescCountry is the schema descriptor for country field.
	organizationDescCountry := organizationFields[6].Descriptor()
	// organization.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	organization.CountryValidator = organizationDescCountry.Validators[0].(func(string) error)
	// organizationDescDateCreated is the schema descriptor for date_created field.
	organizationDescDateCreated := organizationFields[7].Descriptor()
	// organization.DefaultDateCreated holds the default value on creation for the date_created field.
	organization.DefaultDateCreated = organizationDescDateCreated.Default.(func() time.Time)
	// organizationDescLastUpdated is the schema descriptor for last_updated field.
	organizationDescLastUpdated := organizationFields[8].Descriptor()
	// organization.DefaultLastUpdated holds the default value on creation for the last_updated field.
	organization.DefaultLastUpdated = organizationDescLastUpdated.Default.(func() time.Time)
	// organization.UpdateDefaultLastUpdated holds the default value on update for the last_updated field.
	organization.UpdateDefaultLastUpdated = organizationDescLastUpdated.UpdateDefault.(func() time.Time)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationFields[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() string)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescSku is the schema descriptor for sku field.
	productDescSku := productFields[1].Descriptor()
	// product.SkuValidator is a validator for the "sku" field. It is called by the builders before save.
	product.SkuValidator = productDescSku.Validators[0].(func(string) error)
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[2].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescInstallInstr is the schema descriptor for install_instr field.
	productDescInstallInstr := productFields[3].Descriptor()
	// product.DefaultInstallInstr holds the default value on creation for the install_instr field.
	product.DefaultInstallInstr = productDescInstallInstr.Default.(string)
	// productDescLicenseCount is the schema descriptor for license_count field.
	productDescLicenseCount := productFields[4].Descriptor()
	// product.DefaultLicenseCount holds the default value on creation for the license_count field.
	product.DefaultLicenseCount = productDescLicenseCount.Default.(int)
	// productDescDateCreated is the schema descriptor for date_created field.
	productDescDateCreated := productFields[5].Descriptor()
	// product.DefaultDateCreated holds the default value on creation for the date_created field.
	product.DefaultDateCreated = productDescDateCreated.Default.(func() time.Time)
	// productDescLastUpdated is the schema descriptor for last_updated field.
	productDescLastUpdated := productFields[6].Descriptor()
	// product.DefaultLastUpdated holds the default value on creation for the last_updated field.
	product.DefaultLastUpdated = productDescLastUpdated.Default.(func() time.Time)
	// product.UpdateDefaultLastUpdated holds the default value on update for the last_updated field.
	product.UpdateDefaultLastUpdated = productDescLastUpdated.UpdateDefault.(func() time.Time)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescMail is the schema descriptor for mail field.
	userDescMail := userFields[2].Descriptor()
	// user.MailValidator is a validator for the "mail" field. It is called by the builders before save.
	user.MailValidator = userDescMail.Validators[0].(func(string) error)
	// userDescPasswordHash is the schema descriptor for password_hash field.
	userDescPasswordHash := userFields[3].Descriptor()
	// user.PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	user.PasswordHashValidator = userDescPasswordHash.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
}
