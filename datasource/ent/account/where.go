// Code generated by ent, DO NOT EDIT.

package account

import (
	"SWTAC/datasource/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// OpenCode applies equality check predicate on the "open_code" field. It's identical to OpenCodeEQ.
func OpenCode(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldOpenCode, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCategory, v))
}

// OpenCodeEQ applies the EQ predicate on the "open_code" field.
func OpenCodeEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldOpenCode, v))
}

// OpenCodeNEQ applies the NEQ predicate on the "open_code" field.
func OpenCodeNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldOpenCode, v))
}

// OpenCodeIn applies the In predicate on the "open_code" field.
func OpenCodeIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldOpenCode, vs...))
}

// OpenCodeNotIn applies the NotIn predicate on the "open_code" field.
func OpenCodeNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldOpenCode, vs...))
}

// OpenCodeGT applies the GT predicate on the "open_code" field.
func OpenCodeGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldOpenCode, v))
}

// OpenCodeGTE applies the GTE predicate on the "open_code" field.
func OpenCodeGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldOpenCode, v))
}

// OpenCodeLT applies the LT predicate on the "open_code" field.
func OpenCodeLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldOpenCode, v))
}

// OpenCodeLTE applies the LTE predicate on the "open_code" field.
func OpenCodeLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldOpenCode, v))
}

// OpenCodeContains applies the Contains predicate on the "open_code" field.
func OpenCodeContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldOpenCode, v))
}

// OpenCodeHasPrefix applies the HasPrefix predicate on the "open_code" field.
func OpenCodeHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldOpenCode, v))
}

// OpenCodeHasSuffix applies the HasSuffix predicate on the "open_code" field.
func OpenCodeHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldOpenCode, v))
}

// OpenCodeEqualFold applies the EqualFold predicate on the "open_code" field.
func OpenCodeEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldOpenCode, v))
}

// OpenCodeContainsFold applies the ContainsFold predicate on the "open_code" field.
func OpenCodeContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldOpenCode, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldCategory, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}