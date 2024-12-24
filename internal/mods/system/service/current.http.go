/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// CurrentAPIHTTPService is a login service.
type CurrentAPIHTTPService struct {
	pb.UnimplementedCurrentAPIServer

	client pb.CurrentAPIHTTPClient
}

func (s CurrentAPIHTTPService) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	return s.client.GetCurrentUser(ctx, request)
}

func (s CurrentAPIHTTPService) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	return s.client.CurrentLogout(ctx, request)
}

func (s CurrentAPIHTTPService) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return s.client.UpdateCurrentUserPassword(ctx, request)
}

func (s CurrentAPIHTTPService) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	return s.client.UpdateCurrentUser(ctx, request)
}

func (s CurrentAPIHTTPService) ListCurrentMenus(ctx context.Context, request *pb.ListCurrentMenusRequest) (*pb.ListCurrentMenusResponse, error) {
	return s.client.ListCurrentMenus(ctx, request)
}

func (s CurrentAPIHTTPService) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	return s.client.ListCurrentRoles(ctx, request)
}

func (s CurrentAPIHTTPService) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	return s.client.UpdateCurrentSetting(ctx, request)
}

//func (l CurrentAPIHTTPService) mustEmbedUnimplementedCurrentAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewCurrentAPIHTTPService new a login service.
func NewCurrentAPIHTTPService(client pb.CurrentAPIHTTPClient) *CurrentAPIHTTPService {
	return &CurrentAPIHTTPService{client: client}
}

// NewCurrentHTTPServer new a login service.
func NewCurrentHTTPServer(client pb.CurrentAPIHTTPClient) pb.CurrentAPIServer {
	return &CurrentAPIHTTPService{client: client}
}

var _ pb.CurrentAPIServer = (*CurrentAPIHTTPService)(nil)
