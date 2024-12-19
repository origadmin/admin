/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuAPIGINRPCService is a menu service.
type MenuAPIGINRPCService struct {
	resp.Response

	client pb.MenuAPIClient
}

func (s MenuAPIGINRPCService) CreateMenu(context transhttp.Context, request *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	response, err := s.client.CreateMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
	return response, nil
}

func (s MenuAPIGINRPCService) DeleteMenu(context transhttp.Context, request *pb.DeleteMenuRequest) (*pb.DeleteMenuResponse, error) {
	response, err := s.client.DeleteMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Empty,
	})
	return response, nil
}

func (s MenuAPIGINRPCService) GetMenu(context transhttp.Context, request *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	response, err := s.client.GetMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
	return response, nil
}

func (s MenuAPIGINRPCService) ListMenus(context transhttp.Context, request *pb.ListMenusRequest) (*pb.ListMenusResponse, error) {
	response, err := s.client.ListMenus(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Total:   response.TotalSize,
		Data:    response.Menus,
	})
	return response, nil
}

func (s MenuAPIGINRPCService) UpdateMenu(context transhttp.Context, request *pb.UpdateMenuRequest) (*pb.UpdateMenuResponse, error) {
	response, err := s.client.UpdateMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
	return response, nil
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
