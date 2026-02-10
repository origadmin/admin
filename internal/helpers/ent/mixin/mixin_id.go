/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin is the mixin package
package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"origadmin/application/admin/internal/helpers/i18n"
	"origadmin/application/admin/internal/helpers/idutil"
)

// IDBuilder defines the interface for building ID fields with various configurations.
// It can also provide an ent.Mixin representation of itself.
type IDBuilder interface {
	Comment(text string) IDBuilder
	CommentKey(key string) IDBuilder
	Immutable() IDBuilder
	Optional() IDBuilder
	Unique() IDBuilder

	OptionalFK(name string) ent.Field
	FK(name string) ent.Field
	PK(name string) ent.Field

	// Mixin provider method
	Mixin() ent.Mixin
}

// idMixin implements the IDBuilder interface and ent.Mixin.
// It holds the configuration state for building an ID field.
type idMixin struct {
	mixin.Schema
	key                  string
	commentKey           string
	commentStr           string
	optional             bool
	positive             bool
	unique               bool
	immutable            bool
	useDefault           bool
	defaultFunc          func() int64
	useCustomIDGenerator bool
	useAlias             bool
}

// NewIDBuilder creates a new instance of the IDBuilder.
func NewIDBuilder() IDBuilder {
	return idMixin{}
}

func (obj idMixin) Field() ent.Field {
	builder := field.Int64(obj.key)
	if obj.useDefault {
		builder = builder.Default(0)
	}
	if obj.positive {
		builder = builder.Positive()
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
		obj.useCustomIDGenerator = false
	}
	if !obj.useCustomIDGenerator {
		builder = builder.Annotations(entsql.Annotation{
			Incremental: &obj.useCustomIDGenerator,
		})
	}
	return builder
}

// Fields of the mixin.
func (obj idMixin) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

// Mixin returns the mixin itself, as it satisfies the ent.Mixin interface.
func (obj idMixin) Mixin() ent.Mixin {
	return obj
}

func (obj idMixin) FK(name string) ent.Field {
	obj.key = name
	obj.positive = true
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.foreign_key.comment"
	}
	return obj.Field()
}

func (obj idMixin) PK(name string) ent.Field {
	obj.key = name
	obj.unique = true
	obj.positive = true
	obj.immutable = true
	obj.defaultFunc = idutil.Gen
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.primary_key.comment"
	}
	return obj.Field()
}

func (obj idMixin) OptionalFK(name string) ent.Field {
	obj.key = name
	obj.optional = true
	if obj.commentKey == "" && obj.commentStr == "" {
		obj.commentKey = "field.optional_key.comment"
	}
	return obj.Field()
}

// Comment sets a direct comment string.
func (obj idMixin) Comment(text string) IDBuilder {
	obj.commentStr = text
	return obj
}

// CommentKey sets a key for i18n translation.
func (obj idMixin) CommentKey(key string) IDBuilder {
	obj.commentKey = key
	return obj
}

func (obj idMixin) UserDefaultFunc(f func() int64) IDBuilder {
	obj.defaultFunc = f
	obj.useCustomIDGenerator = true
	return obj
}

// Immutable sets the field as immutable.
func (obj idMixin) Immutable() IDBuilder {
	obj.immutable = true
	return obj
}

// Optional sets the field as optional.
func (obj idMixin) Optional() IDBuilder {
	obj.optional = true
	return obj
}

// Unique sets the field as unique.
func (obj idMixin) Unique() IDBuilder {
	obj.unique = true
	return obj
}
