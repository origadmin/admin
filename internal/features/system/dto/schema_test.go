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
		if isIgnored(field.Name, ignoreFields) {
			continue
		}
		protoFields[strings.ToLower(field.Name)] = field
	}

	// Check 1: Ent -> Proto
	for name, entField := range entFields {
		// Special handling for ID field which might be named differently or handled by mixin
		if name == "id" {
			if _, ok := protoFields["id"]; ok {
				continue
			}
		}

		_, exists := protoFields[name]
		if !exists {
			assert.Fail(t, fmt.Sprintf("[%s] Field mismatch: '%s' (%s) exists in Ent but missing in Proto",
				entType.Name(), entField.Name, entField.Type))
		}
	}

	// Check 2: Proto -> Ent
	for name, protoField := range protoFields {
		// Special handling for ID field
		if name == "id" {
			if _, ok := entFields["id"]; ok {
				continue
			}
		}

		_, exists := entFields[name]
		if !exists {
			assert.Fail(t, fmt.Sprintf("[%s] Field mismatch: '%s' (%s) exists in Proto but missing in Ent",
				protoType.Name(), protoField.Name, protoField.Type))
		}
	}
}

func isIgnored(fieldName string, ignoreList []string) bool {
	for _, ignored := range ignoreList {
		if strings.EqualFold(fieldName, ignored) {
			return true
		}
	}
	return false
}

func TestSchemaProtoConsistency(t *testing.T) {
	// Common fields to ignore in both Ent and Proto
	commonIgnores := []string{
		"Edges",
		"config",
		"DeleteTime",
		"XXX_NoUnkeyedLiteral",
		"XXX_unrecognized",
		"XXX_sizecache",
	}

	t.Run("User", func(t *testing.T) {
		specificIgnores := append(commonIgnores, "EncryptedPassword", "Salt", "Token", "IsSystem", "RoleIds", "Roles")
		validateFields(t, &ent.User{}, &types.User{}, specificIgnores)
	})

	t.Run("Role", func(t *testing.T) {
		specificIgnores := append(commonIgnores, "views", "users", "resources", "ResourceIds", "permissions",
			"PermissionIds")
		validateFields(t, &ent.Role{}, &types.Role{}, specificIgnores)
	})

	t.Run("Resource", func(t *testing.T) {
		specificIgnores := append(commonIgnores, "VersionID", "LastSyncVersionID", "children", "parent", "PermissionIds", "permissions")
		validateFields(t, &ent.Resource{}, &types.Resource{}, specificIgnores)
	})

	t.Run("View", func(t *testing.T) {
		specificIgnores := append(commonIgnores, "children", "parent", "resources", "roles")
		validateFields(t, &ent.View{}, &types.View{}, specificIgnores)
	})
}
