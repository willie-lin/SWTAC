package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "accounts"},
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("username").NotEmpty().Unique().Immutable(),
		field.String("email").Unique(),
		field.String("phone").Unique(),
		field.String("password").Sensitive().MinLen(8).MaxLen(120),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("users", User.Type).Ref("accounts"),
	}
}
func (Account) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("username", "email", "phone").Unique(),
	}
}
