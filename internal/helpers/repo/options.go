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

// PagingModeRequest defines the contract for any request that can specify a pagination mode.
type PagingModeRequest interface {
	GetPagingMode() string
}

// CountingRequest defines the contract for any request that supports "count-only" mode.
type CountingRequest interface {
	GetOnlyCount() bool
}

// KeywordRequest defines the contract for any request that supports keyword-based search.
type KeywordRequest interface {
	GetKeyword() string
}

// SortingRequest defines the contract for any request that supports sorting
type SortingRequest interface {
	GetSorting() []string
}

// QueryOption holds common query options like pagination and ordering.
// It is intended to be embedded in more specific query option structs.
type QueryOption struct {
	Page          int
	PageSize      int
	PageToken     string
	PagingMode    string
	OnlyCount     bool
	Keyword       string
	OrderBy       []string
	ReadMask      *fieldmaskpb.FieldMask // Use FieldMask for field selection
	SortFromToken bool
}

// QueryOptionFromRequest creates a QueryOption with common details
// extracted from any request that satisfies the supported contracts.
func QueryOptionFromRequest(req interface{}) QueryOption {
	opt := QueryOption{}

	if r, ok := req.(PaginatingRequest); ok {
		opt.Page = int(r.GetPage())
		opt.PageSize = int(r.GetPageSize())
	}

	if r, ok := req.(TokenPaginatingRequest); ok {
		opt.PageToken = r.GetPageToken()
	}

	if r, ok := req.(PagingModeRequest); ok {
		opt.PagingMode = r.GetPagingMode()
	}

	if r, ok := req.(CountingRequest); ok {
		opt.OnlyCount = r.GetOnlyCount()
	}

	if r, ok := req.(KeywordRequest); ok {
		opt.Keyword = r.GetKeyword()
	}

	if r, ok := req.(SortingRequest); ok {
		opt.OrderBy = r.GetSorting()
	}

	if r, ok := req.(interface{ GetReadMask() *fieldmaskpb.FieldMask }); ok {
		opt.ReadMask = r.GetReadMask()
	}

	return opt
}

type UpdateOption struct {
	UpdateMask *fieldmaskpb.FieldMask
}

func UpdateOptionFromRequest(req interface{}) UpdateOption {
	opt := UpdateOption{}
	if r, ok := req.(interface{ GetUpdateMask() *fieldmaskpb.FieldMask }); ok {
		opt.UpdateMask = r.GetUpdateMask()
	}
	return opt
}

// FirstOrDefault safely retrieves the first non-nil option from a slice of option pointers.
// If the slice is empty or all its elements are nil, it returns a new, non-nil, zero-value instance of the option type.
// This prevents panics from nil pointer dereferences.
func FirstOrDefault[T any](opts ...*T) *T {
	for _, opt := range opts {
		if opt != nil {
			return opt
		}
	}
	return new(T)
}
