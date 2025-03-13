package models

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin untuk timestamps dan soft delete
type TimeMixin struct {
	mixin.Schema
}

// Fields untuk CreatedAt, UpdatedAt, dan DeletedAt
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(), // Tidak bisa diubah setelah dibuat
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now), // Selalu update saat ada perubahan
		field.Time("deleted_at").
			Optional().
			Nillable(), // Bisa null untuk soft delete
	}
}

func (TimeMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if m.Op().Is(ent.OpDelete) {
					if err := m.SetField("deleted_at", time.Now()); err != nil {
						return nil, err
					}
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}
