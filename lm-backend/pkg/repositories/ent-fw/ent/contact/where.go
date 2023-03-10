// Code generated by ent, DO NOT EDIT.

package contact

import (
	"license-manager/pkg/repositories/ent-fw/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Contact {
	return predicate.Contact(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Contact {
	return predicate.Contact(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Contact {
	return predicate.Contact(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Contact {
	return predicate.Contact(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldName, v))
}

// Mail applies equality check predicate on the "mail" field. It's identical to MailEQ.
func Mail(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldMail, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Contact {
	return predicate.Contact(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Contact {
	return predicate.Contact(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Contact {
	return predicate.Contact(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Contact {
	return predicate.Contact(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Contact {
	return predicate.Contact(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Contact {
	return predicate.Contact(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Contact {
	return predicate.Contact(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Contact {
	return predicate.Contact(sql.FieldContainsFold(FieldName, v))
}

// MailEQ applies the EQ predicate on the "mail" field.
func MailEQ(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEQ(FieldMail, v))
}

// MailNEQ applies the NEQ predicate on the "mail" field.
func MailNEQ(v string) predicate.Contact {
	return predicate.Contact(sql.FieldNEQ(FieldMail, v))
}

// MailIn applies the In predicate on the "mail" field.
func MailIn(vs ...string) predicate.Contact {
	return predicate.Contact(sql.FieldIn(FieldMail, vs...))
}

// MailNotIn applies the NotIn predicate on the "mail" field.
func MailNotIn(vs ...string) predicate.Contact {
	return predicate.Contact(sql.FieldNotIn(FieldMail, vs...))
}

// MailGT applies the GT predicate on the "mail" field.
func MailGT(v string) predicate.Contact {
	return predicate.Contact(sql.FieldGT(FieldMail, v))
}

// MailGTE applies the GTE predicate on the "mail" field.
func MailGTE(v string) predicate.Contact {
	return predicate.Contact(sql.FieldGTE(FieldMail, v))
}

// MailLT applies the LT predicate on the "mail" field.
func MailLT(v string) predicate.Contact {
	return predicate.Contact(sql.FieldLT(FieldMail, v))
}

// MailLTE applies the LTE predicate on the "mail" field.
func MailLTE(v string) predicate.Contact {
	return predicate.Contact(sql.FieldLTE(FieldMail, v))
}

// MailContains applies the Contains predicate on the "mail" field.
func MailContains(v string) predicate.Contact {
	return predicate.Contact(sql.FieldContains(FieldMail, v))
}

// MailHasPrefix applies the HasPrefix predicate on the "mail" field.
func MailHasPrefix(v string) predicate.Contact {
	return predicate.Contact(sql.FieldHasPrefix(FieldMail, v))
}

// MailHasSuffix applies the HasSuffix predicate on the "mail" field.
func MailHasSuffix(v string) predicate.Contact {
	return predicate.Contact(sql.FieldHasSuffix(FieldMail, v))
}

// MailEqualFold applies the EqualFold predicate on the "mail" field.
func MailEqualFold(v string) predicate.Contact {
	return predicate.Contact(sql.FieldEqualFold(FieldMail, v))
}

// MailContainsFold applies the ContainsFold predicate on the "mail" field.
func MailContainsFold(v string) predicate.Contact {
	return predicate.Contact(sql.FieldContainsFold(FieldMail, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
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
func Not(p predicate.Contact) predicate.Contact {
	return predicate.Contact(func(s *sql.Selector) {
		p(s.Not())
	})
}
