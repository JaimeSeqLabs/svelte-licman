package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Claims holds the schema definition for the Claims entity.
type Claims struct {
	ent.Schema
}

// Fields of the Claims.
func (Claims) Fields() []ent.Field {
	return []ent.Field {
		field.Text("key").NotEmpty(),
		field.Text("value").NotEmpty(),
	}
}

// Edges of the Claims.
func (Claims) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("claimer", Credentials.Type).
			Ref("claims").
			Unique(),
	}
}
