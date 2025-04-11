/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db implements the functions, types, and interfaces for the module.
package db

import (
	"context"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/toolkits/net/pagination"
)

type QueryPager[T any] interface {
	Limit(int) T
	Offset(int) T
}

type QueryCounter[T any] interface {
	Count(ctx context.Context) (int, error)
}

type FieldSelector[T any] interface {
	Select(...string) T
	Omit(...string) T
}

func PaginationQuery[Q QueryPager[Q]](query Q, in pagination.PageRequest, paging bool) Q {
	if !paging {
		return NoPageQuery(query, in)
	}
	return PageQuery(query, in)
}

func NoPageQuery[Q QueryPager[Q]](query Q, in pagination.PageSizeGetter) Q {
	pageSize := in.GetPageSize()
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}
	return query
}

func handleTokenPagination[Q QueryPager[Q]](query Q, token string) Q {
	// TODO: 实现游标分页逻辑
	// 示例伪代码：
	// decodedToken := decodeToken(token)
	// query = query.Where(...).Order(...).Limit(...)
	return query
}

func PageQuery[Q QueryPager[Q]](query Q, in pagination.PageRequest) Q {
	pageSize := in.GetPageSize()
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}
	token := in.GetPageToken()
	if token != "" {
		return handleTokenPagination(query, token)
	}
	current := in.GetCurrent()
	if current > 0 {
		return query.Offset(int((current - 1) * pageSize))
	}

	return query
}

func PageCount[Q QueryCounter[Q]](ctx context.Context, query Q) (int32, error) {
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

func OrderBy[T ~func(*sql.Selector)](fields []string, orders ...T) []T {
	for _, field := range fields {
		parts := strings.Split(field, ",")
		fieldName := parts[0]
		var orderOpt sql.OrderTermOption

		if len(parts) > 1 {
			switch strings.ToLower(parts[1]) {
			case "desc":
				orderOpt = sql.OrderDesc()
			default:
				orderOpt = sql.OrderAsc()
			}
		} else {
			orderOpt = sql.OrderAsc()
		}

		orders = append(orders, sql.OrderByField(fieldName, orderOpt).ToFunc())
	}
	return orders
}
