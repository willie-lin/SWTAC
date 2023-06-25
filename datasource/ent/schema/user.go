package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) MiXin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("name"),
		field.String("nickname").
			Unique(),
		field.Int("age"),
		field.String("city"),
		field.String("introduction"),
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
