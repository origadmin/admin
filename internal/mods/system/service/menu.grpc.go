/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// MenuServiceServer is a menu service.
type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer

	client pb.MenuServiceClient
}

func (s MenuServiceServer) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	return s.client.ListMenus(ctx, request)
}

func (s MenuServiceServer) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return s.client.GetMenu(ctx, request)
}

func (s MenuServiceServer) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return s.client.CreateMenu(ctx, request)
}

func (s MenuServiceServer) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return s.client.UpdateMenu(ctx, request)
}

func (s MenuServiceServer) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return s.client.DeleteMenu(ctx, request)
}

//func (m MenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewMenuServiceServer new a menu service.
func NewMenuServiceServer(client pb.MenuServiceClient) *MenuServiceServer {
	return &MenuServiceServer{client: client}
}

// NewMenuServiceServerPB new a menu service.
func NewMenuServiceServerPB(client pb.MenuServiceClient) pb.MenuServiceServer {
	return &MenuServiceServer{client: client}
}

var _ pb.MenuServiceServer = (*MenuServiceServer)(nil)
