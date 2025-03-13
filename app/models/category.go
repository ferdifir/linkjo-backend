package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Categories schema definition
type Categories struct {
	ent.Schema
}

// Fields of the Categories.
func (Categories) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique(),
		field.String("name").
			NotEmpty().
			MaxLen(255),
		field.Uint("parent_id").
			Optional(),
	}
}

// Edges of the Categories.
func (Categories) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Categories.Type).
			From("parent").
			Unique().
			Field("parent_id"),
	}
}
