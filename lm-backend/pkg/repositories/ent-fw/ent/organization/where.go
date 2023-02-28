// Code generated by ent, DO NOT EDIT.

package organization

import (
	"license-manager/pkg/repositories/ent-fw/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// Contact applies equality check predicate on the "contact" field. It's identical to ContactEQ.
func Contact(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldContact, v))
}

// Mail applies equality check predicate on the "mail" field. It's identical to MailEQ.
func Mail(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldMail, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAddress, v))
}

// Zipcode applies equality check predicate on the "zipcode" field. It's identical to ZipcodeEQ.
func Zipcode(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldZipcode, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCountry, v))
}

// DateCreated applies equality check predicate on the "date_created" field. It's identical to DateCreatedEQ.
func DateCreated(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateCreated, v))
}

// LastUpdated applies equality check predicate on the "last_updated" field. It's identical to LastUpdatedEQ.
func LastUpdated(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLastUpdated, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldName, v))
}

// ContactEQ applies the EQ predicate on the "contact" field.
func ContactEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldContact, v))
}

// ContactNEQ applies the NEQ predicate on the "contact" field.
func ContactNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldContact, v))
}

// ContactIn applies the In predicate on the "contact" field.
func ContactIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldContact, vs...))
}

// ContactNotIn applies the NotIn predicate on the "contact" field.
func ContactNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldContact, vs...))
}

// ContactGT applies the GT predicate on the "contact" field.
func ContactGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldContact, v))
}

// ContactGTE applies the GTE predicate on the "contact" field.
func ContactGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldContact, v))
}

// ContactLT applies the LT predicate on the "contact" field.
func ContactLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldContact, v))
}

// ContactLTE applies the LTE predicate on the "contact" field.
func ContactLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldContact, v))
}

// ContactContains applies the Contains predicate on the "contact" field.
func ContactContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldContact, v))
}

// ContactHasPrefix applies the HasPrefix predicate on the "contact" field.
func ContactHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldContact, v))
}

// ContactHasSuffix applies the HasSuffix predicate on the "contact" field.
func ContactHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldContact, v))
}

// ContactEqualFold applies the EqualFold predicate on the "contact" field.
func ContactEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldContact, v))
}

// ContactContainsFold applies the ContainsFold predicate on the "contact" field.
func ContactContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldContact, v))
}

// MailEQ applies the EQ predicate on the "mail" field.
func MailEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldMail, v))
}

// MailNEQ applies the NEQ predicate on the "mail" field.
func MailNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldMail, v))
}

// MailIn applies the In predicate on the "mail" field.
func MailIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldMail, vs...))
}

// MailNotIn applies the NotIn predicate on the "mail" field.
func MailNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldMail, vs...))
}

// MailGT applies the GT predicate on the "mail" field.
func MailGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldMail, v))
}

// MailGTE applies the GTE predicate on the "mail" field.
func MailGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldMail, v))
}

// MailLT applies the LT predicate on the "mail" field.
func MailLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldMail, v))
}

// MailLTE applies the LTE predicate on the "mail" field.
func MailLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldMail, v))
}

// MailContains applies the Contains predicate on the "mail" field.
func MailContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldMail, v))
}

// MailHasPrefix applies the HasPrefix predicate on the "mail" field.
func MailHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldMail, v))
}

// MailHasSuffix applies the HasSuffix predicate on the "mail" field.
func MailHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldMail, v))
}

// MailEqualFold applies the EqualFold predicate on the "mail" field.
func MailEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldMail, v))
}

// MailContainsFold applies the ContainsFold predicate on the "mail" field.
func MailContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldMail, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldAddress, v))
}

// ZipcodeEQ applies the EQ predicate on the "zipcode" field.
func ZipcodeEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldZipcode, v))
}

// ZipcodeNEQ applies the NEQ predicate on the "zipcode" field.
func ZipcodeNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldZipcode, v))
}

