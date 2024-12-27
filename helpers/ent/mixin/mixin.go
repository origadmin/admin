/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and interfaces for the module.
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/helpers/i18n"
)

type Ider interface {
	Comment(key string) Ider
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
	return []ent.Field{
		field.String("create_author").
			Comment(i18n.Text("create_author:comment")).
			Default(""),
		field.String("update_author").
			Comment(i18n.Text("update_author:comment")).
			Default(""),
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
	return []ent.Field{
		field.String("manager_id").
			Comment(i18n.Text("manager_id:comment")).
			MaxLen(36).
			Optional(),
		field.String("manager_name").
			Comment(i18n.Text("manager_name:comment")).
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
			Comment(i18n.Text("create_time:comment")).
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
			Comment(i18n.Text("update_time:comment")).
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
			Comment(i18n.Text("delete_time:comment")).
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

// SoftDeleteSchema schema to include control and time fields.
type SoftDeleteSchema = DeleteSchema
