/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and contracts for the module.
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ZeroTime represents the zero value for time.Time.
var ZeroTime = time.Time{}

// innerID is the default ID builder instance.
// It is initialized using the NewIDBuilder factory function.
var innerID = NewIDBuilder()

// Comment sets a direct comment string.
func Comment(text string) IDBuilder {
	return innerID.Comment(text)
}

// CommentKey sets a key for i18n translation.
// Replaces the old I18nComment function logic but keeps the intent clearer.
func CommentKey(key string) IDBuilder {
	return innerID.CommentKey(key)
}

// PK sets the field as a primary key.
func PK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return innerID.PK(name)
	}
	// Assuming shortcut functions provide direct comments.
	// If i18n key is needed, use NewIDBuilder().CommentKey(...).PK(...)
	return innerID.Comment(comment[0]).PK(name)
}

// FK sets the field as a foreign key.
func FK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return innerID.FK(name)
	}
	return innerID.Comment(comment[0]).FK(name)
}

// OptionalFK sets the field as an optional foreign key.
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

// FieldPK returns an ID field with a maximum length of 36 characters.
func FieldPK(name string) ent.Field {
	return NewIDBuilder().PK(name)
}

// FieldFK returns an ID field with a maximum length of 36 characters.
func FieldFK(name string) ent.Field {
	return NewIDBuilder().FK(name)
}

// FieldOptional returns an optional string field with a maximum length of 36 characters.
func FieldOptional(name string) ent.Field {
	// Create an optional string field with the given name and maximum length.
	return NewIDBuilder().OptionalFK(name)
}

// FieldUUIDPK returns an UUID field with a maximum length of 36 characters.
func FieldUUIDPK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return NewUUIDBuilder().PK(name)
	}
	return NewUUIDBuilder().Comment(comment[0]).PK(name)
}

// FieldUUIDFK returns an UUID field with a maximum length of 36 characters.
func FieldUUIDFK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return NewUUIDBuilder().FK(name)
	}
	return NewUUIDBuilder().Comment(comment[0]).FK(name)
}

// FieldUUIDOptional returns an optional UUID field with a maximum length of 36 characters.
func FieldUUIDOptional(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return NewUUIDBuilder().OptionalFK(name)
	}
	return NewUUIDBuilder().Comment(comment[0]).OptionalFK(name)
}

// FieldTime returns a time field with a default value of ZeroTime.
func FieldTime(name string) ent.Field {
	return field.Time(name).
		// Set the default value of the field to ZeroTime.
		Default(func() time.Time {
			return ZeroTime
		})
}
