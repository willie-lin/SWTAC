// Code generated by ent, DO NOT EDIT.

package ent

import (
	"SWTAC/datasource/ent/permission"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCreator sets the "creator" field.
func (pc *PermissionCreate) SetCreator(s string) *PermissionCreate {
	pc.mutation.SetCreator(s)
	return pc
}

// SetNillableCreator sets the "creator" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreator(s *string) *PermissionCreate {
	if s != nil {
		pc.SetCreator(*s)
	}
	return pc
}

// SetEditor sets the "editor" field.
func (pc *PermissionCreate) SetEditor(s string) *PermissionCreate {
	pc.mutation.SetEditor(s)
	return pc
}

// SetNillableEditor sets the "editor" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableEditor(s *string) *PermissionCreate {
	if s != nil {
		pc.SetEditor(*s)
	}
	return pc
}

// SetDeleted sets the "deleted" field.
func (pc *PermissionCreate) SetDeleted(f float64) *PermissionCreate {
	pc.mutation.SetDeleted(f)
	return pc
}

// SetParentID sets the "parent_id" field.
func (pc *PermissionCreate) SetParentID(i int) *PermissionCreate {
	pc.mutation.SetParentID(i)
	return pc
}

// SetCode sets the "code" field.
func (pc *PermissionCreate) SetCode(s string) *PermissionCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetIntro sets the "intro" field.
func (pc *PermissionCreate) SetIntro(s string) *PermissionCreate {
	pc.mutation.SetIntro(s)
	return pc
}

// SetCategory sets the "category" field.
func (pc *PermissionCreate) SetCategory(i int) *PermissionCreate {
	pc.mutation.SetCategory(i)
	return pc
}

// SetURL sets the "url" field.
func (pc *PermissionCreate) SetURL(i int) *PermissionCreate {
	pc.mutation.SetURL(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(u uuid.UUID) *PermissionCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableID(u *uuid.UUID) *PermissionCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := permission.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := permission.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Permission.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Permission.updated_at"`)}
	}
	if v, ok := pc.mutation.Creator(); ok {
		if err := permission.CreatorValidator(v); err != nil {
			return &ValidationError{Name: "creator", err: fmt.Errorf(`ent: validator failed for field "Permission.creator": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Permission.deleted"`)}
	}
	if _, ok := pc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "Permission.parent_id"`)}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Permission.code"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Permission.name"`)}
	}
	if _, ok := pc.mutation.Intro(); !ok {
		return &ValidationError{Name: "intro", err: errors.New(`ent: missing required field "Permission.intro"`)}
	}
	if _, ok := pc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Permission.category"`)}
	}
	if _, ok := pc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Permission.url"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(permission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Creator(); ok {
		_spec.SetField(permission.FieldCreator, field.TypeString, value)
		_node.Creator = value
	}
	if value, ok := pc.mutation.Editor(); ok {
		_spec.SetField(permission.FieldEditor, field.TypeString, value)
		_node.Editor = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.SetField(permission.FieldDeleted, field.TypeFloat64, value)
		_node.Deleted = value
	}
	if value, ok := pc.mutation.ParentID(); ok {
		_spec.SetField(permission.FieldParentID, field.TypeInt, value)
		_node.ParentID = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(permission.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Intro(); ok {
		_spec.SetField(permission.FieldIntro, field.TypeString, value)
		_node.Intro = value
	}
	if value, ok := pc.mutation.Category(); ok {
		_spec.SetField(permission.FieldCategory, field.TypeInt, value)
		_node.Category = value
	}
	if value, ok := pc.mutation.URL(); ok {
		_spec.SetField(permission.FieldURL, field.TypeInt, value)
		_node.URL = value
	}
	return _node, _spec
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
