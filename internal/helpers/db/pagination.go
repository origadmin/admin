/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package db

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"origadmin/application/admin/internal/helpers/repo"
)

// Cursor, EncodeCursor, DecodeCursor, OrderFunc, and counter remain the same as the user's correct version.

type Cursor map[string]interface{}

func EncodeCursor(c Cursor) (string, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(c); err != nil {
		return "", fmt.Errorf("gob encode cursor: %w", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func DecodeCursor(token string) (Cursor, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("base64 decode token: %w", err)
	}
	var c Cursor
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("gob decode cursor: %w", err)
	}
	return c, nil
}

type OrderFunc = func(*sql.Selector)

type paginateable[T any] interface {
	Limit(int) T
	Offset(int) T
	Order(...OrderFunc) T
}

type counter[T any] interface {
	Count(ctx context.Context) (int, error)
}

type queryable interface {
	~func(*sql.Selector)
}

type Where[T queryable] interface {
	Limit(int) Where[T]
	Offset(int) Where[T]
	Order(...OrderFunc) Where[T]
	Where(...T) Where[T]
}

type cursorCallback[T queryable] func(cursor Cursor) T

func applyPageSize(opt *repo.QueryOption) int {
	if opt.NoPaging {
		return repo.HardLimit
	}

	// 2. Determine the final page size for standard pagination.
	pageSize := opt.PageSize
	if pageSize <= 0 {
		pageSize = repo.DefaultPageSize
	}
	if pageSize > repo.MaxPageSize {
		pageSize = repo.MaxPageSize
	}
	return pageSize
}

// Paginate is the single, unified function for applying all pagination logic.
// It correctly handles NoPaging with a hard limit, and standard pagination with default/max sizes.
// It is the caller's responsibility to apply WHERE and ORDER clauses.
func Paginate[P paginateable[P]](query P, opt *repo.QueryOption) P {
	// 1. Handle NoPaging case with a hard security limit.
	pageSize := applyPageSize(opt)
	query = query.Limit(pageSize)

	if opt.NoPaging {
		return query
	}

	// 3. Apply offset only if it's not a token-based pagination request.
	if opt.PageToken == "" && opt.Page > 0 {
		query = query.Offset((opt.Page - 1) * pageSize)
	}

	return query
}

func Token[T queryable](query Where[T], opt *repo.QueryOption, callback cursorCallback[T]) (Where[T], error) {
	// 1. Handle NoPaging case with a hard security limit.
	pageSize := applyPageSize(opt)
	query = query.Limit(pageSize)

	if opt.NoPaging {
		return query, nil
	}

	// 3. Apply cursor-based pagination if a token is provided.
	if opt.PageToken != "" {
		cursor, err := DecodeCursor(opt.PageToken)
		if err != nil {
			return query, fmt.Errorf("decode cursor: %w", err)
		}
		query = query.Where(callback(cursor))
	}

	return query, nil
}

// PageCount remains the same.
func PageCount[Q counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}
