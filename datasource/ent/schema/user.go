package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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

func (User) MiXin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name").Default(""),
		field.String("nickname").
			Unique(),
		field.Int("age").Positive(),
		field.String("city").Optional(),
		field.String("introduction").Optional(),
		field.String("avatar"),
		field.String("email"),
		field.String("phone"),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Index

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),

		// 非唯一的普通索引
		index.Fields("age"),
		// 唯一索引
		index.Fields("name").Unique(),
	}

}
