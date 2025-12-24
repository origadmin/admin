/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and interfaces for the module.
package mixin

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/internal/helpers/i18n"
)

type IDGenerator interface {
	Comment(key string) IDGenerator
	OptionalFK(name string) ent.Field
	FK(name string) ent.Field
	PK(name string) ent.Field
}

// Audit schema to include control and time fields.
type auditMixin struct {
	mixin.Schema
	CreateField string
	UpdateField string
}

func DefaultAudit() ent.Mixin {
	return auditMixin{
		CreateField: "create_author",
		UpdateField: "update_author",
	}
}

// Audit returns a new audit mixin with configurable field names.
func Audit(createField, updateField string) ent.Mixin {
	return auditMixin{
		CreateField: createField,
		UpdateField: updateField,
	}
}

// Fields of the mixin.
func (m auditMixin) Fields() []ent.Field {
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

// Indexes of the mixin.
func (m auditMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(m.CreateField),
		index.Fields(m.UpdateField),
	}
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
	return []ent.Field{
		manager.ToField(),
		field.String("manager_name").
			Comment(i18n.Text("manager_name.field.comment")).
			Default(""),
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

// DeleteMixin schema to include control and time fields.
type DeleteMixin struct {
	mixin.Schema
}

// Fields of the Model.
func (m DeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Comment(i18n.Text("delete_time.field.comment")).
			Optional().
			Nillable(),
	}
}

// Indexes of the mixin.
func (m DeleteMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("delete_time"),
	}
}

var (
	ModelMixin = []ent.Mixin{
		innerID,
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}
	AuditModelMixin = []ent.Mixin{
		innerID,
		DefaultAudit(),
		DefaultCreateMixin(),
		DefaultUpdateMixin(),
	}
)

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

func IsSkipSoftDelete(ctx context.Context) bool {
	v, _ := ctx.Value(softDeleteKey{}).(bool)
	return v
}
