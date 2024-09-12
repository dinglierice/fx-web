// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"fx-web/internal/ent/psconfig"
	"fx-web/internal/ent/psstrategy"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PsConfigCreate is the builder for creating a PsConfig entity.
type PsConfigCreate struct {
	config
	mutation *PsConfigMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pcc *PsConfigCreate) SetCreatedAt(t time.Time) *PsConfigCreate {
	pcc.mutation.SetCreatedAt(t)
	return pcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillableCreatedAt(t *time.Time) *PsConfigCreate {
	if t != nil {
		pcc.SetCreatedAt(*t)
	}
	return pcc
}

// SetUpdatedAt sets the "updated_at" field.
func (pcc *PsConfigCreate) SetUpdatedAt(t time.Time) *PsConfigCreate {
	pcc.mutation.SetUpdatedAt(t)
	return pcc
}

// SetPsScene sets the "ps_scene" field.
func (pcc *PsConfigCreate) SetPsScene(s string) *PsConfigCreate {
	pcc.mutation.SetPsScene(s)
	return pcc
}

// SetPsFilter sets the "ps_filter" field.
func (pcc *PsConfigCreate) SetPsFilter(u uint64) *PsConfigCreate {
	pcc.mutation.SetPsFilter(u)
	return pcc
}

// SetPsMessage sets the "ps_message" field.
func (pcc *PsConfigCreate) SetPsMessage(u uint64) *PsConfigCreate {
	pcc.mutation.SetPsMessage(u)
	return pcc
}

// SetNillablePsMessage sets the "ps_message" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillablePsMessage(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetPsMessage(*u)
	}
	return pcc
}

// SetPsEvent sets the "ps_event" field.
func (pcc *PsConfigCreate) SetPsEvent(u uint64) *PsConfigCreate {
	pcc.mutation.SetPsEvent(u)
	return pcc
}

// SetNillablePsEvent sets the "ps_event" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillablePsEvent(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetPsEvent(*u)
	}
	return pcc
}

// SetPsFeature sets the "ps_feature" field.
func (pcc *PsConfigCreate) SetPsFeature(u uint64) *PsConfigCreate {
	pcc.mutation.SetPsFeature(u)
	return pcc
}

// SetNillablePsFeature sets the "ps_feature" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillablePsFeature(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetPsFeature(*u)
	}
	return pcc
}

// SetPsStrategy sets the "ps_strategy" field.
func (pcc *PsConfigCreate) SetPsStrategy(u uint64) *PsConfigCreate {
	pcc.mutation.SetPsStrategy(u)
	return pcc
}

// SetNillablePsStrategy sets the "ps_strategy" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillablePsStrategy(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetPsStrategy(*u)
	}
	return pcc
}

// SetOwnerID sets the "owner_id" field.
func (pcc *PsConfigCreate) SetOwnerID(u uint64) *PsConfigCreate {
	pcc.mutation.SetOwnerID(u)
	return pcc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillableOwnerID(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetOwnerID(*u)
	}
	return pcc
}

// SetManagers sets the "managers" field.
func (pcc *PsConfigCreate) SetManagers(s string) *PsConfigCreate {
	pcc.mutation.SetManagers(s)
	return pcc
}

// SetNillableManagers sets the "managers" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillableManagers(s *string) *PsConfigCreate {
	if s != nil {
		pcc.SetManagers(*s)
	}
	return pcc
}

