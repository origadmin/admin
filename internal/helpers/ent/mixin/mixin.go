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

type IDGenerator interface {
	Comment(key string) IDGenerator
	OptionalFK(name string) ent.Field
	FK(name string) ent.Field
	PK(name string) ent.Field
}

// auditFields defines only the fields and indexes for auditing, without any hooks.
// This allows models to include audit fields without enabling automatic updates.
type auditFields struct {
	mixin.Schema
	CreateField string
	UpdateField string
}

// Fields of the auditFields mixin.
func (m auditFields) Fields() []ent.Field {
	auditCreate := innerID
	auditCreate.Key = m.CreateField
	auditCreate.CommentKey = i18n.Text("create_author.field.comment")
	auditCreate.UseDefault = true
	auditCreate.Optional = true
	auditUpdate := innerID
	auditUpdate.Key = m.UpdateField
	auditUpdate.CommentKey = i18n.Text("update_author.field.comment")
	auditUpdate.UseDefault = true
	auditUpdate.Optional = true
	return []ent.Field{
		auditCreate.ToField(),
		auditUpdate.ToField(),
	}
}

// Indexes of the auditFields mixin.
func (m auditFields) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.CreateField),
		index.Fields(m.UpdateField),
	}
}

// AuditFields returns a mixin that includes only the audit fields (create_author, update_author)
// and their indexes, without any automatic update hooks.
func AuditFields(createField, updateField string) ent.Mixin {
	return auditFields{
		CreateField: createField,
		UpdateField: updateField,
	}
}

// DefaultAuditFields returns a new audit mixin with default field names, without hooks.
func DefaultAuditFields() ent.Mixin {
	return AuditFields("create_author", "update_author")
}

// auditMixin composes auditFields and adds an automatic update hook.
// This provides the full-featured auditing capability.
type auditMixin struct {
	auditFields
}

// AuditHook is a hook that sets the create_author and update_author fields
// by extracting the user ID from the context via contextutil.GetUserID.
func AuditHook(createField, updateField string) ent.Hook {
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
// It uses the AuditHook to automatically set the author fields during create and update operations.
// This hook relies on a user identifier being present in the `context.Context`
// and uses the encapsulated `contextutil.GetUserID` function to retrieve it.
func (m auditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook(m.CreateField, m.UpdateField),
	}
}

// AuditWithHook returns a mixin that includes audit fields and an automatic update hook.
func AuditWithHook(createField, updateField string) ent.Mixin {
	return auditMixin{
		auditFields: auditFields{
			CreateField: createField,
			UpdateField: updateField,
		},
	}
}

// DefaultAuditWithHook returns a new audit mixin with default field names and the update hook.
func DefaultAuditWithHook() ent.Mixin {
	return AuditWithHook("create_author", "update_author")
}

// ManagerSchema schema to include control and time fields.
type ManagerSchema struct {
	mixin.Schema
}

// Fields of the Model.
func (ManagerSchema) Fields() []ent.Field {
	manager := innerID
	manager.Key = "manager_id"
	manager.CommentKey = i18n.Text("manager_id.field.comment")
	manager.Optional = true
	manager.UseDefault = true
	if manager.UseAlias {
		return []ent.Field{
			manager.ToField(),
			field.String("manager_name").
				Comment(i18n.Text("manager_name.field.comment")).
				Default(""),
		}
	}
	return []ent.Field{
		manager.ToField(),
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
		innerID,
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}

	// AuditFieldsModelMixin provides the basic model fields plus audit fields, but without automatic update hooks.
	AuditFieldsModelMixin = []ent.Mixin{
		innerID,
		DefaultAuditFields(),
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}

	// AuditModelMixin provides the full suite: basic fields, audit fields, and automatic update hooks.
	AuditModelMixin = []ent.Mixin{
		innerID,
		DefaultAuditWithHook(),
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}
)
