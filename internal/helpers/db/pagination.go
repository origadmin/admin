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

func applyPageSize(opt *repo.QueryOption) int {
	if opt == nil {
		return repo.DefaultPageSize
	}

	if opt.NoPaging {
		return repo.HardLimit
	}

	pageSize := opt.PageSize
	if pageSize <= 0 {
		pageSize = repo.DefaultPageSize
	}
	if pageSize > repo.MaxPageSize {
		pageSize = repo.MaxPageSize
	}
	return pageSize
}

type cursorCallback[T selectable] func(cursor Cursor) T

// Paginate is the single, unified function for applying pagination logic.
// It ONLY handles Limit and Offset based on the provided options.
// All WHERE and ORDER clauses are the responsibility of the caller in the DAL layer.
func Paginate[R any, W selectable, O selectable, P paginateable[P, W, O, R]](query P, opt *repo.QueryOption,
	callbacks ...cursorCallback[W]) P {
	// 1. Handle NoPaging case with a hard security limit.
	// 2. Determine the final page size for standard pagination.
	pageSize := applyPageSize(opt)
	query = query.Limit(pageSize)

	if opt.NoPaging {
		return query
	}

	// 3. Apply offset only if it's not a token-based pagination request.
	if opt.PageToken == "" && opt.Page > 0 {
		query = query.Offset((opt.Page - 1) * pageSize)
	}
	// 4. Apply cursor-based pagination if a token is provided.
	if opt.PageToken != "" {
		cursor, err := DecodeCursor(opt.PageToken)
		if err != nil {
			return query
		}
		for _, cb := range callbacks {
			query = query.Where(cb(cursor))
		}
	}

	return query
}

// PageCount 执行count查询并返回结果
func PageCount[Q counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Clone().Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

func Query[R any, W selectable, O selectable, P paginateable[P, W, O, R]](ctx context.Context, query P,
	o *repo.QueryOption, callbacks ...cursorCallback[W]) ([]R, int32, error) {

	// 如果只需要计数，直接执行count查询
	if o != nil && o.OnlyCount {
		count, err := PageCount(ctx, query)
		if err != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", err)
		}
		return nil, count, nil
	}

	// 先克隆原始查询用于count查询（必须在分页之前）
	var count int32
	var countErr error

	// 只有在非cursor分页时才执行count查询
	if o == nil || o.PageToken == "" {
		count, countErr = PageCount(ctx, query)
		if countErr != nil {
			return nil, 0, fmt.Errorf("count query failed: %w", countErr)
		}
	}

	// 对原始查询应用分页
	query = Paginate(query, o, callbacks...)
	result, err := query.All(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("data query failed: %w", err)
	}

	return result, count, nil
}
