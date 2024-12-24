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
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/menuresource"
	"origadmin/application/admin/internal/mods/system/dto"
)

type currentRepo struct {
	db *Data
}

func (repo currentRepo) ListResources(ctx context.Context, in *dto.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.MenuResource(ctx).Query()
	return currentPageQuery(ctx, query, in, option)
}

// NewCurrentRepo .
func NewCurrentRepo(db *Data, logger log.KLogger) dto.CurrentRepo {
	return &currentRepo{
		db: db,
	}
}

func currentPageQuery(ctx context.Context, query *ent.MenuResourceQuery, in *pb.ListResourcesRequest, option dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = currentQueryPage(query, in)
	query = currentQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func currentQueryPage(query *ent.MenuResourceQuery, in *pb.ListResourcesRequest) *ent.MenuResourceQuery {
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

func currentQueryOptions(query *ent.MenuResourceQuery, option dto.ResourceQueryOption) *ent.MenuResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).MenuResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).MenuResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(currentOrderBy(option.OrderFields)...)
	}
	return query
}

func currentOrderBy(fields []string, opts ...sql.OrderTermOption) []menuresource.OrderOption {
	var orders []menuresource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
