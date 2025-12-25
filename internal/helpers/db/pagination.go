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

// paginateable defines an interface for queries that can be paginated.
type paginateable[T any, W selectable, O selectable, R any] interface {
	counter[T]
	cloneable[T]
	filterable[T, W]
	orderable[T, O]
	queryable[R]
	Limit(int) T
	Offset(int) T
}

type orderable[T any, O selectable] interface {
	Order(...O) T
}

type filterable[T any, W selectable] interface {
	Where(...W) T
}

type cloneable[T any] interface {
	Clone() T
}

// counter defines an interface for queries that can count their results.
type counter[T any] interface {
	cloneable[T]
	Count(ctx context.Context) (int, error)
}

type queryable[T any] interface {
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

type cursorCallback[T selectable] func(cursor Cursor) T

// Paginate is the single, unified function for applying pagination logic.
// It ONLY handles Limit and Offset based on the provided options.
// All WHERE and ORDER clauses are the responsibility of the caller in the DAL layer.
// Note: Page boundary validation is handled in the Query function.
func Paginate[R any, W selectable, O selectable, P paginateable[P, W, O, R]](query P, opt *repo.QueryOption,
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

// PageCount executes count query and returns the result
func PageCount[Q counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

func Query[R any, W selectable, O selectable, P paginateable[P, W, O, R]](ctx context.Context, query P,
	o *repo.QueryOption, callbacks ...cursorCallback[W]) ([]R, int32, error) {

	// Unified initialization and validation of options
	o = normalizeQueryOption(o)

	// If only count is needed, execute count query directly
	if o.OnlyCount {
		count, err := PageCount(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
		return nil, count, nil
	}

	// Clone original query for count query (must be before pagination)
	var count int32
	var countErr error

	// Execute count query only for non-cursor pagination
	if o.PageToken == "" {
		count, countErr = PageCount(ctx, query)
		if countErr != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", countErr)
		}

		// Check if requested page exceeds total pages
		if count > 0 && o.PageSize > 0 {
			totalPages := (count + int32(o.PageSize) - 1) / int32(o.PageSize) // round up
			if o.Page > int(totalPages) {
				// Page number out of range, return empty result
				return []R{}, count, nil
			}
		} else if count == 0 {
			// No data, return empty result
			return []R{}, count, nil
		}
	}

	// Apply pagination to original query
	query = Paginate(query, o, callbacks...)
	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}

	return result, count, nil
}
