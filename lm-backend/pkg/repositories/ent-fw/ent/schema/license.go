package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return []ent.Field {
		
		field.String("id").DefaultFunc(uuid.NewString),
		
		field.String("features"),
		field.String("status"), // TODO: use enums here
		field.String("version"),

		field.String("note"),
		field.String("contact"),
		field.String("mail"),

		field.String("secret").
			StructTag(`json:"-"`), // ensure this is not exported to json

		field.Time("expiration_date"),
		field.Time("activation_date"),
		
		field.Time("last_accessed").
			Optional(),
		field.String("last_access_IP").
			Optional(),
		field.Int("access_count").
			Default(0),
		
		field.Time("date_created").
			Default(time.Now),
		field.Time("last_updated").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the License.
func (License) Edges() []ent.Edge {
	return []ent.Edge {
		
		edge.To("license_products", Product.Type).
			Required(),

		edge.From("owner_org", Organization.Type).
			Ref("licenses").
			Unique(),
	}
}
