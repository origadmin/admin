/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// RoleServiceBiz is a RolePB use case.
type RoleServiceBiz struct {
	dao     dto.RoleRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz RoleServiceBiz) ListRoles(ctx context.Context, in *pb.ListRolesRequest, opts ...grpc.CallOption) (*pb.ListRolesResponse, error) {
	var option dto.RoleQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListRoles")
	option.IncludePermissions = true
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListRolesResponse(result, in, total)
}

func (biz RoleServiceBiz) GetRole(ctx context.Context, in *pb.GetRoleRequest, opts ...grpc.CallOption) (*pb.GetRoleResponse, error) {
	var option dto.RoleQueryOption
	if err := option.FromGetRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("GetRole")
	result, err := biz.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetRoleResponse{
		Role: result,
	}, nil
}

func (biz RoleServiceBiz) CreateRole(ctx context.Context, in *pb.CreateRoleRequest, opts ...grpc.CallOption) (*pb.CreateRoleResponse, error) {
	var option dto.RoleUpdateOption
	if err := option.FromCreateRequest(in); err != nil {
		return nil, err
	}
	log.Info("CreateRole")
	result, err := biz.dao.Create(ctx, in.Role, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateRoleResponse{
		Role: result,
	}, nil
}

func (biz RoleServiceBiz) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest, opts ...grpc.CallOption) (*pb.UpdateRoleResponse, error) {
	//var option dto.UpdateRoleOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateRole")
	result, err := biz.dao.Update(ctx, in.Role)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRoleResponse{
		Role: result,
	}, nil
}

func (biz RoleServiceBiz) DeleteRole(ctx context.Context, in *pb.DeleteRoleRequest, opts ...grpc.CallOption) (*pb.DeleteRoleResponse, error) {
	//var option dto.DeleteRoleOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := biz.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	log.Info("DeleteRole")
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteRoleResponse{}, nil
}

// NewRoleServiceBiz new a RolePB use case.
func NewRoleServiceBiz(repo dto.RoleRepo, logger log.KLogger) *RoleServiceBiz {
	return &RoleServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
