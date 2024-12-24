/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package db implements the functions, types, and interfaces for the module.
package db

import (
	pb "origadmin/application/admin/api/v1/services/system"
)

type Queryable[T any] interface {
	Limit(int) Queryable[T]
	Offset(int) Queryable[T]
}

func QueryPage[Q Queryable[T], T any](query Queryable[T], in *pb.ListResourcesRequest) Queryable[T] {
	pageSize := in.PageSize
	if pageSize > 0 {
		query = query.Limit(int(pageSize))
	}
	current := in.Current
	if current > 0 {
		query = query.Offset(int((current - 1) * pageSize))
	}
	return query
}
