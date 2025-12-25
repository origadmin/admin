package dto

import "time"

// PermissionCondition represents a single condition for a permission.
type PermissionCondition struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

// PermissionAccessControl defines the access control rules for a permission.
type PermissionAccessControl struct {
	Actions    []string          `json:"actions"`
	Conditions map[string]string `json:"conditions"`
	ValidFrom  *time.Time        `json:"valid_from"`
	ValidUntil *time.Time        `json:"valid_until"`
	Attributes map[string]any    `json:"attributes"`
}
