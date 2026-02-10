/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and interfaces for the module.
package mixin

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/internal/helpers/contextutil"
	"origadmin/application/admin/internal/helpers/i18n"
)

// auditFields defines only the fields and indexes for auditing, without any hooks.
// This allows models to include audit fields without enabling automatic updates.
type auditFields struct {
	mixin.Schema
	CreateField string
	UpdateField string
}

// Fields of the auditFields mixin.
func (m auditFields) Fields() []ent.Field {
	// Use the innerID builder to construct fields with specific properties like Immutable.
	// innerID.CommentKey(...) returns a new IDBuilder instance, so it's safe to chain.
	return []ent.Field{
		innerID.CommentKey("create_author.field.comment").Immutable().OptionalFK(m.CreateField),
		innerID.CommentKey("update_author.field.comment").OptionalFK(m.UpdateField),
	}
}

// Indexes of the auditFields mixin.
func (m auditFields) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.CreateField),
		index.Fields(m.UpdateField),
	}
}

// AuditMixin returns a mixin that includes only the audit fields (create_author, update_author)
// and their indexes, without any automatic update hooks.
func AuditMixin(createField, updateField string) ent.Mixin {
	return auditFields{
		CreateField: createField,
		UpdateField: updateField,
	}
}

// DefaultAuditMixin returns a new audit mixin with default field names, without hooks.
func DefaultAuditMixin() ent.Mixin {
	return AuditMixin("create_author", "update_author")
}

// auditMixin composes auditFields and adds an automatic update hook.
// This provides the full-featured auditing capability.
type auditMixin struct {
	auditFields
}

// AuditMixinHook is a hook that sets the create_author and update_author fields
// by extracting the user ID from the context via contextutil.GetUserID.
func AuditMixinHook(createField, updateField string) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			// Skip if not a Create or Update operation.
			if !m.Op().Is(ent.OpCreate | ent.OpUpdate) {
				return next.Mutate(ctx, m)
			}

			userID, err := contextutil.GetUserID(ctx)
			if err != nil {
				// If the error indicates no principal was found, proceed without setting audit fields.
				// This allows for anonymous or system-level actions.
				if errors.Is(err, contextutil.ErrNoPrincipalInContext) {
					return next.Mutate(ctx, m)
				}
				// For other errors (e.g., parsing), fail the mutation to prevent data corruption.
				return nil, err
			}

			if m.Op().Is(ent.OpCreate) {
				if err := m.SetField(createField, userID); err != nil {
					return nil, err
				}
			}
			if err := m.SetField(updateField, userID); err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}

// Hooks of the mixin.
// It uses the AuditMixinHook to automatically set the author fields during create and update operations.
// This hook relies on a user identifier being present in the `context.Context`
// and uses the encapsulated `contextutil.GetUserID` function to retrieve it.
func (m auditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditMixinHook(m.CreateField, m.UpdateField),
	}
}

// AuditMixinWithHook returns a mixin that includes audit fields and an automatic update hook.
func AuditMixinWithHook(createField, updateField string) ent.Mixin {
	return auditMixin{
		auditFields: auditFields{
			CreateField: createField,
			UpdateField: updateField,
		},
	}
}

// DefaultAuditMixinWithHook returns a new audit mixin with default field names and the update hook.
func DefaultAuditMixinWithHook() ent.Mixin {
	return AuditMixinWithHook("create_author", "update_author")
}

// ManagerSchema schema to include control and time fields.
type ManagerSchema struct {
	mixin.Schema
}

// Fields of the Model.
func (ManagerSchema) Fields() []ent.Field {
	return []ent.Field{
		innerID.CommentKey("manager_id.field.comment").OptionalFK("manager_id"),
	}
}

// Indexes of the mixin.
func (ManagerSchema) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("manager_id"),
	}
}

// createUpdateMixin schema to include control and time fields.
type createUpdateMixin struct {
	mixin.Schema
	UpdateField string
	CreateField string
}

func DefaultCreateUpdateMixin() ent.Mixin {
	return createUpdateMixin{
		UpdateField: "update_time",
		CreateField: "create_time",
	}
}

func CreateUpdateMixin(updateField, createField string) ent.Mixin {
	return createUpdateMixin{
		UpdateField: updateField,
		CreateField: createField,
	}
}

// Fields of the mixin.
func (m createUpdateMixin) Fields() []ent.Field {
	return append(
		CreateMixin(m.CreateField).Fields(),
		UpdateMixin(m.UpdateField).Fields()...,
	)
}

// Indexes of the mixin.
func (m createUpdateMixin) Indexes() []ent.Index {
	return append(
		CreateMixin(m.CreateField).Indexes(),
		UpdateMixin(m.UpdateField).Indexes()...,
	)
}

// createMixin schema to include control and time fields.
type createMixin struct {
	mixin.Schema
	CreateField string
}

func DefaultCreateMixin() ent.Mixin {
	return createMixin{
		CreateField: "create_time",
	}
}

func CreateMixin(fieldName string) ent.Mixin {
	return createMixin{
		CreateField: fieldName,
	}
}

// Fields of the mixin.
func (m createMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(m.CreateField).
			Comment(i18n.Text("create_time.field.comment")).
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the mixin.
func (m createMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.CreateField),
	}
}

// updateMixin schema to include control and time fields.
type updateMixin struct {
	mixin.Schema
	UpdateField string
}

func DefaultUpdateMixin() ent.Mixin {
	return updateMixin{
		UpdateField: "update_time",
	}
}
func UpdateMixin(fieldName string) ent.Mixin {
	return updateMixin{
		UpdateField: fieldName,
	}
}

// Fields of the mixin.
func (m updateMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(m.UpdateField).
			Comment(i18n.Text("update_time.field.comment")).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Indexes of the mixin.
func (m updateMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.UpdateField),
	}
}

var (
	// ModelMixin provides a basic set of fields for standard models.
	ModelMixin = []ent.Mixin{
		innerID.Mixin(), // Use the Mixin() method to get the ent.Mixin
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}

	// AuditModelMixin provides the basic model fields plus audit fields, but without automatic update hooks.
	AuditModelMixin = []ent.Mixin{
		innerID.Mixin(),
		DefaultAuditMixin(),
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}

	// AuditHookModelMixin provides the full suite: basic fields, audit fields, and automatic update hooks.
	AuditHookModelMixin = []ent.Mixin{
		innerID.Mixin(),
		DefaultAuditMixinWithHook(),
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}
)
