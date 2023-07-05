package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "role"},
	}
}

func (Role) MiXin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("parent_id"),
		field.String("code"),
		field.String("name"),
		field.String("intro"),
		//field.String("icon"),
		//field.String("description"),
		//field.Int("status").Default(1),
		//field.Int("sort").Default(0),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("role"),
		edge.To("permission", Permission.Type),
		edge.From("user_group", UserGroup.Type).Ref("role"),
	}
}

func (Role) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		//index.Fields("age"),
		// 唯一索引
		index.Fields("parent_id", "code").Unique(),
	}

}
