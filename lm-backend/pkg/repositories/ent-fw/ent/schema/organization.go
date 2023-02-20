package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(uuid.NewString),
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("location").
			NotEmpty(),
		field.String("contact_id").
			Optional(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("contact", Contact.Type).
			Unique().
			Field("contact_id"),

		edge.To("licenses", License.Type),

	}
}
