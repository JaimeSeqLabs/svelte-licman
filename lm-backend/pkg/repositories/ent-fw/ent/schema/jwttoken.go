package schema

import (
	"license-manager/pkg/domain"

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
			JSON("claims", domain.Claims{}),
		field.
			String("issuer_id"),
	}
}

// Edges of the JwtToken.
func (JwtToken) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("issuer", User.Type).
			Ref("issued").
			Unique().
			Required().
			Field("issuer_id"),
	}
}
