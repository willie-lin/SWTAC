// Code generated by ent, DO NOT EDIT.

package usergroup

import (
	"SWTAC/datasource/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreator, v))
}

// Editor applies equality check predicate on the "editor" field. It's identical to EditorEQ.
func Editor(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldEditor, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldDeleted, v))
}

// ParentID applies equality check predicate on the "parent_id" field. It's identical to ParentIDEQ.
func ParentID(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldParentID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCode, v))
}

// Intro applies equality check predicate on the "intro" field. It's identical to IntroEQ.
func Intro(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldIntro, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldCreator, v))
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldCreator, v))
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldCreator, v))
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldCreator, v))
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldCreator, v))
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldCreator, v))
}

// EditorEQ applies the EQ predicate on the "editor" field.
func EditorEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldEditor, v))
}

// EditorNEQ applies the NEQ predicate on the "editor" field.
func EditorNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldEditor, v))
}

// EditorIn applies the In predicate on the "editor" field.
func EditorIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldEditor, vs...))
}

// EditorNotIn applies the NotIn predicate on the "editor" field.
func EditorNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldEditor, vs...))
}

// EditorGT applies the GT predicate on the "editor" field.
func EditorGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldEditor, v))
}

// EditorGTE applies the GTE predicate on the "editor" field.
func EditorGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldEditor, v))
}

// EditorLT applies the LT predicate on the "editor" field.
func EditorLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldEditor, v))
}

// EditorLTE applies the LTE predicate on the "editor" field.
func EditorLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldEditor, v))
}

// EditorContains applies the Contains predicate on the "editor" field.
func EditorContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldEditor, v))
}

// EditorHasPrefix applies the HasPrefix predicate on the "editor" field.
func EditorHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldEditor, v))
}

// EditorHasSuffix applies the HasSuffix predicate on the "editor" field.
func EditorHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldEditor, v))
}

// EditorEqualFold applies the EqualFold predicate on the "editor" field.
func EditorEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldEditor, v))
}

// EditorContainsFold applies the ContainsFold predicate on the "editor" field.
func EditorContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldEditor, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldDeleted, v))
}

// DeletedIn applies the In predicate on the "deleted" field.
func DeletedIn(vs ...float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldDeleted, vs...))
}

// DeletedNotIn applies the NotIn predicate on the "deleted" field.
func DeletedNotIn(vs ...float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldDeleted, vs...))
}

// DeletedGT applies the GT predicate on the "deleted" field.
func DeletedGT(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldDeleted, v))
}

// DeletedGTE applies the GTE predicate on the "deleted" field.
func DeletedGTE(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldDeleted, v))
}

// DeletedLT applies the LT predicate on the "deleted" field.
func DeletedLT(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldDeleted, v))
}

// DeletedLTE applies the LTE predicate on the "deleted" field.
func DeletedLTE(v float64) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldDeleted, v))
}

// ParentIDEQ applies the EQ predicate on the "parent_id" field.
func ParentIDEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldParentID, v))
}

// ParentIDNEQ applies the NEQ predicate on the "parent_id" field.
func ParentIDNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldParentID, v))
}

// ParentIDIn applies the In predicate on the "parent_id" field.
func ParentIDIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldParentID, vs...))
}

// ParentIDNotIn applies the NotIn predicate on the "parent_id" field.
func ParentIDNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldParentID, vs...))
}

// ParentIDGT applies the GT predicate on the "parent_id" field.
func ParentIDGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldParentID, v))
}

// ParentIDGTE applies the GTE predicate on the "parent_id" field.
func ParentIDGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldParentID, v))
}

// ParentIDLT applies the LT predicate on the "parent_id" field.
func ParentIDLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldParentID, v))
}

// ParentIDLTE applies the LTE predicate on the "parent_id" field.
func ParentIDLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldParentID, v))
}

// ParentIDContains applies the Contains predicate on the "parent_id" field.
func ParentIDContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldParentID, v))
}

// ParentIDHasPrefix applies the HasPrefix predicate on the "parent_id" field.
func ParentIDHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldParentID, v))
}

// ParentIDHasSuffix applies the HasSuffix predicate on the "parent_id" field.
func ParentIDHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldParentID, v))
}

// ParentIDEqualFold applies the EqualFold predicate on the "parent_id" field.
func ParentIDEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldParentID, v))
}

// ParentIDContainsFold applies the ContainsFold predicate on the "parent_id" field.
func ParentIDContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldParentID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldName, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldCode, v))
}

// IntroEQ applies the EQ predicate on the "intro" field.
func IntroEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEQ(FieldIntro, v))
}

// IntroNEQ applies the NEQ predicate on the "intro" field.
func IntroNEQ(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNEQ(FieldIntro, v))
}

// IntroIn applies the In predicate on the "intro" field.
func IntroIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldIn(FieldIntro, vs...))
}

// IntroNotIn applies the NotIn predicate on the "intro" field.
func IntroNotIn(vs ...string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldNotIn(FieldIntro, vs...))
}

// IntroGT applies the GT predicate on the "intro" field.
func IntroGT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGT(FieldIntro, v))
}

// IntroGTE applies the GTE predicate on the "intro" field.
func IntroGTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldGTE(FieldIntro, v))
}

// IntroLT applies the LT predicate on the "intro" field.
func IntroLT(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLT(FieldIntro, v))
}

// IntroLTE applies the LTE predicate on the "intro" field.
func IntroLTE(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldLTE(FieldIntro, v))
}

// IntroContains applies the Contains predicate on the "intro" field.
func IntroContains(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContains(FieldIntro, v))
}

// IntroHasPrefix applies the HasPrefix predicate on the "intro" field.
func IntroHasPrefix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasPrefix(FieldIntro, v))
}

// IntroHasSuffix applies the HasSuffix predicate on the "intro" field.
func IntroHasSuffix(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldHasSuffix(FieldIntro, v))
}

// IntroEqualFold applies the EqualFold predicate on the "intro" field.
func IntroEqualFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldEqualFold(FieldIntro, v))
}

// IntroContainsFold applies the ContainsFold predicate on the "intro" field.
func IntroContainsFold(v string) predicate.UserGroup {
	return predicate.UserGroup(sql.FieldContainsFold(FieldIntro, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoles applies the HasEdge predicate on the "roles" edge.
func HasRoles() predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RolesTable, RolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRolesWith applies the HasEdge predicate on the "roles" edge with a given conditions (other predicates).
func HasRolesWith(preds ...predicate.Role) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		step := newRolesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
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
func Not(p predicate.UserGroup) predicate.UserGroup {
	return predicate.UserGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
