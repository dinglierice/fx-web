package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type PsConfig struct {
	ent.Schema
}

// Fields of the PsConfig.
func (PsConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			StorageKey("ps_id").
			Positive().
			Unique().
			Immutable().
			Comment("主键"),
		field.String("ps_scene").
			NotEmpty().
			Unique().
			Comment("场景"),
		field.Uint64("ps_filter").
			Comment("过滤策略"),
		field.Uint64("ps_message").
			Optional().
			Nillable().
			Comment("消息策略"),
		field.Uint64("ps_event").
			Optional().
			Nillable().
			Comment("时间策略"),
		field.Uint64("ps_feature").
			Optional().
			Nillable().
			Comment("特征策略"),
		field.Uint64("ps_strategy").
			Optional().
			Nillable().
			Comment("脚本"),
		field.Uint64("owner_id").
			Optional().
			Nillable(),
		field.String("managers").
			Optional().
			Nillable(),
		field.Uint64("update_user").
			Optional().
			Nillable(),
	}
}

// Edges of the PsConfig.
func (PsConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("strategy", PsStrategy.Type).
			Unique().
			Field("ps_strategy").
			Comment("关联的策略脚本"),
	}
}

// Mixin of the PsConfig.
func (PsConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Indexes of the PsConfig.
func (PsConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ps_scene").Unique(),
	}
}

func (PsConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ps_configs"},
	}
}
