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

func (biz CurrentBiz) CurrentLogout(ctx context.Context, in *pb.CurrentLogoutRequest, opts ...grpc.CallOption) (*pb.CurrentLogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) UpdateCurrentUserPassword(ctx context.Context, in *pb.UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) ListCurrentMenus(ctx context.Context, in *pb.ListCurrentMenusRequest, opts ...grpc.CallOption) (*pb.ListCurrentMenusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) UpdateCurrentRoles(ctx context.Context, in *pb.UpdateCurrentRolesRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentRolesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (biz CurrentBiz) UpdateCurrentSetting(ctx context.Context, in *pb.UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*pb.UpdateCurrentSettingResponse, error) {
	//TODO implement me
	panic("implement me")
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
