/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

// PositionNode position.table.comment
type PositionNode struct {
	PositionPB
	DepartmentKeyword string `json:"department_keyword,omitempty"`
}
