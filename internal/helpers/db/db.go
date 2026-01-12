/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common helpers for database operations,
// currently focusing on query construction for Ent.
package db

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ColumnValidator is a function type for validating if a string is a valid column name for an entity.
type ColumnValidator func(string) bool

// processMask is an unexported helper that contains the common logic for processing a FieldMask.
// It normalizes, validates, and filters the paths in the mask, returning a slice of valid column names.
func processMask(mask *fieldmaskpb.FieldMask, validator ColumnValidator, messageType proto.Message) []string {
	if mask == nil {
		return nil
	}
	// Normalize must be called before IsValid.
	mask.Normalize()
	if !mask.IsValid(messageType) {
		return nil
	}

	var cols []string
	for _, path := range mask.GetPaths() {
		if validator(path) {
			cols = append(cols, path)
		}
	}
	return cols
}

// SelectFields parses a ReadMask and returns a slice of column names for selection.
// CRITICAL: It automatically adds the primary key (idField) to the selection if other fields are selected,
// which is essential for ent to hydrate the model correctly.
func SelectFields[Q Selector[S], S any](s Q, mask *fieldmaskpb.FieldMask, validator ColumnValidator,
	idField string,
	messageType proto.Message) S {
	selectCols := processMask(mask, validator, messageType)

	// If any columns are selected, always ensure the ID field is also selected.
	if len(selectCols) > 0 {
		hasID := false
		for _, col := range selectCols {
			if col == idField {
				hasID = true
				break
			}
		}
		if !hasID {
			selectCols = append(selectCols, idField)
		}
	}
	if len(selectCols) > 0 {
		return s.Select(selectCols...)
	}
	return s.Select()
}

// UpdateFields parses an UpdateMask and returns a slice of column names for a partial update.
// It will NOT add the ID field.
func UpdateFields(mask *fieldmaskpb.FieldMask, validator ColumnValidator, messageType proto.Message) []string {
	// The primary key should never be included in an update operation.
	// processMask already ensures this by simply validating columns, and we don't add the idField here.
	return processMask(mask, validator, messageType)
}
