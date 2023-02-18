package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// JwtToken holds the schema definition for the JwtToken entity.
type JwtToken struct {
	ent.Schema
}

// Fields of the JwtToken.
func (JwtToken) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").DefaultFunc(uuid.NewString),
		field.
			String("token").
			NotEmpty().
			Unique(),
		field.
			Bool("revoked").
			Default(false),
		field.
			JSON("claims", map[string]any{}),
	}
}

// Edges of the JwtToken.
func (JwtToken) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("issuer", User.Type).
			Ref("issued").
			Unique().
			Required(),
	}
}
