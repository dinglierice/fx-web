package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields TODO 部分字段Mixin化
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("id").
			Positive().
			Immutable().
			StructTag(`json:"id,omitempty"`),
		field.Time("created_at").
			Optional().
			Nillable(),
		field.Time("updated_at").
			Optional().
			Nillable(),
		field.Time("deleted_at").
			Optional().
			Nillable(),
		field.String("user_name").
			Unique().
			MaxLen(256),
		field.String("email").
			Optional().
			MaxLen(256),
		field.String("password_digest").
			Optional().
			MaxLen(256),
		field.String("nick_name").
			Optional().
			MaxLen(256),
		field.String("status").
			Optional().
			MaxLen(256),
		field.String("avatar").
			Optional().
			MaxLen(1000),
		field.String("money").
			Optional().
			MaxLen(256),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
