/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements reusable schema components for Ent.
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

const (
	// Default field names for mixins.
	DefaultCreateAuthorField = "create_author"
	DefaultUpdateAuthorField = "update_author"
	DefaultOwnerField        = "owner_id"
	DefaultManagerField      = "manager_id"
	DefaultCreateTimeField   = "create_time"
	DefaultUpdateTimeField   = "update_time"
)

// FieldMixin defines an interface for mixins that wrap a single field.
// It extends ent.Mixin with methods to access the underlying field and index.
type FieldMixin interface {
	ent.Mixin
	Field() ent.Field
	Index() ent.Index
}

// fieldMixin is a generic unexported mixin for a single field.
type fieldMixin struct {
	mixin.Schema
	fieldName string
}

// dualFieldMixin is a generic unexported mixin for two fields.
type dualFieldMixin struct {
	mixin.Schema
	createField string
	updateField string
}

// auditFields defines the fields and indexes for auditing.
type auditFields struct {
	dualFieldMixin
}

func (m auditFields) Fields() []ent.Field {
	return []ent.Field{
		innerID.CommentKey("create_author.field.comment").Immutable().OptionalFK(m.createField),
		innerID.CommentKey("update_author.field.comment").OptionalFK(m.updateField),
	}
}

func (m auditFields) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.createField),
		index.Fields(m.updateField),
	}
}

// AuditMixin returns a composite mixin for audit fields.
func AuditMixin(createField, updateField string) ent.Mixin {
	return auditFields{
		dualFieldMixin: dualFieldMixin{createField: createField, updateField: updateField},
	}
}

func DefaultAuditMixin() ent.Mixin {
	return AuditMixin(DefaultCreateAuthorField, DefaultUpdateAuthorField)
}

// auditMixin composes auditFields and adds an automatic update hook.
type auditMixin struct{ auditFields }

func (m auditMixin) Hooks() []ent.Hook {
	return []ent.Hook{AuditMixinHook(m.createField, m.updateField)}
}

func AuditMixinHook(createField, updateField string) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if !m.Op().Is(ent.OpCreate | ent.OpUpdate) {
				return next.Mutate(ctx, m)
			}
			userID, err := contextutil.GetUserID(ctx)
			if err != nil {
				if errors.Is(err, contextutil.ErrNoPrincipalInContext) {
					return next.Mutate(ctx, m)
				}
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

func AuditMixinWithHook(createField, updateField string) ent.Mixin {
	return auditMixin{
		auditFields: auditFields{
			dualFieldMixin: dualFieldMixin{createField: createField, updateField: updateField},
		},
	}
}

func DefaultAuditMixinWithHook() ent.Mixin {
	return AuditMixinWithHook(DefaultCreateAuthorField, DefaultUpdateAuthorField)
}

// ownerMixin defines the owner_id field.
type ownerMixin struct{ fieldMixin }

func (m ownerMixin) Field() ent.Field {
	return innerID.CommentKey("owner_id.field.comment").Immutable().OptionalFK(m.fieldName)
}
func (m ownerMixin) Index() ent.Index     { return index.Fields(m.fieldName) }
func (m ownerMixin) Fields() []ent.Field  { return []ent.Field{m.Field()} }
func (m ownerMixin) Indexes() []ent.Index { return []ent.Index{m.Index()} }

func OwnerMixin(ownerField string) FieldMixin {
	return ownerMixin{fieldMixin: fieldMixin{fieldName: ownerField}}
}

func DefaultOwnerMixin() FieldMixin {
	return OwnerMixin(DefaultOwnerField)
}

// managerMixin defines the manager_id field.
type managerMixin struct{ fieldMixin }

func (m managerMixin) Field() ent.Field {
	return innerID.CommentKey("manager_id.field.comment").OptionalFK(m.fieldName)
}
func (m managerMixin) Index() ent.Index     { return index.Fields(m.fieldName) }
func (m managerMixin) Fields() []ent.Field  { return []ent.Field{m.Field()} }
func (m managerMixin) Indexes() []ent.Index { return []ent.Index{m.Index()} }

// createUpdateMixin is a composite mixin for create and update timestamps.
type createUpdateMixin struct{ dualFieldMixin }

func CreateUpdateMixin(updateField, createField string) ent.Mixin {
	return createUpdateMixin{
		dualFieldMixin: dualFieldMixin{createField: createField, updateField: updateField},
	}
}

func DefaultCreateUpdateMixin() ent.Mixin {
	return CreateUpdateMixin(DefaultUpdateTimeField, DefaultCreateTimeField)
}

func (m createUpdateMixin) Fields() []ent.Field {
	return []ent.Field{
		CreateMixin(m.createField).Field(),
		UpdateMixin(m.updateField).Field(),
	}
}

func (m createUpdateMixin) Indexes() []ent.Index {
	return []ent.Index{
		CreateMixin(m.createField).Index(),
		UpdateMixin(m.updateField).Index(),
	}
}

// createMixin defines the create_time field.
type createMixin struct{ fieldMixin }

func CreateMixin(fieldName string) FieldMixin {
	return createMixin{fieldMixin: fieldMixin{fieldName: fieldName}}
}

func DefaultCreateMixin() FieldMixin {
	return CreateMixin(DefaultCreateTimeField)
}

func (m createMixin) Field() ent.Field {
	return field.Time(m.fieldName).
		Comment(i18n.Text("create_time.field.comment")).
		Default(time.Now).
		Immutable()
}
func (m createMixin) Index() ent.Index     { return index.Fields(m.fieldName) }
func (m createMixin) Fields() []ent.Field  { return []ent.Field{m.Field()} }
func (m createMixin) Indexes() []ent.Index { return []ent.Index{m.Index()} }

// updateMixin defines the update_time field.
type updateMixin struct{ fieldMixin }

func UpdateMixin(fieldName string) FieldMixin {
	return updateMixin{fieldMixin: fieldMixin{fieldName: fieldName}}
}

func DefaultUpdateMixin() FieldMixin {
	return UpdateMixin(DefaultUpdateTimeField)
}

func (m updateMixin) Field() ent.Field {
	return field.Time(m.fieldName).
		Comment(i18n.Text("update_time.field.comment")).
		Default(time.Now).
		UpdateDefault(time.Now)
}
func (m updateMixin) Index() ent.Index     { return index.Fields(m.fieldName) }
func (m updateMixin) Fields() []ent.Field  { return []ent.Field{m.Field()} }
func (m updateMixin) Indexes() []ent.Index { return []ent.Index{m.Index()} }

var (
	// ModelMixin provides a basic set of fields for standard models.
	ModelMixin = []ent.Mixin{
		innerID.Mixin(),
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
