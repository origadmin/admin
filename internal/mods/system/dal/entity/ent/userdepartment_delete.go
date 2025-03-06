// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userdepartment"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserDepartmentDelete is the builder for deleting a UserDepartment entity.
type UserDepartmentDelete struct {
	config
	hooks    []Hook
	mutation *UserDepartmentMutation
}

// Where appends a list predicates to the UserDepartmentDelete builder.
func (udd *UserDepartmentDelete) Where(ps ...predicate.UserDepartment) *UserDepartmentDelete {
	udd.mutation.Where(ps...)
	return udd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (udd *UserDepartmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, udd.sqlExec, udd.mutation, udd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (udd *UserDepartmentDelete) ExecX(ctx context.Context) int {
	n, err := udd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (udd *UserDepartmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userdepartment.Table, sqlgraph.NewFieldSpec(userdepartment.FieldID, field.TypeInt))
	if ps := udd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, udd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	udd.mutation.done = true
	return affected, err
}

// UserDepartmentDeleteOne is the builder for deleting a single UserDepartment entity.
type UserDepartmentDeleteOne struct {
	udd *UserDepartmentDelete
}

// Where appends a list predicates to the UserDepartmentDelete builder.
func (uddo *UserDepartmentDeleteOne) Where(ps ...predicate.UserDepartment) *UserDepartmentDeleteOne {
	uddo.udd.mutation.Where(ps...)
	return uddo
}

// Exec executes the deletion query.
func (uddo *UserDepartmentDeleteOne) Exec(ctx context.Context) error {
	n, err := uddo.udd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userdepartment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uddo *UserDepartmentDeleteOne) ExecX(ctx context.Context) {
	if err := uddo.Exec(ctx); err != nil {
		panic(err)
	}
}
