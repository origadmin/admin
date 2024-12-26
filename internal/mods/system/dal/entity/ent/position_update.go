// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userposition"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks     []Hook
	mutation  *PositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *PositionUpdate) SetUpdateTime(t time.Time) *PositionUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetName sets the "name" field.
func (pu *PositionUpdate) SetName(s string) *PositionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableName(s *string) *PositionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *PositionUpdate) SetDescription(s string) *PositionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableDescription(s *string) *PositionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetDepartmentID sets the "department_id" field.
func (pu *PositionUpdate) SetDepartmentID(i int) *PositionUpdate {
	pu.mutation.SetDepartmentID(i)
	return pu
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableDepartmentID(i *int) *PositionUpdate {
	if i != nil {
		pu.SetDepartmentID(*i)
	}
	return pu
}

// ClearDepartmentID clears the value of the "department_id" field.
func (pu *PositionUpdate) ClearDepartmentID() *PositionUpdate {
	pu.mutation.ClearDepartmentID()
	return pu
}

// SetDepartment sets the "department" edge to the Department entity.
func (pu *PositionUpdate) SetDepartment(d *Department) *PositionUpdate {
	return pu.SetDepartmentID(d.ID)
}

// AddUserPositionIDs adds the "user_positions" edge to the UserPosition entity by IDs.
func (pu *PositionUpdate) AddUserPositionIDs(ids ...int) *PositionUpdate {
	pu.mutation.AddUserPositionIDs(ids...)
	return pu
}

// AddUserPositions adds the "user_positions" edges to the UserPosition entity.
func (pu *PositionUpdate) AddUserPositions(u ...*UserPosition) *PositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserPositionIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// ClearDepartment clears the "department" edge to the Department entity.
func (pu *PositionUpdate) ClearDepartment() *PositionUpdate {
	pu.mutation.ClearDepartment()
	return pu
}

// ClearUserPositions clears all "user_positions" edges to the UserPosition entity.
func (pu *PositionUpdate) ClearUserPositions() *PositionUpdate {
	pu.mutation.ClearUserPositions()
	return pu
}

// RemoveUserPositionIDs removes the "user_positions" edge to UserPosition entities by IDs.
func (pu *PositionUpdate) RemoveUserPositionIDs(ids ...int) *PositionUpdate {
	pu.mutation.RemoveUserPositionIDs(ids...)
	return pu
}

// RemoveUserPositions removes "user_positions" edges to UserPosition entities.
func (pu *PositionUpdate) RemoveUserPositions(u ...*UserPosition) *PositionUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserPositionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PositionUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := position.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PositionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Description(); ok {
		if err := position.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Position.description": %w`, err)}
		}
	}
	if v, ok := pu.mutation.DepartmentID(); ok {
		if err := position.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Position.department_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PositionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(position.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.DepartmentTable,
			Columns: []string{position.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.DepartmentTable,
			Columns: []string{position.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.UserPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUserPositionsIDs(); len(nodes) > 0 && !pu.mutation.UserPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PositionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "update_time" field.
func (puo *PositionUpdateOne) SetUpdateTime(t time.Time) *PositionUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetName sets the "name" field.
func (puo *PositionUpdateOne) SetName(s string) *PositionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableName(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *PositionUpdateOne) SetDescription(s string) *PositionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableDescription(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetDepartmentID sets the "department_id" field.
func (puo *PositionUpdateOne) SetDepartmentID(i int) *PositionUpdateOne {
	puo.mutation.SetDepartmentID(i)
	return puo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableDepartmentID(i *int) *PositionUpdateOne {
	if i != nil {
		puo.SetDepartmentID(*i)
	}
	return puo
}

// ClearDepartmentID clears the value of the "department_id" field.
func (puo *PositionUpdateOne) ClearDepartmentID() *PositionUpdateOne {
	puo.mutation.ClearDepartmentID()
	return puo
}

// SetDepartment sets the "department" edge to the Department entity.
func (puo *PositionUpdateOne) SetDepartment(d *Department) *PositionUpdateOne {
	return puo.SetDepartmentID(d.ID)
}

// AddUserPositionIDs adds the "user_positions" edge to the UserPosition entity by IDs.
func (puo *PositionUpdateOne) AddUserPositionIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.AddUserPositionIDs(ids...)
	return puo
}

// AddUserPositions adds the "user_positions" edges to the UserPosition entity.
func (puo *PositionUpdateOne) AddUserPositions(u ...*UserPosition) *PositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserPositionIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// ClearDepartment clears the "department" edge to the Department entity.
func (puo *PositionUpdateOne) ClearDepartment() *PositionUpdateOne {
	puo.mutation.ClearDepartment()
	return puo
}

// ClearUserPositions clears all "user_positions" edges to the UserPosition entity.
func (puo *PositionUpdateOne) ClearUserPositions() *PositionUpdateOne {
	puo.mutation.ClearUserPositions()
	return puo
}

// RemoveUserPositionIDs removes the "user_positions" edge to UserPosition entities by IDs.
func (puo *PositionUpdateOne) RemoveUserPositionIDs(ids ...int) *PositionUpdateOne {
	puo.mutation.RemoveUserPositionIDs(ids...)
	return puo
}

// RemoveUserPositions removes "user_positions" edges to UserPosition entities.
func (puo *PositionUpdateOne) RemoveUserPositions(u ...*UserPosition) *PositionUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserPositionIDs(ids...)
}

// Where appends a list predicates to the PositionUpdate builder.
func (puo *PositionUpdateOne) Where(ps ...predicate.Position) *PositionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PositionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := position.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PositionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Description(); ok {
		if err := position.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Position.description": %w`, err)}
		}
	}
	if v, ok := puo.mutation.DepartmentID(); ok {
		if err := position.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Position.department_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PositionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PositionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.SetField(position.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(position.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DepartmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.DepartmentTable,
			Columns: []string{position.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.DepartmentTable,
			Columns: []string{position.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.UserPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUserPositionsIDs(); len(nodes) > 0 && !puo.mutation.UserPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UserPositionsTable,
			Columns: []string{position.UserPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userposition.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}

// SetPosition set the Position
func (pu *PositionUpdate) SetPosition(input *Position, fields ...string) *PositionUpdate {
	m := pu.mutation
	if len(fields) == 0 {
		fields = position.OmitColumns(position.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return pu
}

// SetPositionWithZero set the Position
func (pu *PositionUpdate) SetPositionWithZero(input *Position, fields ...string) *PositionUpdate {
	m := pu.mutation
	if len(fields) == 0 {
		fields = position.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return pu
}

// SetPosition set the Position
func (puo *PositionUpdateOne) SetPosition(input *Position, fields ...string) *PositionUpdateOne {
	m := puo.mutation
	if len(fields) == 0 {
		fields = position.OmitColumns(position.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return puo
}

// SetPositionWithZero set the Position
func (puo *PositionUpdateOne) SetPositionWithZero(input *Position, fields ...string) *PositionUpdateOne {
	m := puo.mutation
	if len(fields) == 0 {
		fields = position.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return puo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (puo *PositionUpdateOne) Omit(fields ...string) *PositionUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	puo.fields = []string(nil)
	for _, col := range position.Columns {
		if _, ok := omits[col]; !ok {
			puo.fields = append(puo.fields, col)
		}
	}
	return puo
}
