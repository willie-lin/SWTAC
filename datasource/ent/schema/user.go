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

// Mixin MiXin Mixin User
func (User) Mixin() []ent.Mixin {
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
		//// Immutable 函数告诉我们，生下来后你的出生年月就定了，不能改变。你明明是半老徐娘就不能说自己芳龄十八。
		//field.Time("created_at").Immutable().Default(time.Now),
		//field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user_groups", UserGroup.Type).Ref("users"),
		edge.To("roles", Role.Type),
		edge.To("accounts", Account.Type),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		index.Fields("age"),
		// 唯一索引
		index.Fields("id", "username").Unique(),
	}

}