// ZipcodeIn applies the In predicate on the "zipcode" field.
func ZipcodeIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldZipcode, vs...))
}

// ZipcodeNotIn applies the NotIn predicate on the "zipcode" field.
func ZipcodeNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldZipcode, vs...))
}

// ZipcodeGT applies the GT predicate on the "zipcode" field.
func ZipcodeGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldZipcode, v))
}

// ZipcodeGTE applies the GTE predicate on the "zipcode" field.
func ZipcodeGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldZipcode, v))
}

// ZipcodeLT applies the LT predicate on the "zipcode" field.
func ZipcodeLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldZipcode, v))
}

// ZipcodeLTE applies the LTE predicate on the "zipcode" field.
func ZipcodeLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldZipcode, v))
}

// ZipcodeContains applies the Contains predicate on the "zipcode" field.
func ZipcodeContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldZipcode, v))
}

// ZipcodeHasPrefix applies the HasPrefix predicate on the "zipcode" field.
func ZipcodeHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldZipcode, v))
}

// ZipcodeHasSuffix applies the HasSuffix predicate on the "zipcode" field.
func ZipcodeHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldZipcode, v))
}

// ZipcodeEqualFold applies the EqualFold predicate on the "zipcode" field.
func ZipcodeEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldZipcode, v))
}

// ZipcodeContainsFold applies the ContainsFold predicate on the "zipcode" field.
func ZipcodeContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldZipcode, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldCountry, v))
}

// DateCreatedEQ applies the EQ predicate on the "date_created" field.
func DateCreatedEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDateCreated, v))
}

// DateCreatedNEQ applies the NEQ predicate on the "date_created" field.
func DateCreatedNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDateCreated, v))
}

// DateCreatedIn applies the In predicate on the "date_created" field.
func DateCreatedIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDateCreated, vs...))
}

// DateCreatedNotIn applies the NotIn predicate on the "date_created" field.
func DateCreatedNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDateCreated, vs...))
}

// DateCreatedGT applies the GT predicate on the "date_created" field.
func DateCreatedGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDateCreated, v))
}

// DateCreatedGTE applies the GTE predicate on the "date_created" field.
func DateCreatedGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDateCreated, v))
}

// DateCreatedLT applies the LT predicate on the "date_created" field.
func DateCreatedLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDateCreated, v))
}

// DateCreatedLTE applies the LTE predicate on the "date_created" field.
func DateCreatedLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDateCreated, v))
}

// LastUpdatedEQ applies the EQ predicate on the "last_updated" field.
func LastUpdatedEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLastUpdated, v))
}

// LastUpdatedNEQ applies the NEQ predicate on the "last_updated" field.
func LastUpdatedNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldLastUpdated, v))
}

// LastUpdatedIn applies the In predicate on the "last_updated" field.
func LastUpdatedIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldLastUpdated, vs...))
}

// LastUpdatedNotIn applies the NotIn predicate on the "last_updated" field.
func LastUpdatedNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldLastUpdated, vs...))
}

// LastUpdatedGT applies the GT predicate on the "last_updated" field.
func LastUpdatedGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldLastUpdated, v))
}

// LastUpdatedGTE applies the GTE predicate on the "last_updated" field.
func LastUpdatedGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldLastUpdated, v))
}

// LastUpdatedLT applies the LT predicate on the "last_updated" field.
func LastUpdatedLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldLastUpdated, v))
}

// LastUpdatedLTE applies the LTE predicate on the "last_updated" field.
func LastUpdatedLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldLastUpdated, v))
}

// HasLicenses applies the HasEdge predicate on the "licenses" edge.
func HasLicenses() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LicensesTable, LicensesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLicensesWith applies the HasEdge predicate on the "licenses" edge with a given conditions (other predicates).
func HasLicensesWith(preds ...predicate.License) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LicensesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LicensesTable, LicensesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		p(s.Not())
	})
}
