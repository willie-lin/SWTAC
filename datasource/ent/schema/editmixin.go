package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type EditMixin struct {
	mixin.Schema
}

func (EditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("creator"),
		field.String("editor"),
		field.Float("deleted").SchemaType(map[string]string{
			dialect.MySQL:    "decimal(1,0)",
			dialect.Postgres: "numeric",
		}),
	}

}
