/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dto"
)

type authRepo struct {
	db *Data
}

func (repo authRepo) ListAuthResources(ctx context.Context, in *dto.ListAuthResourcesRequest, options ...dto.AuthResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.AuthResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Resource(ctx).Query()
	return authResourcePageQuery(ctx, query, in, option)
}

// NewAuthRepo .
func NewAuthRepo(db *Data, logger log.KLogger) dto.AuthRepo {
	return &authRepo{
		db: db,
	}
}

func authResourcePageQuery(ctx context.Context, query *ent.ResourceQuery, in *pb.ListAuthResourcesRequest, option dto.AuthResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = authResourceQueryPage(query, in)
	query = authResourceQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func authResourceQueryPage(query *ent.ResourceQuery, in *pb.ListAuthResourcesRequest) *ent.ResourceQuery {
	if in.NoPaging {
		pageSize := in.PageSize
		if pageSize > 0 {
			query = query.Limit(int(pageSize))
		}
		return query
	}

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

func authResourceQueryOptions(query *ent.ResourceQuery, option dto.AuthResourceQueryOption) *ent.ResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).ResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).ResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(resourceOrderBy(option.OrderFields)...)
	}
	return query
}

func authResourceOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
