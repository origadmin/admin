/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db provides common, generic helpers for database operations.
// Its scope is strictly limited to functionalities that are truly generic
// and do not depend on schema-specific types like predicates.
package db

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"

	"origadmin/application/admin/internal/helpers/repo"
)

const (
	// PagingModeCursor indicates cursor-based pagination (for infinite scrolling).
	PagingModeCursor = "cursor"
	// PagingModeOffset indicates traditional offset-based pagination.
	PagingModeOffset = "offset"
	// PagingModeNone indicates no pagination, returning all available records up to a hard limit.
	PagingModeNone = "none"
)

// Identifiable defines the contract for any entity that has a retrievable ID.
type Identifiable interface {
	GetId() int64
}

// GenerateNextPageToken creates a pagination token for cursor-based pagination.
// It accepts the raw request interface and extracts necessary options internally.
func GenerateNextPageToken[T Identifiable](results []T, req interface{}) (string, error) {
	opt := repo.QueryOptionFromRequest(req)

	pagingMode := opt.PagingMode
	if pagingMode == "" {
		pagingMode = PagingModeOffset // Default to offset for backward compatibility
	}

	if pagingMode != PagingModeCursor {
		return "", nil
	}

	if len(results) > 0 && len(results) == opt.PageSize {
		last := results[len(results)-1]

		// Default sort order if not provided
		sortField := "id"
		sortDesc := true
		if len(opt.OrderBy) > 0 {
			parts := strings.Fields(opt.OrderBy[0])
			sortField = parts[0]
			if len(parts) > 1 && strings.ToLower(parts[1]) == "asc" {
				sortDesc = false
			}
		}

		cursor := Cursor{ID: last.GetId(), Field: sortField, Desc: sortDesc}
		token, err := EncodeCursor(cursor)
		if err != nil {
			return "", fmt.Errorf("failed to encode next page token: %w", err)
		}
		return token, nil
	}

	return "", nil
}

type Cursor struct {
	ID    int64  `json:"id"`
	Field string `json:"field"`
	Desc  bool   `json:"desc"`
}

func EncodeCursor(c Cursor) (string, error) {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(c); err != nil {
		return "", fmt.Errorf("gob encode cursor: %w", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func DecodeCursor(token string) (Cursor, error) {
	var c Cursor
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return c, fmt.Errorf("base64 decode token: %w", err)
	}

	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&c); err != nil {
		return c, fmt.Errorf("gob decode cursor: %w", err)
	}
	return c, nil
}

// Pageable defines the interface for queries that can be paginated.
type Pageable[T any, W any, O any, R any] interface {
	Counter[T]
	Cloner[T]
	Filterer[T, W]
	Orderer[T, O]
	Fetcher[R]
	Limit(int) T
	Offset(int) T
}

type Selector interface {
	~func(*sql.Selector)
}

// Orderer defines the interface for queries that can be ordered.
type Orderer[T any, O any] interface {
	Order(...O) T
}

// Filterer defines the interface for queries that can be filtered.
type Filterer[T any, W any] interface {
	Where(...W) T
}

// Cloner defines the interface for objects that can clone themselves.
type Cloner[T any] interface {
	Clone() T
}

// Counter defines the interface for queries that can count their results.
type Counter[T any] interface {
	Cloner[T]
	Count(ctx context.Context) (int, error)
}

// Fetcher defines the interface for queries that can execute and fetch results.
type Fetcher[T any] interface {
	All(ctx context.Context) ([]T, error)
	Only(ctx context.Context) (T, error)
}

// normalizeQueryOption handles unified initialization and validation of QueryOption
func normalizeQueryOption(opt *repo.QueryOption) *repo.QueryOption {
	if opt == nil {
		return &repo.QueryOption{
			Page:     1,
			PageSize: repo.DefaultPageSize,
		}
	}

	// Normalize page number
	if opt.Page <= 0 {
		opt.Page = 1
	}

	// Normalize page size
	if opt.PagingMode == PagingModeNone {
		if opt.PageSize <= 0 || opt.PageSize > repo.HardLimit {
			opt.PageSize = repo.HardLimit
		}
	} else {
		if opt.PageSize <= 0 {
			opt.PageSize = repo.DefaultPageSize
		}
		if opt.PageSize > repo.MaxPageSize {
			opt.PageSize = repo.MaxPageSize
		}
	}

	return opt
}

func applyPageSize(opt *repo.QueryOption) int {
	if opt == nil {
		return repo.DefaultPageSize
	}

	if opt.PagingMode == PagingModeNone {
		return repo.HardLimit
	}

	return opt.PageSize
}

type cursorCallback[T any] func(cursor Cursor) T

// Paginate is the single, unified function for applying pagination logic.
func Paginate[R any, W any, O Selector, P Pageable[P, W, O, R]](query P, opt *repo.QueryOption,
	callbacks ...cursorCallback[W]) P {
	if opt == nil {
		return query.Limit(repo.DefaultPageSize)
	}

	pageSize := applyPageSize(opt)
	query = query.Limit(pageSize)

	pagingMode := opt.PagingMode
	if pagingMode == "" {
		pagingMode = PagingModeOffset
	}

	switch pagingMode {
	case PagingModeCursor:
		if opt.PageToken != "" {
			cursor, err := DecodeCursor(opt.PageToken)
			if err != nil {
				query = query.Limit(0)
			} else {
				opt.SortFromToken = true // Mark that sorting is now dictated by the token
				query = query.Order(OrderByField[O](cursor.Field, cursor.Desc))
				for _, cb := range callbacks {
					query = query.Where(cb(cursor))
				}
			}
		}
	case PagingModeOffset:
		query = query.Offset((opt.Page - 1) * pageSize)
	case PagingModeNone:
		// NoPaging is handled by applyPageSize, so nothing more to do here.
	}

	return query
}

// CountTotal executes the count query and returns the total number of records.
func CountTotal[Q Counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

// Find executes the query with pagination options and returns the results and total count.
func Find[R any, W any, O Selector, P Pageable[P, W, O, R]](ctx context.Context, query P,
	o *repo.QueryOption, callbacks ...cursorCallback[W]) ([]R, int32, error) {

	o = normalizeQueryOption(o)

	if o.OnlyCount {
		count, err := CountTotal(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
		return nil, count, nil
	}

	var count int32
	var err error

	if o.PagingMode != PagingModeCursor {
		count, err = CountTotal(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}

		if count == 0 {
			return nil, 0, nil
		}

		if o.PageSize > 0 && o.PagingMode == PagingModeOffset {
			totalPages := (int(count) + o.PageSize - 1) / o.PageSize
			if o.Page > totalPages {
				return nil, count, nil
			}
		}
	}

	query = Paginate(query, o, callbacks...)

	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}

	return result, count, nil
}
