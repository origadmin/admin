// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/rolemenu"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleMenuUpdate is the builder for updating RoleMenu entities.
type RoleMenuUpdate struct {
	config
	hooks     []Hook
	mutation  *RoleMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RoleMenuUpdate builder.
func (rmu *RoleMenuUpdate) Where(ps ...predicate.RoleMenu) *RoleMenuUpdate {
	rmu.mutation.Where(ps...)
	return rmu
}

// SetUpdateTime sets the "update_time" field.
func (rmu *RoleMenuUpdate) SetUpdateTime(t time.Time) *RoleMenuUpdate {
	rmu.mutation.SetUpdateTime(t)
	return rmu
}

// SetRoleID sets the "role_id" field.
func (rmu *RoleMenuUpdate) SetRoleID(s string) *RoleMenuUpdate {
	rmu.mutation.SetRoleID(s)
	return rmu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rmu *RoleMenuUpdate) SetNillableRoleID(s *string) *RoleMenuUpdate {
	if s != nil {
		rmu.SetRoleID(*s)
	}
	return rmu
}

// SetMenuID sets the "menu_id" field.
func (rmu *RoleMenuUpdate) SetMenuID(s string) *RoleMenuUpdate {
	rmu.mutation.SetMenuID(s)
	return rmu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (rmu *RoleMenuUpdate) SetNillableMenuID(s *string) *RoleMenuUpdate {
	if s != nil {
		rmu.SetMenuID(*s)
	}
	return rmu
}

// SetRole sets the "role" edge to the Role entity.
func (rmu *RoleMenuUpdate) SetRole(r *Role) *RoleMenuUpdate {
	return rmu.SetRoleID(r.ID)
}

// SetMenu sets the "menu" edge to the Menu entity.
func (rmu *RoleMenuUpdate) SetMenu(m *Menu) *RoleMenuUpdate {
	return rmu.SetMenuID(m.ID)
}

// Mutation returns the RoleMenuMutation object of the builder.
func (rmu *RoleMenuUpdate) Mutation() *RoleMenuMutation {
	return rmu.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rmu *RoleMenuUpdate) ClearRole() *RoleMenuUpdate {
	rmu.mutation.ClearRole()
	return rmu
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (rmu *RoleMenuUpdate) ClearMenu() *RoleMenuUpdate {
	rmu.mutation.ClearMenu()
	return rmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rmu *RoleMenuUpdate) Save(ctx context.Context) (int, error) {
	rmu.defaults()
	return withHooks(ctx, rmu.sqlSave, rmu.mutation, rmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmu *RoleMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := rmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rmu *RoleMenuUpdate) Exec(ctx context.Context) error {
	_, err := rmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmu *RoleMenuUpdate) ExecX(ctx context.Context) {
	if err := rmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmu *RoleMenuUpdate) defaults() {
	if _, ok := rmu.mutation.UpdateTime(); !ok {
		v := rolemenu.UpdateDefaultUpdateTime()
		rmu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmu *RoleMenuUpdate) check() error {
	if v, ok := rmu.mutation.RoleID(); ok {
		if err := rolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.role_id": %w`, err)}
		}
	}
	if v, ok := rmu.mutation.MenuID(); ok {
		if err := rolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.menu_id": %w`, err)}
		}
	}
	if rmu.mutation.RoleCleared() && len(rmu.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleMenu.role"`)
	}
	if rmu.mutation.MenuCleared() && len(rmu.mutation.MenuIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleMenu.menu"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rmu *RoleMenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleMenuUpdate {
	rmu.modifiers = append(rmu.modifiers, modifiers...)
	return rmu
}

func (rmu *RoleMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rolemenu.Table, rolemenu.Columns, sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeString))
	if ps := rmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmu.mutation.UpdateTime(); ok {
		_spec.SetField(rolemenu.FieldUpdateTime, field.TypeTime, value)
	}
	if rmu.mutation.RoleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RoleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmu.mutation.MenuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.MenuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, rmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rmu.mutation.done = true
	return n, nil
}

// RoleMenuUpdateOne is the builder for updating a single RoleMenu entity.
type RoleMenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RoleMenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (rmuo *RoleMenuUpdateOne) SetUpdateTime(t time.Time) *RoleMenuUpdateOne {
	rmuo.mutation.SetUpdateTime(t)
	return rmuo
}

// SetRoleID sets the "role_id" field.
func (rmuo *RoleMenuUpdateOne) SetRoleID(s string) *RoleMenuUpdateOne {
	rmuo.mutation.SetRoleID(s)
	return rmuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rmuo *RoleMenuUpdateOne) SetNillableRoleID(s *string) *RoleMenuUpdateOne {
	if s != nil {
		rmuo.SetRoleID(*s)
	}
	return rmuo
}

// SetMenuID sets the "menu_id" field.
func (rmuo *RoleMenuUpdateOne) SetMenuID(s string) *RoleMenuUpdateOne {
	rmuo.mutation.SetMenuID(s)
	return rmuo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (rmuo *RoleMenuUpdateOne) SetNillableMenuID(s *string) *RoleMenuUpdateOne {
	if s != nil {
		rmuo.SetMenuID(*s)
	}
	return rmuo
}

// SetRole sets the "role" edge to the Role entity.
func (rmuo *RoleMenuUpdateOne) SetRole(r *Role) *RoleMenuUpdateOne {
	return rmuo.SetRoleID(r.ID)
}

// SetMenu sets the "menu" edge to the Menu entity.
func (rmuo *RoleMenuUpdateOne) SetMenu(m *Menu) *RoleMenuUpdateOne {
	return rmuo.SetMenuID(m.ID)
}

// Mutation returns the RoleMenuMutation object of the builder.
func (rmuo *RoleMenuUpdateOne) Mutation() *RoleMenuMutation {
	return rmuo.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rmuo *RoleMenuUpdateOne) ClearRole() *RoleMenuUpdateOne {
	rmuo.mutation.ClearRole()
	return rmuo
}

// ClearMenu clears the "menu" edge to the Menu entity.
func (rmuo *RoleMenuUpdateOne) ClearMenu() *RoleMenuUpdateOne {
	rmuo.mutation.ClearMenu()
	return rmuo
}

// Where appends a list predicates to the RoleMenuUpdate builder.
func (rmuo *RoleMenuUpdateOne) Where(ps ...predicate.RoleMenu) *RoleMenuUpdateOne {
	rmuo.mutation.Where(ps...)
	return rmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rmuo *RoleMenuUpdateOne) Select(field string, fields ...string) *RoleMenuUpdateOne {
	rmuo.fields = append([]string{field}, fields...)
	return rmuo
}

// Save executes the query and returns the updated RoleMenu entity.
func (rmuo *RoleMenuUpdateOne) Save(ctx context.Context) (*RoleMenu, error) {
	rmuo.defaults()
	return withHooks(ctx, rmuo.sqlSave, rmuo.mutation, rmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmuo *RoleMenuUpdateOne) SaveX(ctx context.Context) *RoleMenu {
	node, err := rmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rmuo *RoleMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := rmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmuo *RoleMenuUpdateOne) ExecX(ctx context.Context) {
	if err := rmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmuo *RoleMenuUpdateOne) defaults() {
	if _, ok := rmuo.mutation.UpdateTime(); !ok {
		v := rolemenu.UpdateDefaultUpdateTime()
		rmuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmuo *RoleMenuUpdateOne) check() error {
	if v, ok := rmuo.mutation.RoleID(); ok {
		if err := rolemenu.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.role_id": %w`, err)}
		}
	}
	if v, ok := rmuo.mutation.MenuID(); ok {
		if err := rolemenu.MenuIDValidator(v); err != nil {
			return &ValidationError{Name: "menu_id", err: fmt.Errorf(`ent: validator failed for field "RoleMenu.menu_id": %w`, err)}
		}
	}
	if rmuo.mutation.RoleCleared() && len(rmuo.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleMenu.role"`)
	}
	if rmuo.mutation.MenuCleared() && len(rmuo.mutation.MenuIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RoleMenu.menu"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (rmuo *RoleMenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RoleMenuUpdateOne {
	rmuo.modifiers = append(rmuo.modifiers, modifiers...)
	return rmuo
}

func (rmuo *RoleMenuUpdateOne) sqlSave(ctx context.Context) (_node *RoleMenu, err error) {
	if err := rmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rolemenu.Table, rolemenu.Columns, sqlgraph.NewFieldSpec(rolemenu.FieldID, field.TypeString))
	id, ok := rmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoleMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolemenu.FieldID)
		for _, f := range fields {
			if !rolemenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rmuo.mutation.UpdateTime(); ok {
		_spec.SetField(rolemenu.FieldUpdateTime, field.TypeTime, value)
	}
	if rmuo.mutation.RoleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RoleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmuo.mutation.MenuCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.MenuIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(rmuo.modifiers...)
	_node = &RoleMenu{config: rmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rmuo.mutation.done = true
	return _node, nil
}

// SetRoleMenu set the RoleMenu
func (rmu *RoleMenuUpdate) SetRoleMenu(input *RoleMenu, fields ...string) *RoleMenuUpdate {
	m := rmu.mutation
	_ = m.SetFields(input, fields...)
	return rmu
}

// SetRoleMenu set the RoleMenu
func (rmuo *RoleMenuUpdateOne) SetRoleMenu(input *RoleMenu, fields ...string) *RoleMenuUpdateOne {
	m := rmuo.mutation
	_ = m.SetFields(input, fields...)
	return rmuo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (rmuo *RoleMenuUpdateOne) Omit(fields ...string) *RoleMenuUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	rmuo.fields = []string(nil)
	for _, col := range rolemenu.Columns {
		if _, ok := omits[col]; !ok {
			rmuo.fields = append(rmuo.fields, col)
		}
	}
	return rmuo
}