/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"context"

	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// CurrentServiceClientBiz is a Current use case.
type CurrentServiceClientBiz struct {
	dao     dto.CurrentRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz CurrentServiceClientBiz) ListCurrentResources(ctx context.Context, in *pb.ListCurrentResourcesRequest, opts ...grpc.CallOption) (*pb.ListCurrentResourcesResponse, error) {
	return biz.dao.ListCurrentResources(ctx, in)
}

func (biz CurrentServiceClientBiz) RefreshCurrentToken(ctx context.Context, in *pb.RefreshCurrentTokenRequest, opts ...grpc.CallOption) (*pb.RefreshCurrentTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentServiceClientBiz) GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest, opts ...grpc.CallOption) (*pb.GetCurrentUserResponse, error) {
	return biz.dao.GetCurrentUser(ctx, in)
}

func (biz CurrentServiceClientBiz) ListCurrentRoles(ctx context.Context, in *pb.ListCurrentRolesRequest, opts ...grpc.CallOption) (*pb.ListCurrentRolesResponse, error) {
	return biz.dao.ListCurrentRoles(ctx, in)
}

func (biz CurrentServiceClientBiz) CurrentLogout(ctx context.Context, in *pb.CurrentLogoutRequest, opts ...grpc.CallOption) (*pb.CurrentLogoutResponse, error) {
	return &pb.CurrentLogoutResponse{}, nil
}

func (biz CurrentServiceClientBiz) UpdateCurrentUserPassword(ctx context.Context, in *pb.UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return biz.dao.UpdateCurrentUserPassword(ctx, in)
}

func (biz CurrentServiceClientBiz) UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserResponse, error) {
	return biz.dao.UpdateCurrentUser(ctx, in)
}

func (biz CurrentServiceClientBiz) UpdateCurrentSetting(ctx context.Context, in *pb.UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentSettingResponse, error) {
	return &pb.UpdateCurrentSettingResponse{}, nil
}

// NewCurrentServiceClientBiz new a Current use case.
func NewCurrentServiceClientBiz(repo dto.CurrentRepo, logger log.KLogger) *CurrentServiceClientBiz {
	return &CurrentServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewCurrentServiceClient new a Current use case.
func NewCurrentServiceClient(repo dto.CurrentRepo, logger log.KLogger) pb.CurrentServiceClient {
	return &CurrentServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.CurrentServiceClient = (*CurrentServiceClientBiz)(nil)
