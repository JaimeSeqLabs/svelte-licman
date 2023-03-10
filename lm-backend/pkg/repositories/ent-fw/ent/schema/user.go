package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").DefaultFunc(uuid.NewString),
		field.String("username").NotEmpty(),
		field.String("mail").NotEmpty().Unique(),
		field.String("password_hash").NotEmpty(),
		field.JSON("claims", map[string]any{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("issued", JwtToken.Type),
	}
}
