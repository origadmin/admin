/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and interfaces for the module.
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/internal/helpers/i18n"
)

// ZeroTime represents the zero value for time.Time.
var ZeroTime = time.Time{}
var innerID = ID{}

func Comment(key string) IDGenerator {
	return innerID.Comment(key)
}

func I18nComment(key string) IDGenerator {
	return innerID.Comment(i18n.Text(key))
}

func PK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return innerID.PK(name)
	}
	return innerID.Comment(comment[0]).PK(name)
}

func FK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return innerID.FK(name)
	}
	return innerID.Comment(comment[0]).FK(name)
}

func OptionalFK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return innerID.OptionalFK(name)
	}
	return innerID.Comment(comment[0]).OptionalFK(name)
}

// TimeOptional returns a time field with a default value of ZeroTime.
func TimeOptional(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return field.Time(name).Optional()
	}
	// Create a time field with the given name and a default value of ZeroTime.
	return field.Time(name).
		Comment(comment[0]).
		Optional()
}

// Time returns a time field with a default value of ZeroTime.
func Time(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return FieldTime(name)
	}
	// Create a time field with the given name and a default value of ZeroTime.
	return field.Time(name).
		Comment(comment[0]).
		// Set the default value of the field to ZeroTime.
		Default(func() time.Time {
			return ZeroTime
		})
}

// FieldIndex returns a field with index
func FieldIndex(name string) ent.Field {
	return field.Int(name).Unique()

}
func FieldPK(name string) ent.Field {
	return ID{}.PK(name)
}

func FieldFK(name string) ent.Field {
	return ID{}.FK(name)
}

// FieldOptional returns an optional string field with a maximum length of 36 characters.
func FieldOptional(name string) ent.Field {
	// Create an optional string field with the given name and maximum length.
	return ID{}.OptionalFK(name)
}

func FieldUUIDPK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return UUID{}.PK(name)
	}
	// Create an optional string field with the given name and maximum length.
	return UUID{}.Comment(comment[0]).PK(name)
}

func FieldUUIDFK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return UUID{}.FK(name)
	}
	// Create an optional string field with the given name and maximum length.
	return UUID{}.Comment(comment[0]).FK(name)
}

func FieldUUIDOptional(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return UUID{}.OptionalFK(name)
	}
	// Create an optional string field with the given name and maximum length.
	return UUID{}.Comment(comment[0]).OptionalFK(name)
}

// FieldTime returns a time field with a default value of ZeroTime.
func FieldTime(name string) ent.Field {
	return field.Time(name).
		// Set the default value of the field to ZeroTime.
		Default(func() time.Time {
			return ZeroTime
		})
}
