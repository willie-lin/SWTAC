// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/account"
	"SWTAC/datasource/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetOpenCode sets the "open_code" field.
func (ac *AccountCreate) SetOpenCode(s string) *AccountCreate {
	ac.mutation.SetOpenCode(s)
	return ac
}

// SetCategory sets the "category" field.
func (ac *AccountCreate) SetCategory(s string) *AccountCreate {
	ac.mutation.SetCategory(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(i int) *AccountCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AccountCreate) SetUserID(id int) *AccountCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ac *AccountCreate) SetNillableUserID(id *int) *AccountCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AccountCreate) SetUser(u *User) *AccountCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.OpenCode(); !ok {
		return &ValidationError{Name: "open_code", err: errors.New(`ent: missing required field "Account.open_code"`)}
	}
	if _, ok := ac.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Account.category"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.OpenCode(); ok {
		_spec.SetField(account.FieldOpenCode, field.TypeString, value)
		_node.OpenCode = value
	}
	if value, ok := ac.mutation.Category(); ok {
		_spec.SetField(account.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_accounts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
