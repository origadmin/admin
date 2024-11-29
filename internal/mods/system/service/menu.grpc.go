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

// NewMenuAPIService new a menu service.
func NewMenuAPIService(client pb.MenuAPIClient) *MenuAPIService {
	return &MenuAPIService{client: client}
}

// NewMenuAPIServer new a menu service.
func NewMenuAPIServer(client pb.MenuAPIClient) pb.MenuAPIServer {
	return &MenuAPIService{client: client}
}

func (m MenuAPIService) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	return m.client.ListMenus(ctx, request)
}

func (m MenuAPIService) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return m.client.GetMenu(ctx, request)
}

func (m MenuAPIService) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return m.client.CreateMenu(ctx, request)
}

func (m MenuAPIService) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return m.client.UpdateMenu(ctx, request)
}

func (m MenuAPIService) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return m.client.DeleteMenu(ctx, request)
}

//func (m MenuAPIService) mustEmbedUnimplementedMenuAPIServer() {
//	//TODO implement me
//	panic("implement me")
//}

var _ pb.MenuAPIServer = (*MenuAPIService)(nil)
