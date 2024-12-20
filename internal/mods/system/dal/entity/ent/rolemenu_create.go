// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMenuCreate is the builder for creating a RoleMenu entity.
type RoleMenuCreate struct {
	config
	mutation *RoleMenuMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rmc *RoleMenuCreate) SetCreateTime(t time.Time) *RoleMenuCreate {
	rmc.mutation.SetCreateTime(t)
	return rmc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rmc *RoleMenuCreate) SetNillableCreateTime(t *time.Time) *RoleMenuCreate {
	if t != nil {
		rmc.SetCreateTime(*t)
	}
	return rmc
}

// SetUpdateTime sets the "update_time" field.
func (rmc *RoleMenuCreate) SetUpdateTime(t time.Time) *RoleMenuCreate {
	rmc.mutation.SetUpdateTime(t)
	return rmc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rmc *RoleMenuCreate) SetNillableUpdateTime(t *time.Time) *RoleMenuCreate {
	if t != nil {
		rmc.SetUpdateTime(*t)
	}
	return rmc
}

// SetRoleID sets the "role_id" field.
func (rmc *RoleMenuCreate) SetRoleID(s string) *RoleMenuCreate {
	rmc.mutation.SetRoleID(s)
	return rmc
}

// SetMenuID sets the "menu_id" field.
func (rmc *RoleMenuCreate) SetMenuID(s string) *RoleMenuCreate {
	rmc.mutation.SetMenuID(s)
	return rmc
}

// SetID sets the "id" field.
func (rmc *RoleMenuCreate) SetID(s string) *RoleMenuCreate {
	rmc.mutation.SetID(s)
	return rmc
}

// SetRole sets the "role" edge to the Role entity.
func (rmc *RoleMenuCreate) SetRole(r *Role) *RoleMenuCreate {
	return rmc.SetRoleID(r.ID)
}

// SetMenu sets the "menu" edge to the Menu entity.
func (rmc *RoleMenuCreate) SetMenu(m *Menu) *RoleMenuCreate {
	return rmc.SetMenuID(m.ID)
}

// Mutation returns the RoleMenuMutation object of the builder.
func (rmc *RoleMenuCreate) Mutation() *RoleMenuMutation {
	return rmc.mutation
}

// Save creates the RoleMenu in the database.
func (rmc *RoleMenuCreate) Save(ctx context.Context) (*RoleMenu, error) {
	rmc.defaults()
	return withHooks(ctx, rmc.sqlSave, rmc.mutation, rmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rmc *RoleMenuCreate) SaveX(ctx context.Context) *RoleMenu {
	v, err := rmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmc *RoleMenuCreate) Exec(ctx context.Context) error {
	_, err := rmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmc *RoleMenuCreate) ExecX(ctx context.Context) {
	if err := rmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmc *RoleMenuCreate) defaults() {
	if _, ok := rmc.mutation.CreateTime(); !ok {
		v := rolemenu.DefaultCreateTime()
		rmc.mutation.SetCreateTime(v)
	}
	if _, ok := rmc.mutation.UpdateTime(); !ok {
		v := rolemenu.DefaultUpdateTime()
		rmc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmc *RoleMenuCreate) check() error {
	if _, ok := rmc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "RoleMenu.create_time"`)}
	}
	if _, ok := rmc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "RoleMenu.update_time"`)}
	}
	if _, ok := rmc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "RoleMenu.role_id"`)}
	}
	if v, ok := rmc.mutation.RoleID(); ok {
		if err := rolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.role_id": %w`, err)}
		}
	}
	if _, ok := rmc.mutation.MenuID(); !ok {
		return &ValidationError{Name: "menu_id", err: errors.New(`ent: missing required field "RoleMenu.menu_id"`)}
	}
	if v, ok := rmc.mutation.MenuID(); ok {
		if err := rolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.menu_id": %w`, err)}
		}
	}
	if v, ok := rmc.mutation.ID(); ok {
		if err := rolemenu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.id": %w`, err)}
		}
	}
	if len(rmc.mutation.RoleIDs()) == 0 {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "RoleMenu.role"`)}
	}
	if len(rmc.mutation.MenuIDs()) == 0 {
		return &ValidationError{Name: "menu", err: errors.New(`ent: missing required edge "RoleMenu.menu"`)}
	}
	return nil
}

func (rmc *RoleMenuCreate) sqlSave(ctx context.Context) (*RoleMenu, error) {
	if err := rmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RoleMenu.ID type: %T", _spec.ID.Value)
		}
	}
	rmc.mutation.id = &_node.ID
	rmc.mutation.done = true
	return _node, nil
}

func (rmc *RoleMenuCreate) createSpec() (*RoleMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleMenu{config: rmc.config}
		_spec = sqlgraph.NewCreateSpec(rolemenu.Table, sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeString))
	)
	if id, ok := rmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rmc.mutation.CreateTime(); ok {
		_spec.SetField(rolemenu.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rmc.mutation.UpdateTime(); ok {
		_spec.SetField(rolemenu.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if nodes := rmc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rolemenu.RoleTable,
			Columns: []string{rolemenu.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rmc.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   rolemenu.MenuTable,
			Columns: []string{rolemenu.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MenuID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetRoleMenu set the RoleMenu
func (rmc *RoleMenuCreate) SetRoleMenu(input *RoleMenu, fields ...string) *RoleMenuCreate {
	m := rmc.mutation
	_ = m.SetFields(input, fields...)
	return rmc
}

// RoleMenuCreateBulk is the builder for creating many RoleMenu entities in bulk.
type RoleMenuCreateBulk struct {
	config
	err      error
	builders []*RoleMenuCreate
}

// Save creates the RoleMenu entities in the database.
func (rmcb *RoleMenuCreateBulk) Save(ctx context.Context) ([]*RoleMenu, error) {
	if rmcb.err != nil {
		return nil, rmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rmcb.builders))
	nodes := make([]*RoleMenu, len(rmcb.builders))
	mutators := make([]Mutator, len(rmcb.builders))
	for i := range rmcb.builders {
		func(i int, root context.Context) {
			builder := rmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMenuMutation)
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
					_, err = mutators[i+1].Mutate(root, rmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rmcb *RoleMenuCreateBulk) SaveX(ctx context.Context) []*RoleMenu {
	v, err := rmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmcb *RoleMenuCreateBulk) Exec(ctx context.Context) error {
	_, err := rmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmcb *RoleMenuCreateBulk) ExecX(ctx context.Context) {
	if err := rmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
