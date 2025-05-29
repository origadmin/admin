/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuServiceBridge is a menu service.
type MenuServiceBridge struct {
	pb.UnimplementedMenuServiceServer

	client pb.MenuServiceClient
}

func (s MenuServiceBridge) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateMenu(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

func (s MenuServiceBridge) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	_, err := s.client.DeleteMenu(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    nil,
	})
	return nil, nil
}

func (s MenuServiceBridge) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetMenu(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

func (s MenuServiceBridge) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListMenus(ctx, request)
	if err != nil {
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Menus...),
	})
	return nil, nil
}

func (s MenuServiceBridge) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateMenu(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response.Menu),
	})
	return nil, nil
}

// NewMenuServiceBridge new a menu service.
func NewMenuServiceBridge(client pb.MenuServiceClient) *MenuServiceBridge {
	return &MenuServiceBridge{client: client}
}

// NewMenuServiceBridgePB new a menu service.
func NewMenuServiceBridgePB(client pb.MenuServiceClient) pb.MenuServiceServer {
	return &MenuServiceBridge{client: client}
}
func NewMenuServiceBridgeClient(client *service.GRPCClient) pb.MenuServiceServer {
	cli := pb.NewMenuServiceClient(client)
	return NewMenuServiceBridge(cli)
}

var _ pb.MenuServiceServer = (*MenuServiceBridge)(nil)
