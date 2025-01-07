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

	"origadmin/application/admin/helpers/i18n"
)

type IDGenerator interface {
	Comment(key string) IDGenerator
	OP(name string) ent.Field
	FK(name string) ent.Field
	PK(name string) ent.Field
}

// Audit schema to include control and time fields.
type Audit struct {
	mixin.Schema
}

// Fields of the mixin.
func (Audit) Fields() []ent.Field {
	auditCreate := _id
	auditCreate.Key = "create_author"
	auditCreate.CommentKey = i18n.Text("create_author.field.comment")
	auditCreate.UseDefault = true
	auditCreate.Optional = true
	auditUpdate := _id
	auditUpdate.Key = "update_author"
	auditUpdate.CommentKey = i18n.Text("update_author.field.comment")
	auditUpdate.UseDefault = true
	auditUpdate.Optional = true
	return []ent.Field{
		auditCreate.ToField(),
		auditUpdate.ToField(),
	}
}

// Indexes of the mixin.
func (Audit) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_author"),
		index.Fields("update_author"),
	}
}

// ManagerSchema schema to include control and time fields.
type ManagerSchema struct {
	mixin.Schema
}

// Fields of the Model.
func (ManagerSchema) Fields() []ent.Field {
	manager := _id
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

// CreateUpdateSchema schema to include control and time fields.
type CreateUpdateSchema struct {
	mixin.Schema
}

// Fields of the mixin.
func (CreateUpdateSchema) Fields() []ent.Field {
	return append(
		CreateSchema{}.Fields(),
		UpdateSchema{}.Fields()...,
	)
}

// Indexes of the mixin.
func (CreateUpdateSchema) Indexes() []ent.Index {
	return append(
		CreateSchema{}.Indexes(),
		UpdateSchema{}.Indexes()...,
	)
}

// CreateSchema schema to include control and time fields.
type CreateSchema struct {
	mixin.Schema
}

// Fields of the mixin.
func (CreateSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Comment(i18n.Text("create_time.field.comment")).
			Default(time.Now).
			Immutable(),
	}
}

// Indexes of the mixin.
func (CreateSchema) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
	}
}

// UpdateSchema schema to include control and time fields.
type UpdateSchema struct {
	mixin.Schema
}

// Fields of the mixin.
func (UpdateSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Comment(i18n.Text("update_time.field.comment")).
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Indexes of the mixin.
func (UpdateSchema) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("update_time"),
	}
}

// DeleteSchema schema to include control and time fields.
type DeleteSchema struct {
	mixin.Schema
}

// Fields of the Model.
func (DeleteSchema) Fields() []ent.Field {
	return []ent.Field{
		field.Time("delete_time").
			Comment(i18n.Text("delete_time.field.comment")).
			Optional().
			Nillable(),
	}
}

// Indexes of the mixin.
func (DeleteSchema) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("delete_time"),
	}
}

var (
	ModelMixin = []ent.Mixin{
		_id,
		CreateSchema{},
		UpdateSchema{},
	}
	AuditModelMixin = []ent.Mixin{
		_id,
		Audit{},
		CreateSchema{},
		UpdateSchema{},
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
