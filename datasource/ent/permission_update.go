// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/permission"
	"SWTAC/datasource/ent/predicate"
	"SWTAC/datasource/ent/role"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionMutation
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetCreator sets the "creator" field.
func (pu *PermissionUpdate) SetCreator(s string) *PermissionUpdate {
	pu.mutation.SetCreator(s)
	return pu
}

// SetEditor sets the "editor" field.
func (pu *PermissionUpdate) SetEditor(s string) *PermissionUpdate {
	pu.mutation.SetEditor(s)
	return pu
}

// SetDeleted sets the "deleted" field.
func (pu *PermissionUpdate) SetDeleted(f float64) *PermissionUpdate {
	pu.mutation.ResetDeleted()
	pu.mutation.SetDeleted(f)
	return pu
}

// AddDeleted adds f to the "deleted" field.
func (pu *PermissionUpdate) AddDeleted(f float64) *PermissionUpdate {
	pu.mutation.AddDeleted(f)
	return pu
}

// SetParentID sets the "parent_id" field.
func (pu *PermissionUpdate) SetParentID(i int) *PermissionUpdate {
	pu.mutation.ResetParentID()
	pu.mutation.SetParentID(i)
	return pu
}

// AddParentID adds i to the "parent_id" field.
func (pu *PermissionUpdate) AddParentID(i int) *PermissionUpdate {
	pu.mutation.AddParentID(i)
	return pu
}

// SetCode sets the "code" field.
func (pu *PermissionUpdate) SetCode(s string) *PermissionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetIntro sets the "intro" field.
func (pu *PermissionUpdate) SetIntro(s string) *PermissionUpdate {
	pu.mutation.SetIntro(s)
	return pu
}

// SetCategory sets the "category" field.
func (pu *PermissionUpdate) SetCategory(i int) *PermissionUpdate {
	pu.mutation.ResetCategory()
	pu.mutation.SetCategory(i)
	return pu
}

// AddCategory adds i to the "category" field.
func (pu *PermissionUpdate) AddCategory(i int) *PermissionUpdate {
	pu.mutation.AddCategory(i)
	return pu
}

// SetURL sets the "url" field.
func (pu *PermissionUpdate) SetURL(i int) *PermissionUpdate {
	pu.mutation.ResetURL()
	pu.mutation.SetURL(i)
	return pu
}

// AddURL adds i to the "url" field.
func (pu *PermissionUpdate) AddURL(i int) *PermissionUpdate {
	pu.mutation.AddURL(i)
	return pu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Creator(); ok {
		_spec.SetField(permission.FieldCreator, field.TypeString, value)
	}
	if value, ok := pu.mutation.Editor(); ok {
		_spec.SetField(permission.FieldEditor, field.TypeString, value)
	}
	if value, ok := pu.mutation.Deleted(); ok {
		_spec.SetField(permission.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedDeleted(); ok {
		_spec.AddField(permission.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedParentID(); ok {
		_spec.AddField(permission.FieldParentID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(permission.FieldCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Intro(); ok {
		_spec.SetField(permission.FieldIntro, field.TypeString, value)
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.SetField(permission.FieldCategory, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedCategory(); ok {
		_spec.AddField(permission.FieldCategory, field.TypeInt, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(permission.FieldURL, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedURL(); ok {
		_spec.AddField(permission.FieldURL, field.TypeInt, value)
	}
	if pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetCreator sets the "creator" field.
func (puo *PermissionUpdateOne) SetCreator(s string) *PermissionUpdateOne {
	puo.mutation.SetCreator(s)
	return puo
}

// SetEditor sets the "editor" field.
func (puo *PermissionUpdateOne) SetEditor(s string) *PermissionUpdateOne {
	puo.mutation.SetEditor(s)
	return puo
}

// SetDeleted sets the "deleted" field.
func (puo *PermissionUpdateOne) SetDeleted(f float64) *PermissionUpdateOne {
	puo.mutation.ResetDeleted()
	puo.mutation.SetDeleted(f)
	return puo
}

// AddDeleted adds f to the "deleted" field.
func (puo *PermissionUpdateOne) AddDeleted(f float64) *PermissionUpdateOne {
	puo.mutation.AddDeleted(f)
	return puo
}

// SetParentID sets the "parent_id" field.
func (puo *PermissionUpdateOne) SetParentID(i int) *PermissionUpdateOne {
	puo.mutation.ResetParentID()
	puo.mutation.SetParentID(i)
	return puo
}

// AddParentID adds i to the "parent_id" field.
func (puo *PermissionUpdateOne) AddParentID(i int) *PermissionUpdateOne {
	puo.mutation.AddParentID(i)
	return puo
}

// SetCode sets the "code" field.
func (puo *PermissionUpdateOne) SetCode(s string) *PermissionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetIntro sets the "intro" field.
func (puo *PermissionUpdateOne) SetIntro(s string) *PermissionUpdateOne {
	puo.mutation.SetIntro(s)
	return puo
}

// SetCategory sets the "category" field.
func (puo *PermissionUpdateOne) SetCategory(i int) *PermissionUpdateOne {
	puo.mutation.ResetCategory()
	puo.mutation.SetCategory(i)
	return puo
}

// AddCategory adds i to the "category" field.
func (puo *PermissionUpdateOne) AddCategory(i int) *PermissionUpdateOne {
	puo.mutation.AddCategory(i)
	return puo
}

// SetURL sets the "url" field.
func (puo *PermissionUpdateOne) SetURL(i int) *PermissionUpdateOne {
	puo.mutation.ResetURL()
	puo.mutation.SetURL(i)
	return puo
}

// AddURL adds i to the "url" field.
func (puo *PermissionUpdateOne) AddURL(i int) *PermissionUpdateOne {
	puo.mutation.AddURL(i)
	return puo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Creator(); ok {
		_spec.SetField(permission.FieldCreator, field.TypeString, value)
	}
	if value, ok := puo.mutation.Editor(); ok {
		_spec.SetField(permission.FieldEditor, field.TypeString, value)
	}
	if value, ok := puo.mutation.Deleted(); ok {
		_spec.SetField(permission.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedDeleted(); ok {
		_spec.AddField(permission.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedParentID(); ok {
		_spec.AddField(permission.FieldParentID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(permission.FieldCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Intro(); ok {
		_spec.SetField(permission.FieldIntro, field.TypeString, value)
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.SetField(permission.FieldCategory, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedCategory(); ok {
		_spec.AddField(permission.FieldCategory, field.TypeInt, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(permission.FieldURL, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedURL(); ok {
		_spec.AddField(permission.FieldURL, field.TypeInt, value)
	}
	if puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
