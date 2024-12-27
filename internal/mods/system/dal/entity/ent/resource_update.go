// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks     []Hook
	mutation  *ResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdateTime sets the "update_time" field.
func (ru *ResourceUpdate) SetUpdateTime(t time.Time) *ResourceUpdate {
	ru.mutation.SetUpdateTime(t)
	return ru
}

// SetMethod sets the "method" field.
func (ru *ResourceUpdate) SetMethod(s string) *ResourceUpdate {
	ru.mutation.SetMethod(s)
	return ru
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableMethod(s *string) *ResourceUpdate {
	if s != nil {
		ru.SetMethod(*s)
	}
	return ru
}

// SetOperation sets the "operation" field.
func (ru *ResourceUpdate) SetOperation(s string) *ResourceUpdate {
	ru.mutation.SetOperation(s)
	return ru
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableOperation(s *string) *ResourceUpdate {
	if s != nil {
		ru.SetOperation(*s)
	}
	return ru
}

// SetPath sets the "path" field.
func (ru *ResourceUpdate) SetPath(s string) *ResourceUpdate {
	ru.mutation.SetPath(s)
	return ru
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillablePath(s *string) *ResourceUpdate {
	if s != nil {
		ru.SetPath(*s)
	}
	return ru
}

// SetMenuID sets the "menu_id" field.
func (ru *ResourceUpdate) SetMenuID(i int64) *ResourceUpdate {
	ru.mutation.SetMenuID(i)
	return ru
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableMenuID(i *int64) *ResourceUpdate {
	if i != nil {
		ru.SetMenuID(*i)
	}
	return ru
}

// ClearMenuID clears the value of the "menu_id" field.
func (ru *ResourceUpdate) ClearMenuID() *ResourceUpdate {
	ru.mutation.ClearMenuID()
	return ru
}

// SetMenu sets the "menu" edge to the Menu entity.
func (ru *ResourceUpdate) SetMenu(m *Menu) *ResourceUpdate {
	return ru.SetMenuID(m.ID)
}

// AddPermissionIDs adds the "permission" edge to the Permission entity by IDs.
func (ru *ResourceUpdate) AddPermissionIDs(ids ...int64) *ResourceUpdate {
	ru.mutation.AddPermissionIDs(ids...)
	return ru
}

// AddPermission adds the "permission" edges to the Permission entity.
func (ru *ResourceUpdate) AddPermission(p ...*Permission) *ResourceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermissionIDs(ids...)
}

// AddPermissionResourceIDs adds the "permission_resources" edge to the PermissionResource entity by IDs.
func (ru *ResourceUpdate) AddPermissionResourceIDs(ids ...int64) *ResourceUpdate {
	ru.mutation.AddPermissionResourceIDs(ids...)
	return ru
}

// AddPermissionResources adds the "permission_resources" edges to the PermissionResource entity.
func (ru *ResourceUpdate) AddPermissionResources(p ...*PermissionResource) *ResourceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddPermissionResourceIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (ru *ResourceUpdate) ClearMenu() *ResourceUpdate {
	ru.mutation.ClearMenu()
	return ru
}

// ClearPermission clears all "permission" edges to the Permission entity.
func (ru *ResourceUpdate) ClearPermission() *ResourceUpdate {
	ru.mutation.ClearPermission()
	return ru
}

// RemovePermissionIDs removes the "permission" edge to Permission entities by IDs.
func (ru *ResourceUpdate) RemovePermissionIDs(ids ...int64) *ResourceUpdate {
	ru.mutation.RemovePermissionIDs(ids...)
	return ru
}

// RemovePermission removes "permission" edges to Permission entities.
func (ru *ResourceUpdate) RemovePermission(p ...*Permission) *ResourceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermissionIDs(ids...)
}

// ClearPermissionResources clears all "permission_resources" edges to the PermissionResource entity.
func (ru *ResourceUpdate) ClearPermissionResources() *ResourceUpdate {
	ru.mutation.ClearPermissionResources()
	return ru
}

// RemovePermissionResourceIDs removes the "permission_resources" edge to PermissionResource entities by IDs.
func (ru *ResourceUpdate) RemovePermissionResourceIDs(ids ...int64) *ResourceUpdate {
	ru.mutation.RemovePermissionResourceIDs(ids...)
	return ru
}

// RemovePermissionResources removes "permission_resources" edges to PermissionResource entities.
func (ru *ResourceUpdate) RemovePermissionResources(p ...*PermissionResource) *ResourceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemovePermissionResourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ResourceUpdate) defaults() {
	if _, ok := ru.mutation.UpdateTime(); !ok {
		v := resource.UpdateDefaultUpdateTime()
		ru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResourceUpdate) check() error {
	if v, ok := ru.mutation.Method(); ok {
		if err := resource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Resource.method": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Operation(); ok {
		if err := resource.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "Resource.operation": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if v, ok := ru.mutation.MenuID(); ok {
		if err := resource.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "Resource.menu_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *ResourceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdateTime(); ok {
		_spec.SetField(resource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Method(); ok {
		_spec.SetField(resource.FieldMethod, field.TypeString, value)
	}
	if value, ok := ru.mutation.Operation(); ok {
		_spec.SetField(resource.FieldOperation, field.TypeString, value)
	}
	if value, ok := ru.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
	}
	if ru.mutation.MenuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MenuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PermissionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermissionIDs(); len(nodes) > 0 && !ru.mutation.PermissionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermissionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PermissionResourcesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedPermissionResourcesIDs(); len(nodes) > 0 && !ru.mutation.PermissionResourcesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PermissionResourcesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (ruo *ResourceUpdateOne) SetUpdateTime(t time.Time) *ResourceUpdateOne {
	ruo.mutation.SetUpdateTime(t)
	return ruo
}

// SetMethod sets the "method" field.
func (ruo *ResourceUpdateOne) SetMethod(s string) *ResourceUpdateOne {
	ruo.mutation.SetMethod(s)
	return ruo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableMethod(s *string) *ResourceUpdateOne {
	if s != nil {
		ruo.SetMethod(*s)
	}
	return ruo
}

// SetOperation sets the "operation" field.
func (ruo *ResourceUpdateOne) SetOperation(s string) *ResourceUpdateOne {
	ruo.mutation.SetOperation(s)
	return ruo
}

// SetNillableOperation sets the "operation" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableOperation(s *string) *ResourceUpdateOne {
	if s != nil {
		ruo.SetOperation(*s)
	}
	return ruo
}

// SetPath sets the "path" field.
func (ruo *ResourceUpdateOne) SetPath(s string) *ResourceUpdateOne {
	ruo.mutation.SetPath(s)
	return ruo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillablePath(s *string) *ResourceUpdateOne {
	if s != nil {
		ruo.SetPath(*s)
	}
	return ruo
}

// SetMenuID sets the "menu_id" field.
func (ruo *ResourceUpdateOne) SetMenuID(i int64) *ResourceUpdateOne {
	ruo.mutation.SetMenuID(i)
	return ruo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableMenuID(i *int64) *ResourceUpdateOne {
	if i != nil {
		ruo.SetMenuID(*i)
	}
	return ruo
}

// ClearMenuID clears the value of the "menu_id" field.
func (ruo *ResourceUpdateOne) ClearMenuID() *ResourceUpdateOne {
	ruo.mutation.ClearMenuID()
	return ruo
}

// SetMenu sets the "menu" edge to the Menu entity.
func (ruo *ResourceUpdateOne) SetMenu(m *Menu) *ResourceUpdateOne {
	return ruo.SetMenuID(m.ID)
}

// AddPermissionIDs adds the "permission" edge to the Permission entity by IDs.
func (ruo *ResourceUpdateOne) AddPermissionIDs(ids ...int64) *ResourceUpdateOne {
	ruo.mutation.AddPermissionIDs(ids...)
	return ruo
}

// AddPermission adds the "permission" edges to the Permission entity.
func (ruo *ResourceUpdateOne) AddPermission(p ...*Permission) *ResourceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermissionIDs(ids...)
}

// AddPermissionResourceIDs adds the "permission_resources" edge to the PermissionResource entity by IDs.
func (ruo *ResourceUpdateOne) AddPermissionResourceIDs(ids ...int64) *ResourceUpdateOne {
	ruo.mutation.AddPermissionResourceIDs(ids...)
	return ruo
}

// AddPermissionResources adds the "permission_resources" edges to the PermissionResource entity.
func (ruo *ResourceUpdateOne) AddPermissionResources(p ...*PermissionResource) *ResourceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddPermissionResourceIDs(ids...)
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (ruo *ResourceUpdateOne) ClearMenu() *ResourceUpdateOne {
	ruo.mutation.ClearMenu()
	return ruo
}

// ClearPermission clears all "permission" edges to the Permission entity.
func (ruo *ResourceUpdateOne) ClearPermission() *ResourceUpdateOne {
	ruo.mutation.ClearPermission()
	return ruo
}

// RemovePermissionIDs removes the "permission" edge to Permission entities by IDs.
func (ruo *ResourceUpdateOne) RemovePermissionIDs(ids ...int64) *ResourceUpdateOne {
	ruo.mutation.RemovePermissionIDs(ids...)
	return ruo
}

// RemovePermission removes "permission" edges to Permission entities.
func (ruo *ResourceUpdateOne) RemovePermission(p ...*Permission) *ResourceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermissionIDs(ids...)
}

// ClearPermissionResources clears all "permission_resources" edges to the PermissionResource entity.
func (ruo *ResourceUpdateOne) ClearPermissionResources() *ResourceUpdateOne {
	ruo.mutation.ClearPermissionResources()
	return ruo
}

// RemovePermissionResourceIDs removes the "permission_resources" edge to PermissionResource entities by IDs.
func (ruo *ResourceUpdateOne) RemovePermissionResourceIDs(ids ...int64) *ResourceUpdateOne {
	ruo.mutation.RemovePermissionResourceIDs(ids...)
	return ruo
}

// RemovePermissionResources removes "permission_resources" edges to PermissionResource entities.
func (ruo *ResourceUpdateOne) RemovePermissionResources(p ...*PermissionResource) *ResourceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemovePermissionResourceIDs(ids...)
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ruo *ResourceUpdateOne) Where(ps ...predicate.Resource) *ResourceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourceUpdateOne) Select(field string, fields ...string) *ResourceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resource entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ResourceUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateTime(); !ok {
		v := resource.UpdateDefaultUpdateTime()
		ruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResourceUpdateOne) check() error {
	if v, ok := ruo.mutation.Method(); ok {
		if err := resource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "Resource.method": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Operation(); ok {
		if err := resource.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "Resource.operation": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Path(); ok {
		if err := resource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Resource.path": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.MenuID(); ok {
		if err := resource.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "Resource.menu_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *ResourceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ResourceUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(resource.Table, resource.Columns, sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Resource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for _, f := range fields {
			if !resource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdateTime(); ok {
		_spec.SetField(resource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Method(); ok {
		_spec.SetField(resource.FieldMethod, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Operation(); ok {
		_spec.SetField(resource.FieldOperation, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Path(); ok {
		_spec.SetField(resource.FieldPath, field.TypeString, value)
	}
	if ruo.mutation.MenuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MenuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PermissionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermissionIDs(); len(nodes) > 0 && !ruo.mutation.PermissionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermissionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PermissionResourcesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedPermissionResourcesIDs(); len(nodes) > 0 && !ruo.mutation.PermissionResourcesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PermissionResourcesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}

// SetResource set the Resource
func (ru *ResourceUpdate) SetResource(input *Resource, fields ...string) *ResourceUpdate {
	m := ru.mutation
	if len(fields) == 0 {
		fields = resource.OmitColumns(resource.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return ru
}

// SetResourceWithZero set the Resource
func (ru *ResourceUpdate) SetResourceWithZero(input *Resource, fields ...string) *ResourceUpdate {
	m := ru.mutation
	if len(fields) == 0 {
		fields = resource.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return ru
}

// SetResource set the Resource
func (ruo *ResourceUpdateOne) SetResource(input *Resource, fields ...string) *ResourceUpdateOne {
	m := ruo.mutation
	if len(fields) == 0 {
		fields = resource.OmitColumns(resource.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return ruo
}

// SetResourceWithZero set the Resource
func (ruo *ResourceUpdateOne) SetResourceWithZero(input *Resource, fields ...string) *ResourceUpdateOne {
	m := ruo.mutation
	if len(fields) == 0 {
		fields = resource.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return ruo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ruo *ResourceUpdateOne) Omit(fields ...string) *ResourceUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	ruo.fields = []string(nil)
	for _, col := range resource.Columns {
		if _, ok := omits[col]; !ok {
			ruo.fields = append(ruo.fields, col)
		}
	}
	return ruo
}
