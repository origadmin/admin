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

// PersonalServiceBiz is a Personal use case.
type PersonalServiceBiz struct {
	dao     dto.PersonalRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (biz PersonalServiceBiz) ListPersonalResources(ctx context.Context, in *pb.ListPersonalResourcesRequest, opts ...grpc.CallOption) (*pb.ListPersonalResourcesResponse, error) {
	return biz.dao.ListPersonalResources(ctx, in)
}

func (biz PersonalServiceBiz) RefreshPersonalToken(ctx context.Context, in *pb.RefreshPersonalTokenRequest, opts ...grpc.CallOption) (*pb.RefreshPersonalTokenResponse, error) {
	//return biz.dao.RefreshPersonalToken(ctx, in)
	return &pb.RefreshPersonalTokenResponse{}, nil
}

func (biz PersonalServiceBiz) GetPersonalProfile(ctx context.Context, in *pb.GetPersonalProfileRequest, opts ...grpc.CallOption) (*pb.GetPersonalProfileResponse, error) {
	return biz.dao.GetPersonalProfile(ctx, in)
}

func (biz PersonalServiceBiz) ListPersonalRoles(ctx context.Context, in *pb.ListPersonalRolesRequest, opts ...grpc.CallOption) (*pb.ListPersonalRolesResponse, error) {
	return biz.dao.ListPersonalRoles(ctx, in)
}

func (biz PersonalServiceBiz) PersonalLogout(ctx context.Context, in *pb.PersonalLogoutRequest, opts ...grpc.CallOption) (*pb.PersonalLogoutResponse, error) {
	return &pb.PersonalLogoutResponse{}, nil
}

func (biz PersonalServiceBiz) UpdatePersonalPassword(ctx context.Context, in *pb.UpdatePersonalPasswordRequest, opts ...grpc.CallOption) (*pb.UpdatePersonalPasswordResponse, error) {
	return biz.dao.UpdatePersonalPassword(ctx, in)
}

func (biz PersonalServiceBiz) UpdatePersonalProfile(ctx context.Context, in *pb.UpdatePersonalProfileRequest, opts ...grpc.CallOption) (*pb.UpdatePersonalProfileResponse, error) {
	return biz.dao.UpdatePersonalProfile(ctx, in)
}

func (biz PersonalServiceBiz) UpdatePersonalSetting(ctx context.Context, in *pb.UpdatePersonalSettingRequest, opts ...grpc.CallOption) (*pb.UpdatePersonalSettingResponse, error) {
	return &pb.UpdatePersonalSettingResponse{}, nil
}

// NewPersonalServiceBiz new a Personal use case.
func NewPersonalServiceBiz(repo dto.PersonalRepo, logger log.KLogger) *PersonalServiceBiz {
	return &PersonalServiceBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}
