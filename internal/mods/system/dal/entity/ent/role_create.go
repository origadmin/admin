// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/departmentrole"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolepermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleCreate is the builder for creating a Role entity.
type RoleCreate struct {
	config
	mutation *RoleMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (rc *RoleCreate) SetCreateTime(t time.Time) *RoleCreate {
	rc.mutation.SetCreateTime(t)
	return rc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rc *RoleCreate) SetNillableCreateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetCreateTime(*t)
	}
	return rc
}

// SetUpdateTime sets the "update_time" field.
func (rc *RoleCreate) SetUpdateTime(t time.Time) *RoleCreate {
	rc.mutation.SetUpdateTime(t)
	return rc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (rc *RoleCreate) SetNillableUpdateTime(t *time.Time) *RoleCreate {
	if t != nil {
		rc.SetUpdateTime(*t)
	}
	return rc
}

// SetKeyword sets the "keyword" field.
func (rc *RoleCreate) SetKeyword(s string) *RoleCreate {
	rc.mutation.SetKeyword(s)
	return rc
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (rc *RoleCreate) SetNillableKeyword(s *string) *RoleCreate {
	if s != nil {
		rc.SetKeyword(*s)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *RoleCreate) SetName(s string) *RoleCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (rc *RoleCreate) SetNillableName(s *string) *RoleCreate {
	if s != nil {
		rc.SetName(*s)
	}
	return rc
}

// SetDescription sets the "description" field.
func (rc *RoleCreate) SetDescription(s string) *RoleCreate {
	rc.mutation.SetDescription(s)
	return rc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (rc *RoleCreate) SetNillableDescription(s *string) *RoleCreate {
	if s != nil {
		rc.SetDescription(*s)
	}
	return rc
}

// SetType sets the "type" field.
func (rc *RoleCreate) SetType(i int8) *RoleCreate {
	rc.mutation.SetType(i)
	return rc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (rc *RoleCreate) SetNillableType(i *int8) *RoleCreate {
	if i != nil {
		rc.SetType(*i)
	}
	return rc
}

// SetSequence sets the "sequence" field.
func (rc *RoleCreate) SetSequence(i int) *RoleCreate {
	rc.mutation.SetSequence(i)
	return rc
}

// SetStatus sets the "status" field.
func (rc *RoleCreate) SetStatus(i int8) *RoleCreate {
	rc.mutation.SetStatus(i)
	return rc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (rc *RoleCreate) SetNillableStatus(i *int8) *RoleCreate {
	if i != nil {
		rc.SetStatus(*i)
	}
	return rc
}

// SetIsSystem sets the "is_system" field.
func (rc *RoleCreate) SetIsSystem(b bool) *RoleCreate {
	rc.mutation.SetIsSystem(b)
	return rc
}

// SetNillableIsSystem sets the "is_system" field if the given value is not nil.
func (rc *RoleCreate) SetNillableIsSystem(b *bool) *RoleCreate {
	if b != nil {
		rc.SetIsSystem(*b)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoleCreate) SetID(s string) *RoleCreate {
	rc.mutation.SetID(s)
	return rc
}

// AddMenuIDs adds the "menus" edge to the Menu entity by IDs.
func (rc *RoleCreate) AddMenuIDs(ids ...string) *RoleCreate {
	rc.mutation.AddMenuIDs(ids...)
	return rc
}

// AddMenus adds the "menus" edges to the Menu entity.
func (rc *RoleCreate) AddMenus(m ...*Menu) *RoleCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMenuIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (rc *RoleCreate) AddUserIDs(ids ...string) *RoleCreate {
	rc.mutation.AddUserIDs(ids...)
	return rc
}

// AddUsers adds the "users" edges to the User entity.
func (rc *RoleCreate) AddUsers(u ...*User) *RoleCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (rc *RoleCreate) AddPermissionIDs(ids ...string) *RoleCreate {
	rc.mutation.AddPermissionIDs(ids...)
	return rc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (rc *RoleCreate) AddPermissions(p ...*Permission) *RoleCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddPermissionIDs(ids...)
}

// AddDepartmentIDs adds the "departments" edge to the Department entity by IDs.
func (rc *RoleCreate) AddDepartmentIDs(ids ...string) *RoleCreate {
	rc.mutation.AddDepartmentIDs(ids...)
	return rc
}

// AddDepartments adds the "departments" edges to the Department entity.
func (rc *RoleCreate) AddDepartments(d ...*Department) *RoleCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return rc.AddDepartmentIDs(ids...)
}

// AddRoleMenuIDs adds the "role_menus" edge to the RoleMenu entity by IDs.
func (rc *RoleCreate) AddRoleMenuIDs(ids ...int) *RoleCreate {
	rc.mutation.AddRoleMenuIDs(ids...)
	return rc
}

// AddRoleMenus adds the "role_menus" edges to the RoleMenu entity.
func (rc *RoleCreate) AddRoleMenus(r ...*RoleMenu) *RoleCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRoleMenuIDs(ids...)
}

// AddUserRoleIDs adds the "user_roles" edge to the UserRole entity by IDs.
func (rc *RoleCreate) AddUserRoleIDs(ids ...int) *RoleCreate {
	rc.mutation.AddUserRoleIDs(ids...)
	return rc
}

// AddUserRoles adds the "user_roles" edges to the UserRole entity.
func (rc *RoleCreate) AddUserRoles(u ...*UserRole) *RoleCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return rc.AddUserRoleIDs(ids...)
}

// AddRolePermissionIDs adds the "role_permissions" edge to the RolePermission entity by IDs.
func (rc *RoleCreate) AddRolePermissionIDs(ids ...int) *RoleCreate {
	rc.mutation.AddRolePermissionIDs(ids...)
	return rc
}

// AddRolePermissions adds the "role_permissions" edges to the RolePermission entity.
func (rc *RoleCreate) AddRolePermissions(r ...*RolePermission) *RoleCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRolePermissionIDs(ids...)
}

// AddDepartmentRoleIDs adds the "department_roles" edge to the DepartmentRole entity by IDs.
func (rc *RoleCreate) AddDepartmentRoleIDs(ids ...int) *RoleCreate {
	rc.mutation.AddDepartmentRoleIDs(ids...)
	return rc
}

// AddDepartmentRoles adds the "department_roles" edges to the DepartmentRole entity.
func (rc *RoleCreate) AddDepartmentRoles(d ...*DepartmentRole) *RoleCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return rc.AddDepartmentRoleIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (rc *RoleCreate) Mutation() *RoleMutation {
	return rc.mutation
}

// Save creates the Role in the database.
func (rc *RoleCreate) Save(ctx context.Context) (*Role, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoleCreate) SaveX(ctx context.Context) *Role {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoleCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoleCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoleCreate) defaults() {
	if _, ok := rc.mutation.CreateTime(); !ok {
		v := role.DefaultCreateTime()
		rc.mutation.SetCreateTime(v)
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		v := role.DefaultUpdateTime()
		rc.mutation.SetUpdateTime(v)
	}
	if _, ok := rc.mutation.Keyword(); !ok {
		v := role.DefaultKeyword
		rc.mutation.SetKeyword(v)
	}
	if _, ok := rc.mutation.Name(); !ok {
		v := role.DefaultName
		rc.mutation.SetName(v)
	}
	if _, ok := rc.mutation.Description(); !ok {
		v := role.DefaultDescription
		rc.mutation.SetDescription(v)
	}
	if _, ok := rc.mutation.GetType(); !ok {
		v := role.DefaultType
		rc.mutation.SetType(v)
	}
	if _, ok := rc.mutation.Status(); !ok {
		v := role.DefaultStatus
		rc.mutation.SetStatus(v)
	}
	if _, ok := rc.mutation.IsSystem(); !ok {
		v := role.DefaultIsSystem
		rc.mutation.SetIsSystem(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoleCreate) check() error {
	if _, ok := rc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Role.create_time"`)}
	}
	if _, ok := rc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Role.update_time"`)}
	}
	if _, ok := rc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Role.keyword"`)}
	}
	if v, ok := rc.mutation.Keyword(); ok {
		if err := role.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Role.keyword": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Role.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Role.description"`)}
	}
	if v, ok := rc.mutation.Description(); ok {
		if err := role.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Role.description": %w`, err)}
		}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Role.type"`)}
	}
	if _, ok := rc.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`ent: missing required field "Role.sequence"`)}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Role.status"`)}
	}
	if _, ok := rc.mutation.IsSystem(); !ok {
		return &ValidationError{Name: "is_system", err: errors.New(`ent: missing required field "Role.is_system"`)}
	}
	if v, ok := rc.mutation.ID(); ok {
		if err := role.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Role.id": %w`, err)}
		}
	}
	return nil
}

