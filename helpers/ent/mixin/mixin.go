/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin is the mixin package
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// UUID schema to include control and time fields.
type UUID struct {
	mixin.Schema
}

// Fields of the mixin.
func (obj UUID) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

func (obj UUID) FK(name string) ent.Field {
	return field.String(name).
		MaxLen(36)
}

func (UUID) PK(name string) ent.Field {
	return field.String(name).
		MaxLen(36).
		Unique().
		Immutable()
}

func (UUID) OP(name string) ent.Field {
	return field.String(name).
		MaxLen(36).
		Optional()
}

// Indexes of the mixin.
func (UUID) Indexes() []ent.Index {
	return []ent.Index{}
}

type ID struct {
	mixin.Schema
}

// Fields of the mixin.
func (obj ID) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

func (ID) FK(name string) ent.Field {
	return field.Int(name).
		Positive()
}

func (ID) PK(name string) ent.Field {
	return field.Int(name).
		Positive().
		Unique().
		Immutable()
}

func (ID) OP(name string) ent.Field {
	return field.Int(name).
		Positive().
		Optional()
}

// Audit schema to include control and time fields.
type Audit struct {
	mixin.Schema
}

// Fields of the mixin.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.String("create_author").Default(""),
		field.String("update_author").Default(""),
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
			MaxLen(36).
			Optional(),
		field.String("manager_name").
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
		ID{},
		CreateSchema{},
		UpdateSchema{},
	}
	AuditModelMixin = []ent.Mixin{
		ID{},
		Audit{},
		CreateSchema{},
		UpdateSchema{},
	}
)

// SoftDeleteSchema schema to include control and time fields.
type SoftDeleteSchema = DeleteSchema
