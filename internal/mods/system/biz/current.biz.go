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

// CurrentBiz is a Current use case.
type CurrentBiz struct {
	dao     dto.CurrentRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz CurrentBiz) RefreshCurrentToken(ctx context.Context, in *pb.RefreshCurrentTokenRequest, opts ...grpc.CallOption) (*pb.RefreshCurrentTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest, opts ...grpc.CallOption) (*pb.GetCurrentUserResponse, error) {
	return biz.dao.GetCurrentUser(ctx, in)
}

func (biz CurrentBiz) ListCurrentRoles(ctx context.Context, in *pb.ListCurrentRolesRequest, opts ...grpc.CallOption) (*pb.ListCurrentRolesResponse, error) {
	return biz.dao.ListCurrentRoles(ctx, in)
}

func (biz CurrentBiz) CurrentLogout(ctx context.Context, in *pb.CurrentLogoutRequest, opts ...grpc.CallOption) (*pb.CurrentLogoutResponse, error) {
	return &pb.CurrentLogoutResponse{}, nil
}

func (biz CurrentBiz) UpdateCurrentUserPassword(ctx context.Context, in *pb.UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return biz.dao.UpdateCurrentUserPassword(ctx, in)
}

func (biz CurrentBiz) UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserResponse, error) {
	return biz.dao.UpdateCurrentUser(ctx, in)
}

func (biz CurrentBiz) ListCurrentMenus(ctx context.Context, in *pb.ListCurrentMenusRequest, opts ...grpc.CallOption) (*pb.ListCurrentMenusResponse, error) {
	return biz.dao.ListCurrentMenus(ctx, in)
}

func (biz CurrentBiz) UpdateCurrentSetting(ctx context.Context, in *pb.UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentSettingResponse, error) {
	return &pb.UpdateCurrentSettingResponse{}, nil
}

// NewCurrentBiz new a Current use case.
func NewCurrentBiz(repo dto.CurrentRepo, logger log.KLogger) *CurrentBiz {
	return &CurrentBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewCurrentClient new a Current use case.
func NewCurrentClient(repo dto.CurrentRepo, logger log.KLogger) pb.CurrentAPIClient {
	return &CurrentBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.CurrentAPIClient = (*CurrentBiz)(nil)
