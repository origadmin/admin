/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package mixin implements the functions, types, and interfaces for the module.
package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"

	"origadmin/application/admin/helpers/i18n"
)

// ZeroTime represents the zero value for time.Time.
var ZeroTime = time.Time{}
var _id = ID{}

func Comment(key string) IDGenerator {
	return _id.Comment(key)
}

func I18nComment(key string) IDGenerator {
	return _id.Comment(i18n.Text(key))
}

func PK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return _id.PK(name)
	}
	return _id.Comment(comment[0]).PK(name)
}

func FK(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return _id.FK(name)
	}
	return _id.Comment(comment[0]).FK(name)
}

func OP(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return _id.OP(name)
	}
	return _id.Comment(comment[0]).OP(name)
}

// TimeOP returns a time field with a default value of ZeroTime and a custom schema type for MySQL.
func TimeOP(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return field.Time(name).
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			})
	}
	// Create a time field with the given name and a default value of ZeroTime.
	return field.Time(name).
		Comment(comment[0]).
		Optional().
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		})
}

// Time returns a time field with a default value of ZeroTime and a custom schema type for MySQL.
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
		}).
		// Set the schema type of the field to "datetime" for MySQL dialect.
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
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

// FieldOP returns an optional string field with a maximum length of 36 characters.
func FieldOP(name string) ent.Field {
	// Create an optional string field with the given name and maximum length.
	return ID{}.OP(name)
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

func FieldUUIDOP(name string, comment ...string) ent.Field {
	if len(comment) == 0 {
		return UUID{}.OP(name)
	}
	// Create an optional string field with the given name and maximum length.
	return UUID{}.Comment(comment[0]).OP(name)
}

// FieldTime returns a time field with a default value of ZeroTime and a custom schema type for MySQL.
func FieldTime(name string) ent.Field {
	return field.Time(name).
		// Set the default value of the field to ZeroTime.
		Default(func() time.Time {
			return ZeroTime
		}).
		// Set the schema type of the field to "datetime" for MySQL dialect.
		SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		})
}
