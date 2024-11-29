// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/userrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRoleDelete is the builder for deleting a UserRole entity.
type UserRoleDelete struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// Where appends a list predicates to the UserRoleDelete builder.
func (urd *UserRoleDelete) Where(ps ...predicate.UserRole) *UserRoleDelete {
	urd.mutation.Where(ps...)
	return urd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (urd *UserRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, urd.sqlExec, urd.mutation, urd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (urd *UserRoleDelete) ExecX(ctx context.Context) int {
	n, err := urd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (urd *UserRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userrole.Table, sqlgraph.NewFieldSpec(userrole.FieldID, field.TypeString))
	if ps := urd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, urd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	urd.mutation.done = true
	return affected, err
}

// UserRoleDeleteOne is the builder for deleting a single UserRole entity.
type UserRoleDeleteOne struct {
	urd *UserRoleDelete
}

// Where appends a list predicates to the UserRoleDelete builder.
func (urdo *UserRoleDeleteOne) Where(ps ...predicate.UserRole) *UserRoleDeleteOne {
	urdo.urd.mutation.Where(ps...)
	return urdo
}

// Exec executes the deletion query.
func (urdo *UserRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := urdo.urd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (urdo *UserRoleDeleteOne) ExecX(ctx context.Context) {
	if err := urdo.Exec(ctx); err != nil {
		panic(err)
	}
}
