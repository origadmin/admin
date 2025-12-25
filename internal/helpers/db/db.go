/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common helpers for database operations,
// currently focusing on query construction for Ent.
package db

import (
	"entgo.io/ent/dialect/sql"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// selectable is a generic constraint for types that can be used as order functions in Ent.
// It's an alias for any type whose underlying type is func(*sql.Selector).
type selectable interface {
	~func(*sql.Selector)
}

// ColumnValidator is a function type for validating if a string is a valid column name for an entity.
type ColumnValidator func(string) bool

func SelectFields(mask *fieldmaskpb.FieldMask, validator ColumnValidator, messageType proto.Message) []string {
	var selectCols []string
	if mask != nil {
		mask.Normalize()
		if mask.IsValid(messageType) {
			for _, path := range mask.GetPaths() {
				if validator(path) {
					selectCols = append(selectCols, path)
				}
			}
		}
	}
	return selectCols
}

func UpdateFields(mask *fieldmaskpb.FieldMask, validator ColumnValidator, messageType proto.Message) []string {
	var updateCols []string
	if mask == nil || !mask.IsValid(messageType) {
		return updateCols
	}
	for _, path := range mask.GetPaths() {
		if validator(path) {
			updateCols = append(updateCols, path)
		}
	}
	return updateCols
}
