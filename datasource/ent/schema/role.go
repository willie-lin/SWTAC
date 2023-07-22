package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "roles"},
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
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
		//edge.From("users", User.Type).Ref("roles"),
		//edge.To("permissions", Permission.Type),
		//edge.From("groups", UserGroup.Type).Ref("roles"),
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
