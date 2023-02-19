package schema

import (
	"time"

	"entgo.io/ent"
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
		field.Int("license_count").
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
	return nil
}
