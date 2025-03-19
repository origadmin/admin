/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/biz"
)

// PersonalServiceServer is a login service.
type PersonalServiceServer struct {
	pb.UnimplementedPersonalServiceServer

	client *biz.PersonalServiceBiz
}

func (s PersonalServiceServer) GetPersonalProfile(ctx context.Context, request *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error) {
	return s.client.GetPersonalProfile(ctx, request)
}
func (s PersonalServiceServer) PersonalLogout(ctx context.Context, request *pb.PersonalLogoutRequest) (*pb.PersonalLogoutResponse, error) {
	return s.client.PersonalLogout(ctx, request)
}

func (s PersonalServiceServer) UpdatePersonalPassword(ctx context.Context, request *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error) {
	return s.client.UpdatePersonalPassword(ctx, request)
}

func (s PersonalServiceServer) UpdatePersonalProfile(ctx context.Context, request *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error) {
	return s.client.UpdatePersonalProfile(ctx, request)
}

func (s PersonalServiceServer) ListPersonalResources(ctx context.Context, request *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error) {
	return s.client.ListPersonalResources(ctx, request)
}

func (s PersonalServiceServer) ListPersonalRoles(ctx context.Context, request *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error) {
	return s.client.ListPersonalRoles(ctx, request)
}

func (s PersonalServiceServer) UpdatePersonalSetting(ctx context.Context, request *pb.UpdatePersonalSettingRequest) (*pb.UpdatePersonalSettingResponse, error) {
	return s.client.UpdatePersonalSetting(ctx, request)
}

//func (l PersonalServiceServer) mustEmbedUnimplementedPersonalServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewPersonalServiceServer new a login service.
func NewPersonalServiceServer(client *biz.PersonalServiceBiz) *PersonalServiceServer {
	return &PersonalServiceServer{client: client}
}

// NewPersonalServiceServerPB new a login service.
func NewPersonalServiceServerPB(client *biz.PersonalServiceBiz) pb.PersonalServiceServer {
	return &PersonalServiceServer{client: client}
}

var _ pb.PersonalServiceServer = (*PersonalServiceServer)(nil)
