package mixin

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// SoftDeleteMixin implements the soft-delete pattern for a schema.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Comment("Time of soft-delete").
			Optional().
			Nillable(),
	}
}

// Hooks of the SoftDeleteMixin.
func (SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		softDeleteHook(),
	}
}

// Interceptors of the SoftDeleteMixin.
func (SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		softDeleteInterceptor(),
	}
}

// softDeleteHook intercepts DELETE operations and converts them to UPDATEs.
func softDeleteHook() ent.Hook {
	// Define an interface for mutations that support soft-delete.
	// This relies on structural typing and code generation.
	type softDeleter interface {
		SetOp(ent.Op)
		SetDeleteTime(time.Time)
	}

	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Skip if not a DELETE operation or if soft-delete is skipped.
			if !m.Op().Is(ent.OpDelete|ent.OpDeleteOne) || IsSkipSoftDelete(ctx) {
				return next.Mutate(ctx, m)
			}

			// Check if the mutation implements the softDeleter interface.
			mx, ok := m.(softDeleter)
			if !ok {
				return nil, errors.New("ent: mutation does not support soft-delete")
			}

			// Change the operation to UPDATE and set the delete_time.
			mx.SetOp(ent.OpUpdate)
			mx.SetDeleteTime(time.Now())

			// Proceed with the mutation, which is now an update.
			return next.Mutate(ctx, m)
		})
	}
}

// softDeleteInterceptor filters out soft-deleted records from queries.
func softDeleteInterceptor() ent.Interceptor {
	// Define an interface for queries that support WhereP.
	type queryWither interface {
		WhereP(...func(*sql.Selector))
	}

	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			// Skip if soft-delete is skipped for this query.
			if IsSkipSoftDelete(ctx) {
				return next.Query(ctx, query)
			}

			// Check if the query supports the WhereP method.
			q, ok := query.(queryWither)
			if !ok {
				return next.Query(ctx, query)
			}

			// Add the WHERE clause to filter out soft-deleted records.
			q.WhereP(func(s *sql.Selector) {
				s.Where(sql.IsNull(s.C("delete_time")))
			})

			return next.Query(ctx, query)
		})
	})
}
