// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceCreate is the builder for creating a Resource entity.
type ResourceCreate struct {
	config
	mutation *ResourceMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *ResourceCreate) SetCreateTime(t time.Time) *ResourceCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableCreateTime(t *time.Time) *ResourceCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *ResourceCreate) SetUpdateTime(t time.Time) *ResourceCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableUpdateTime(t *time.Time) *ResourceCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetMethod sets the "method" field.
func (rc *ResourceCreate) SetMethod(s string) *ResourceCreate {
	rc.mutation.SetMethod(s)
	return rc
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableMethod(s *string) *ResourceCreate {
	if s != nil {
		rc.SetMethod(*s)
	}
	return rc
}

// SetOperation sets the "operation" field.
func (rc *ResourceCreate) SetOperation(s string) *ResourceCreate {
	rc.mutation.SetOperation(s)
	return rc
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableOperation(s *string) *ResourceCreate {
	if s != nil {
		rc.SetOperation(*s)
	}
	return rc
}

// SetPath sets the "path" field.
func (rc *ResourceCreate) SetPath(s string) *ResourceCreate {
	rc.mutation.SetPath(s)
	return rc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (rc *ResourceCreate) SetNillablePath(s *string) *ResourceCreate {
	if s != nil {
		rc.SetPath(*s)
	}
	return rc
}

// SetMenuID sets the "menu_id" field.
func (rc *ResourceCreate) SetMenuID(i int64) *ResourceCreate {
	rc.mutation.SetMenuID(i)
	return rc
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableMenuID(i *int64) *ResourceCreate {
	if i != nil {
		rc.SetMenuID(*i)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ResourceCreate) SetID(i int64) *ResourceCreate {
	rc.mutation.SetID(i)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ResourceCreate) SetNillableID(i *int64) *ResourceCreate {
	if i != nil {
		rc.SetID(*i)
	}
	return rc
}

// SetMenu sets the "menu" edge to the Menu entity.
func (rc *ResourceCreate) SetMenu(m *Menu) *ResourceCreate {
	return rc.SetMenuID(m.ID)
}

// AddPermissionIDs adds the "permission" edge to the Permission entity by IDs.
func (rc *ResourceCreate) AddPermissionIDs(ids ...int64) *ResourceCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermission adds the "permission" edges to the Permission entity.
func (rc *ResourceCreate) AddPermission(p ...*Permission) *ResourceCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// AddPermissionResourceIDs adds the "permission_resources" edge to the PermissionResource entity by IDs.
func (rc *ResourceCreate) AddPermissionResourceIDs(ids ...int64) *ResourceCreate {
	rc.mutation.AddPermissionResourceIDs(ids...)
	return rc
}

// AddPermissionResources adds the "permission_resources" edges to the PermissionResource entity.
func (rc *ResourceCreate) AddPermissionResources(p ...*PermissionResource) *ResourceCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionResourceIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (rc *ResourceCreate) Mutation() *ResourceMutation {
	return rc.mutation
}

// Save creates the Resource in the database.
func (rc *ResourceCreate) Save(ctx context.Context) (*Resource, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResourceCreate) SaveX(ctx context.Context) *Resource {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResourceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResourceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResourceCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := resource.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := resource.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.Method(); !ok {
		v := resource.DefaultMethod
		rc.mutation.SetMethod(v)
	}
	if _, ok := rc.mutation.Operation(); !ok {
		v := resource.DefaultOperation
		rc.mutation.SetOperation(v)
	}
	if _, ok := rc.mutation.Path(); !ok {
		v := resource.DefaultPath
		rc.mutation.SetPath(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := resource.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResourceCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Resource.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Resource.update_time"`)}
	}
	if _, ok := rc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "Resource.method"`)}
	}
	if v, ok := rc.mutation.Method(); ok {
		if err := resource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Resource.method": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "Resource.operation"`)}
	}
	if v, ok := rc.mutation.Operation(); ok {
		if err := resource.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "Resource.operation": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Resource.path"`)}
	}
	if v, ok := rc.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if v, ok := rc.mutation.MenuID(); ok {
		if err := resource.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "Resource.menu_id": %w`, err)}
		}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := resource.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Resource.id": %w`, err)}
		}
	}
	return nil
}

func (rc *ResourceCreate) sqlSave(ctx context.Context) (*Resource, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ResourceCreate) createSpec() (*Resource, *sqlgraph.CreateSpec) {
	var (
		_node = &Resource{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(resource.Table, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(resource.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(resource.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.Method(); ok {
		_spec.SetField(resource.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := rc.mutation.Operation(); ok {
		_spec.SetField(resource.FieldOperation, field.TypeString, value)
		_node.Operation = value
	}
	if value, ok := rc.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := rc.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.MenuTable,
			Columns: []string{resource.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MenuID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   resource.PermissionTable,
			Columns: resource.PermissionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &PermissionResourceCreate{config: rc.config, mutation: newPermissionResourceMutation(rc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   resource.PermissionResourcesTable,
			Columns: []string{resource.PermissionResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permissionresource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetResource set the Resource
func (rc *ResourceCreate) SetResource(input *Resource, fields ...string) *ResourceCreate {
	m := rc.mutation
	if len(fields) == 0 {
		fields = resource.Columns
	}
	_ = m.SetFields(input, fields...)
	return rc
}

// SetResourceWithZero set the Resource
func (rc *ResourceCreate) SetResourceWithZero(input *Resource, fields ...string) *ResourceCreate {
	m := rc.mutation
	if len(fields) == 0 {
		fields = resource.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return rc
}

// ResourceCreateBulk is the builder for creating many Resource entities in bulk.
type ResourceCreateBulk struct {
	config
	err      error
	builders []*ResourceCreate
}

// Save creates the Resource entities in the database.
func (rcb *ResourceCreateBulk) Save(ctx context.Context) ([]*Resource, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Resource, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResourceCreateBulk) SaveX(ctx context.Context) []*Resource {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResourceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResourceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
