/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// MenuAPIHTTPService is a menu service.
type MenuAPIHTTPService struct {
	pb.UnimplementedMenuAPIServer

	client pb.MenuAPIHTTPClient
}

func (s MenuAPIHTTPService) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return s.client.CreateMenu(ctx, request)
}

func (s MenuAPIHTTPService) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return s.client.DeleteMenu(ctx, request)
}

func (s MenuAPIHTTPService) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return s.client.GetMenu(ctx, request)
}

func (s MenuAPIHTTPService) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	return s.client.ListMenus(ctx, request)
}

func (s MenuAPIHTTPService) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return s.client.UpdateMenu(ctx, request)
}

//func (m MenuAPIHTTPService) mustEmbedUnimplementedMenuAPIHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewMenuAPIHTTPService new a menu service.
func NewMenuAPIHTTPService(client pb.MenuAPIHTTPClient) *MenuAPIHTTPService {
	return &MenuAPIHTTPService{client: client}
}

// NewMenuAPIHTTPServer new a menu service.
func NewMenuAPIHTTPServer(client pb.MenuAPIHTTPClient) pb.MenuAPIServer {
	return &MenuAPIHTTPService{client: client}
}

var _ pb.MenuAPIServer = (*MenuAPIHTTPService)(nil)
