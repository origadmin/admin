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

type authRepo struct {
	db *Data
}

func (repo authRepo) ListResources(ctx context.Context, in *dto.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.MenuResource(ctx).Query()
	return resourcePageQuery(ctx, query, in, option)
}

// NewAuthRepo .
func NewAuthRepo(db *Data, logger log.KLogger) dto.AuthRepo {
	return &authRepo{
		db: db,
	}
}

func resourcePageQuery(ctx context.Context, query *ent.MenuResourceQuery, in *pb.ListResourcesRequest, option dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = resourceQueryPage(query, in)
	query = resourceQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func resourceQueryPage(query *ent.MenuResourceQuery, in *pb.ListResourcesRequest) *ent.MenuResourceQuery {
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

func resourceQueryOptions(query *ent.MenuResourceQuery, option dto.ResourceQueryOption) *ent.MenuResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).MenuResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).MenuResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(resourceOrderBy(option.OrderFields)...)
	}
	return query
}

func resourceOrderBy(fields []string, opts ...sql.OrderTermOption) []menuresource.OrderOption {
	var orders []menuresource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
