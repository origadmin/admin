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

	"origadmin/application/admin/helpers/i18n"
	"origadmin/application/admin/helpers/id"
)

type ID struct {
	mixin.Schema
	Key                  string
	CommentKey           string
	Optional             bool
	Positive             bool
	Unique               bool
	Immutable            bool
	UseDefault           bool
	DefaultFunc          func() int64
	UseCustomIDGenerator bool
}

func (obj ID) ToField() ent.Field {
	builder := field.Int64(obj.Key)
	if obj.UseDefault {
		builder = builder.Default(0)
	}
	if obj.Positive {
		builder = builder.Positive()
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
		obj.UseCustomIDGenerator = true
	}
	if obj.UseCustomIDGenerator {
		builder = builder.Annotations(entsql.Annotation{
			Incremental: &obj.UseCustomIDGenerator,
		})
	}
	return builder
}

// Fields of the mixin.
func (obj ID) Fields() []ent.Field {
	return []ent.Field{
		obj.PK("id"),
	}
}

func (obj ID) FK(name string) ent.Field {
	obj.Key = name
	obj.Positive = true
	obj.DefaultFunc = id.Gen
	if obj.CommentKey == "" {
		obj.CommentKey = "field:foreign_key:comment"
	}
	return obj.ToField()
}

func (obj ID) PK(name string) ent.Field {
	obj.Key = name
	obj.Unique = true
	obj.Positive = true
	obj.Immutable = true
	obj.DefaultFunc = id.Gen
	if obj.CommentKey == "" {
		obj.CommentKey = "field:primary_key:comment"
	}
	return obj.ToField()
}

func (obj ID) OP(name string) ent.Field {
	obj.Key = name
	obj.Positive = true
	obj.Optional = true
	obj.DefaultFunc = id.Gen
	if obj.CommentKey == "" {
		obj.CommentKey = "field:optional_key:comment"
	}
	return obj.ToField()
}

func (obj ID) Comment(key string) IDGenerator {
	return ID{
		CommentKey: key,
	}
}

func (obj ID) UserDefaultFunc(f func() int64) IDGenerator {
	return ID{
		DefaultFunc:          f,
		UseCustomIDGenerator: true,
	}
}
