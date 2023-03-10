package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").
			DefaultFunc(uuid.NewString),
		field.String("sku").
			NotEmpty().
			Unique(),
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("install_instr").
			Default(""),
		field.Int("license_count"). // TODO: remove, this can be obtained from edges
			Default(0),

		field.Time("date_created").
			Default(time.Now),
		field.Time("last_updated").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("license", License.Type).
			Ref("license_products"),
	}
}
