/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and contracts for the module.
package builder

import (
	"regexp"

	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Int64Builder defines the interface for building int64 fields with various configurations.
type Int64Builder interface {
	Unique() Int64Builder
	Range(i, j int64) Int64Builder
	Min(i int64) Int64Builder
	Max(i int64) Int64Builder
	Positive() Int64Builder
	Negative() Int64Builder
	NonNegative() Int64Builder
	Default(i int64) Int64Builder
	DefaultFunc(fn any) Int64Builder
	UpdateDefault(fn any) Int64Builder
	Nillable() Int64Builder
	Comment(c string) Int64Builder
	Optional() Int64Builder
	Immutable() Int64Builder
	StructTag(s string) Int64Builder
	Validate(fn func(int64) error) Int64Builder
	StorageKey(key string) Int64Builder
	SchemaType(types map[string]string) Int64Builder
	GoType(typ any) Int64Builder
	ValueScanner(vs any) Int64Builder
	Annotations(annotations ...schema.Annotation) Int64Builder
	Deprecated(reason ...string) Int64Builder
	Descriptor() *field.Descriptor
}

// StringBuilder defines the interface for building string fields with various configurations.
type StringBuilder interface {
	Unique() StringBuilder
	Sensitive() StringBuilder
	Match(re *regexp.Regexp) StringBuilder
	MinLen(i int) StringBuilder
	MinRuneLen(i int) StringBuilder
	NotEmpty() StringBuilder
	MaxLen(i int) StringBuilder
	MaxRuneLen(i int) StringBuilder
	Validate(fn func(string) error) StringBuilder
	Default(s string) StringBuilder
	DefaultFunc(fn any) StringBuilder
	Nillable() StringBuilder
	Optional() StringBuilder
	Immutable() StringBuilder
	Comment(c string) StringBuilder
	StructTag(s string) StringBuilder
	StorageKey(key string) StringBuilder
	SchemaType(types map[string]string) StringBuilder
	GoType(typ any) StringBuilder
	ValueScanner(vs any) StringBuilder
	Annotations(annotations ...schema.Annotation) StringBuilder
	Deprecated(reason ...string) StringBuilder
	Descriptor() *field.Descriptor
}