// SetUpdateUser sets the "update_user" field.
func (pcc *PsConfigCreate) SetUpdateUser(u uint64) *PsConfigCreate {
	pcc.mutation.SetUpdateUser(u)
	return pcc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (pcc *PsConfigCreate) SetNillableUpdateUser(u *uint64) *PsConfigCreate {
	if u != nil {
		pcc.SetUpdateUser(*u)
	}
	return pcc
}

// SetID sets the "id" field.
func (pcc *PsConfigCreate) SetID(u uint64) *PsConfigCreate {
	pcc.mutation.SetID(u)
	return pcc
}

// SetStrategyID sets the "strategy" edge to the PsStrategy entity by ID.
func (pcc *PsConfigCreate) SetStrategyID(id uint64) *PsConfigCreate {
	pcc.mutation.SetStrategyID(id)
	return pcc
}

// SetNillableStrategyID sets the "strategy" edge to the PsStrategy entity by ID if the given value is not nil.
func (pcc *PsConfigCreate) SetNillableStrategyID(id *uint64) *PsConfigCreate {
	if id != nil {
		pcc = pcc.SetStrategyID(*id)
	}
	return pcc
}

// SetStrategy sets the "strategy" edge to the PsStrategy entity.
func (pcc *PsConfigCreate) SetStrategy(p *PsStrategy) *PsConfigCreate {
	return pcc.SetStrategyID(p.ID)
}

// Mutation returns the PsConfigMutation object of the builder.
func (pcc *PsConfigCreate) Mutation() *PsConfigMutation {
	return pcc.mutation
}

// Save creates the PsConfig in the database.
func (pcc *PsConfigCreate) Save(ctx context.Context) (*PsConfig, error) {
	pcc.defaults()
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *PsConfigCreate) SaveX(ctx context.Context) *PsConfig {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *PsConfigCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *PsConfigCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *PsConfigCreate) defaults() {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		v := psconfig.DefaultCreatedAt()
		pcc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *PsConfigCreate) check() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PsConfig.created_at"`)}
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PsConfig.updated_at"`)}
	}
	if _, ok := pcc.mutation.PsScene(); !ok {
		return &ValidationError{Name: "ps_scene", err: errors.New(`ent: missing required field "PsConfig.ps_scene"`)}
	}
	if v, ok := pcc.mutation.PsScene(); ok {
		if err := psconfig.PsSceneValidator(v); err != nil {
			return &ValidationError{Name: "ps_scene", err: fmt.Errorf(`ent: validator failed for field "PsConfig.ps_scene": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.PsFilter(); !ok {
		return &ValidationError{Name: "ps_filter", err: errors.New(`ent: missing required field "PsConfig.ps_filter"`)}
	}
	if v, ok := pcc.mutation.ID(); ok {
		if err := psconfig.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "PsConfig.id": %w`, err)}
		}
	}
	return nil
}

func (pcc *PsConfigCreate) sqlSave(ctx context.Context) (*PsConfig, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *PsConfigCreate) createSpec() (*PsConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &PsConfig{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(psconfig.Table, sqlgraph.NewFieldSpec(psconfig.FieldID, field.TypeUint64))
	)
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pcc.mutation.CreatedAt(); ok {
		_spec.SetField(psconfig.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pcc.mutation.UpdatedAt(); ok {
		_spec.SetField(psconfig.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pcc.mutation.PsScene(); ok {
		_spec.SetField(psconfig.FieldPsScene, field.TypeString, value)
		_node.PsScene = value
	}
	if value, ok := pcc.mutation.PsFilter(); ok {
		_spec.SetField(psconfig.FieldPsFilter, field.TypeUint64, value)
		_node.PsFilter = value
	}
	if value, ok := pcc.mutation.PsMessage(); ok {
		_spec.SetField(psconfig.FieldPsMessage, field.TypeUint64, value)
		_node.PsMessage = &value
	}
	if value, ok := pcc.mutation.PsEvent(); ok {
		_spec.SetField(psconfig.FieldPsEvent, field.TypeUint64, value)
		_node.PsEvent = &value
	}
	if value, ok := pcc.mutation.PsFeature(); ok {
		_spec.SetField(psconfig.FieldPsFeature, field.TypeUint64, value)
		_node.PsFeature = &value
	}
	if value, ok := pcc.mutation.OwnerID(); ok {
		_spec.SetField(psconfig.FieldOwnerID, field.TypeUint64, value)
		_node.OwnerID = &value
	}
	if value, ok := pcc.mutation.Managers(); ok {
		_spec.SetField(psconfig.FieldManagers, field.TypeString, value)
		_node.Managers = &value
	}
	if value, ok := pcc.mutation.UpdateUser(); ok {
		_spec.SetField(psconfig.FieldUpdateUser, field.TypeUint64, value)
		_node.UpdateUser = &value
	}
	if nodes := pcc.mutation.StrategyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   psconfig.StrategyTable,
			Columns: []string{psconfig.StrategyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(psstrategy.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PsStrategy = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PsConfigCreateBulk is the builder for creating many PsConfig entities in bulk.
type PsConfigCreateBulk struct {
	config
	err      error
	builders []*PsConfigCreate
}

// Save creates the PsConfig entities in the database.
func (pccb *PsConfigCreateBulk) Save(ctx context.Context) ([]*PsConfig, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*PsConfig, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PsConfigMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *PsConfigCreateBulk) SaveX(ctx context.Context) []*PsConfig {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *PsConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *PsConfigCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
