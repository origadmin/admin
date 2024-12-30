// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PositionCreate) SetCreateTime(t time.Time) *PositionCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreateTime(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PositionCreate) SetUpdateTime(t time.Time) *PositionCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PositionCreate) SetNillableUpdateTime(t *time.Time) *PositionCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PositionCreate) SetName(s string) *PositionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PositionCreate) SetDescription(s string) *PositionCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetDepartmentID sets the "department_id" field.
func (pc *PositionCreate) SetDepartmentID(i int64) *PositionCreate {
	pc.mutation.SetDepartmentID(i)
	return pc
}

// SetID sets the "id" field.
func (pc *PositionCreate) SetID(i int64) *PositionCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PositionCreate) SetNillableID(i *int64) *PositionCreate {
	if i != nil {
		pc.SetID(*i)
	}
	return pc
}

// SetDepartment sets the "department" edge to the Department entity.
func (pc *PositionCreate) SetDepartment(d *Department) *PositionCreate {
	return pc.SetDepartmentID(d.ID)
}

// AddUserPositionIDs adds the "user_positions" edge to the UserPosition entity by IDs.
func (pc *PositionCreate) AddUserPositionIDs(ids ...int64) *PositionCreate {
	pc.mutation.AddUserPositionIDs(ids...)
	return pc
}

// AddUserPositions adds the "user_positions" edges to the UserPosition entity.
func (pc *PositionCreate) AddUserPositions(u ...*UserPosition) *PositionCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserPositionIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PositionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PositionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PositionCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := position.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := position.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := position.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PositionCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Position.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Position.update_time"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Position.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Position.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := position.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Position.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.DepartmentID(); !ok {
		return &ValidationError{Name: "department_id", err: errors.New(`ent: missing required field "Position.department_id"`)}
	}
	if v, ok := pc.mutation.DepartmentID(); ok {
		if err := position.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Position.department_id": %w`, err)}
		}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := position.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Position.id": %w`, err)}
		}
	}
	if len(pc.mutation.DepartmentIDs()) == 0 {
		return &ValidationError{Name: "department", err: errors.New(`ent: missing required edge "Position.department"`)}
	}
	return nil
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		_node = &Position{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(position.Table, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(position.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(position.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.DepartmentTable,
			Columns: []string{position.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DepartmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.UserPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetPosition set the Position
func (pc *PositionCreate) SetPosition(input *Position, fields ...string) *PositionCreate {
	m := pc.mutation
	if len(fields) == 0 {
		fields = position.Columns
	}
	_ = m.SetFields(input, fields...)
	return pc
}

// SetPositionWithZero set the Position
func (pc *PositionCreate) SetPositionWithZero(input *Position, fields ...string) *PositionCreate {
	m := pc.mutation
	if len(fields) == 0 {
		fields = position.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return pc
}

// PositionCreateBulk is the builder for creating many Position entities in bulk.
type PositionCreateBulk struct {
	config
	err      error
	builders []*PositionCreate
}

// Save creates the Position entities in the database.
func (pcb *PositionCreateBulk) Save(ctx context.Context) ([]*Position, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Position, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PositionMutation)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PositionCreateBulk) SaveX(ctx context.Context) []*Position {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PositionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PositionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
