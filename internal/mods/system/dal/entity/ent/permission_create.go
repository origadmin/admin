// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolepermission"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PermissionCreate) SetCreateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PermissionCreate) SetUpdateTime(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdateTime(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PermissionCreate) SetName(s string) *PermissionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableName(s *string) *PermissionCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetKeyword sets the "keyword" field.
func (pc *PermissionCreate) SetKeyword(s string) *PermissionCreate {
	pc.mutation.SetKeyword(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PermissionCreate) SetDescription(s string) *PermissionCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDescription(s *string) *PermissionCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetDataScope sets the "data_scope" field.
func (pc *PermissionCreate) SetDataScope(s string) *PermissionCreate {
	pc.mutation.SetDataScope(s)
	return pc
}

// SetNillableDataScope sets the "data_scope" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableDataScope(s *string) *PermissionCreate {
	if s != nil {
		pc.SetDataScope(*s)
	}
	return pc
}

// SetDataRules sets the "data_rules" field.
func (pc *PermissionCreate) SetDataRules(m []map[string]string) *PermissionCreate {
	pc.mutation.SetDataRules(m)
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(i int64) *PermissionCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableID(i *int64) *PermissionCreate {
	if i != nil {
		pc.SetID(*i)
	}
	return pc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pc *PermissionCreate) AddRoleIDs(ids ...int64) *PermissionCreate {
	pc.mutation.AddRoleIDs(ids...)
	return pc
}

// AddRoles adds the "roles" edges to the Role entity.
func (pc *PermissionCreate) AddRoles(r ...*Role) *PermissionCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRoleIDs(ids...)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (pc *PermissionCreate) AddResourceIDs(ids ...int64) *PermissionCreate {
	pc.mutation.AddResourceIDs(ids...)
	return pc
}

// AddResources adds the "resources" edges to the Resource entity.
func (pc *PermissionCreate) AddResources(r ...*Resource) *PermissionCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddResourceIDs(ids...)
}

// AddRolePermissionIDs adds the "role_permissions" edge to the RolePermission entity by IDs.
func (pc *PermissionCreate) AddRolePermissionIDs(ids ...int64) *PermissionCreate {
	pc.mutation.AddRolePermissionIDs(ids...)
	return pc
}

// AddRolePermissions adds the "role_permissions" edges to the RolePermission entity.
func (pc *PermissionCreate) AddRolePermissions(r ...*RolePermission) *PermissionCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRolePermissionIDs(ids...)
}

// AddPermissionResourceIDs adds the "permission_resources" edge to the PermissionResource entity by IDs.
func (pc *PermissionCreate) AddPermissionResourceIDs(ids ...int64) *PermissionCreate {
	pc.mutation.AddPermissionResourceIDs(ids...)
	return pc
}

// AddPermissionResources adds the "permission_resources" edges to the PermissionResource entity.
func (pc *PermissionCreate) AddPermissionResources(p ...*PermissionResource) *PermissionCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPermissionResourceIDs(ids...)
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
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := permission.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := permission.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := permission.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Description(); !ok {
		v := permission.DefaultDescription
		pc.mutation.SetDescription(v)
	}
	if _, ok := pc.mutation.DataScope(); !ok {
		v := permission.DefaultDataScope
		pc.mutation.SetDataScope(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := permission.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Permission.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Permission.update_time"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Permission.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := permission.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Permission.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Permission.keyword"`)}
	}
	if v, ok := pc.mutation.Keyword(); ok {
		if err := permission.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Permission.keyword": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Permission.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := permission.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Permission.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.DataScope(); !ok {
		return &ValidationError{Name: "data_scope", err: errors.New(`ent: missing required field "Permission.data_scope"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := permission.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Permission.id": %w`, err)}
		}
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(permission.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(permission.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Keyword(); ok {
		_spec.SetField(permission.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.DataScope(); ok {
		_spec.SetField(permission.FieldDataScope, field.TypeString, value)
		_node.DataScope = value
	}
	if value, ok := pc.mutation.DataRules(); ok {
		_spec.SetField(permission.FieldDataRules, field.TypeJSON, value)
		_node.DataRules = value
	}
	if nodes := pc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RolePermissionCreate{config: pc.config, mutation: newRolePermissionMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   permission.ResourcesTable,
			Columns: permission.ResourcesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &PermissionResourceCreate{config: pc.config, mutation: newPermissionResourceMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RolePermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   permission.RolePermissionsTable,
			Columns: []string{permission.RolePermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rolepermission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PermissionResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   permission.PermissionResourcesTable,
			Columns: []string{permission.PermissionResourcesColumn},
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

// SetPermission set the Permission
func (pc *PermissionCreate) SetPermission(input *Permission, fields ...string) *PermissionCreate {
	m := pc.mutation
	if len(fields) == 0 {
		fields = permission.Columns
	}
	_ = m.SetFields(input, fields...)
	return pc
}

// SetPermissionWithZero set the Permission
func (pc *PermissionCreate) SetPermissionWithZero(input *Permission, fields ...string) *PermissionCreate {
	m := pc.mutation
	if len(fields) == 0 {
		fields = permission.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return pc
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	err      error
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
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
