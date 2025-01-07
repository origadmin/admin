// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/permissionmenu"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionMenuDelete is the builder for deleting a PermissionMenu entity.
type PermissionMenuDelete struct {
	config
	hooks    []Hook
	mutation *PermissionMenuMutation
}

// Where appends a list predicates to the PermissionMenuDelete builder.
func (pmd *PermissionMenuDelete) Where(ps ...predicate.PermissionMenu) *PermissionMenuDelete {
	pmd.mutation.Where(ps...)
	return pmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pmd *PermissionMenuDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pmd.sqlExec, pmd.mutation, pmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pmd *PermissionMenuDelete) ExecX(ctx context.Context) int {
	n, err := pmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pmd *PermissionMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(permissionmenu.Table, sqlgraph.NewFieldSpec(permissionmenu.FieldID, field.TypeInt64))
	if ps := pmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pmd.mutation.done = true
	return affected, err
}

// PermissionMenuDeleteOne is the builder for deleting a single PermissionMenu entity.
type PermissionMenuDeleteOne struct {
	pmd *PermissionMenuDelete
}

// Where appends a list predicates to the PermissionMenuDelete builder.
func (pmdo *PermissionMenuDeleteOne) Where(ps ...predicate.PermissionMenu) *PermissionMenuDeleteOne {
	pmdo.pmd.mutation.Where(ps...)
	return pmdo
}

// Exec executes the deletion query.
func (pmdo *PermissionMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := pmdo.pmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{permissionmenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pmdo *PermissionMenuDeleteOne) ExecX(ctx context.Context) {
	if err := pmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
