// Code generated by ent, DO NOT EDIT.

package organization

import (
	"license-manager/pkg/repositories/ent-fw/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLocation, v))
}

// ContactID applies equality check predicate on the "contact_id" field. It's identical to ContactIDEQ.
func ContactID(v int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldContactID, v))
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

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldLocation, v))
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldLocation, v))
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldLocation, vs...))
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldLocation, vs...))
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldLocation, v))
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldLocation, v))
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldLocation, v))
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldLocation, v))
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldLocation, v))
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldLocation, v))
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldLocation, v))
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldLocation, v))
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldLocation, v))
}

// ContactIDEQ applies the EQ predicate on the "contact_id" field.
func ContactIDEQ(v int) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldContactID, v))
}

// ContactIDNEQ applies the NEQ predicate on the "contact_id" field.
func ContactIDNEQ(v int) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldContactID, v))
}

// ContactIDIn applies the In predicate on the "contact_id" field.
func ContactIDIn(vs ...int) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldContactID, vs...))
}

// ContactIDNotIn applies the NotIn predicate on the "contact_id" field.
func ContactIDNotIn(vs ...int) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldContactID, vs...))
}

// ContactIDIsNil applies the IsNil predicate on the "contact_id" field.
func ContactIDIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldContactID))
}

// ContactIDNotNil applies the NotNil predicate on the "contact_id" field.
func ContactIDNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldContactID))
}

// HasContact applies the HasEdge predicate on the "contact" edge.
func HasContact() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ContactTable, ContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactWith applies the HasEdge predicate on the "contact" edge with a given conditions (other predicates).
func HasContactWith(preds ...predicate.Contact) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ContactInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ContactTable, ContactColumn),
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