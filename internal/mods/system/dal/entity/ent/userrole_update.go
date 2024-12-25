// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRoleUpdate is the builder for updating UserRole entities.
type UserRoleUpdate struct {
	config
	hooks     []Hook
	mutation  *UserRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserRoleUpdate builder.
func (uru *UserRoleUpdate) Where(ps ...predicate.UserRole) *UserRoleUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetUserID sets the "user_id" field.
func (uru *UserRoleUpdate) SetUserID(s string) *UserRoleUpdate {
	uru.mutation.SetUserID(s)
	return uru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uru *UserRoleUpdate) SetNillableUserID(s *string) *UserRoleUpdate {
	if s != nil {
		uru.SetUserID(*s)
	}
	return uru
}

// SetRoleID sets the "role_id" field.
func (uru *UserRoleUpdate) SetRoleID(s string) *UserRoleUpdate {
	uru.mutation.SetRoleID(s)
	return uru
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uru *UserRoleUpdate) SetNillableRoleID(s *string) *UserRoleUpdate {
	if s != nil {
		uru.SetRoleID(*s)
	}
	return uru
}

// SetUser sets the "user" edge to the User entity.
func (uru *UserRoleUpdate) SetUser(u *User) *UserRoleUpdate {
	return uru.SetUserID(u.ID)
}

// SetRole sets the "role" edge to the Role entity.
func (uru *UserRoleUpdate) SetRole(r *Role) *UserRoleUpdate {
	return uru.SetRoleID(r.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uru *UserRoleUpdate) Mutation() *UserRoleMutation {
	return uru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uru *UserRoleUpdate) ClearUser() *UserRoleUpdate {
	uru.mutation.ClearUser()
	return uru
}

// ClearRole clears the "role" edge to the Role entity.
func (uru *UserRoleUpdate) ClearRole() *UserRoleUpdate {
	uru.mutation.ClearRole()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uru.sqlSave, uru.mutation, uru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRoleUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRoleUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uru *UserRoleUpdate) check() error {
	if v, ok := uru.mutation.UserID(); ok {
		if err := userrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.user_id": %w`, err)}
		}
	}
	if v, ok := uru.mutation.RoleID(); ok {
		if err := userrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.role_id": %w`, err)}
		}
	}
	if uru.mutation.UserCleared() && len(uru.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserRole.user"`)
	}
	if uru.mutation.RoleCleared() && len(uru.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserRole.role"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uru *UserRoleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRoleUpdate {
	uru.modifiers = append(uru.modifiers, modifiers...)
	return uru
}

func (uru *UserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userrole.Table, userrole.Columns, sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeString))
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uru.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
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
	_spec.AddModifiers(uru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uru.mutation.done = true
	return n, nil
}

// UserRoleUpdateOne is the builder for updating a single UserRole entity.
type UserRoleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserRoleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUserID sets the "user_id" field.
func (uruo *UserRoleUpdateOne) SetUserID(s string) *UserRoleUpdateOne {
	uruo.mutation.SetUserID(s)
	return uruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uruo *UserRoleUpdateOne) SetNillableUserID(s *string) *UserRoleUpdateOne {
	if s != nil {
		uruo.SetUserID(*s)
	}
	return uruo
}

