// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionResourceUpdate is the builder for updating PermissionResource entities.
type PermissionResourceUpdate struct {
	config
	hooks     []Hook
	mutation  *PermissionResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PermissionResourceUpdate builder.
func (pru *PermissionResourceUpdate) Where(ps ...predicate.PermissionResource) *PermissionResourceUpdate {
	pru.mutation.Where(ps...)
	return pru
}

// SetPermissionID sets the "permission_id" field.
func (pru *PermissionResourceUpdate) SetPermissionID(i int64) *PermissionResourceUpdate {
	pru.mutation.SetPermissionID(i)
	return pru
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (pru *PermissionResourceUpdate) SetNillablePermissionID(i *int64) *PermissionResourceUpdate {
	if i != nil {
		pru.SetPermissionID(*i)
	}
	return pru
}

// SetResourceID sets the "resource_id" field.
func (pru *PermissionResourceUpdate) SetResourceID(i int64) *PermissionResourceUpdate {
	pru.mutation.SetResourceID(i)
	return pru
}

// SetNillableResourceID sets the "resource_id" field if the given value is not nil.
func (pru *PermissionResourceUpdate) SetNillableResourceID(i *int64) *PermissionResourceUpdate {
	if i != nil {
		pru.SetResourceID(*i)
	}
	return pru
}

// SetPermission sets the "permission" edge to the Permission entity.
func (pru *PermissionResourceUpdate) SetPermission(p *Permission) *PermissionResourceUpdate {
	return pru.SetPermissionID(p.ID)
}

// SetResource sets the "resource" edge to the Resource entity.
func (pru *PermissionResourceUpdate) SetResource(r *Resource) *PermissionResourceUpdate {
	return pru.SetResourceID(r.ID)
}

// Mutation returns the PermissionResourceMutation object of the builder.
func (pru *PermissionResourceUpdate) Mutation() *PermissionResourceMutation {
	return pru.mutation
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (pru *PermissionResourceUpdate) ClearPermission() *PermissionResourceUpdate {
	pru.mutation.ClearPermission()
	return pru
}

// ClearResource clears the "resource" edge to the Resource entity.
func (pru *PermissionResourceUpdate) ClearResource() *PermissionResourceUpdate {
	pru.mutation.ClearResource()
	return pru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pru *PermissionResourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pru.sqlSave, pru.mutation, pru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pru *PermissionResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := pru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pru *PermissionResourceUpdate) Exec(ctx context.Context) error {
	_, err := pru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pru *PermissionResourceUpdate) ExecX(ctx context.Context) {
	if err := pru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pru *PermissionResourceUpdate) check() error {
	if v, ok := pru.mutation.PermissionID(); ok {
		if err := permissionresource.PermissionIDValidator(v); err != nil {
			return &ValidationError{Name: "permission_id", err: fmt.Errorf(`ent: validator failed for field "PermissionResource.permission_id": %w`, err)}
		}
	}
	if v, ok := pru.mutation.ResourceID(); ok {
		if err := permissionresource.ResourceIDValidator(v); err != nil {
			return &ValidationError{Name: "resource_id", err: fmt.Errorf(`ent: validator failed for field "PermissionResource.resource_id": %w`, err)}
		}
	}
	if pru.mutation.PermissionCleared() && len(pru.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PermissionResource.permission"`)
	}
	if pru.mutation.ResourceCleared() && len(pru.mutation.ResourceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PermissionResource.resource"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pru *PermissionResourceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionResourceUpdate {
	pru.modifiers = append(pru.modifiers, modifiers...)
	return pru
}

func (pru *PermissionResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permissionresource.Table, permissionresource.Columns, sqlgraph.NewFieldSpec(permissionresource.FieldID, field.TypeInt64))
	if ps := pru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pru.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.PermissionTable,
			Columns: []string{permissionresource.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.PermissionTable,
			Columns: []string{permissionresource.PermissionColumn},
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
	if pru.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.ResourceTable,
			Columns: []string{permissionresource.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pru.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.ResourceTable,
			Columns: []string{permissionresource.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissionresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pru.mutation.done = true
	return n, nil
}

// PermissionResourceUpdateOne is the builder for updating a single PermissionResource entity.
type PermissionResourceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PermissionResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPermissionID sets the "permission_id" field.
func (pruo *PermissionResourceUpdateOne) SetPermissionID(i int64) *PermissionResourceUpdateOne {
	pruo.mutation.SetPermissionID(i)
	return pruo
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (pruo *PermissionResourceUpdateOne) SetNillablePermissionID(i *int64) *PermissionResourceUpdateOne {
	if i != nil {
		pruo.SetPermissionID(*i)
	}
	return pruo
}

// SetResourceID sets the "resource_id" field.
func (pruo *PermissionResourceUpdateOne) SetResourceID(i int64) *PermissionResourceUpdateOne {
	pruo.mutation.SetResourceID(i)
	return pruo
}

// SetNillableResourceID sets the "resource_id" field if the given value is not nil.
func (pruo *PermissionResourceUpdateOne) SetNillableResourceID(i *int64) *PermissionResourceUpdateOne {
	if i != nil {
		pruo.SetResourceID(*i)
	}
	return pruo
}

// SetPermission sets the "permission" edge to the Permission entity.
func (pruo *PermissionResourceUpdateOne) SetPermission(p *Permission) *PermissionResourceUpdateOne {
	return pruo.SetPermissionID(p.ID)
}

// SetResource sets the "resource" edge to the Resource entity.
func (pruo *PermissionResourceUpdateOne) SetResource(r *Resource) *PermissionResourceUpdateOne {
	return pruo.SetResourceID(r.ID)
}

// Mutation returns the PermissionResourceMutation object of the builder.
func (pruo *PermissionResourceUpdateOne) Mutation() *PermissionResourceMutation {
	return pruo.mutation
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (pruo *PermissionResourceUpdateOne) ClearPermission() *PermissionResourceUpdateOne {
	pruo.mutation.ClearPermission()
	return pruo
}

// ClearResource clears the "resource" edge to the Resource entity.
func (pruo *PermissionResourceUpdateOne) ClearResource() *PermissionResourceUpdateOne {
	pruo.mutation.ClearResource()
	return pruo
}

// Where appends a list predicates to the PermissionResourceUpdate builder.
func (pruo *PermissionResourceUpdateOne) Where(ps ...predicate.PermissionResource) *PermissionResourceUpdateOne {
	pruo.mutation.Where(ps...)
	return pruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pruo *PermissionResourceUpdateOne) Select(field string, fields ...string) *PermissionResourceUpdateOne {
	pruo.fields = append([]string{field}, fields...)
	return pruo
}

// Save executes the query and returns the updated PermissionResource entity.
func (pruo *PermissionResourceUpdateOne) Save(ctx context.Context) (*PermissionResource, error) {
	return withHooks(ctx, pruo.sqlSave, pruo.mutation, pruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pruo *PermissionResourceUpdateOne) SaveX(ctx context.Context) *PermissionResource {
	node, err := pruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pruo *PermissionResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := pruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pruo *PermissionResourceUpdateOne) ExecX(ctx context.Context) {
	if err := pruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pruo *PermissionResourceUpdateOne) check() error {
	if v, ok := pruo.mutation.PermissionID(); ok {
		if err := permissionresource.PermissionIDValidator(v); err != nil {
			return &ValidationError{Name: "permission_id", err: fmt.Errorf(`ent: validator failed for field "PermissionResource.permission_id": %w`, err)}
		}
	}
	if v, ok := pruo.mutation.ResourceID(); ok {
		if err := permissionresource.ResourceIDValidator(v); err != nil {
			return &ValidationError{Name: "resource_id", err: fmt.Errorf(`ent: validator failed for field "PermissionResource.resource_id": %w`, err)}
		}
	}
	if pruo.mutation.PermissionCleared() && len(pruo.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PermissionResource.permission"`)
	}
	if pruo.mutation.ResourceCleared() && len(pruo.mutation.ResourceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PermissionResource.resource"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pruo *PermissionResourceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionResourceUpdateOne {
	pruo.modifiers = append(pruo.modifiers, modifiers...)
	return pruo
}

func (pruo *PermissionResourceUpdateOne) sqlSave(ctx context.Context) (_node *PermissionResource, err error) {
	if err := pruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permissionresource.Table, permissionresource.Columns, sqlgraph.NewFieldSpec(permissionresource.FieldID, field.TypeInt64))
	id, ok := pruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PermissionResource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionresource.FieldID)
		for _, f := range fields {
			if !permissionresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permissionresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pruo.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.PermissionTable,
			Columns: []string{permissionresource.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.PermissionTable,
			Columns: []string{permissionresource.PermissionColumn},
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
	if pruo.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.ResourceTable,
			Columns: []string{permissionresource.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pruo.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permissionresource.ResourceTable,
			Columns: []string{permissionresource.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pruo.modifiers...)
	_node = &PermissionResource{config: pruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissionresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pruo.mutation.done = true
	return _node, nil
}

// SetPermissionResource set the PermissionResource
func (pru *PermissionResourceUpdate) SetPermissionResource(input *PermissionResource, fields ...string) *PermissionResourceUpdate {
	m := pru.mutation
	if len(fields) == 0 {
		fields = permissionresource.OmitColumns(permissionresource.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return pru
}

// SetPermissionResourceWithZero set the PermissionResource
func (pru *PermissionResourceUpdate) SetPermissionResourceWithZero(input *PermissionResource, fields ...string) *PermissionResourceUpdate {
	m := pru.mutation
	if len(fields) == 0 {
		fields = permissionresource.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return pru
}

// SetPermissionResource set the PermissionResource
func (pruo *PermissionResourceUpdateOne) SetPermissionResource(input *PermissionResource, fields ...string) *PermissionResourceUpdateOne {
	m := pruo.mutation
	if len(fields) == 0 {
		fields = permissionresource.OmitColumns(permissionresource.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return pruo
}

// SetPermissionResourceWithZero set the PermissionResource
func (pruo *PermissionResourceUpdateOne) SetPermissionResourceWithZero(input *PermissionResource, fields ...string) *PermissionResourceUpdateOne {
	m := pruo.mutation
	if len(fields) == 0 {
		fields = permissionresource.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return pruo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (pruo *PermissionResourceUpdateOne) Omit(fields ...string) *PermissionResourceUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	pruo.fields = []string(nil)
	for _, col := range permissionresource.Columns {
		if _, ok := omits[col]; !ok {
			pruo.fields = append(pruo.fields, col)
		}
	}
	return pruo
}
