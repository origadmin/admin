/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package schema implements the functions, types, and interfaces for the module.
package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"origadmin/application/admin/helpers/ent/mixin"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/hook"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/intercept"
)

// SoftDelete is schema to include control and time fields.
type SoftDelete mixin.DeleteSchema

// Interceptors of the SoftDeleteMixin.
func (s SoftDelete) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if mixin.IsSkipSoftDelete(ctx) {
				return nil
			}
			s.P(q)
			return nil
		}),
	}
}

func (s SoftDelete) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				// Skip soft-delete, means delete the entity permanently.
				if mixin.IsSkipSoftDelete(ctx) {
					return next.Mutate(ctx, m)
				}
				mx, ok := m.(interface {
					SetOp(ent.Op)
					SetDeleteTime(time.Time)
				})
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				mx.SetOp(ent.OpUpdate)
				mx.SetDeleteTime(time.Now())
				return next.Mutate(ctx, m)
			})
		},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (s SoftDelete) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(s.Fields()[0].Descriptor().Name),
	)
}
