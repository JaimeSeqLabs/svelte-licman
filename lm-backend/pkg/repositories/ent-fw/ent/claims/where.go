// Code generated by ent, DO NOT EDIT.

package claims

import (
	"license-manager/pkg/repositories/ent-fw/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Claims {
	return predicate.Claims(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Claims {
	return predicate.Claims(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Claims {
	return predicate.Claims(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Claims {
	return predicate.Claims(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Claims {
	return predicate.Claims(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Claims {
	return predicate.Claims(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Claims {
	return predicate.Claims(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Claims {
	return predicate.Claims(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Claims {
	return predicate.Claims(sql.FieldLTE(FieldID, id))
}

// Values applies equality check predicate on the "values" field. It's identical to ValuesEQ.
func Values(v string) predicate.Claims {
	return predicate.Claims(sql.FieldEQ(FieldValues, v))
}

// ValuesEQ applies the EQ predicate on the "values" field.
func ValuesEQ(v string) predicate.Claims {
	return predicate.Claims(sql.FieldEQ(FieldValues, v))
}

// ValuesNEQ applies the NEQ predicate on the "values" field.
func ValuesNEQ(v string) predicate.Claims {
	return predicate.Claims(sql.FieldNEQ(FieldValues, v))
}

// ValuesIn applies the In predicate on the "values" field.
func ValuesIn(vs ...string) predicate.Claims {
	return predicate.Claims(sql.FieldIn(FieldValues, vs...))
}

// ValuesNotIn applies the NotIn predicate on the "values" field.
func ValuesNotIn(vs ...string) predicate.Claims {
	return predicate.Claims(sql.FieldNotIn(FieldValues, vs...))
}

// ValuesGT applies the GT predicate on the "values" field.
func ValuesGT(v string) predicate.Claims {
	return predicate.Claims(sql.FieldGT(FieldValues, v))
}

// ValuesGTE applies the GTE predicate on the "values" field.
func ValuesGTE(v string) predicate.Claims {
	return predicate.Claims(sql.FieldGTE(FieldValues, v))
}

// ValuesLT applies the LT predicate on the "values" field.
func ValuesLT(v string) predicate.Claims {
	return predicate.Claims(sql.FieldLT(FieldValues, v))
}

// ValuesLTE applies the LTE predicate on the "values" field.
func ValuesLTE(v string) predicate.Claims {
	return predicate.Claims(sql.FieldLTE(FieldValues, v))
}

// ValuesContains applies the Contains predicate on the "values" field.
func ValuesContains(v string) predicate.Claims {
	return predicate.Claims(sql.FieldContains(FieldValues, v))
}

// ValuesHasPrefix applies the HasPrefix predicate on the "values" field.
func ValuesHasPrefix(v string) predicate.Claims {
	return predicate.Claims(sql.FieldHasPrefix(FieldValues, v))
}

// ValuesHasSuffix applies the HasSuffix predicate on the "values" field.
func ValuesHasSuffix(v string) predicate.Claims {
	return predicate.Claims(sql.FieldHasSuffix(FieldValues, v))
}

// ValuesEqualFold applies the EqualFold predicate on the "values" field.
func ValuesEqualFold(v string) predicate.Claims {
	return predicate.Claims(sql.FieldEqualFold(FieldValues, v))
}

// ValuesContainsFold applies the ContainsFold predicate on the "values" field.
func ValuesContainsFold(v string) predicate.Claims {
	return predicate.Claims(sql.FieldContainsFold(FieldValues, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Claims) predicate.Claims {
	return predicate.Claims(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Claims) predicate.Claims {
	return predicate.Claims(func(s *sql.Selector) {
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
func Not(p predicate.Claims) predicate.Claims {
	return predicate.Claims(func(s *sql.Selector) {
		p(s.Not())
	})
}
