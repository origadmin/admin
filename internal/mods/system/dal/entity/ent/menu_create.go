// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menupermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (mc *MenuCreate) SetCreateTime(t time.Time) *MenuCreate {
	mc.mutation.SetCreateTime(t)
	return mc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreateTime(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetCreateTime(*t)
	}
	return mc
}

// SetUpdateTime sets the "update_time" field.
func (mc *MenuCreate) SetUpdateTime(t time.Time) *MenuCreate {
	mc.mutation.SetUpdateTime(t)
	return mc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdateTime(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetUpdateTime(*t)
	}
	return mc
}

// SetKeyword sets the "keyword" field.
func (mc *MenuCreate) SetKeyword(s string) *MenuCreate {
	mc.mutation.SetKeyword(s)
	return mc
}

// SetNillableKeyword sets the "keyword" field if the given value is not nil.
func (mc *MenuCreate) SetNillableKeyword(s *string) *MenuCreate {
	if s != nil {
		mc.SetKeyword(*s)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mc *MenuCreate) SetNillableName(s *string) *MenuCreate {
	if s != nil {
		mc.SetName(*s)
	}
	return mc
}

// SetI18nKey sets the "i18n_key" field.
func (mc *MenuCreate) SetI18nKey(s string) *MenuCreate {
	mc.mutation.SetI18nKey(s)
	return mc
}

// SetNillableI18nKey sets the "i18n_key" field if the given value is not nil.
func (mc *MenuCreate) SetNillableI18nKey(s *string) *MenuCreate {
	if s != nil {
		mc.SetI18nKey(*s)
	}
	return mc
}

// SetDescription sets the "description" field.
func (mc *MenuCreate) SetDescription(s string) *MenuCreate {
	mc.mutation.SetDescription(s)
	return mc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDescription(s *string) *MenuCreate {
	if s != nil {
		mc.SetDescription(*s)
	}
	return mc
}

// SetType sets the "type" field.
func (mc *MenuCreate) SetType(s string) *MenuCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mc *MenuCreate) SetNillableType(s *string) *MenuCreate {
	if s != nil {
		mc.SetType(*s)
	}
	return mc
}

// SetIcon sets the "icon" field.
func (mc *MenuCreate) SetIcon(s string) *MenuCreate {
	mc.mutation.SetIcon(s)
	return mc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIcon(s *string) *MenuCreate {
	if s != nil {
		mc.SetIcon(*s)
	}
	return mc
}

// SetPath sets the "path" field.
func (mc *MenuCreate) SetPath(s string) *MenuCreate {
	mc.mutation.SetPath(s)
	return mc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mc *MenuCreate) SetNillablePath(s *string) *MenuCreate {
	if s != nil {
		mc.SetPath(*s)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MenuCreate) SetStatus(i int8) *MenuCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MenuCreate) SetNillableStatus(i *int8) *MenuCreate {
	if i != nil {
		mc.SetStatus(*i)
	}
	return mc
}

// SetParentPath sets the "parent_path" field.
func (mc *MenuCreate) SetParentPath(s string) *MenuCreate {
	mc.mutation.SetParentPath(s)
	return mc
}

// SetNillableParentPath sets the "parent_path" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentPath(s *string) *MenuCreate {
	if s != nil {
		mc.SetParentPath(*s)
	}
	return mc
}

// SetSequence sets the "sequence" field.
func (mc *MenuCreate) SetSequence(i int) *MenuCreate {
	mc.mutation.SetSequence(i)
	return mc
}

// SetNillableSequence sets the "sequence" field if the given value is not nil.
func (mc *MenuCreate) SetNillableSequence(i *int) *MenuCreate {
	if i != nil {
		mc.SetSequence(*i)
	}
	return mc
}

// SetProperties sets the "properties" field.
func (mc *MenuCreate) SetProperties(s string) *MenuCreate {
	mc.mutation.SetProperties(s)
	return mc
}

// SetNillableProperties sets the "properties" field if the given value is not nil.
func (mc *MenuCreate) SetNillableProperties(s *string) *MenuCreate {
	if s != nil {
		mc.SetProperties(*s)
	}
	return mc
}

// SetParentID sets the "parent_id" field.
func (mc *MenuCreate) SetParentID(i int64) *MenuCreate {
	mc.mutation.SetParentID(i)
	return mc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentID(i *int64) *MenuCreate {
	if i != nil {
		mc.SetParentID(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(i int64) *MenuCreate {
	mc.mutation.SetID(i)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableID(i *int64) *MenuCreate {
	if i != nil {
		mc.SetID(*i)
	}
	return mc
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mc *MenuCreate) AddChildIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddChildIDs(ids...)
	return mc
}

// AddChildren adds the "children" edges to the Menu entity.
func (mc *MenuCreate) AddChildren(m ...*Menu) *MenuCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddChildIDs(ids...)
}

// SetParent sets the "parent" edge to the Menu entity.
func (mc *MenuCreate) SetParent(m *Menu) *MenuCreate {
	return mc.SetParentID(m.ID)
}

// AddResourceIDs adds the "resources" edge to the Resource entity by IDs.
func (mc *MenuCreate) AddResourceIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddResourceIDs(ids...)
	return mc
}

// AddResources adds the "resources" edges to the Resource entity.
func (mc *MenuCreate) AddResources(r ...*Resource) *MenuCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddResourceIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (mc *MenuCreate) AddRoleIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddRoleIDs(ids...)
	return mc
}

// AddRoles adds the "roles" edges to the Role entity.
func (mc *MenuCreate) AddRoles(r ...*Role) *MenuCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoleIDs(ids...)
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (mc *MenuCreate) AddPermissionIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddPermissionIDs(ids...)
	return mc
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (mc *MenuCreate) AddPermissions(p ...*Permission) *MenuCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddPermissionIDs(ids...)
}

// AddRoleMenuIDs adds the "role_menus" edge to the RoleMenu entity by IDs.
func (mc *MenuCreate) AddRoleMenuIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddRoleMenuIDs(ids...)
	return mc
}

