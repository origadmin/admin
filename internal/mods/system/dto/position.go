/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

// PositionNode position.table.comment
type PositionNode struct {
	PositionPB
	//Id                int64  `json:"id,omitempty"`
	//CreateTime        int64  `json:"create_time,omitempty"`
	//UpdateTime        int64  `json:"update_time,omitempty"`
	//Name              string `json:"name,omitempty"`
	//Keyword           string `json:"keyword,omitempty"`
	//Description       string `json:"description,omitempty"`
	//DepartmentId      int64  `json:"department_id,omitempty"`
	DepartmentKeyword string `json:"department_keyword,omitempty"`
}
