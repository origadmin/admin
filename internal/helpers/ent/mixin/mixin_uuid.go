/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin is the mixin package
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/internal/helpers/i18n"
)

// uuidMixin implements the IDBuilder interface for UUID (string) fields.
type uuidMixin struct {
	mixin.Schema
	key         string
	maxLen      int
	commentKey  string
	commentStr  string
	optional    bool
	positive    bool
	unique      bool
	immutable   bool
	useDefault  bool
	defaultFunc func() string
}

// NewUUIDBuilder creates a new instance of the IDBuilder backed by UUID logic.
func NewUUIDBuilder() IDBuilder {
	return uuidMixin{}
}

func (obj uuidMixin) Field() ent.Field {
	builder := field.String(obj.key)
	if obj.useDefault {
		builder = builder.Default("")
	}
	if obj.unique {
		builder = builder.Unique()
	}
	if obj.immutable {
		builder = builder.Immutable()
	}
	if obj.optional {
		builder = builder.Optional()
	}
	// Prioritize direct comment string over i18n key
	if obj.commentStr != "" {
		builder = builder.Comment(obj.commentStr)
	} else if obj.commentKey != "" {
		builder = builder.Comment(i18n.Text(obj.commentKey))
	}

	if obj.defaultFunc != nil {
		builder = builder.DefaultFunc(obj.defaultFunc)
		// string will not be incremented by the database.
	}
	if obj.maxLen > 0 {
		builder = builder.MaxLen(obj.maxLen)
	} else {
		builder = builder.MaxLen(36)
	}
	return builder
}

// Comment sets a direct comment string.
func (obj uuidMixin) Comment(text string) IDBuilder {
	obj.commentStr = text
	return obj
}

// CommentKey sets a key for i18n translation.
func (obj uuidMixin) CommentKey(key string) IDBuilder {
	obj.commentKey = key
	return obj
}

// Fields of the mixin.
func (obj uuidMixin) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

// Mixin returns the mixin itself, as it satisfies the ent.Mixin interface.
func (obj uuidMixin) Mixin() ent.Mixin {
	return obj
}

func (obj uuidMixin) FK(name string) ent.Field {
	obj.key = name
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.foreign_key.comment"
	}
	return obj.Field()
}

func (obj uuidMixin) PK(name string) ent.Field {
	obj.key = name
	obj.unique = true
	obj.immutable = true
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.primary_key.comment"
	}
	return obj.Field()
}

func (obj uuidMixin) OptionalFK(name string) ent.Field {
	obj.key = name
	obj.optional = true
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.optional_key.comment"
	}
	return obj.Field()
}

// Immutable sets the field as immutable.
func (obj uuidMixin) Immutable() IDBuilder {
	obj.immutable = true
	return obj
}

// Optional sets the field as optional.
func (obj uuidMixin) Optional() IDBuilder {
	obj.optional = true
	return obj
}

// Unique sets the field as unique.
func (obj uuidMixin) Unique() IDBuilder {
	obj.unique = true
	return obj
}

// UserDefaultFunc sets a custom default function.
// Note: This is specific to uuidMixin and not part of the IDBuilder interface.
func (obj uuidMixin) UserDefaultFunc(f func() string) uuidMixin {
	obj.defaultFunc = f
	return obj
}
