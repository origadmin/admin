/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and contracts for the module.
package dto

import (
	"strings"

	"origadmin/application/admin/internal/data/entity/ent/view"
)

const (
	ViewTypeRoot     = view.TypeT
	ViewTypeGroup    = view.TypeG
	ViewTypeMenu     = view.TypeM
	ViewTypeLink     = view.TypeL
	ViewTypePage     = view.TypeP
	ViewTypeButton   = view.TypeB
	ViewTypeElement  = view.TypeE
	ViewTypeRedirect = view.TypeR
	ViewTypeUnknown  = view.TypeU
)

type ViewType = view.Type

// ViewTypeName returns the name of the resource type
func ViewTypeName(str ViewType) string {
	switch str {
	case ViewTypeMenu:
		return "Menu"
	case ViewTypePage:
		return "Page"
	case ViewTypeButton:
		return "Button"
	case ViewTypeElement:
		return "Element"
	case ViewTypeRedirect:
		return "Redirect"
	case ViewTypeRoot:
		return "Root"
	case ViewTypeGroup:
		return "Group"
	case ViewTypeLink:
		return "Link"
	default:
		return "Unknown"
	}
}

// ViewTypeCode returns the code of the resource type
func ViewTypeCode(s string) ViewType {
	switch strings.ToLower(s) {
	case "menu":
		return ViewTypeMenu
	case "page":
		return ViewTypePage
	case "button":
		return ViewTypeButton
	case "redirect":
		return ViewTypeRedirect
	case "root":
		return ViewTypeRoot
	case "group":
		return ViewTypeGroup
	case "link":
		return ViewTypeLink
	case "element":
		return ViewTypeElement
	default:
		return ViewTypeUnknown
	}
}
