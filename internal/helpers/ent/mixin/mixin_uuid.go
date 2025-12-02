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
	Key         string
	MaxLen      int
	CommentKey  string
	Optional    bool
	Positive    bool
	Unique      bool
	Immutable   bool
	UseDefault  bool
	DefaultFunc func() string
}

func (obj UUID) ToField() ent.Field {
	builder := field.String(obj.Key)
	if obj.UseDefault {
		builder = builder.Default("")
	}
	if obj.Unique {
		builder = builder.Unique()
	}
	if obj.Immutable {
		builder = builder.Immutable()
	}
	if obj.Optional {
		builder = builder.Optional()
	}
	if obj.CommentKey != "" {
		builder = builder.Comment(i18n.Text(obj.CommentKey))
	}
	if obj.DefaultFunc != nil {
		builder = builder.DefaultFunc(obj.DefaultFunc)
		// string will not be incremented by the database.
	}
	if obj.MaxLen > 0 {
		builder = builder.MaxLen(obj.MaxLen)
	} else {
		builder = builder.MaxLen(36)
	}
	return builder
}
func (obj UUID) Comment(key string) IDGenerator {
	return UUID{
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
	obj.Key = name
	if obj.CommentKey == "" {
		obj.CommentKey = "field.foreign_key.comment"
	}
	return obj.ToField()
}

func (obj UUID) PK(name string) ent.Field {
	obj.Key = name
	obj.Unique = true
	obj.Immutable = true
	if obj.CommentKey == "" {
		obj.CommentKey = "field.primary_key.comment"
	}
	return obj.ToField()
}

func (obj UUID) OP(name string) ent.Field {
	obj.Key = name
	obj.Optional = true
	if obj.CommentKey == "" {
		obj.CommentKey = "field.optional_key.comment"
	}
	return obj.ToField()
}

func (obj UUID) UserDefaultFunc(f func() string) IDGenerator {
	return UUID{
		DefaultFunc: f,
	}
}