// AddRoleMenus adds the "role_menus" edges to the RoleMenu entity.
func (mc *MenuCreate) AddRoleMenus(r ...*RoleMenu) *MenuCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoleMenuIDs(ids...)
}

// AddMenuPermissionIDs adds the "menu_permissions" edge to the MenuPermission entity by IDs.
func (mc *MenuCreate) AddMenuPermissionIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddMenuPermissionIDs(ids...)
	return mc
}

// AddMenuPermissions adds the "menu_permissions" edges to the MenuPermission entity.
func (mc *MenuCreate) AddMenuPermissions(m ...*MenuPermission) *MenuCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMenuPermissionIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.CreateTime(); !ok {
		v := menu.DefaultCreateTime()
		mc.mutation.SetCreateTime(v)
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		v := menu.DefaultUpdateTime()
		mc.mutation.SetUpdateTime(v)
	}
	if _, ok := mc.mutation.Keyword(); !ok {
		v := menu.DefaultKeyword
		mc.mutation.SetKeyword(v)
	}
	if _, ok := mc.mutation.Name(); !ok {
		v := menu.DefaultName
		mc.mutation.SetName(v)
	}
	if _, ok := mc.mutation.I18nKey(); !ok {
		v := menu.DefaultI18nKey
		mc.mutation.SetI18nKey(v)
	}
	if _, ok := mc.mutation.Description(); !ok {
		v := menu.DefaultDescription
		mc.mutation.SetDescription(v)
	}
	if _, ok := mc.mutation.GetType(); !ok {
		v := menu.DefaultType
		mc.mutation.SetType(v)
	}
	if _, ok := mc.mutation.Icon(); !ok {
		v := menu.DefaultIcon
		mc.mutation.SetIcon(v)
	}
	if _, ok := mc.mutation.Path(); !ok {
		v := menu.DefaultPath
		mc.mutation.SetPath(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := menu.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.ParentPath(); !ok {
		v := menu.DefaultParentPath
		mc.mutation.SetParentPath(v)
	}
	if _, ok := mc.mutation.Sequence(); !ok {
		v := menu.DefaultSequence
		mc.mutation.SetSequence(v)
	}
	if _, ok := mc.mutation.Properties(); !ok {
		v := menu.DefaultProperties
		mc.mutation.SetProperties(v)
	}
	if _, ok := mc.mutation.ParentID(); !ok {
		v := menu.DefaultParentID()
		mc.mutation.SetParentID(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := menu.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Menu.create_time"`)}
	}
	if _, ok := mc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Menu.update_time"`)}
	}
	if _, ok := mc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Menu.keyword"`)}
	}
	if v, ok := mc.mutation.Keyword(); ok {
		if err := menu.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Menu.keyword": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Menu.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Menu.name": %w`, err)}
		}
	}
	if _, ok := mc.mutation.I18nKey(); !ok {
		return &ValidationError{Name: "i18n_key", err: errors.New(`ent: missing required field "Menu.i18n_key"`)}
	}
	if v, ok := mc.mutation.I18nKey(); ok {
		if err := menu.I18nKeyValidator(v); err != nil {
			return &ValidationError{Name: "i18n_key", err: fmt.Errorf(`ent: validator failed for field "Menu.i18n_key": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Menu.description"`)}
	}
	if v, ok := mc.mutation.Description(); ok {
		if err := menu.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Menu.description": %w`, err)}
		}
	}
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Menu.type"`)}
	}
	if _, ok := mc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "Menu.icon"`)}
	}
	if v, ok := mc.mutation.Icon(); ok {
		if err := menu.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "Menu.icon": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Menu.path"`)}
	}
	if v, ok := mc.mutation.Path(); ok {
		if err := menu.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Menu.path": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Menu.status"`)}
	}
	if _, ok := mc.mutation.ParentPath(); !ok {
		return &ValidationError{Name: "parent_path", err: errors.New(`ent: missing required field "Menu.parent_path"`)}
	}
	if v, ok := mc.mutation.ParentPath(); ok {
		if err := menu.ParentPathValidator(v); err != nil {
			return &ValidationError{Name: "parent_path", err: fmt.Errorf(`ent: validator failed for field "Menu.parent_path": %w`, err)}
		}
	}
	if _, ok := mc.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`ent: missing required field "Menu.sequence"`)}
	}
	if v, ok := mc.mutation.ParentID(); ok {
		if err := menu.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Menu.parent_id": %w`, err)}
		}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := menu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Menu.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(menu.Table, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreateTime(); ok {
		_spec.SetField(menu.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mc.mutation.UpdateTime(); ok {
		_spec.SetField(menu.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := mc.mutation.Keyword(); ok {
		_spec.SetField(menu.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.I18nKey(); ok {
		_spec.SetField(menu.FieldI18nKey, field.TypeString, value)
		_node.I18nKey = value
	}
	if value, ok := mc.mutation.Description(); ok {
		_spec.SetField(menu.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := mc.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.ParentPath(); ok {
		_spec.SetField(menu.FieldParentPath, field.TypeString, value)
		_node.ParentPath = value
	}
	if value, ok := mc.mutation.Sequence(); ok {
		_spec.SetField(menu.FieldSequence, field.TypeInt, value)
		_node.Sequence = value
	}
	if value, ok := mc.mutation.Properties(); ok {
		_spec.SetField(menu.FieldProperties, field.TypeString, value)
		_node.Properties = value
	}
	if nodes := mc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ResourcesTable,
			Columns: []string{menu.ResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoleMenuCreate{config: mc.config, mutation: newRoleMenuMutation(mc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.PermissionsTable,
			Columns: menu.PermissionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &MenuPermissionCreate{config: mc.config, mutation: newMenuPermissionMutation(mc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RoleMenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   menu.RoleMenusTable,
			Columns: []string{menu.RoleMenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MenuPermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   menu.MenuPermissionsTable,
			Columns: []string{menu.MenuPermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menupermission.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetMenu set the Menu
func (mc *MenuCreate) SetMenu(input *Menu, fields ...string) *MenuCreate {
	m := mc.mutation
	if len(fields) == 0 {
		fields = menu.Columns
	}
	_ = m.SetFields(input, fields...)
	return mc
}

// SetMenuWithZero set the Menu
func (mc *MenuCreate) SetMenuWithZero(input *Menu, fields ...string) *MenuCreate {
	m := mc.mutation
	if len(fields) == 0 {
		fields = menu.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return mc
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	err      error
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
