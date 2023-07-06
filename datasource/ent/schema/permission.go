package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permissions"},
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Int("parent_id"),
		field.String("code"),
		field.String("name"),
		field.String("intro"),
		field.Int("category"),
		field.Int("url"),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("roles", Role.Type).Ref("roles"),
		edge.From("roles", Role.Type).Ref("permissions"),
		//edge.To("permissions", Permission.Type),
		//edge.To("roles", Role.Type),
	}
}

func (Permission) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		//index.Fields("age"),
		// 唯一索引
		index.Fields("parent_id", "code").Unique(),
	}

}
