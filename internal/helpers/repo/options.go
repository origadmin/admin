/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package repo provides common helpers for repository implementations,
// focusing on abstracting common query patterns like pagination.
package repo

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// DefaultPageSize is the page size used when the client does not specify one.
	DefaultPageSize = 10
	// MaxPageSize is the maximum page size allowed for normal pagination.
	MaxPageSize = 100
	// HardLimit is the absolute maximum number of records that can be returned in a single query,
	// typically used when NoPaging is requested.
	HardLimit = 1000
)

// PaginatingRequest defines the contract for any request that supports offset-based pagination.
type PaginatingRequest interface {
	GetPage() int32
	GetPageSize() int32
}

// TokenPaginatingRequest defines the contract for any request that supports token-based pagination.
type TokenPaginatingRequest interface {
	GetPageToken() string
}

// NoPagingRequest defines the contract for any request that can disable pagination.
type NoPagingRequest interface {
	GetNoPaging() bool
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
	PageToken string
	NoPaging  bool
	OnlyCount bool
	Keyword   string
	OrderBy   []string
	ReadMask  *fieldmaskpb.FieldMask // Use FieldMask for field selection
}

// OptionFromRequest creates a QueryOption with common details
// extracted from any request that satisfies the supported interfaces.
func OptionFromRequest(req interface{}) QueryOption {
	opt := QueryOption{}

	if r, ok := req.(PaginatingRequest); ok {
		opt.Page = int(r.GetPage())
		opt.PageSize = int(r.GetPageSize())
	}

	if r, ok := req.(TokenPaginatingRequest); ok {
		opt.PageToken = r.GetPageToken()
	}

	if r, ok := req.(NoPagingRequest); ok {
		opt.NoPaging = r.GetNoPaging()
	}

	if r, ok := req.(CountingRequest); ok {
		opt.OnlyCount = r.GetOnlyCount()
	}

	if r, ok := req.(KeywordRequest); ok {
		opt.Keyword = r.GetKeyword()
	}

	// This is a generic helper. The ReadMask should be populated from the specific
	// request type in the service layer, as the field name (`read_mask`) can vary.
	// Example in service layer:
	// if r, ok := req.(interface{ GetReadMask() *fieldmaskpb.FieldMask }); ok {
	// 	opt.ReadMask = r.GetReadMask()
	// }

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
