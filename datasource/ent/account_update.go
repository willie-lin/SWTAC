// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/account"
	"SWTAC/datasource/ent/predicate"
	"SWTAC/datasource/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetPhone sets the "phone" field.
func (au *AccountUpdate) SetPhone(s string) *AccountUpdate {
	au.mutation.SetPhone(s)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (au *AccountUpdate) AddUserIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.AddUserIDs(ids...)
	return au
}

// AddUsers adds the "users" edges to the User entity.
func (au *AccountUpdate) AddUsers(u ...*User) *AccountUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.AddUserIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (au *AccountUpdate) ClearUsers() *AccountUpdate {
	au.mutation.ClearUsers()
	return au
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (au *AccountUpdate) RemoveUserIDs(ids ...uuid.UUID) *AccountUpdate {
	au.mutation.RemoveUserIDs(ids...)
	return au
}

// RemoveUsers removes "users" edges to User entities.
func (au *AccountUpdate) RemoveUsers(u ...*User) *AccountUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return au.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := au.mutation.Phone(); ok {
		_spec.SetField(account.FieldPhone, field.TypeString, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedUsersIDs(); len(nodes) > 0 && !au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetPhone sets the "phone" field.
func (auo *AccountUpdateOne) SetPhone(s string) *AccountUpdateOne {
	auo.mutation.SetPhone(s)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (auo *AccountUpdateOne) AddUserIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.AddUserIDs(ids...)
	return auo
}

// AddUsers adds the "users" edges to the User entity.
func (auo *AccountUpdateOne) AddUsers(u ...*User) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.AddUserIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (auo *AccountUpdateOne) ClearUsers() *AccountUpdateOne {
	auo.mutation.ClearUsers()
	return auo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (auo *AccountUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *AccountUpdateOne {
	auo.mutation.RemoveUserIDs(ids...)
	return auo
}

// RemoveUsers removes "users" edges to User entities.
func (auo *AccountUpdateOne) RemoveUsers(u ...*User) *AccountUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return auo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := auo.mutation.Phone(); ok {
		_spec.SetField(account.FieldPhone, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   account.UsersTable,
			Columns: account.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
