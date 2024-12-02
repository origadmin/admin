/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/runtime/context"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// RolesBiz is a RolePB use case.
type RolesBiz struct {
	dao     dto.RoleRepo
	limiter pagination.PageLimiter
	log     *log.Helper
}

func (m RolesBiz) ListRoles(ctx context.Context, in *pb.ListRolesRequest, opts ...grpc.CallOption) (*pb.ListRolesResponse, error) {
	var option dto.RoleQueryOption
	if err := option.FromListRequest(in, m.limiter); err != nil {
		return nil, err
	}
	log.Info("ListRoles")
	result, total, err := m.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ListRoleResponse(result, in, total)
}

func (m RolesBiz) GetRole(ctx context.Context, in *pb.GetRoleRequest, opts ...grpc.CallOption) (*pb.GetRoleResponse, error) {
	var option dto.RoleQueryOption
	if err := option.FromGetRequest(in, m.limiter); err != nil {
		return nil, err
	}
	log.Info("GetRole")
	result, err := m.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetRoleResponse{
		Role: result,
	}, nil
}

func (m RolesBiz) CreateRole(ctx context.Context, in *pb.CreateRoleRequest, opts ...grpc.CallOption) (*pb.CreateRoleResponse, error) {
	var option dto.RoleQueryOption
	if err := option.FromCreateRequest(in, m.limiter); err != nil {
		return nil, err
	}
	log.Info("CreateRole")
	result, err := m.dao.Create(ctx, in.Role, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateRoleResponse{
		Role: result,
	}, nil
}

func (m RolesBiz) UpdateRole(ctx context.Context, in *pb.UpdateRoleRequest, opts ...grpc.CallOption) (*pb.UpdateRoleResponse, error) {
	//var option dto.UpdateRoleOption
	//if err := option.FromListRequest(in, m.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateRole")
	result, err := m.dao.Update(ctx, in.Role)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateRoleResponse{
		Role: result,
	}, nil
}

func (m RolesBiz) DeleteRole(ctx context.Context, in *pb.DeleteRoleRequest, opts ...grpc.CallOption) (*pb.DeleteRoleResponse, error) {
	//var option dto.DeleteRoleOption
	//if err := option.FromListRequest(in, m.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := m.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	log.Info("DeleteRole")
	if err := m.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteRoleResponse{}, nil
}

// NewRolesBiz new a RolePB use case.
func NewRolesBiz(repo dto.RoleRepo, logger log.Logger) *RolesBiz {
	return &RolesBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewRolesClient new a RolePB use case.
func NewRolesClient(repo dto.RoleRepo, logger log.Logger) pb.RoleAPIClient {
	return &RolesBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.RoleAPIClient = (*RolesBiz)(nil)
