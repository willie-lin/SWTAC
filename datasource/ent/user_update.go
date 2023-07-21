// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/account"
	"SWTAC/datasource/ent/predicate"
	"SWTAC/datasource/ent/role"
	"SWTAC/datasource/ent/user"
	"SWTAC/datasource/ent/usergroup"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetCreator sets the "creator" field.
func (uu *UserUpdate) SetCreator(s string) *UserUpdate {
	uu.mutation.SetCreator(s)
	return uu
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreator(s *string) *UserUpdate {
	if s != nil {
		uu.SetCreator(*s)
	}
	return uu
}

// ClearCreator clears the value of the "creator" field.
func (uu *UserUpdate) ClearCreator() *UserUpdate {
	uu.mutation.ClearCreator()
	return uu
}

// SetEditor sets the "editor" field.
func (uu *UserUpdate) SetEditor(s string) *UserUpdate {
	uu.mutation.SetEditor(s)
	return uu
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEditor(s *string) *UserUpdate {
	if s != nil {
		uu.SetEditor(*s)
	}
	return uu
}

// ClearEditor clears the value of the "editor" field.
func (uu *UserUpdate) ClearEditor() *UserUpdate {
	uu.mutation.ClearEditor()
	return uu
}

// SetDeleted sets the "deleted" field.
func (uu *UserUpdate) SetDeleted(f float64) *UserUpdate {
	uu.mutation.ResetDeleted()
	uu.mutation.SetDeleted(f)
	return uu
}

// AddDeleted adds f to the "deleted" field.
func (uu *UserUpdate) AddDeleted(f float64) *UserUpdate {
	uu.mutation.AddDeleted(f)
	return uu
}

// SetNickname sets the "nickname" field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.mutation.SetNickname(s)
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetCity sets the "city" field.
func (uu *UserUpdate) SetCity(s string) *UserUpdate {
	uu.mutation.SetCity(s)
	return uu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCity(s *string) *UserUpdate {
	if s != nil {
		uu.SetCity(*s)
	}
	return uu
}

// ClearCity clears the value of the "city" field.
func (uu *UserUpdate) ClearCity() *UserUpdate {
	uu.mutation.ClearCity()
	return uu
}

// SetIntroduction sets the "introduction" field.
func (uu *UserUpdate) SetIntroduction(s string) *UserUpdate {
	uu.mutation.SetIntroduction(s)
	return uu
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIntroduction(s *string) *UserUpdate {
	if s != nil {
		uu.SetIntroduction(*s)
	}
	return uu
}

// ClearIntroduction clears the value of the "introduction" field.
func (uu *UserUpdate) ClearIntroduction() *UserUpdate {
	uu.mutation.ClearIntroduction()
	return uu
}

// SetState sets the "state" field.
func (uu *UserUpdate) SetState(i int) *UserUpdate {
	uu.mutation.ResetState()
	uu.mutation.SetState(i)
	return uu
}

// AddState adds i to the "state" field.
func (uu *UserUpdate) AddState(i int) *UserUpdate {
	uu.mutation.AddState(i)
	return uu
}

// AddGroupIDs adds the "groups" edge to the UserGroup entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the UserGroup entity.
func (uu *UserUpdate) AddGroups(u ...*UserGroup) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uu *UserUpdate) AddRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddRoleIDs(ids...)
	return uu
}

// AddRoles adds the "roles" edges to the Role entity.
func (uu *UserUpdate) AddRoles(r ...*Role) *UserUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRoleIDs(ids...)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (uu *UserUpdate) AddAccountIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddAccountIDs(ids...)
	return uu
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (uu *UserUpdate) AddAccounts(a ...*Account) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroups clears all "groups" edges to the UserGroup entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to UserGroup entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to UserGroup entities.
func (uu *UserUpdate) RemoveGroups(u ...*UserGroup) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uu *UserUpdate) ClearRoles() *UserUpdate {
	uu.mutation.ClearRoles()
	return uu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uu *UserUpdate) RemoveRoleIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveRoleIDs(ids...)
	return uu
}