func (rc *RoleCreate) sqlSave(ctx context.Context) (*Role, error) {
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
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Role.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RoleCreate) createSpec() (*Role, *sqlgraph.CreateSpec) {
	var (
		_node = &Role{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(role.Table, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreateTime(); ok {
		_spec.SetField(role.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := rc.mutation.UpdateTime(); ok {
		_spec.SetField(role.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := rc.mutation.Keyword(); ok {
		_spec.SetField(role.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Description(); ok {
		_spec.SetField(role.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.SetField(role.FieldType, field.TypeInt8, value)
		_node.Type = value
	}
	if value, ok := rc.mutation.Sequence(); ok {
		_spec.SetField(role.FieldSequence, field.TypeInt, value)
		_node.Sequence = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.IsSystem(); ok {
		_spec.SetField(role.FieldIsSystem, field.TypeBool, value)
		_node.IsSystem = value
	}
	if nodes := rc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: role.MenusPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.UsersTable,
			Columns: role.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: role.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   role.DepartmentsTable,
			Columns: role.DepartmentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RoleMenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.RoleMenusTable,
			Columns: []string{role.RoleMenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.UserRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.UserRolesTable,
			Columns: []string{role.UserRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.RolePermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.RolePermissionsTable,
			Columns: []string{role.RolePermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rolepermission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.DepartmentRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   role.DepartmentRolesTable,
			Columns: []string{role.DepartmentRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetRole set the Role
func (rc *RoleCreate) SetRole(input *Role, fields ...string) *RoleCreate {
	m := rc.mutation
	if len(fields) == 0 {
		fields = role.Columns
	}
	_ = m.SetFields(input, fields...)
	return rc
}

// SetRoleWithZero set the Role
func (rc *RoleCreate) SetRoleWithZero(input *Role, fields ...string) *RoleCreate {
	m := rc.mutation
	if len(fields) == 0 {
		fields = role.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return rc
}

// RoleCreateBulk is the builder for creating many Role entities in bulk.
type RoleCreateBulk struct {
	config
	err      error
	builders []*RoleCreate
}

// Save creates the Role entities in the database.
func (rcb *RoleCreateBulk) Save(ctx context.Context) ([]*Role, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Role, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleMutation)
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
func (rcb *RoleCreateBulk) SaveX(ctx context.Context) []*Role {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoleCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoleCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
