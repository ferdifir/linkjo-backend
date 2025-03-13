package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id").
			Unique(),
		field.String("owner_name").
			NotEmpty(),
		field.String("outlet_name").
			NotEmpty(),
		field.String("email").
			Unique().
			NotEmpty(),
		field.String("password").
			Sensitive(),
		field.String("phone").
			Unique().
			NotEmpty(),
		field.String("role").
			Default("user"),
		field.String("photo").
			Default("default.png"),
		field.String("address").
			Optional(),
		field.String("city").
			Optional(),
		field.Float("latitude").
			Default(0),
		field.Float("longitude").
			Default(0),
		field.Bool("is_active").
			Default(true),
	}
}

// Mixin for created_at, updated_at, and soft delete (deleted_at).
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
