package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// JwtToken holds the schema definition for the JwtToken entity.
type JwtToken struct {
	ent.Schema
}

// Fields of the JwtToken.
func (JwtToken) Fields() []ent.Field {
	return []ent.Field {
		field.String("token").NotEmpty().Unique(),
		field.Bool("revoked").Default(false),
	}
}

// Edges of the JwtToken.
func (JwtToken) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("claims", Claims.Type).Required(),
	}
}