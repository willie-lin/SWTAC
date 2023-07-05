// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/predicate"
	"SWTAC/datasource/ent/role"
	"SWTAC/datasource/ent/user"
	"SWTAC/datasource/ent/usergroup"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupUpdate is the builder for updating UserGroup entities.
type UserGroupUpdate struct {
	config
	hooks    []Hook
	mutation *UserGroupMutation
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (ugu *UserGroupUpdate) Where(ps ...predicate.UserGroup) *UserGroupUpdate {
	ugu.mutation.Where(ps...)
	return ugu
}

// SetParentID sets the "parent_id" field.
func (ugu *UserGroupUpdate) SetParentID(s string) *UserGroupUpdate {
	ugu.mutation.SetParentID(s)
	return ugu
}

// SetName sets the "name" field.
func (ugu *UserGroupUpdate) SetName(s string) *UserGroupUpdate {
	ugu.mutation.SetName(s)
	return ugu
}

// SetCode sets the "code" field.
func (ugu *UserGroupUpdate) SetCode(s string) *UserGroupUpdate {
	ugu.mutation.SetCode(s)
	return ugu
}

// SetIntro sets the "intro" field.
func (ugu *UserGroupUpdate) SetIntro(s string) *UserGroupUpdate {
	ugu.mutation.SetIntro(s)
	return ugu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ugu *UserGroupUpdate) AddUserIDs(ids ...int) *UserGroupUpdate {
	ugu.mutation.AddUserIDs(ids...)
	return ugu
}

// AddUsers adds the "users" edges to the User entity.
func (ugu *UserGroupUpdate) AddUsers(u ...*User) *UserGroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ugu.AddUserIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (ugu *UserGroupUpdate) AddRoleIDs(ids ...int) *UserGroupUpdate {
	ugu.mutation.AddRoleIDs(ids...)
	return ugu
}

// AddRoles adds the "roles" edges to the Role entity.
func (ugu *UserGroupUpdate) AddRoles(r ...*Role) *UserGroupUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ugu.AddRoleIDs(ids...)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugu *UserGroupUpdate) Mutation() *UserGroupMutation {
	return ugu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (ugu *UserGroupUpdate) ClearUsers() *UserGroupUpdate {
	ugu.mutation.ClearUsers()
	return ugu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (ugu *UserGroupUpdate) RemoveUserIDs(ids ...int) *UserGroupUpdate {
	ugu.mutation.RemoveUserIDs(ids...)
	return ugu
}

// RemoveUsers removes "users" edges to User entities.
func (ugu *UserGroupUpdate) RemoveUsers(u ...*User) *UserGroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ugu.RemoveUserIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (ugu *UserGroupUpdate) ClearRoles() *UserGroupUpdate {
	ugu.mutation.ClearRoles()
	return ugu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (ugu *UserGroupUpdate) RemoveRoleIDs(ids ...int) *UserGroupUpdate {
	ugu.mutation.RemoveRoleIDs(ids...)
	return ugu
}

// RemoveRoles removes "roles" edges to Role entities.
func (ugu *UserGroupUpdate) RemoveRoles(r ...*Role) *UserGroupUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ugu.RemoveRoleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ugu *UserGroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ugu.sqlSave, ugu.mutation, ugu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ugu *UserGroupUpdate) SaveX(ctx context.Context) int {
	affected, err := ugu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ugu *UserGroupUpdate) Exec(ctx context.Context) error {
	_, err := ugu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugu *UserGroupUpdate) ExecX(ctx context.Context) {
	if err := ugu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ugu *UserGroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(usergroup.Table, usergroup.Columns, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	if ps := ugu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ugu.mutation.ParentID(); ok {
		_spec.SetField(usergroup.FieldParentID, field.TypeString, value)
	}
	if value, ok := ugu.mutation.Name(); ok {
		_spec.SetField(usergroup.FieldName, field.TypeString, value)
	}
	if value, ok := ugu.mutation.Code(); ok {
		_spec.SetField(usergroup.FieldCode, field.TypeString, value)
	}
	if value, ok := ugu.mutation.Intro(); ok {
		_spec.SetField(usergroup.FieldIntro, field.TypeString, value)
	}
	if ugu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ugu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !ugu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ugu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ugu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ugu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !ugu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
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
	if nodes := ugu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, ugu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ugu.mutation.done = true
	return n, nil
}

// UserGroupUpdateOne is the builder for updating a single UserGroup entity.
type UserGroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserGroupMutation
}

// SetParentID sets the "parent_id" field.
func (uguo *UserGroupUpdateOne) SetParentID(s string) *UserGroupUpdateOne {
	uguo.mutation.SetParentID(s)
	return uguo
}

// SetName sets the "name" field.
func (uguo *UserGroupUpdateOne) SetName(s string) *UserGroupUpdateOne {
	uguo.mutation.SetName(s)
	return uguo
}

// SetCode sets the "code" field.
func (uguo *UserGroupUpdateOne) SetCode(s string) *UserGroupUpdateOne {
	uguo.mutation.SetCode(s)
	return uguo
}

// SetIntro sets the "intro" field.
func (uguo *UserGroupUpdateOne) SetIntro(s string) *UserGroupUpdateOne {
	uguo.mutation.SetIntro(s)
	return uguo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (uguo *UserGroupUpdateOne) AddUserIDs(ids ...int) *UserGroupUpdateOne {
	uguo.mutation.AddUserIDs(ids...)
	return uguo
}

// AddUsers adds the "users" edges to the User entity.
func (uguo *UserGroupUpdateOne) AddUsers(u ...*User) *UserGroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uguo.AddUserIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uguo *UserGroupUpdateOne) AddRoleIDs(ids ...int) *UserGroupUpdateOne {
	uguo.mutation.AddRoleIDs(ids...)
	return uguo
}

// AddRoles adds the "roles" edges to the Role entity.
func (uguo *UserGroupUpdateOne) AddRoles(r ...*Role) *UserGroupUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uguo.AddRoleIDs(ids...)
}

// Mutation returns the UserGroupMutation object of the builder.
func (uguo *UserGroupUpdateOne) Mutation() *UserGroupMutation {
	return uguo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (uguo *UserGroupUpdateOne) ClearUsers() *UserGroupUpdateOne {
	uguo.mutation.ClearUsers()
	return uguo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (uguo *UserGroupUpdateOne) RemoveUserIDs(ids ...int) *UserGroupUpdateOne {
	uguo.mutation.RemoveUserIDs(ids...)
	return uguo
}

// RemoveUsers removes "users" edges to User entities.
func (uguo *UserGroupUpdateOne) RemoveUsers(u ...*User) *UserGroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uguo.RemoveUserIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uguo *UserGroupUpdateOne) ClearRoles() *UserGroupUpdateOne {
	uguo.mutation.ClearRoles()
	return uguo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uguo *UserGroupUpdateOne) RemoveRoleIDs(ids ...int) *UserGroupUpdateOne {
	uguo.mutation.RemoveRoleIDs(ids...)
	return uguo
}

// RemoveRoles removes "roles" edges to Role entities.
func (uguo *UserGroupUpdateOne) RemoveRoles(r ...*Role) *UserGroupUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uguo.RemoveRoleIDs(ids...)
}

// Where appends a list predicates to the UserGroupUpdate builder.
func (uguo *UserGroupUpdateOne) Where(ps ...predicate.UserGroup) *UserGroupUpdateOne {
	uguo.mutation.Where(ps...)
	return uguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uguo *UserGroupUpdateOne) Select(field string, fields ...string) *UserGroupUpdateOne {
	uguo.fields = append([]string{field}, fields...)
	return uguo
}

// Save executes the query and returns the updated UserGroup entity.
func (uguo *UserGroupUpdateOne) Save(ctx context.Context) (*UserGroup, error) {
	return withHooks(ctx, uguo.sqlSave, uguo.mutation, uguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) SaveX(ctx context.Context) *UserGroup {
	node, err := uguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uguo *UserGroupUpdateOne) Exec(ctx context.Context) error {
	_, err := uguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uguo *UserGroupUpdateOne) ExecX(ctx context.Context) {
	if err := uguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uguo *UserGroupUpdateOne) sqlSave(ctx context.Context) (_node *UserGroup, err error) {
	_spec := sqlgraph.NewUpdateSpec(usergroup.Table, usergroup.Columns, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	id, ok := uguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserGroup.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usergroup.FieldID)
		for _, f := range fields {
			if !usergroup.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usergroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uguo.mutation.ParentID(); ok {
		_spec.SetField(usergroup.FieldParentID, field.TypeString, value)
	}
	if value, ok := uguo.mutation.Name(); ok {
		_spec.SetField(usergroup.FieldName, field.TypeString, value)
	}
	if value, ok := uguo.mutation.Code(); ok {
		_spec.SetField(usergroup.FieldCode, field.TypeString, value)
	}
	if value, ok := uguo.mutation.Intro(); ok {
		_spec.SetField(usergroup.FieldIntro, field.TypeString, value)
	}
	if uguo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uguo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !uguo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uguo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.UsersTable,
			Columns: usergroup.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uguo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uguo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uguo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
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
	if nodes := uguo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   usergroup.RolesTable,
			Columns: usergroup.RolesPrimaryKey,
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
	_node = &UserGroup{config: uguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usergroup.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uguo.mutation.done = true
	return _node, nil
}
