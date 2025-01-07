/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// PersonalServiceHTTPServer is a login service.
type PersonalServiceHTTPServer struct {
	pb.UnimplementedPersonalServiceServer

	client pb.PersonalServiceHTTPClient
}

func (s PersonalServiceHTTPServer) GetPersonalProfile(ctx context.Context, request *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error) {
	return s.client.GetPersonalProfile(ctx, request)
}

func (s PersonalServiceHTTPServer) PersonalLogout(ctx context.Context, request *pb.PersonalLogoutRequest) (*pb.PersonalLogoutResponse, error) {
	return s.client.PersonalLogout(ctx, request)
}

func (s PersonalServiceHTTPServer) UpdatePersonalPassword(ctx context.Context, request *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error) {
	return s.client.UpdatePersonalPassword(ctx, request)
}

func (s PersonalServiceHTTPServer) UpdatePersonalProfile(ctx context.Context, request *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error) {
	return s.client.UpdatePersonalProfile(ctx, request)
}

func (s PersonalServiceHTTPServer) ListPersonalResources(ctx context.Context, request *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error) {
	return s.client.ListPersonalResources(ctx, request)
}

func (s PersonalServiceHTTPServer) ListPersonalRoles(ctx context.Context, request *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error) {
	return s.client.ListPersonalRoles(ctx, request)
}

func (s PersonalServiceHTTPServer) UpdatePersonalSetting(ctx context.Context, request *pb.UpdatePersonalSettingRequest) (*pb.UpdatePersonalSettingResponse, error) {
	return s.client.UpdatePersonalSetting(ctx, request)
}

//func (l PersonalServiceHTTPServer) mustEmbedUnimplementedPersonalServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewPersonalServiceHTTPServer new a login service.
func NewPersonalServiceHTTPServer(client pb.PersonalServiceHTTPClient) *PersonalServiceHTTPServer {
	return &PersonalServiceHTTPServer{client: client}
}

// NewPersonalServiceHTTPServerPB new a login service.
func NewPersonalServiceHTTPServerPB(client pb.PersonalServiceHTTPClient) pb.PersonalServiceHTTPServer {
	return &PersonalServiceHTTPServer{client: client}
}

var _ pb.PersonalServiceServer = (*PersonalServiceHTTPServer)(nil)
