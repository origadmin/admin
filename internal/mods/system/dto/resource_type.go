/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/schema"
)

const (
	ResourceTypeUnknown   = schema.ResourceTypeUnknown
	ResourceTypeMenu      = schema.ResourceTypeMenu
	ResourceTypePage      = schema.ResourceTypePage
	ResourceTypeButton    = schema.ResourceTypeButton
	ResourceTypeAPI       = schema.ResourceTypeAPI
	ResourceTypeButtonAPI = schema.ResourceTypeButtonAPI
	ResourceTypeRedirect  = schema.ResourceTypeRedirect
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
	case ResourceTypeButtonAPI:
		return "ButtonAPI"
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
	case "ButtonAPI":
		return ResourceTypeButtonAPI
	default:
		return ResourceTypeUnknown
	}
}
