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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Mixin MiXin Mixin User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("nickname").Unique(),
		field.String("avatar").Default("https://example.com/default-avatar.png").Optional(),
		field.Int("age").Positive().Min(0).Max(150).Default(1).Optional(),
		field.Enum("gender").Values("male", "female", "other").Optional(),
		field.String("city").Optional(),
		field.String("introduction").Optional(),
		field.Text("description").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
		edge.From("groups", Group.Type).Ref("users"),
		edge.From("roles", Role.Type).Ref("users"),
	}
}
func (User) Index() []ent.Index {
	return []ent.Index{
		// 非唯一的普通索引
		// 唯一索引
		index.Fields("id").Unique(),
	}
}
