/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// MenuAPIService is a menu service.
type MenuAPIService struct {
	pb.UnimplementedMenuAPIServer

	client pb.MenuAPIClient
}

func (s MenuAPIService) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	return s.client.ListMenus(ctx, request)
}

func (s MenuAPIService) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return s.client.GetMenu(ctx, request)
}

func (s MenuAPIService) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return s.client.CreateMenu(ctx, request)
}

func (s MenuAPIService) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return s.client.UpdateMenu(ctx, request)
}

func (s MenuAPIService) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return s.client.DeleteMenu(ctx, request)
}

//func (m MenuAPIService) mustEmbedUnimplementedMenuAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewMenuAPIService new a menu service.
func NewMenuAPIService(client pb.MenuAPIClient) *MenuAPIService {
	return &MenuAPIService{client: client}
}

// NewMenuAPIServer new a menu service.
func NewMenuAPIServer(client pb.MenuAPIClient) pb.MenuAPIServer {
	return &MenuAPIService{client: client}
}

var _ pb.MenuAPIServer = (*MenuAPIService)(nil)
