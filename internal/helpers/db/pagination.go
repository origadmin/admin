/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package db

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"entgo.io/ent/dialect/sql"

	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/helpers/repo"
)

// Cursor represents the data encoded in a pagination token.
type Cursor struct {
	ID int64 `json:"id"`
}

// EncodeCursor creates a base64-encoded token from a cursor.
func EncodeCursor(c *Cursor) (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DecodeCursor parses a base64-encoded token into a Cursor struct.
func DecodeCursor(token string) (*Cursor, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	var c Cursor
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, err
	}
	return &c, nil
}

type OrderFunc = func(*sql.Selector)

type WherePredicate interface {
	~func(*sql.Selector)
}

// paginateable defines the minimal interface for a query builder to support pagination.
// All Ent query builders satisfy this interface.
type paginateable[T any] interface {
	Limit(int) T
	Offset(int) T
	Order(...OrderFunc) T
}

// counter defines an interface for queries that can count their results.
type counter[T any] interface {
	Count(ctx context.Context) (int, error)
}

// Pagination applies pagination logic to a query builder.
// It is the caller's responsibility to apply the correct WHERE clause for cursor-based pagination *before* calling this function.
func Pagination[P paginateable[P]](query P, opt *repo.QueryOption, idField string) P {
	// Apply the limit first, as it's common to both pagination types.
	if opt.PageSize > 0 {
		query = query.Limit(opt.PageSize)
	}

	if opt.PageToken != "" {
		// If a token is used, the WHERE clause is assumed to be already applied by the caller.
		// We just need to enforce the correct ordering for the cursor to work.
		query = query.Order(ent.Asc(idField))
	} else if opt.Page > 0 {
		// For offset-based pagination, apply the offset.
		query = query.Offset((opt.Page - 1) * opt.PageSize)
	}

	return query
}

// PageCount executes a count query and returns the total number of records.
func PageCount[Q counter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}
