/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin is the mixin package
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/helpers/i18n"
)

// UUID schema to include control and time fields.
type UUID struct {
	mixin.Schema
}

func (obj UUID) Comment(key string) Ider {
	return CommentedUUID{
		CommentKey: key,
	}
}

// Fields of the mixin.
func (obj UUID) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

func (obj UUID) FK(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text("foreign_key:comment")).
		MaxLen(36)
}

func (UUID) PK(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text("primary_key:comment")).
		MaxLen(36).
		Unique().
		Immutable()
}

func (UUID) OP(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text("optional_key:comment")).
		MaxLen(36).
		Optional()
}

// Indexes of the mixin.
func (UUID) Indexes() []ent.Index {
	return []ent.Index{}
}

type CommentedUUID struct {
	CommentKey string
	mixin.Schema
}

func (c CommentedUUID) Comment(key string) Ider {
	return c
}

func (c CommentedUUID) OP(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text(c.CommentKey)).
		MaxLen(36).
		Optional()
}

func (c CommentedUUID) FK(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text(c.CommentKey)).
		MaxLen(36)
}

func (c CommentedUUID) PK(name string) ent.Field {
	return field.String(name).
		Comment(i18n.Text(c.CommentKey)).
		MaxLen(36).
		Unique().
		Immutable()
}
