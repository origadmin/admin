/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/securityx"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/resource"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/user"
	"origadmin/application/admin/internal/mods/system/dto"
)

type currentRepo struct {
	db *Data
}

func (repo currentRepo) GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	userid := securityx.GetUserID(ctx)
	if userid == "" {
		return nil, dto.ErrUserNotFound
	}
	if userid == "admin" {
		return &pb.GetCurrentUserResponse{
			User: &pb.User{
				Id:       0,
				Uuid:     "admin",
				Username: "admin",
				Email:    "admin",
				Phone:    "admin",
				Avatar:   "https://raw.githubusercontent.com/OrigAdmin/OrigAdmin/master/origadmin.png",
				Nickname: "admin",
				Status:   1,
			},
		}, nil
	}
	userObj, err := repo.db.User(ctx).Query().Where(user.UUID(userid)).First(ctx)
	if err != nil {
		return nil, dto.ErrUserNotFound
	}
	return &pb.GetCurrentUserResponse{
		User: dto.ConvertUser2PB(userObj),
	}, nil
}

func (repo currentRepo) ListCurrentRoles(ctx context.Context, in *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo currentRepo) UpdateCurrentUserPassword(ctx context.Context, in *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo currentRepo) UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo currentRepo) ListCurrentResources(ctx context.Context, in *pb.ListCurrentResourcesRequest) (*pb.ListCurrentResourcesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo currentRepo) ListResources(ctx context.Context, in *dto.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Resource(ctx).Query()
	return currentPageQuery(ctx, query, in, option)
}

// NewCurrentRepo .
func NewCurrentRepo(db *Data, logger log.KLogger) dto.CurrentRepo {
	return &currentRepo{
		db: db,
	}
}

func currentPageQuery(ctx context.Context, query *ent.ResourceQuery, in *pb.ListResourcesRequest, option dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = currentQueryPage(query, in)
	query = currentQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func currentQueryPage(query *ent.ResourceQuery, in *pb.ListResourcesRequest) *ent.ResourceQuery {
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

func currentQueryOptions(query *ent.ResourceQuery, option dto.ResourceQueryOption) *ent.ResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).ResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).ResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(currentOrderBy(option.OrderFields)...)
	}
	return query
}

func currentOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
