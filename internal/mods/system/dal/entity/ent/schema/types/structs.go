/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package types implements the functions, types, and interfaces for the module.
package types

import (
	"time"
)

type PermissionCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type PermissionAccessControl struct {
	Actions    []string          `json:"actions"`
	Conditions map[string]string `json:"conditions"`
	ValidFrom  *time.Time        `json:"valid_from"`
	ValidUntil *time.Time        `json:"valid_until"`
	Attributes map[string]any    `json:"attributes"`
}
