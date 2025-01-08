/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema"
)

const (
	ResourceTypeRoot     = schema.ResourceTypeRoot
	ResourceTypeGroup    = schema.ResourceTypeGroup
	ResourceTypeMenu     = schema.ResourceTypeMenu
	ResourceTypePage     = schema.ResourceTypePage
	ResourceTypeButton   = schema.ResourceTypeButton
	ResourceTypeAPI      = schema.ResourceTypeAPI
	ResourceTypeRedirect = schema.ResourceTypeRedirect
	ResourceTypeUnknown  = schema.ResourceTypeUnknown
)

// ResourceTypeName returns the name of the resource type
func ResourceTypeName(str string) string {
	switch str {
	case ResourceTypeMenu:
		return "Menu"
	case ResourceTypePage:
		return "Page"
	case ResourceTypeButton:
		return "Button"
	case ResourceTypeAPI:
		return "API"
	case ResourceTypeRedirect:
		return "Redirect"
	case ResourceTypeRoot:
		return "ROOT"
	case ResourceTypeGroup:
		return "Group"
	default:
		return "Unknown"
	}
}

// ResourceTypeCode returns the code of the resource type
func ResourceTypeCode(s string) string {
	switch s {
	case "Menu":
		return ResourceTypeMenu
	case "Page":
		return ResourceTypePage
	case "Button":
		return ResourceTypeButton
	case "API":
		return ResourceTypeAPI
	case "Redirect":
		return ResourceTypeRedirect
	case "ROOT":
		return ResourceTypeRoot
	case "Group":
		return ResourceTypeGroup
	default:
		return ResourceTypeUnknown
	}
}
