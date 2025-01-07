/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/context"

	pb "origadmin/application/admin/api/v1/services/system"
)

// MenuServiceHTTPServer is a menu service.
type MenuServiceHTTPServer struct {
	pb.UnimplementedMenuServiceServer

	client pb.MenuServiceHTTPClient
}

func (s MenuServiceHTTPServer) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	return s.client.CreateMenu(ctx, request)
}

func (s MenuServiceHTTPServer) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	return s.client.DeleteMenu(ctx, request)
}

func (s MenuServiceHTTPServer) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	return s.client.GetMenu(ctx, request)
}

func (s MenuServiceHTTPServer) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	return s.client.ListMenus(ctx, request)
}

func (s MenuServiceHTTPServer) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	return s.client.UpdateMenu(ctx, request)
}

//func (m MenuServiceHTTPServer) mustEmbedUnimplementedMenuServiceHTTPServer() {
//	//TODO implement me
//	panic("implement me")
//}

// NewMenuServiceHTTPServer new a menu service.
func NewMenuServiceHTTPServer(client pb.MenuServiceHTTPClient) *MenuServiceHTTPServer {
	return &MenuServiceHTTPServer{client: client}
}

// NewMenuServiceHTTPServerPB new a menu service.
func NewMenuServiceHTTPServerPB(client pb.MenuServiceHTTPClient) pb.MenuServiceHTTPServer {
	return &MenuServiceHTTPServer{client: client}
}

var _ pb.MenuServiceServer = (*MenuServiceHTTPServer)(nil)