// RemoveRoles removes "roles" edges to Role entities.
func (uu *UserUpdate) RemoveRoles(r ...*Role) *UserUpdate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRoleIDs(ids...)
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (uu *UserUpdate) ClearAccounts() *UserUpdate {
	uu.mutation.ClearAccounts()
	return uu
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (uu *UserUpdate) RemoveAccountIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveAccountIDs(ids...)
	return uu
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (uu *UserUpdate) RemoveAccounts(a ...*Account) *UserUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Creator(); ok {
		if err := user.CreatorValidator(v); err != nil {
			return &ValidationError{Name: "creator", err: fmt.Errorf(`ent: validator failed for field "User.creator": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Creator(); ok {
		_spec.SetField(user.FieldCreator, field.TypeString, value)
	}
	if uu.mutation.CreatorCleared() {
		_spec.ClearField(user.FieldCreator, field.TypeString)
	}
	if value, ok := uu.mutation.Editor(); ok {
		_spec.SetField(user.FieldEditor, field.TypeString, value)
	}
	if uu.mutation.EditorCleared() {
		_spec.ClearField(user.FieldEditor, field.TypeString)
	}
	if value, ok := uu.mutation.Deleted(); ok {
		_spec.SetField(user.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedDeleted(); ok {
		_spec.AddField(user.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.City(); ok {
		_spec.SetField(user.FieldCity, field.TypeString, value)
	}
	if uu.mutation.CityCleared() {
		_spec.ClearField(user.FieldCity, field.TypeString)
	}
	if value, ok := uu.mutation.Introduction(); ok {
		_spec.SetField(user.FieldIntroduction, field.TypeString, value)
	}
	if uu.mutation.IntroductionCleared() {
		_spec.ClearField(user.FieldIntroduction, field.TypeString)
	}
	if value, ok := uu.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedState(); ok {
		_spec.AddField(user.FieldState, field.TypeInt, value)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !uu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetCreator sets the "creator" field.
func (uuo *UserUpdateOne) SetCreator(s string) *UserUpdateOne {
	uuo.mutation.SetCreator(s)
	return uuo
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreator(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCreator(*s)
	}
	return uuo
}

// ClearCreator clears the value of the "creator" field.
func (uuo *UserUpdateOne) ClearCreator() *UserUpdateOne {
	uuo.mutation.ClearCreator()
	return uuo
}

// SetEditor sets the "editor" field.
func (uuo *UserUpdateOne) SetEditor(s string) *UserUpdateOne {
	uuo.mutation.SetEditor(s)
	return uuo
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEditor(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEditor(*s)
	}
	return uuo
}

// ClearEditor clears the value of the "editor" field.
func (uuo *UserUpdateOne) ClearEditor() *UserUpdateOne {
	uuo.mutation.ClearEditor()
	return uuo
}

// SetDeleted sets the "deleted" field.
func (uuo *UserUpdateOne) SetDeleted(f float64) *UserUpdateOne {
	uuo.mutation.ResetDeleted()
	uuo.mutation.SetDeleted(f)
	return uuo
}

// AddDeleted adds f to the "deleted" field.
func (uuo *UserUpdateOne) AddDeleted(f float64) *UserUpdateOne {
	uuo.mutation.AddDeleted(f)
	return uuo
}

// SetNickname sets the "nickname" field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.mutation.SetNickname(s)
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetCity sets the "city" field.
func (uuo *UserUpdateOne) SetCity(s string) *UserUpdateOne {
	uuo.mutation.SetCity(s)
	return uuo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCity(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCity(*s)
	}
	return uuo
}

// ClearCity clears the value of the "city" field.
func (uuo *UserUpdateOne) ClearCity() *UserUpdateOne {
	uuo.mutation.ClearCity()
	return uuo
}

// SetIntroduction sets the "introduction" field.
func (uuo *UserUpdateOne) SetIntroduction(s string) *UserUpdateOne {
	uuo.mutation.SetIntroduction(s)
	return uuo
}

// SetNillableIntroduction sets the "introduction" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIntroduction(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIntroduction(*s)
	}
	return uuo
}

// ClearIntroduction clears the value of the "introduction" field.
func (uuo *UserUpdateOne) ClearIntroduction() *UserUpdateOne {
	uuo.mutation.ClearIntroduction()
	return uuo
}

// SetState sets the "state" field.
func (uuo *UserUpdateOne) SetState(i int) *UserUpdateOne {
	uuo.mutation.ResetState()
	uuo.mutation.SetState(i)
	return uuo
}

// AddState adds i to the "state" field.
func (uuo *UserUpdateOne) AddState(i int) *UserUpdateOne {
	uuo.mutation.AddState(i)
	return uuo
}

// AddGroupIDs adds the "groups" edge to the UserGroup entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the UserGroup entity.
func (uuo *UserUpdateOne) AddGroups(u ...*UserGroup) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (uuo *UserUpdateOne) AddRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddRoleIDs(ids...)
	return uuo
}

// AddRoles adds the "roles" edges to the Role entity.
func (uuo *UserUpdateOne) AddRoles(r ...*Role) *UserUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRoleIDs(ids...)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (uuo *UserUpdateOne) AddAccountIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddAccountIDs(ids...)
	return uuo
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (uuo *UserUpdateOne) AddAccounts(a ...*Account) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAccountIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroups clears all "groups" edges to the UserGroup entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to UserGroup entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to UserGroup entities.
func (uuo *UserUpdateOne) RemoveGroups(u ...*UserGroup) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role entity.
func (uuo *UserUpdateOne) ClearRoles() *UserUpdateOne {
	uuo.mutation.ClearRoles()
	return uuo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (uuo *UserUpdateOne) RemoveRoleIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveRoleIDs(ids...)
	return uuo
}

// RemoveRoles removes "roles" edges to Role entities.
func (uuo *UserUpdateOne) RemoveRoles(r ...*Role) *UserUpdateOne {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRoleIDs(ids...)
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (uuo *UserUpdateOne) ClearAccounts() *UserUpdateOne {
	uuo.mutation.ClearAccounts()
	return uuo
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (uuo *UserUpdateOne) RemoveAccountIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveAccountIDs(ids...)
	return uuo
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (uuo *UserUpdateOne) RemoveAccounts(a ...*Account) *UserUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAccountIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Creator(); ok {
		if err := user.CreatorValidator(v); err != nil {
			return &ValidationError{Name: "creator", err: fmt.Errorf(`ent: validator failed for field "User.creator": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Nickname(); ok {
		if err := user.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`ent: validator failed for field "User.nickname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Creator(); ok {
		_spec.SetField(user.FieldCreator, field.TypeString, value)
	}
	if uuo.mutation.CreatorCleared() {
		_spec.ClearField(user.FieldCreator, field.TypeString)
	}
	if value, ok := uuo.mutation.Editor(); ok {
		_spec.SetField(user.FieldEditor, field.TypeString, value)
	}
	if uuo.mutation.EditorCleared() {
		_spec.ClearField(user.FieldEditor, field.TypeString)
	}
	if value, ok := uuo.mutation.Deleted(); ok {
		_spec.SetField(user.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedDeleted(); ok {
		_spec.AddField(user.FieldDeleted, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.City(); ok {
		_spec.SetField(user.FieldCity, field.TypeString, value)
	}
	if uuo.mutation.CityCleared() {
		_spec.ClearField(user.FieldCity, field.TypeString)
	}
	if value, ok := uuo.mutation.Introduction(); ok {
		_spec.SetField(user.FieldIntroduction, field.TypeString, value)
	}
	if uuo.mutation.IntroductionCleared() {
		_spec.ClearField(user.FieldIntroduction, field.TypeString)
	}
	if value, ok := uuo.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedState(); ok {
		_spec.AddField(user.FieldState, field.TypeInt, value)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !uuo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.RolesTable,
			Columns: user.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !uuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AccountsTable,
			Columns: []string{user.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
