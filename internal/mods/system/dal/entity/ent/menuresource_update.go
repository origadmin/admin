// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menuresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuResourceUpdate is the builder for updating MenuResource entities.
type MenuResourceUpdate struct {
	config
	hooks     []Hook
	mutation  *MenuResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MenuResourceUpdate builder.
func (mru *MenuResourceUpdate) Where(ps ...predicate.MenuResource) *MenuResourceUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetUpdateTime sets the "update_time" field.
func (mru *MenuResourceUpdate) SetUpdateTime(t time.Time) *MenuResourceUpdate {
	mru.mutation.SetUpdateTime(t)
	return mru
}

// SetMenuID sets the "menu_id" field.
func (mru *MenuResourceUpdate) SetMenuID(s string) *MenuResourceUpdate {
	mru.mutation.SetMenuID(s)
	return mru
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (mru *MenuResourceUpdate) SetNillableMenuID(s *string) *MenuResourceUpdate {
	if s != nil {
		mru.SetMenuID(*s)
	}
	return mru
}

// ClearMenuID clears the value of the "menu_id" field.
func (mru *MenuResourceUpdate) ClearMenuID() *MenuResourceUpdate {
	mru.mutation.ClearMenuID()
	return mru
}

// SetMethod sets the "method" field.
func (mru *MenuResourceUpdate) SetMethod(s string) *MenuResourceUpdate {
	mru.mutation.SetMethod(s)
	return mru
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (mru *MenuResourceUpdate) SetNillableMethod(s *string) *MenuResourceUpdate {
	if s != nil {
		mru.SetMethod(*s)
	}
	return mru
}

// SetPath sets the "path" field.
func (mru *MenuResourceUpdate) SetPath(s string) *MenuResourceUpdate {
	mru.mutation.SetPath(s)
	return mru
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mru *MenuResourceUpdate) SetNillablePath(s *string) *MenuResourceUpdate {
	if s != nil {
		mru.SetPath(*s)
	}
	return mru
}

// SetMenu sets the "menu" edge to the Menu entity.
func (mru *MenuResourceUpdate) SetMenu(m *Menu) *MenuResourceUpdate {
	return mru.SetMenuID(m.ID)
}

// Mutation returns the MenuResourceMutation object of the builder.
func (mru *MenuResourceUpdate) Mutation() *MenuResourceMutation {
	return mru.mutation
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (mru *MenuResourceUpdate) ClearMenu() *MenuResourceUpdate {
	mru.mutation.ClearMenu()
	return mru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MenuResourceUpdate) Save(ctx context.Context) (int, error) {
	mru.defaults()
	return withHooks(ctx, mru.sqlSave, mru.mutation, mru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MenuResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MenuResourceUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MenuResourceUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mru *MenuResourceUpdate) defaults() {
	if _, ok := mru.mutation.UpdateTime(); !ok {
		v := menuresource.UpdateDefaultUpdateTime()
		mru.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mru *MenuResourceUpdate) check() error {
	if v, ok := mru.mutation.MenuID(); ok {
		if err := menuresource.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "MenuResource.menu_id": %w`, err)}
		}
	}
	if v, ok := mru.mutation.Method(); ok {
		if err := menuresource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "MenuResource.method": %w`, err)}
		}
	}
	if v, ok := mru.mutation.Path(); ok {
		if err := menuresource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "MenuResource.path": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mru *MenuResourceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuResourceUpdate {
	mru.modifiers = append(mru.modifiers, modifiers...)
	return mru
}

func (mru *MenuResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(menuresource.Table, menuresource.Columns, sqlgraph.NewFieldSpec(menuresource.FieldID, field.TypeString))
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.UpdateTime(); ok {
		_spec.SetField(menuresource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mru.mutation.Method(); ok {
		_spec.SetField(menuresource.FieldMethod, field.TypeString, value)
	}
	if value, ok := mru.mutation.Path(); ok {
		_spec.SetField(menuresource.FieldPath, field.TypeString, value)
	}
	if mru.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuresource.MenuTable,
			Columns: []string{menuresource.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuresource.MenuTable,
			Columns: []string{menuresource.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mru.mutation.done = true
	return n, nil
}

// MenuResourceUpdateOne is the builder for updating a single MenuResource entity.
type MenuResourceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MenuResourceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (mruo *MenuResourceUpdateOne) SetUpdateTime(t time.Time) *MenuResourceUpdateOne {
	mruo.mutation.SetUpdateTime(t)
	return mruo
}

// SetMenuID sets the "menu_id" field.
func (mruo *MenuResourceUpdateOne) SetMenuID(s string) *MenuResourceUpdateOne {
	mruo.mutation.SetMenuID(s)
	return mruo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (mruo *MenuResourceUpdateOne) SetNillableMenuID(s *string) *MenuResourceUpdateOne {
	if s != nil {
		mruo.SetMenuID(*s)
	}
	return mruo
}

// ClearMenuID clears the value of the "menu_id" field.
func (mruo *MenuResourceUpdateOne) ClearMenuID() *MenuResourceUpdateOne {
	mruo.mutation.ClearMenuID()
	return mruo
}

// SetMethod sets the "method" field.
func (mruo *MenuResourceUpdateOne) SetMethod(s string) *MenuResourceUpdateOne {
	mruo.mutation.SetMethod(s)
	return mruo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (mruo *MenuResourceUpdateOne) SetNillableMethod(s *string) *MenuResourceUpdateOne {
	if s != nil {
		mruo.SetMethod(*s)
	}
	return mruo
}

// SetPath sets the "path" field.
func (mruo *MenuResourceUpdateOne) SetPath(s string) *MenuResourceUpdateOne {
	mruo.mutation.SetPath(s)
	return mruo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mruo *MenuResourceUpdateOne) SetNillablePath(s *string) *MenuResourceUpdateOne {
	if s != nil {
		mruo.SetPath(*s)
	}
	return mruo
}

// SetMenu sets the "menu" edge to the Menu entity.
func (mruo *MenuResourceUpdateOne) SetMenu(m *Menu) *MenuResourceUpdateOne {
	return mruo.SetMenuID(m.ID)
}

// Mutation returns the MenuResourceMutation object of the builder.
func (mruo *MenuResourceUpdateOne) Mutation() *MenuResourceMutation {
	return mruo.mutation
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (mruo *MenuResourceUpdateOne) ClearMenu() *MenuResourceUpdateOne {
	mruo.mutation.ClearMenu()
	return mruo
}

// Where appends a list predicates to the MenuResourceUpdate builder.
func (mruo *MenuResourceUpdateOne) Where(ps ...predicate.MenuResource) *MenuResourceUpdateOne {
	mruo.mutation.Where(ps...)
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MenuResourceUpdateOne) Select(field string, fields ...string) *MenuResourceUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MenuResource entity.
func (mruo *MenuResourceUpdateOne) Save(ctx context.Context) (*MenuResource, error) {
	mruo.defaults()
	return withHooks(ctx, mruo.sqlSave, mruo.mutation, mruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MenuResourceUpdateOne) SaveX(ctx context.Context) *MenuResource {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MenuResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MenuResourceUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mruo *MenuResourceUpdateOne) defaults() {
	if _, ok := mruo.mutation.UpdateTime(); !ok {
		v := menuresource.UpdateDefaultUpdateTime()
		mruo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mruo *MenuResourceUpdateOne) check() error {
	if v, ok := mruo.mutation.MenuID(); ok {
		if err := menuresource.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "MenuResource.menu_id": %w`, err)}
		}
	}
	if v, ok := mruo.mutation.Method(); ok {
		if err := menuresource.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "MenuResource.method": %w`, err)}
		}
	}
	if v, ok := mruo.mutation.Path(); ok {
		if err := menuresource.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "MenuResource.path": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mruo *MenuResourceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuResourceUpdateOne {
	mruo.modifiers = append(mruo.modifiers, modifiers...)
	return mruo
}

func (mruo *MenuResourceUpdateOne) sqlSave(ctx context.Context) (_node *MenuResource, err error) {
	if err := mruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(menuresource.Table, menuresource.Columns, sqlgraph.NewFieldSpec(menuresource.FieldID, field.TypeString))
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MenuResource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menuresource.FieldID)
		for _, f := range fields {
			if !menuresource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menuresource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.UpdateTime(); ok {
		_spec.SetField(menuresource.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := mruo.mutation.Method(); ok {
		_spec.SetField(menuresource.FieldMethod, field.TypeString, value)
	}
	if value, ok := mruo.mutation.Path(); ok {
		_spec.SetField(menuresource.FieldPath, field.TypeString, value)
	}
	if mruo.mutation.MenuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuresource.MenuTable,
			Columns: []string{menuresource.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.MenuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuresource.MenuTable,
			Columns: []string{menuresource.MenuColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mruo.modifiers...)
	_node = &MenuResource{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menuresource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mruo.mutation.done = true
	return _node, nil
}

// SetMenuResource set the MenuResource
func (mru *MenuResourceUpdate) SetMenuResource(input *MenuResource, fields ...string) *MenuResourceUpdate {
	m := mru.mutation
	_ = m.SetFields(input, fields...)
	return mru
}

// SetMenuResource set the MenuResource
func (mruo *MenuResourceUpdateOne) SetMenuResource(input *MenuResource, fields ...string) *MenuResourceUpdateOne {
	m := mruo.mutation
	_ = m.SetFields(input, fields...)
	return mruo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (mruo *MenuResourceUpdateOne) Omit(fields ...string) *MenuResourceUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	mruo.fields = []string(nil)
	for _, col := range menuresource.Columns {
		if _, ok := omits[col]; !ok {
			mruo.fields = append(mruo.fields, col)
		}
	}
	return mruo
}