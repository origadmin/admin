/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// CurrentServiceServer is a login service.
type CurrentServiceServer struct {
	pb.UnimplementedCurrentServiceServer

	client pb.CurrentServiceClient
}

func (s CurrentServiceServer) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	return s.client.GetCurrentUser(ctx, request)
}
func (s CurrentServiceServer) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	return s.client.CurrentLogout(ctx, request)
}

func (s CurrentServiceServer) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return s.client.UpdateCurrentUserPassword(ctx, request)
}

func (s CurrentServiceServer) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	return s.client.UpdateCurrentUser(ctx, request)
}

func (s CurrentServiceServer) ListCurrentResources(ctx context.Context, request *pb.ListCurrentResourcesRequest) (*pb.ListCurrentResourcesResponse, error) {
	return s.client.ListCurrentResources(ctx, request)
}

func (s CurrentServiceServer) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	return s.client.ListCurrentRoles(ctx, request)
}

func (s CurrentServiceServer) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	return s.client.UpdateCurrentSetting(ctx, request)
}

//func (l CurrentServiceServer) mustEmbedUnimplementedCurrentServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewCurrentServiceServer new a login service.
func NewCurrentServiceServer(client pb.CurrentServiceClient) *CurrentServiceServer {
	return &CurrentServiceServer{client: client}
}

// NewCurrentServiceServerPB new a login service.
func NewCurrentServiceServerPB(client pb.CurrentServiceClient) pb.CurrentServiceServer {
	return &CurrentServiceServer{client: client}
}

var _ pb.CurrentServiceServer = (*CurrentServiceServer)(nil)