// SetRoleID sets the "role_id" field.
func (uruo *UserRoleUpdateOne) SetRoleID(s string) *UserRoleUpdateOne {
	uruo.mutation.SetRoleID(s)
	return uruo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (uruo *UserRoleUpdateOne) SetNillableRoleID(s *string) *UserRoleUpdateOne {
	if s != nil {
		uruo.SetRoleID(*s)
	}
	return uruo
}

// SetUser sets the "user" edge to the User entity.
func (uruo *UserRoleUpdateOne) SetUser(u *User) *UserRoleUpdateOne {
	return uruo.SetUserID(u.ID)
}

// SetRole sets the "role" edge to the Role entity.
func (uruo *UserRoleUpdateOne) SetRole(r *Role) *UserRoleUpdateOne {
	return uruo.SetRoleID(r.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uruo *UserRoleUpdateOne) Mutation() *UserRoleMutation {
	return uruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uruo *UserRoleUpdateOne) ClearUser() *UserRoleUpdateOne {
	uruo.mutation.ClearUser()
	return uruo
}

// ClearRole clears the "role" edge to the Role entity.
func (uruo *UserRoleUpdateOne) ClearRole() *UserRoleUpdateOne {
	uruo.mutation.ClearRole()
	return uruo
}

// Where appends a list predicates to the UserRoleUpdate builder.
func (uruo *UserRoleUpdateOne) Where(ps ...predicate.UserRole) *UserRoleUpdateOne {
	uruo.mutation.Where(ps...)
	return uruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserRoleUpdateOne) Select(field string, fields ...string) *UserRoleUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserRole entity.
func (uruo *UserRoleUpdateOne) Save(ctx context.Context) (*UserRole, error) {
	return withHooks(ctx, uruo.sqlSave, uruo.mutation, uruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) SaveX(ctx context.Context) *UserRole {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uruo *UserRoleUpdateOne) check() error {
	if v, ok := uruo.mutation.UserID(); ok {
		if err := userrole.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.user_id": %w`, err)}
		}
	}
	if v, ok := uruo.mutation.RoleID(); ok {
		if err := userrole.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "UserRole.role_id": %w`, err)}
		}
	}
	if uruo.mutation.UserCleared() && len(uruo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserRole.user"`)
	}
	if uruo.mutation.RoleCleared() && len(uruo.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserRole.role"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uruo *UserRoleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRoleUpdateOne {
	uruo.modifiers = append(uruo.modifiers, modifiers...)
	return uruo
}

func (uruo *UserRoleUpdateOne) sqlSave(ctx context.Context) (_node *UserRole, err error) {
	if err := uruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userrole.Table, userrole.Columns, sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeString))
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrole.FieldID)
		for _, f := range fields {
			if !userrole.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if uruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uruo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
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
	_spec.AddModifiers(uruo.modifiers...)
	_node = &UserRole{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uruo.mutation.done = true
	return _node, nil
}

// SetUserRole set the UserRole
func (uru *UserRoleUpdate) SetUserRole(input *UserRole, fields ...string) *UserRoleUpdate {
	m := uru.mutation
	if len(fields) == 0 {
		fields = userrole.OmitColumns(userrole.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return uru
}

// SetUserRoleWithZero set the UserRole
func (uru *UserRoleUpdate) SetUserRoleWithZero(input *UserRole, fields ...string) *UserRoleUpdate {
	m := uru.mutation
	if len(fields) == 0 {
		fields = userrole.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return uru
}

// SetUserRole set the UserRole
func (uruo *UserRoleUpdateOne) SetUserRole(input *UserRole, fields ...string) *UserRoleUpdateOne {
	m := uruo.mutation
	if len(fields) == 0 {
		fields = userrole.OmitColumns(userrole.FieldID)
	}
	_ = m.SetFields(input, fields...)
	return uruo
}

// SetUserRoleWithZero set the UserRole
func (uruo *UserRoleUpdateOne) SetUserRoleWithZero(input *UserRole, fields ...string) *UserRoleUpdateOne {
	m := uruo.mutation
	if len(fields) == 0 {
		fields = userrole.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return uruo
}

// Omit allows the unselect one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (uruo *UserRoleUpdateOne) Omit(fields ...string) *UserRoleUpdateOne {
	omits := make(map[string]struct{}, len(fields))
	for i := range fields {
		omits[fields[i]] = struct{}{}
	}
	uruo.fields = []string(nil)
	for _, col := range userrole.Columns {
		if _, ok := omits[col]; !ok {
			uruo.fields = append(uruo.fields, col)
		}
	}
	return uruo
}
