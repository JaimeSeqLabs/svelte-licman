// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"license-manager/pkg/repositories/ent-fw/ent/jwttoken"
	"license-manager/pkg/repositories/ent-fw/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// JwtToken is the model entity for the JwtToken schema.
type JwtToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Revoked holds the value of the "revoked" field.
	Revoked bool `json:"revoked,omitempty"`
	// Claims holds the value of the "claims" field.
	Claims map[string]interface{} `json:"claims,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JwtTokenQuery when eager-loading is set.
	Edges       JwtTokenEdges `json:"edges"`
	user_issued *int
}

// JwtTokenEdges holds the relations/edges for other nodes in the graph.
type JwtTokenEdges struct {
	// Issuer holds the value of the issuer edge.
	Issuer *User `json:"issuer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// IssuerOrErr returns the Issuer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JwtTokenEdges) IssuerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Issuer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Issuer, nil
	}
	return nil, &NotLoadedError{edge: "issuer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JwtToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jwttoken.FieldClaims:
			values[i] = new([]byte)
		case jwttoken.FieldRevoked:
			values[i] = new(sql.NullBool)
		case jwttoken.FieldID:
			values[i] = new(sql.NullInt64)
		case jwttoken.FieldToken:
			values[i] = new(sql.NullString)
		case jwttoken.ForeignKeys[0]: // user_issued
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type JwtToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JwtToken fields.
func (jt *JwtToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jwttoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			jt.ID = int(value.Int64)
		case jwttoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				jt.Token = value.String
			}
		case jwttoken.FieldRevoked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field revoked", values[i])
			} else if value.Valid {
				jt.Revoked = value.Bool
			}
		case jwttoken.FieldClaims:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field claims", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &jt.Claims); err != nil {
					return fmt.Errorf("unmarshal field claims: %w", err)
				}
			}
		case jwttoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_issued", value)
			} else if value.Valid {
				jt.user_issued = new(int)
				*jt.user_issued = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryIssuer queries the "issuer" edge of the JwtToken entity.
func (jt *JwtToken) QueryIssuer() *UserQuery {
	return NewJwtTokenClient(jt.config).QueryIssuer(jt)
}

// Update returns a builder for updating this JwtToken.
// Note that you need to call JwtToken.Unwrap() before calling this method if this JwtToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (jt *JwtToken) Update() *JwtTokenUpdateOne {
	return NewJwtTokenClient(jt.config).UpdateOne(jt)
}

// Unwrap unwraps the JwtToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jt *JwtToken) Unwrap() *JwtToken {
	_tx, ok := jt.config.driver.(*txDriver)
	if !ok {
		panic("ent: JwtToken is not a transactional entity")
	}
	jt.config.driver = _tx.drv
	return jt
}

// String implements the fmt.Stringer.
func (jt *JwtToken) String() string {
	var builder strings.Builder
	builder.WriteString("JwtToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jt.ID))
	builder.WriteString("token=")
	builder.WriteString(jt.Token)
	builder.WriteString(", ")
	builder.WriteString("revoked=")
	builder.WriteString(fmt.Sprintf("%v", jt.Revoked))
	builder.WriteString(", ")
	builder.WriteString("claims=")
	builder.WriteString(fmt.Sprintf("%v", jt.Claims))
	builder.WriteByte(')')
	return builder.String()
}

// JwtTokens is a parsable slice of JwtToken.
type JwtTokens []*JwtToken

func (jt JwtTokens) config(cfg config) {
	for _i := range jt {
		jt[_i].config = cfg
	}
}
