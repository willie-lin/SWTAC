package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
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
		//field.Int("id").Unique(),
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("nickname").Unique(),
		field.String("avatar").Default("https://example.com/default-avatar.png").Optional(),
		field.Int("age").Positive().Min(0).Max(150).Default(1).Optional(),
		field.Enum("gender").Values("male", "female", "other").Optional(),
		field.String("city").Optional(),
		field.String("introduction").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", UserGroup.Type).Ref("users"),
		edge.To("roles", Role.Type),
		edge.To("accounts", Account.Type),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{

		// 非唯一的普通索引
		index.Fields("age"),
		// 唯一索引
		index.Fields("id", "nickname").Unique(),
	}
}

// Hooks of the User.
//func (User) Hooks() []ent.Hook {
//	return []ent.Hook{
//		hook.On(validateAge, ent.OpCreate, ent.OpUpdate),
//	}
//}

// validateAge is a validator that ensures the age is between 0 and 150.
func validateAge(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		age, exists := m.Field("age")
		if !exists || age.(int) >= 0 && age.(int) <= 150 {
			return next.Mutate(ctx, m)
		}
		return nil, fmt.Errorf("invalid age: %d", age)
	})
}
