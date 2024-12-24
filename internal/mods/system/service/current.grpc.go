/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// CurrentAPIService is a login service.
type CurrentAPIService struct {
	pb.UnimplementedCurrentAPIServer

	client pb.CurrentAPIClient
}

func (s CurrentAPIService) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	return s.client.GetCurrentUser(ctx, request)
}
func (s CurrentAPIService) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	return s.client.CurrentLogout(ctx, request)
}

func (s CurrentAPIService) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	return s.client.UpdateCurrentUserPassword(ctx, request)
}

func (s CurrentAPIService) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	return s.client.UpdateCurrentUser(ctx, request)
}

func (s CurrentAPIService) ListCurrentMenus(ctx context.Context, request *pb.ListCurrentMenusRequest) (*pb.ListCurrentMenusResponse, error) {
	return s.client.ListCurrentMenus(ctx, request)
}

func (s CurrentAPIService) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	return s.client.ListCurrentRoles(ctx, request)
}

func (s CurrentAPIService) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	return s.client.UpdateCurrentSetting(ctx, request)
}

//func (l CurrentAPIService) mustEmbedUnimplementedCurrentAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewCurrentAPIService new a login service.
func NewCurrentAPIService(client pb.CurrentAPIClient) *CurrentAPIService {
	return &CurrentAPIService{client: client}
}

// NewCurrentAPIServer new a login service.
func NewCurrentAPIServer(client pb.CurrentAPIClient) pb.CurrentAPIServer {
	return &CurrentAPIService{client: client}
}

var _ pb.CurrentAPIServer = (*CurrentAPIService)(nil)
