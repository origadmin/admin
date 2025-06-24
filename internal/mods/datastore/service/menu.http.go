/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"

	pb "origadmin/application/admin/api/v1/services/system"
)

// MenuServiceHTTPServer is a menu service.
type MenuServiceHTTPServer struct {
	pb.UnimplementedMenuServiceServer

	client pb.MenuServiceHTTPClient
	log    *log.KHelper
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
func NewMenuServiceHTTPServer(client pb.MenuServiceHTTPClient, logger log.KLogger) *MenuServiceHTTPServer {
	return &MenuServiceHTTPServer{
		client: client,
		log:    log.NewHelper(logger),
	}
}

// NewMenuServiceHTTPServerPB new a menu service.
func NewMenuServiceHTTPServerPB(r runtime.Runtime, client pb.MenuServiceHTTPClient) pb.MenuServiceHTTPServer {
	return NewMenuServiceHTTPServer(client, r.WithLogger("module", "service/system"))
}

var _ pb.MenuServiceServer = (*MenuServiceHTTPServer)(nil)
