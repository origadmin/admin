// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/positionpermission"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionPermissionUpdate is the builder for updating PositionPermission entities.
type PositionPermissionUpdate struct {
	config
	hooks     []Hook
	mutation  *PositionPermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PositionPermissionUpdate builder.
func (ppu *PositionPermissionUpdate) Where(ps ...predicate.PositionPermission) *PositionPermissionUpdate {
	ppu.mutation.Where(ps...)
	return ppu
}

// SetPositionID sets the "position_id" field.
func (ppu *PositionPermissionUpdate) SetPositionID(i int64) *PositionPermissionUpdate {
	ppu.mutation.SetPositionID(i)
	return ppu
}

// SetNillablePositionID sets the "position_id" field if the given value is not nil.
func (ppu *PositionPermissionUpdate) SetNillablePositionID(i *int64) *PositionPermissionUpdate {
	if i != nil {
		ppu.SetPositionID(*i)
	}
	return ppu
}

// SetPermissionID sets the "permission_id" field.
func (ppu *PositionPermissionUpdate) SetPermissionID(i int64) *PositionPermissionUpdate {
	ppu.mutation.SetPermissionID(i)
	return ppu
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (ppu *PositionPermissionUpdate) SetNillablePermissionID(i *int64) *PositionPermissionUpdate {
	if i != nil {
		ppu.SetPermissionID(*i)
	}
	return ppu
}

// SetPosition sets the "position" edge to the Position entity.
func (ppu *PositionPermissionUpdate) SetPosition(p *Position) *PositionPermissionUpdate {
	return ppu.SetPositionID(p.ID)
}

// SetPermission sets the "permission" edge to the Permission entity.
func (ppu *PositionPermissionUpdate) SetPermission(p *Permission) *PositionPermissionUpdate {
	return ppu.SetPermissionID(p.ID)
}

// Mutation returns the PositionPermissionMutation object of the builder.
func (ppu *PositionPermissionUpdate) Mutation() *PositionPermissionMutation {
	return ppu.mutation
}

// ClearPosition clears the "position" edge to the Position entity.
func (ppu *PositionPermissionUpdate) ClearPosition() *PositionPermissionUpdate {
	ppu.mutation.ClearPosition()
	return ppu
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (ppu *PositionPermissionUpdate) ClearPermission() *PositionPermissionUpdate {
	ppu.mutation.ClearPermission()
	return ppu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ppu *PositionPermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ppu.sqlSave, ppu.mutation, ppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *PositionPermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *PositionPermissionUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *PositionPermissionUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppu *PositionPermissionUpdate) check() error {
	if v, ok := ppu.mutation.PositionID(); ok {
		if err := positionpermission.PositionIDValidator(v); err != nil {
			return &ValidationError{Name: "position_id", err: fmt.Errorf(`ent: validator failed for field "PositionPermission.position_id": %w`, err)}
		}
	}
	if v, ok := ppu.mutation.PermissionID(); ok {
		if err := positionpermission.PermissionIDValidator(v); err != nil {
			return &ValidationError{Name: "permission_id", err: fmt.Errorf(`ent: validator failed for field "PositionPermission.permission_id": %w`, err)}
		}
	}
	if ppu.mutation.PositionCleared() && len(ppu.mutation.PositionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PositionPermission.position"`)
	}
	if ppu.mutation.PermissionCleared() && len(ppu.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PositionPermission.permission"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ppu *PositionPermissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionPermissionUpdate {
	ppu.modifiers = append(ppu.modifiers, modifiers...)
	return ppu
}

func (ppu *PositionPermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ppu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(positionpermission.Table, positionpermission.Columns, sqlgraph.NewFieldSpec(positionpermission.FieldID, field.TypeInt64))
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ppu.mutation.PositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PositionTable,
			Columns: []string{positionpermission.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PositionTable,
			Columns: []string{positionpermission.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppu.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PermissionTable,
			Columns: []string{positionpermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PermissionTable,
			Columns: []string{positionpermission.PermissionColumn},
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
	_spec.AddModifiers(ppu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{positionpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ppu.mutation.done = true
	return n, nil
}

// PositionPermissionUpdateOne is the builder for updating a single PositionPermission entity.
type PositionPermissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PositionPermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPositionID sets the "position_id" field.
func (ppuo *PositionPermissionUpdateOne) SetPositionID(i int64) *PositionPermissionUpdateOne {
	ppuo.mutation.SetPositionID(i)
	return ppuo
}

// SetNillablePositionID sets the "position_id" field if the given value is not nil.
func (ppuo *PositionPermissionUpdateOne) SetNillablePositionID(i *int64) *PositionPermissionUpdateOne {
	if i != nil {
		ppuo.SetPositionID(*i)
	}
	return ppuo
}

// SetPermissionID sets the "permission_id" field.
func (ppuo *PositionPermissionUpdateOne) SetPermissionID(i int64) *PositionPermissionUpdateOne {
	ppuo.mutation.SetPermissionID(i)
	return ppuo
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (ppuo *PositionPermissionUpdateOne) SetNillablePermissionID(i *int64) *PositionPermissionUpdateOne {
	if i != nil {
		ppuo.SetPermissionID(*i)
	}
	return ppuo
}

// SetPosition sets the "position" edge to the Position entity.
func (ppuo *PositionPermissionUpdateOne) SetPosition(p *Position) *PositionPermissionUpdateOne {
	return ppuo.SetPositionID(p.ID)
}

// SetPermission sets the "permission" edge to the Permission entity.
func (ppuo *PositionPermissionUpdateOne) SetPermission(p *Permission) *PositionPermissionUpdateOne {
	return ppuo.SetPermissionID(p.ID)
}

// Mutation returns the PositionPermissionMutation object of the builder.
func (ppuo *PositionPermissionUpdateOne) Mutation() *PositionPermissionMutation {
	return ppuo.mutation
}

// ClearPosition clears the "position" edge to the Position entity.
func (ppuo *PositionPermissionUpdateOne) ClearPosition() *PositionPermissionUpdateOne {
	ppuo.mutation.ClearPosition()
	return ppuo
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (ppuo *PositionPermissionUpdateOne) ClearPermission() *PositionPermissionUpdateOne {
	ppuo.mutation.ClearPermission()
	return ppuo
}

// Where appends a list predicates to the PositionPermissionUpdate builder.
func (ppuo *PositionPermissionUpdateOne) Where(ps ...predicate.PositionPermission) *PositionPermissionUpdateOne {
	ppuo.mutation.Where(ps...)
	return ppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ppuo *PositionPermissionUpdateOne) Select(field string, fields ...string) *PositionPermissionUpdateOne {
	ppuo.fields = append([]string{field}, fields...)
	return ppuo
}

// Save executes the query and returns the updated PositionPermission entity.
func (ppuo *PositionPermissionUpdateOne) Save(ctx context.Context) (*PositionPermission, error) {
	return withHooks(ctx, ppuo.sqlSave, ppuo.mutation, ppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *PositionPermissionUpdateOne) SaveX(ctx context.Context) *PositionPermission {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *PositionPermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *PositionPermissionUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppuo *PositionPermissionUpdateOne) check() error {
	if v, ok := ppuo.mutation.PositionID(); ok {
		if err := positionpermission.PositionIDValidator(v); err != nil {
			return &ValidationError{Name: "position_id", err: fmt.Errorf(`ent: validator failed for field "PositionPermission.position_id": %w`, err)}
		}
	}
	if v, ok := ppuo.mutation.PermissionID(); ok {
		if err := positionpermission.PermissionIDValidator(v); err != nil {
			return &ValidationError{Name: "permission_id", err: fmt.Errorf(`ent: validator failed for field "PositionPermission.permission_id": %w`, err)}
		}
	}
	if ppuo.mutation.PositionCleared() && len(ppuo.mutation.PositionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PositionPermission.position"`)
	}
	if ppuo.mutation.PermissionCleared() && len(ppuo.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "PositionPermission.permission"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ppuo *PositionPermissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionPermissionUpdateOne {
	ppuo.modifiers = append(ppuo.modifiers, modifiers...)
	return ppuo
}

func (ppuo *PositionPermissionUpdateOne) sqlSave(ctx context.Context) (_node *PositionPermission, err error) {
	if err := ppuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(positionpermission.Table, positionpermission.Columns, sqlgraph.NewFieldSpec(positionpermission.FieldID, field.TypeInt64))
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PositionPermission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, positionpermission.FieldID)
		for _, f := range fields {
			if !positionpermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != positionpermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ppuo.mutation.PositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PositionTable,
			Columns: []string{positionpermission.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PositionTable,
			Columns: []string{positionpermission.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ppuo.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PermissionTable,
			Columns: []string{positionpermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   positionpermission.PermissionTable,
			Columns: []string{positionpermission.PermissionColumn},
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
	_spec.AddModifiers(ppuo.modifiers...)
	_node = &PositionPermission{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{positionpermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ppuo.mutation.done = true
	return _node, nil
}

// SetPositionPermission set the PositionPermission
func (ppu *PositionPermissionUpdate) SetPositionPermission(input *PositionPermission, fields ...string) *PositionPermissionUpdate {
	m := ppu.mutation
	if len(fields) == 0 {
		fields = positionpermission.OmitColumns(positionpermission.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return ppu
}

// SetPositionPermissionWithZero set the PositionPermission
func (ppu *PositionPermissionUpdate) SetPositionPermissionWithZero(input *PositionPermission, fields ...string) *PositionPermissionUpdate {
	m := ppu.mutation
	if len(fields) == 0 {
		fields = positionpermission.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return ppu
}

// SetPositionPermission set the PositionPermission
func (ppuo *PositionPermissionUpdateOne) SetPositionPermission(input *PositionPermission, fields ...string) *PositionPermissionUpdateOne {
	m := ppuo.mutation
	if len(fields) == 0 {
		fields = positionpermission.OmitColumns(positionpermission.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return ppuo
}

// SetPositionPermissionWithZero set the PositionPermission
func (ppuo *PositionPermissionUpdateOne) SetPositionPermissionWithZero(input *PositionPermission, fields ...string) *PositionPermissionUpdateOne {
	m := ppuo.mutation
	if len(fields) == 0 {
		fields = positionpermission.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return ppuo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ppuo *PositionPermissionUpdateOne) Omit(fields ...string) *PositionPermissionUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	ppuo.fields = []string(nil)
	for _, col := range positionpermission.Columns {
		if _, ok := omits[col]; !ok {
			ppuo.fields = append(ppuo.fields, col)
		}
	}
	return ppuo
}
