// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRoleCreate is the builder for creating a UserRole entity.
type UserRoleCreate struct {
	config
	mutation *UserRoleMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (urc *UserRoleCreate) SetUserID(i int64) *UserRoleCreate {
	urc.mutation.SetUserID(i)
	return urc
}

// SetRoleID sets the "role_id" field.
func (urc *UserRoleCreate) SetRoleID(i int64) *UserRoleCreate {
	urc.mutation.SetRoleID(i)
	return urc
}

// SetUser sets the "user" edge to the User entity.
func (urc *UserRoleCreate) SetUser(u *User) *UserRoleCreate {
	return urc.SetUserID(u.ID)
}

// SetRole sets the "role" edge to the Role entity.
func (urc *UserRoleCreate) SetRole(r *Role) *UserRoleCreate {
	return urc.SetRoleID(r.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (urc *UserRoleCreate) Mutation() *UserRoleMutation {
	return urc.mutation
}

// Save creates the UserRole in the database.
func (urc *UserRoleCreate) Save(ctx context.Context) (*UserRole, error) {
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRoleCreate) SaveX(ctx context.Context) *UserRole {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRoleCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRoleCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRoleCreate) check() error {
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserRole.user_id"`)}
	}
	if v, ok := urc.mutation.UserID(); ok {
		if err := userrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.user_id": %w`, err)}
		}
	}
	if _, ok := urc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "UserRole.role_id"`)}
	}
	if v, ok := urc.mutation.RoleID(); ok {
		if err := userrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.role_id": %w`, err)}
		}
	}
	if len(urc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserRole.user"`)}
	}
	if len(urc.mutation.RoleIDs()) == 0 {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required edge "UserRole.role"`)}
	}
	return nil
}

func (urc *UserRoleCreate) sqlSave(ctx context.Context) (*UserRole, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UserRoleCreate) createSpec() (*UserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRole{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrole.Table, sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeInt))
	)
	if nodes := urc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
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
	if nodes := urc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetUserRole set the UserRole
func (urc *UserRoleCreate) SetUserRole(input *UserRole, fields ...string) *UserRoleCreate {
	m := urc.mutation
	if len(fields) == 0 {
		fields = userrole.Columns
	}
	_ = m.SetFields(input, fields...)
	return urc
}

// SetUserRoleWithZero set the UserRole
func (urc *UserRoleCreate) SetUserRoleWithZero(input *UserRole, fields ...string) *UserRoleCreate {
	m := urc.mutation
	if len(fields) == 0 {
		fields = userrole.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return urc
}

// UserRoleCreateBulk is the builder for creating many UserRole entities in bulk.
type UserRoleCreateBulk struct {
	config
	err      error
	builders []*UserRoleCreate
}

// Save creates the UserRole entities in the database.
func (urcb *UserRoleCreateBulk) Save(ctx context.Context) ([]*UserRole, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRole, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) SaveX(ctx context.Context) []*UserRole {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
