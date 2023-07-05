package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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

func (Account) MiXin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		EditMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		//field.Int("user_id"),
		field.String("open_code"),
		field.String("category"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("user", User.Type),
		edge.From("user", User.Type).
			Ref("accounts").
			Unique(),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		// 非唯一约束索引
		//index.Fields("field1", "field2"),
		// 唯一约束索引
		index.Fields("id").
			Edges("user").
			Unique(),
	}
}
