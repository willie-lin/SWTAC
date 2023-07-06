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

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

func (UserGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_groups"},
	}
}

func (UserGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("parent_id"),
		field.String("name"),
		field.String("code"),
		field.String("intro"),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("roles", Role.Type),
	}
}

// Index

func (UserGroup) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		// 唯一索引
		index.Fields("parent_id", "code").Unique(),
	}

}
