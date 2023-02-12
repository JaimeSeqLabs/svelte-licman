package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Credentials holds the schema definition for the Credentials entity.
type Credentials struct {
	ent.Schema
}

// Fields of the Credentials.
func (Credentials) Fields() []ent.Field {
	return []ent.Field {
		field.String("username").NotEmpty(),
		field.String("password_hash").NotEmpty(),
		field.JSON("claims", map[string]any{}),
	}
}

// Edges of the Credentials.
func (Credentials) Edges() []ent.Edge {
	return nil
}
