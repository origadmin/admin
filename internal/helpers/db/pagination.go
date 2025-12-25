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

	"origadmin/application/admin/internal/helpers/repo"
)

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
// It aggregates capabilities for counting, cloning, filtering, ordering, and fetching results.
type Pageable[T any, W any, O any, R any] interface {
	Counter[T]
	Cloner[T]
	Filterer[T, W]
	Orderer[T, O]
	Fetcher[R]
	Limit(int) T
	Offset(int) T
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
	if opt.NoPaging {
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
	// Note: This function now only handles page size, page validation is handled in normalizeQueryOption
	if opt == nil {
		return repo.DefaultPageSize
	}

	if opt.NoPaging {
		return repo.HardLimit
	}

	return opt.PageSize
}

type cursorCallback[T any] func(cursor Cursor) T

// Paginate is the single, unified function for applying pagination logic.
// It ONLY handles Limit and Offset based on the provided options.
// All WHERE and ORDER clauses are the responsibility of the caller in the DAL layer.
// Note: Page boundary validation is handled in the Find function.
func Paginate[R any, W any, O any, P Pageable[P, W, O, R]](query P, opt *repo.QueryOption,
	callbacks ...cursorCallback[W]) P {
	if opt == nil {
		return query.Limit(repo.DefaultPageSize)
	}

	// 1. Handle NoPaging case with a hard security limit.
	// 2. Determine the final page size for standard pagination.
	pageSize := applyPageSize(opt)
	query = query.Limit(pageSize)

	if opt.NoPaging {
		return query
	}

	// 3. Apply offset only if it's not a token-based pagination request.
	if opt.PageToken == "" {
		query = query.Offset((opt.Page - 1) * pageSize)
	}

	// 4. Apply cursor-based pagination if a token is provided.
	if opt.PageToken != "" {
		cursor, err := DecodeCursor(opt.PageToken)
		if err != nil {
			query = query.Limit(0)
		} else {
			query.Order(OrderByField[O](cursor.Field, cursor.Desc))
			for _, cb := range callbacks {
				query = query.Where(cb(cursor))
			}
		}
	}

	return query
}

// CountTotal executes the count query and returns the total number of records.
// It clones the query to avoid side effects on the original query builder.
func CountTotal[Q Counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

// Find executes the query with pagination options and returns the results and total count.
// It handles both offset-based and cursor-based pagination.
// For offset-based pagination, it performs an optimization to skip the data query
// if the total count is 0 or the requested page is out of range.
func Find[R any, W any, O any, P Pageable[P, W, O, R]](ctx context.Context, query P,
	o *repo.QueryOption, callbacks ...cursorCallback[W]) ([]R, int32, error) {

	// Unified initialization and validation of options
	o = normalizeQueryOption(o)

	// Optimization: If only count is needed, execute count query directly
	if o.OnlyCount {
		count, err := CountTotal(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
		return nil, count, nil
	}

	var count int32
	var err error

	// Execute count query only for non-cursor pagination (Offset-based)
	if o.PageToken == "" {
		count, err = CountTotal(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}

		// Optimization: Early exit if no data found
		if count == 0 {
			return []R{}, 0, nil
		}

		// Optimization: Early exit if requested page exceeds total pages
		// Note: o.PageSize is guaranteed to be > 0 by normalizeQueryOption (unless NoPaging is true, where Page is 1)
		if o.PageSize > 0 {
			// Calculate total pages: ceil(count / pageSize)
			totalPages := (int(count) + o.PageSize - 1) / o.PageSize
			if o.Page > totalPages {
				// Page number out of range, return empty result
				return []R{}, count, nil
			}
		}
	}

	// Apply pagination logic (Limit, Offset, Cursor) to the query
	query = Paginate(query, o, callbacks...)

	// Execute the data query
	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}

	return result, count, nil
}
