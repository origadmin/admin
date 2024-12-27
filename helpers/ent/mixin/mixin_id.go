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
	return field.Int64(name).
		Comment(i18n.Text("foreign_key:comment")).
		Positive()
}

func (ID) PK(name string) ent.Field {
	return field.Int64(name).
		Comment(i18n.Text("primary_key:comment")).
		Positive().
		Unique().
		Immutable()
}

func (ID) OP(name string) ent.Field {
	return field.Int64(name).
		Comment(i18n.Text("optional_key:comment")).
		Positive().
		Optional()
}

func (ID) Comment(key string) Ider {
	return CommentedID{
		CommentKey: key,
	}
}

type CommentedID struct {
	CommentKey string
	mixin.Schema
}

func (c CommentedID) Comment(key string) Ider {
	return c
}

func (c CommentedID) OP(name string) ent.Field {
	return field.Int64(name).
		Comment(i18n.Text(c.CommentKey)).
		Positive().
		Optional()
}

func (c CommentedID) FK(name string) ent.Field {
	return field.Int64(name).
		Comment(i18n.Text(c.CommentKey)).
		Positive()
}

func (c CommentedID) PK(name string) ent.Field {
	return field.Int64(name).
		Comment(i18n.Text(c.CommentKey)).
		Positive().
		Unique().
		Immutable()
}
