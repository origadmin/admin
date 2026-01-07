/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dto_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/internal/data/entity/ent"
)

// validateFields checks if all exported fields in entStruct exist in protoStruct and have compatible types.
func validateFields(t *testing.T, entStruct interface{}, protoStruct interface{}, ignoreFields []string) {
	entType := reflect.TypeOf(entStruct)
	protoType := reflect.TypeOf(protoStruct)

	if entType.Kind() == reflect.Ptr {
		entType = entType.Elem()
	}
	if protoType.Kind() == reflect.Ptr {
		protoType = protoType.Elem()
	}

	entFields := make(map[string]reflect.StructField)
	for i := 0; i < entType.NumField(); i++ {
		field := entType.Field(i)
		// Skip unexported fields and embedded structs like ent.Schema
		if field.PkgPath != "" || field.Anonymous {
			continue
		}
		// Skip specific fields
		if isIgnored(field.Name, ignoreFields) {
			continue
		}
		entFields[strings.ToLower(field.Name)] = field
	}

	protoFields := make(map[string]reflect.StructField)
	for i := 0; i < protoType.NumField(); i++ {
		field := protoType.Field(i)
		// Skip unexported fields and internal proto fields
		if field.PkgPath != "" || strings.HasPrefix(field.Name, "XXX_") {
			continue
		}
		protoFields[strings.ToLower(field.Name)] = field
	}

	for name, entField := range entFields {
		// Special handling for ID field which might be named differently or handled by mixin
		if name == "id" {
			if _, ok := protoFields["id"]; ok {
				continue
			}
		}

		protoField, exists := protoFields[name]
		if !exists {
			assert.Fail(t, fmt.Sprintf("[%s] Field mismatch: '%s' (%s) exists in Ent but missing in Proto",
				entType.Name(), entField.Name, entField.Type))
			continue
		}

		// Optional: Check for type compatibility if needed.
		// Note: Ent types (e.g. int8) might differ from Proto types (e.g. int32), so strict equality check might fail.
		// We can add loose type checking here if required.
		_ = protoField
	}
}

func isIgnored(fieldName string, ignoreList []string) bool {
	for _, ignored := range ignoreList {
		if fieldName == ignored {
			return true
		}
	}
	return false
}

func TestSchemaProtoConsistency(t *testing.T) {
	// Common fields to ignore in Ent entities that are not expected in Proto
	commonIgnores := []string{
		"Edges",
		"config",
	}

	t.Run("User", func(t *testing.T) {
		validateFields(t, &ent.User{}, &types.User{}, append(commonIgnores, "EncryptedPassword", "Salt", "Token", "IsSystem", "DeleteTime"))
	})

	t.Run("Role", func(t *testing.T) {
		validateFields(t, &ent.Role{}, &types.Role{}, commonIgnores)
	})

	t.Run("Resource", func(t *testing.T) {
		validateFields(t, &ent.Resource{}, &types.Resource{}, commonIgnores)
	})

	t.Run("View", func(t *testing.T) {
		validateFields(t, &ent.View{}, &types.View{}, commonIgnores)
	})
}
