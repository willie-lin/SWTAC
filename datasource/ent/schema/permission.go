package schema

import "github.com/facebook/ent"

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return nil
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
