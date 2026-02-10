package mixin

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	gen "origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/hook"
	"origadmin/application/admin/internal/data/entity/ent/intercept"
	"origadmin/application/admin/internal/helpers/i18n"
)

// SoftDeleteField holds the name of the soft delete field.
const SoftDeleteField = "delete_time"

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// IsSkipSoftDelete checks if the context is configured to skip soft-delete.
func IsSkipSoftDelete(ctx context.Context) bool {
	v, _ := ctx.Value(softDeleteKey{}).(bool)
	return v
}

// SoftDeleteMixin provides soft-delete capabilities to a schema.
// It adds a `delete_time` field and a hook to intercept delete operations.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(SoftDeleteField).
			Comment(i18n.Text("delete_time.field.comment")).
			Optional().
			Nillable(),
	}
}

// Hooks of the SoftDeleteMixin.
// This hook intercepts delete operations and converts them to update operations
// that set the `delete_time` field.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		SoftDeleteHook(d),
	}
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		SoftDeleteInterceptor(d),
	}
}

// SoftDeleteHook intercepts DELETE operations and converts them to UPDATEs.
func SoftDeleteHook(d SoftDeleteMixin) ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				// Skip soft-delete, means delete the entity permanently.
				if IsSkipSoftDelete(ctx) {
					return next.Mutate(ctx, m)
				}
				mx, ok := m.(interface {
					SetOp(ent.Op)
					Client() *gen.Client
					SetDeleteTime(time.Time)
					WhereP(...func(*sql.Selector))
				})
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				d.P(mx)
				mx.SetOp(ent.OpUpdate)
				mx.SetDeleteTime(time.Now())
				return mx.Client().Mutate(ctx, m)
			})
		},
		ent.OpDeleteOne|ent.OpDelete,
	)
}

type P interface {
	WhereP(...func(*sql.Selector))
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w P) {
	w.WhereP(
		sql.FieldIsNull(SoftDeleteField),
	)
}

// SoftDeleteInterceptor returns a query interceptor that filters out soft-deleted records.
func SoftDeleteInterceptor(d SoftDeleteMixin) ent.Interceptor {
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		if IsSkipSoftDelete(ctx) {
			return nil
		}
		d.P(q)
		return nil
	})
}
