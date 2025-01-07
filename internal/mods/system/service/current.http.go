/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// CurrentServiceHTTPServer is a login service.
type CurrentServiceHTTPServer struct {
	pb.UnimplementedCurrentServiceServer

	client pb.CurrentServiceHTTPClient
}

func (s CurrentServiceHTTPServer) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	return s.client.GetCurrentUser(ctx, request)
}

func (s CurrentServiceHTTPServer) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	return s.client.CurrentLogout(ctx, request)
}

func (s CurrentServiceHTTPServer) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return s.client.UpdateCurrentUserPassword(ctx, request)
}

func (s CurrentServiceHTTPServer) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	return s.client.UpdateCurrentUser(ctx, request)
}

func (s CurrentServiceHTTPServer) ListCurrentResources(ctx context.Context, request *pb.ListCurrentResourcesRequest) (*pb.ListCurrentResourcesResponse, error) {
	return s.client.ListCurrentResources(ctx, request)
}

func (s CurrentServiceHTTPServer) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	return s.client.ListCurrentRoles(ctx, request)
}

func (s CurrentServiceHTTPServer) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	return s.client.UpdateCurrentSetting(ctx, request)
}

//func (l CurrentServiceHTTPServer) mustEmbedUnimplementedCurrentServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewCurrentServiceHTTPServer new a login service.
func NewCurrentServiceHTTPServer(client pb.CurrentServiceHTTPClient) *CurrentServiceHTTPServer {
	return &CurrentServiceHTTPServer{client: client}
}

// NewCurrentServiceHTTPServerPB new a login service.
func NewCurrentServiceHTTPServerPB(client pb.CurrentServiceHTTPClient) pb.CurrentServiceHTTPServer {
	return &CurrentServiceHTTPServer{client: client}
}

var _ pb.CurrentServiceServer = (*CurrentServiceHTTPServer)(nil)
