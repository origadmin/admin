// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/department"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/departmentrole"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/position"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/role"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (dc *DepartmentCreate) SetCreateTime(t time.Time) *DepartmentCreate {
	dc.mutation.SetCreateTime(t)
	return dc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableCreateTime(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetCreateTime(*t)
	}
	return dc
}

// SetUpdateTime sets the "update_time" field.
func (dc *DepartmentCreate) SetUpdateTime(t time.Time) *DepartmentCreate {
	dc.mutation.SetUpdateTime(t)
	return dc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableUpdateTime(t *time.Time) *DepartmentCreate {
	if t != nil {
		dc.SetUpdateTime(*t)
	}
	return dc
}

// SetKeyword sets the "keyword" field.
func (dc *DepartmentCreate) SetKeyword(s string) *DepartmentCreate {
	dc.mutation.SetKeyword(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DepartmentCreate) SetName(s string) *DepartmentCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableName(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetName(*s)
	}
	return dc
}

// SetTreePath sets the "tree_path" field.
func (dc *DepartmentCreate) SetTreePath(s string) *DepartmentCreate {
	dc.mutation.SetTreePath(s)
	return dc
}

// SetNillableTreePath sets the "tree_path" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableTreePath(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetTreePath(*s)
	}
	return dc
}

// SetSequence sets the "sequence" field.
func (dc *DepartmentCreate) SetSequence(i int) *DepartmentCreate {
	dc.mutation.SetSequence(i)
	return dc
}

// SetStatus sets the "status" field.
func (dc *DepartmentCreate) SetStatus(i int8) *DepartmentCreate {
	dc.mutation.SetStatus(i)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableStatus(i *int8) *DepartmentCreate {
	if i != nil {
		dc.SetStatus(*i)
	}
	return dc
}

// SetLevel sets the "level" field.
func (dc *DepartmentCreate) SetLevel(i int) *DepartmentCreate {
	dc.mutation.SetLevel(i)
	return dc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableLevel(i *int) *DepartmentCreate {
	if i != nil {
		dc.SetLevel(*i)
	}
	return dc
}

// SetDescription sets the "description" field.
func (dc *DepartmentCreate) SetDescription(s string) *DepartmentCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableDescription(s *string) *DepartmentCreate {
	if s != nil {
		dc.SetDescription(*s)
	}
	return dc
}

// SetParentID sets the "parent_id" field.
func (dc *DepartmentCreate) SetParentID(i int64) *DepartmentCreate {
	dc.mutation.SetParentID(i)
	return dc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableParentID(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetParentID(*i)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DepartmentCreate) SetID(i int64) *DepartmentCreate {
	dc.mutation.SetID(i)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DepartmentCreate) SetNillableID(i *int64) *DepartmentCreate {
	if i != nil {
		dc.SetID(*i)
	}
	return dc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (dc *DepartmentCreate) AddUserIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddUserIDs(ids...)
	return dc
}

// AddUsers adds the "users" edges to the User entity.
func (dc *DepartmentCreate) AddUsers(u ...*User) *DepartmentCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// AddPositionIDs adds the "positions" edge to the Position entity by IDs.
func (dc *DepartmentCreate) AddPositionIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddPositionIDs(ids...)
	return dc
}

// AddPositions adds the "positions" edges to the Position entity.
func (dc *DepartmentCreate) AddPositions(p ...*Position) *DepartmentCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddPositionIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (dc *DepartmentCreate) AddRoleIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddRoleIDs(ids...)
	return dc
}

// AddRoles adds the "roles" edges to the Role entity.
func (dc *DepartmentCreate) AddRoles(r ...*Role) *DepartmentCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return dc.AddRoleIDs(ids...)
}

// AddChildIDs adds the "children" edge to the Department entity by IDs.
func (dc *DepartmentCreate) AddChildIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Department entity.
func (dc *DepartmentCreate) AddChildren(d ...*Department) *DepartmentCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// SetParent sets the "parent" edge to the Department entity.
func (dc *DepartmentCreate) SetParent(d *Department) *DepartmentCreate {
	return dc.SetParentID(d.ID)
}

// AddUserDepartmentIDs adds the "user_departments" edge to the UserDepartment entity by IDs.
func (dc *DepartmentCreate) AddUserDepartmentIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddUserDepartmentIDs(ids...)
	return dc
}

// AddUserDepartments adds the "user_departments" edges to the UserDepartment entity.
func (dc *DepartmentCreate) AddUserDepartments(u ...*UserDepartment) *DepartmentCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserDepartmentIDs(ids...)
}

// AddDepartmentRoleIDs adds the "department_roles" edge to the DepartmentRole entity by IDs.
func (dc *DepartmentCreate) AddDepartmentRoleIDs(ids ...int64) *DepartmentCreate {
	dc.mutation.AddDepartmentRoleIDs(ids...)
	return dc
}

