// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fx-web/internal/ent/psstrategy"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PsStrategyCreate is the builder for creating a PsStrategy entity.
type PsStrategyCreate struct {
	config
	mutation *PsStrategyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (psc *PsStrategyCreate) SetCreatedAt(t time.Time) *PsStrategyCreate {
	psc.mutation.SetCreatedAt(t)
	return psc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (psc *PsStrategyCreate) SetNillableCreatedAt(t *time.Time) *PsStrategyCreate {
	if t != nil {
		psc.SetCreatedAt(*t)
	}
	return psc
}

// SetUpdatedAt sets the "updated_at" field.
func (psc *PsStrategyCreate) SetUpdatedAt(t time.Time) *PsStrategyCreate {
	psc.mutation.SetUpdatedAt(t)
	return psc
}

// SetOwner sets the "owner" field.
func (psc *PsStrategyCreate) SetOwner(u uint64) *PsStrategyCreate {
	psc.mutation.SetOwner(u)
	return psc
}

// SetScriptContent sets the "script_content" field.
func (psc *PsStrategyCreate) SetScriptContent(s string) *PsStrategyCreate {
	psc.mutation.SetScriptContent(s)
	return psc
}

// SetIsDelete sets the "is_delete" field.
func (psc *PsStrategyCreate) SetIsDelete(i int) *PsStrategyCreate {
	psc.mutation.SetIsDelete(i)
	return psc
}

// SetNillableIsDelete sets the "is_delete" field if the given value is not nil.
func (psc *PsStrategyCreate) SetNillableIsDelete(i *int) *PsStrategyCreate {
	if i != nil {
		psc.SetIsDelete(*i)
	}
	return psc
}

// SetID sets the "id" field.
func (psc *PsStrategyCreate) SetID(u uint64) *PsStrategyCreate {
	psc.mutation.SetID(u)
	return psc
}

// Mutation returns the PsStrategyMutation object of the builder.
func (psc *PsStrategyCreate) Mutation() *PsStrategyMutation {
	return psc.mutation
}

// Save creates the PsStrategy in the database.
func (psc *PsStrategyCreate) Save(ctx context.Context) (*PsStrategy, error) {
	psc.defaults()
	return withHooks(ctx, psc.sqlSave, psc.mutation, psc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psc *PsStrategyCreate) SaveX(ctx context.Context) *PsStrategy {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *PsStrategyCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *PsStrategyCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *PsStrategyCreate) defaults() {
	if _, ok := psc.mutation.CreatedAt(); !ok {
		v := psstrategy.DefaultCreatedAt()
		psc.mutation.SetCreatedAt(v)
	}
	if _, ok := psc.mutation.IsDelete(); !ok {
		v := psstrategy.DefaultIsDelete
		psc.mutation.SetIsDelete(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *PsStrategyCreate) check() error {
	if _, ok := psc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PsStrategy.created_at"`)}
	}
	if _, ok := psc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PsStrategy.updated_at"`)}
	}
	if _, ok := psc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "PsStrategy.owner"`)}
	}
	if _, ok := psc.mutation.ScriptContent(); !ok {
		return &ValidationError{Name: "script_content", err: errors.New(`ent: missing required field "PsStrategy.script_content"`)}
	}
	if _, ok := psc.mutation.IsDelete(); !ok {
		return &ValidationError{Name: "is_delete", err: errors.New(`ent: missing required field "PsStrategy.is_delete"`)}
	}
	if v, ok := psc.mutation.ID(); ok {
		if err := psstrategy.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PsStrategy.id": %w`, err)}
		}
	}
	return nil
}

func (psc *PsStrategyCreate) sqlSave(ctx context.Context) (*PsStrategy, error) {
	if err := psc.check(); err != nil {
		return nil, err
	}
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	psc.mutation.id = &_node.ID
	psc.mutation.done = true
	return _node, nil
}

func (psc *PsStrategyCreate) createSpec() (*PsStrategy, *sqlgraph.CreateSpec) {
	var (
		_node = &PsStrategy{config: psc.config}
		_spec = sqlgraph.NewCreateSpec(psstrategy.Table, sqlgraph.NewFieldSpec(psstrategy.FieldID, field.TypeUint64))
	)
	if id, ok := psc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := psc.mutation.CreatedAt(); ok {
		_spec.SetField(psstrategy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := psc.mutation.UpdatedAt(); ok {
		_spec.SetField(psstrategy.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := psc.mutation.Owner(); ok {
		_spec.SetField(psstrategy.FieldOwner, field.TypeUint64, value)
		_node.Owner = value
	}
	if value, ok := psc.mutation.ScriptContent(); ok {
		_spec.SetField(psstrategy.FieldScriptContent, field.TypeString, value)
		_node.ScriptContent = value
	}
	if value, ok := psc.mutation.IsDelete(); ok {
		_spec.SetField(psstrategy.FieldIsDelete, field.TypeInt, value)
		_node.IsDelete = value
	}
	return _node, _spec
}

// PsStrategyCreateBulk is the builder for creating many PsStrategy entities in bulk.
type PsStrategyCreateBulk struct {
	config
	err      error
	builders []*PsStrategyCreate
}

// Save creates the PsStrategy entities in the database.
func (pscb *PsStrategyCreateBulk) Save(ctx context.Context) ([]*PsStrategy, error) {
	if pscb.err != nil {
		return nil, pscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*PsStrategy, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PsStrategyMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
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
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *PsStrategyCreateBulk) SaveX(ctx context.Context) []*PsStrategy {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *PsStrategyCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *PsStrategyCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}