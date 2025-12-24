/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package repo provides common helpers for repository implementations,
// focusing on abstracting common query patterns like pagination.
package repo

// PaginatingRequest defines the contract for any request that supports pagination.
type PaginatingRequest interface {
	GetPage() int32
	GetPageSize() int32
}

// CountingRequest defines the contract for any request that supports "count-only" mode.
type CountingRequest interface {
	GetOnlyCount() bool
}

// KeywordRequest defines the contract for any request that supports keyword-based search.
type KeywordRequest interface {
	GetKeyword() string
}

// QueryOption holds common query options like pagination and ordering.
// It is intended to be embedded in more specific query option structs.
type QueryOption struct {
	Page      int
	PageSize  int
	OnlyCount bool
	Keyword   string
	OrderBy   []string
}

// OptionFromRequest creates a QueryOption with common details
// extracted from any request that satisfies the PaginatingRequest,
// CountingRequest, or KeywordRequest interfaces.
func OptionFromRequest(req interface{}) QueryOption {
	opt := QueryOption{}

	if r, ok := req.(PaginatingRequest); ok {
		opt.Page = int(r.GetPage())
		opt.PageSize = int(r.GetPageSize())
	}

	if r, ok := req.(CountingRequest); ok {
		opt.OnlyCount = r.GetOnlyCount()
	}

	if r, ok := req.(KeywordRequest); ok {
		opt.Keyword = r.GetKeyword()
	}

	return opt
}

// GetFirstOption safely retrieves the first option from a slice of option pointers.
// If the slice is empty or the first element is nil, it returns a new, non-nil instance of the option type.
func GetFirstOption[T any](opts ...*T) *T {
	if len(opts) > 0 && opts[0] != nil {
		return opts[0]
	}
	return new(T)
}
