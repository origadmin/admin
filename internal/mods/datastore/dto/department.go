/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

type DepartmentNode struct {
	DepartmentPB
	Children         []*DepartmentNode `json:"children"`
	PositionKeywords []string          `json:"position_keywords"`
}
