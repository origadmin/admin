// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPositionCreate is the builder for creating a UserPosition entity.
type UserPositionCreate struct {
	config
	mutation *UserPositionMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (upc *UserPositionCreate) SetUserID(i int64) *UserPositionCreate {
	upc.mutation.SetUserID(i)
	return upc
}

// SetPositionID sets the "position_id" field.
func (upc *UserPositionCreate) SetPositionID(i int64) *UserPositionCreate {
	upc.mutation.SetPositionID(i)
	return upc
}

// SetID sets the "id" field.
func (upc *UserPositionCreate) SetID(i int64) *UserPositionCreate {
	upc.mutation.SetID(i)
	return upc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (upc *UserPositionCreate) SetNillableID(i *int64) *UserPositionCreate {
	if i != nil {
		upc.SetID(*i)
	}
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserPositionCreate) SetUser(u *User) *UserPositionCreate {
	return upc.SetUserID(u.ID)
}

// SetPosition sets the "position" edge to the Position entity.
func (upc *UserPositionCreate) SetPosition(p *Position) *UserPositionCreate {
	return upc.SetPositionID(p.ID)
}

// Mutation returns the UserPositionMutation object of the builder.
func (upc *UserPositionCreate) Mutation() *UserPositionMutation {
	return upc.mutation
}

// Save creates the UserPosition in the database.
func (upc *UserPositionCreate) Save(ctx context.Context) (*UserPosition, error) {
	upc.defaults()
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserPositionCreate) SaveX(ctx context.Context) *UserPosition {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserPositionCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserPositionCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserPositionCreate) defaults() {
	if _, ok := upc.mutation.ID(); !ok {
		v := userposition.DefaultID()
		upc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserPositionCreate) check() error {
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserPosition.user_id"`)}
	}
	if v, ok := upc.mutation.UserID(); ok {
		if err := userposition.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "UserPosition.user_id": %w`, err)}
		}
	}
	if _, ok := upc.mutation.PositionID(); !ok {
		return &ValidationError{Name: "position_id", err: errors.New(`ent: missing required field "UserPosition.position_id"`)}
	}
	if v, ok := upc.mutation.PositionID(); ok {
		if err := userposition.PositionIDValidator(v); err != nil {
			return &ValidationError{Name: "position_id", err: fmt.Errorf(`ent: validator failed for field "UserPosition.position_id": %w`, err)}
		}
	}
	if v, ok := upc.mutation.ID(); ok {
		if err := userposition.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "UserPosition.id": %w`, err)}
		}
	}
	if len(upc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserPosition.user"`)}
	}
	if len(upc.mutation.PositionIDs()) == 0 {
		return &ValidationError{Name: "position", err: errors.New(`ent: missing required edge "UserPosition.position"`)}
	}
	return nil
}

func (upc *UserPositionCreate) sqlSave(ctx context.Context) (*UserPosition, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserPositionCreate) createSpec() (*UserPosition, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPosition{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userposition.Table, sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt64))
	)
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userposition.UserTable,
			Columns: []string{userposition.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := upc.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userposition.PositionTable,
			Columns: []string{userposition.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PositionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetUserPosition set the UserPosition
func (upc *UserPositionCreate) SetUserPosition(input *UserPosition, fields ...string) *UserPositionCreate {
	m := upc.mutation
	if len(fields) == 0 {
		fields = userposition.Columns
	}
	_ = m.SetFields(input, fields...)
	return upc
}

// SetUserPositionWithZero set the UserPosition
func (upc *UserPositionCreate) SetUserPositionWithZero(input *UserPosition, fields ...string) *UserPositionCreate {
	m := upc.mutation
	if len(fields) == 0 {
		fields = userposition.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return upc
}

// UserPositionCreateBulk is the builder for creating many UserPosition entities in bulk.
type UserPositionCreateBulk struct {
	config
	err      error
	builders []*UserPositionCreate
}

// Save creates the UserPosition entities in the database.
func (upcb *UserPositionCreateBulk) Save(ctx context.Context) ([]*UserPosition, error) {
	if upcb.err != nil {
		return nil, upcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserPosition, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPositionMutation)
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
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserPositionCreateBulk) SaveX(ctx context.Context) []*UserPosition {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserPositionCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserPositionCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
