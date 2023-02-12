// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/claims"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Claims is the model entity for the Claims schema.
type Claims struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Values holds the value of the "values" field.
	Values             string `json:"values,omitempty"`
	credentials_claims *int
	jwt_token_claims   *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Claims) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case claims.FieldID:
			values[i] = new(sql.NullInt64)
		case claims.FieldValues:
			values[i] = new(sql.NullString)
		case claims.ForeignKeys[0]: // credentials_claims
			values[i] = new(sql.NullInt64)
		case claims.ForeignKeys[1]: // jwt_token_claims
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Claims", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Claims fields.
func (c *Claims) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case claims.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case claims.FieldValues:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field values", values[i])
			} else if value.Valid {
				c.Values = value.String
			}
		case claims.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field credentials_claims", value)
			} else if value.Valid {
				c.credentials_claims = new(int)
				*c.credentials_claims = int(value.Int64)
			}
		case claims.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field jwt_token_claims", value)
			} else if value.Valid {
				c.jwt_token_claims = new(int)
				*c.jwt_token_claims = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Claims.
// Note that you need to call Claims.Unwrap() before calling this method if this Claims
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Claims) Update() *ClaimsUpdateOne {
	return NewClaimsClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Claims entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Claims) Unwrap() *Claims {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Claims is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Claims) String() string {
	var builder strings.Builder
	builder.WriteString("Claims(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("values=")
	builder.WriteString(c.Values)
	builder.WriteByte(')')
	return builder.String()
}

// ClaimsSlice is a parsable slice of Claims.
type ClaimsSlice []*Claims

func (c ClaimsSlice) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}