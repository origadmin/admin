/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package repo provides common query options for pagination.
package repo

// QueryOption holds common query options like pagination and ordering.
// It is intended to be embedded in more specific query option structs.
type QueryOption struct {
	Page     int
	PageSize int
	OrderBy  []string
}

// IsOption is a marker method to ensure type safety.
func (o *QueryOption) IsOption() {}
