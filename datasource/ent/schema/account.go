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
		EditMixin{},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		//field.Int("user_id"),
		//field.String("open_code"),
		//field.String("category"),
		field.String("username").NotEmpty().Unique().Immutable(),
		field.String("email").NotEmpty().Unique(),
		field.String("phone").NotEmpty().Unique().MinLen(11).MaxLen(11),
		field.String("password").Sensitive().MinLen(8).MinLen(120),
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
