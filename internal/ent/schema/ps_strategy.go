package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type PsStrategy struct {
	ent.Schema
}

// Mixin of the PsStrategy.
func (PsStrategy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the PsStrategy.
func (PsStrategy) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Positive().
			Unique().
			Immutable(),
		field.Uint64("owner"),
		field.Text("script_content"),
		field.Int("is_delete").
			Default(0),
	}
}

// Edges of the PsStrategy.
func (PsStrategy) Edges() []ent.Edge {
	return nil
}
