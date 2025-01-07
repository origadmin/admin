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

// MenuAPIGINRPCService is a menu service.
type MenuAPIGINRPCService struct {
	resp.Response

	client pb.MenuAPIClient
}

func (s MenuAPIGINRPCService) CreateMenu(ctx context.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
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

func (s MenuAPIGINRPCService) DeleteMenu(ctx context.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
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

func (s MenuAPIGINRPCService) GetMenu(ctx context.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
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

func (s MenuAPIGINRPCService) ListMenus(ctx context.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
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

func (s MenuAPIGINRPCService) UpdateMenu(ctx context.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
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

// NewMenuAPIGINRPCService new a menu service.
func NewMenuAPIGINRPCService(client pb.MenuAPIClient) *MenuAPIGINRPCService {
	return &MenuAPIGINRPCService{client: client}
}

// NewMenuAPIAgent new a menu service.
func NewMenuAPIAgent(client pb.MenuAPIClient) pb.MenuAPIAgent {
	return &MenuAPIGINRPCService{client: client}
}
func NewMenuServerAgent(client *service.GRPCClient) pb.MenuAPIAgent {
	cli := pb.NewMenuAPIClient(client)
	return NewMenuAPIAgent(cli)
}

var _ pb.MenuAPIAgent = (*MenuAPIGINRPCService)(nil)
