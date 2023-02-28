package schema

import (
	"time"

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
		field.String("contact"),
		field.String("mail"),
		field.String("address"),
		field.String("zipcode"),
		field.String("country").
			NotEmpty(),
		field.Time("date_created").
			Default(time.Now),
		field.Time("last_updated").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("licenses", License.Type),
	}
}
