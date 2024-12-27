// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/departmentrole"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentRoleDelete is the builder for deleting a DepartmentRole entity.
type DepartmentRoleDelete struct {
	config
	hooks    []Hook
	mutation *DepartmentRoleMutation
}

// Where appends a list predicates to the DepartmentRoleDelete builder.
func (drd *DepartmentRoleDelete) Where(ps ...predicate.DepartmentRole) *DepartmentRoleDelete {
	drd.mutation.Where(ps...)
	return drd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (drd *DepartmentRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, drd.sqlExec, drd.mutation, drd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (drd *DepartmentRoleDelete) ExecX(ctx context.Context) int {
	n, err := drd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (drd *DepartmentRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(departmentrole.Table, sqlgraph.NewFieldSpec(departmentrole.FieldID, field.TypeInt))
	if ps := drd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, drd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	drd.mutation.done = true
	return affected, err
}

// DepartmentRoleDeleteOne is the builder for deleting a single DepartmentRole entity.
type DepartmentRoleDeleteOne struct {
	drd *DepartmentRoleDelete
}

// Where appends a list predicates to the DepartmentRoleDelete builder.
func (drdo *DepartmentRoleDeleteOne) Where(ps ...predicate.DepartmentRole) *DepartmentRoleDeleteOne {
	drdo.drd.mutation.Where(ps...)
	return drdo
}

// Exec executes the deletion query.
func (drdo *DepartmentRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := drdo.drd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{departmentrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (drdo *DepartmentRoleDeleteOne) ExecX(ctx context.Context) {
	if err := drdo.Exec(ctx); err != nil {
		panic(err)
	}
}
