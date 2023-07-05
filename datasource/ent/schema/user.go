package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

func (User) MiXin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("username").Unique(),
		field.String("nickname"),
		field.String("avatar"),
		field.Int("age").Positive(),
		field.String("city").Optional(),
		field.String("introduction").Optional(),
		field.String("email"),
		field.String("phone"),
		field.String("password").Sensitive(),
		field.Int("state"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_group", UserGroup.Type).Ref("user"),
		edge.To("role", Role.Type),
		edge.To("accounts", Account.Type),
	}
}

// Index

func (User) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		index.Fields("age"),
		// 唯一索引
		index.Fields("id", "username").Unique(),
	}

}
