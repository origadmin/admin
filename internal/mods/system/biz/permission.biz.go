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

// PermissionServiceBiz is a PermissionPB use case.
type PermissionServiceBiz struct {
	dao     dto.PermissionRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz PermissionServiceBiz) ListPermissions(ctx context.Context, in *pb.ListPermissionsRequest, opts ...grpc.CallOption) (*pb.ListPermissionsResponse, error) {
	var option dto.PermissionQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	option.IncludeResources = true
	log.Info("ListPermissions")
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ToListPermissionsResponse(result, in, total)
}

func (biz PermissionServiceBiz) GetPermission(ctx context.Context, in *pb.GetPermissionRequest, opts ...grpc.CallOption) (*pb.GetPermissionResponse, error) {
	var option dto.PermissionQueryOption
	if err := option.FromGetRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("GetPermission")
	result, err := biz.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetPermissionResponse{
		Permission: result,
	}, nil
}

func (biz PermissionServiceBiz) CreatePermission(ctx context.Context, in *pb.CreatePermissionRequest, opts ...grpc.CallOption) (*pb.CreatePermissionResponse, error) {
	var option dto.PermissionQueryOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("CreatePermission")
	result, err := biz.dao.Create(ctx, in.Permission, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePermissionResponse{
		Permission: result,
	}, nil
}

func (biz PermissionServiceBiz) UpdatePermission(ctx context.Context, in *pb.UpdatePermissionRequest, opts ...grpc.CallOption) (*pb.UpdatePermissionResponse, error) {
	log.Info("UpdatePermission")
	result, err := biz.dao.Update(ctx, in.Permission)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePermissionResponse{
		Permission: result,
	}, nil
}

func (biz PermissionServiceBiz) DeletePermission(ctx context.Context, in *pb.DeletePermissionRequest, opts ...grpc.CallOption) (*pb.DeletePermissionResponse, error) {
	log.Info("DeletePermission")
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeletePermissionResponse{}, nil
}

// NewPermissionServiceBiz new a PermissionPB use case.
func NewPermissionServiceBiz(repo dto.PermissionRepo, logger log.KLogger) *PermissionServiceBiz {
	return &PermissionServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
