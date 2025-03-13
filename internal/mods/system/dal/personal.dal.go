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

type personalRepo struct {
	db *Data
}

func (repo personalRepo) GetPersonalProfile(ctx context.Context, in *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error) {
	userid := securityx.GetUserID(ctx)
	if userid == "" {
		return nil, dto.ErrUserNotFound
	}
	if userid == "admin" {
		return &pb.GetPersonalProfileResponse{
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
	return &pb.GetPersonalProfileResponse{
		User: dto.ConvertUser2PB(userObj),
	}, nil
}

func (repo personalRepo) ListPersonalRoles(ctx context.Context, in *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo personalRepo) UpdatePersonalPassword(ctx context.Context, in *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo personalRepo) UpdatePersonalProfile(ctx context.Context, in *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (repo personalRepo) ListPersonalResources(ctx context.Context, in *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error) {
	uid := securityx.GetUserID(ctx)
	log.Infof("uid: %+v", uid)
	if uid == "admin" {
		resources, err := repo.db.Resource(ctx).Query().Where(resource.StatusEQ(dto.ResourceStatusEnabled)).All(ctx)
		if err != nil {
			return nil, err
		}
		return &pb.ListPersonalResourcesResponse{
			TotalSize: int64(len(resources)),
			Resources: dto.ConvertResources(resources),
		}, nil
	}
	resources, err := repo.db.User(ctx).Query().Where(user.ID(in.Id)).QueryRoles().QueryPermissions().QueryResources().
		Where(resource.StatusEQ(dto.ResourceStatusEnabled)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ListPersonalResourcesResponse{
		TotalSize: int64(len(resources)),
		Resources: dto.ConvertResources(resources),
	}, nil
}

func (repo personalRepo) ListResources(ctx context.Context, in *dto.ListResourcesRequest, options ...dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	var option dto.ResourceQueryOption
	if len(options) > 0 {
		option = options[0]
	}
	query := repo.db.Resource(ctx).Query()
	return personalPageQuery(ctx, query, in, option)
}

// NewPersonalRepo .
func NewPersonalRepo(db *Data, logger log.KLogger) dto.PersonalRepo {
	return &personalRepo{
		db: db,
	}
}

func personalPageQuery(ctx context.Context, query *ent.ResourceQuery, in *pb.ListResourcesRequest, option dto.ResourceQueryOption) ([]*dto.ResourcePB, int32, error) {
	query = personalQueryPage(query, in)
	query = personalQueryOptions(query, option)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	result, err := query.All(ctx)
	return dto.ConvertResources(result), int32(count), err
}

func personalQueryPage(query *ent.ResourceQuery, in *pb.ListResourcesRequest) *ent.ResourceQuery {
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

func personalQueryOptions(query *ent.ResourceQuery, option dto.ResourceQueryOption) *ent.ResourceQuery {
	if len(option.SelectFields) > 0 {
		query = query.Select(option.SelectFields...).ResourceQuery
	}
	if len(option.OmitFields) > 0 {
		query = query.Omit(option.OmitFields...).ResourceQuery
	}
	if len(option.OrderFields) > 0 {
		query = query.Order(personalOrderBy(option.OrderFields)...)
	}
	return query
}

func personalOrderBy(fields []string, opts ...sql.OrderTermOption) []resource.OrderOption {
	var orders []resource.OrderOption
	for _, field := range fields {
		orders = append(orders, sql.OrderByField(field, opts...).ToFunc())
	}
	return orders
}