// AddDepartmentRoles adds the "department_roles" edges to the DepartmentRole entity.
func (dc *DepartmentCreate) AddDepartmentRoles(d ...*DepartmentRole) *DepartmentCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDepartmentRoleIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DepartmentCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DepartmentCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DepartmentCreate) defaults() {
	if _, ok := dc.mutation.CreateTime(); !ok {
		v := department.DefaultCreateTime()
		dc.mutation.SetCreateTime(v)
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		v := department.DefaultUpdateTime()
		dc.mutation.SetUpdateTime(v)
	}
	if _, ok := dc.mutation.Name(); !ok {
		v := department.DefaultName
		dc.mutation.SetName(v)
	}
	if _, ok := dc.mutation.TreePath(); !ok {
		v := department.DefaultTreePath
		dc.mutation.SetTreePath(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := department.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Level(); !ok {
		v := department.DefaultLevel
		dc.mutation.SetLevel(v)
	}
	if _, ok := dc.mutation.Description(); !ok {
		v := department.DefaultDescription
		dc.mutation.SetDescription(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		v := department.DefaultID()
		dc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DepartmentCreate) check() error {
	if _, ok := dc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Department.create_time"`)}
	}
	if _, ok := dc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Department.update_time"`)}
	}
	if _, ok := dc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Department.keyword"`)}
	}
	if v, ok := dc.mutation.Keyword(); ok {
		if err := department.KeywordValidator(v); err != nil {
			return &ValidationError{Name: "keyword", err: fmt.Errorf(`ent: validator failed for field "Department.keyword": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Department.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := department.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Department.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.TreePath(); !ok {
		return &ValidationError{Name: "tree_path", err: errors.New(`ent: missing required field "Department.tree_path"`)}
	}
	if v, ok := dc.mutation.TreePath(); ok {
		if err := department.TreePathValidator(v); err != nil {
			return &ValidationError{Name: "tree_path", err: fmt.Errorf(`ent: validator failed for field "Department.tree_path": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Sequence(); !ok {
		return &ValidationError{Name: "sequence", err: errors.New(`ent: missing required field "Department.sequence"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Department.status"`)}
	}
	if _, ok := dc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Department.level"`)}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Department.description"`)}
	}
	if v, ok := dc.mutation.Description(); ok {
		if err := department.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Department.description": %w`, err)}
		}
	}
	if v, ok := dc.mutation.ParentID(); ok {
		if err := department.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Department.parent_id": %w`, err)}
		}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := department.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Department.id": %w`, err)}
		}
	}
	return nil
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		_node = &Department{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreateTime(); ok {
		_spec.SetField(department.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dc.mutation.UpdateTime(); ok {
		_spec.SetField(department.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dc.mutation.Keyword(); ok {
		_spec.SetField(department.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.TreePath(); ok {
		_spec.SetField(department.FieldTreePath, field.TypeString, value)
		_node.TreePath = value
	}
	if value, ok := dc.mutation.Sequence(); ok {
		_spec.SetField(department.FieldSequence, field.TypeInt, value)
		_node.Sequence = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(department.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Level(); ok {
		_spec.SetField(department.FieldLevel, field.TypeInt, value)
		_node.Level = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(department.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := dc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: department.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserDepartmentCreate{config: dc.config, mutation: newUserDepartmentMutation(dc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.PositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PositionsTable,
			Columns: []string{department.PositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(position.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   department.RolesTable,
			Columns: department.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &DepartmentRoleCreate{config: dc.config, mutation: newDepartmentRoleMutation(dc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.UserDepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UserDepartmentsTable,
			Columns: []string{department.UserDepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userdepartment.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DepartmentRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.DepartmentRolesTable,
			Columns: []string{department.DepartmentRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(departmentrole.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SetDepartment set the Department
func (dc *DepartmentCreate) SetDepartment(input *Department, fields ...string) *DepartmentCreate {
	m := dc.mutation
	if len(fields) == 0 {
		fields = department.Columns
	}
	_ = m.SetFields(input, fields...)
	return dc
}

// SetDepartmentWithZero set the Department
func (dc *DepartmentCreate) SetDepartmentWithZero(input *Department, fields ...string) *DepartmentCreate {
	m := dc.mutation
	if len(fields) == 0 {
		fields = department.Columns
	}
	_ = m.SetFieldsWithZero(input, fields...)
	return dc
}

// DepartmentCreateBulk is the builder for creating many Department entities in bulk.
type DepartmentCreateBulk struct {
	config
	err      error
	builders []*DepartmentCreate
}

// Save creates the Department entities in the database.
func (dcb *DepartmentCreateBulk) Save(ctx context.Context) ([]*Department, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Department, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DepartmentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) SaveX(ctx context.Context) []*Department {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DepartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DepartmentCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
