// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionresource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionResourceDelete is the builder for deleting a PermissionResource entity.
type PermissionResourceDelete struct {
	config
	hooks    []Hook
	mutation *PermissionResourceMutation
}

// Where appends a list predicates to the PermissionResourceDelete builder.
func (prd *PermissionResourceDelete) Where(ps ...predicate.PermissionResource) *PermissionResourceDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *PermissionResourceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, prd.sqlExec, prd.mutation, prd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *PermissionResourceDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *PermissionResourceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(permissionresource.Table, sqlgraph.NewFieldSpec(permissionresource.FieldID, field.TypeInt64))
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prd.mutation.done = true
	return affected, err
}

// PermissionResourceDeleteOne is the builder for deleting a single PermissionResource entity.
type PermissionResourceDeleteOne struct {
	prd *PermissionResourceDelete
}

// Where appends a list predicates to the PermissionResourceDelete builder.
func (prdo *PermissionResourceDeleteOne) Where(ps ...predicate.PermissionResource) *PermissionResourceDeleteOne {
	prdo.prd.mutation.Where(ps...)
	return prdo
}

// Exec executes the deletion query.
func (prdo *PermissionResourceDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{permissionresource.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *PermissionResourceDeleteOne) ExecX(ctx context.Context) {
	if err := prdo.Exec(ctx); err != nil {
		panic(err)
	}
}
