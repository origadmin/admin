/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuAPIGINRPCService is a menu service.
type MenuAPIGINRPCService struct {
	resp.Response

	client pb.MenuAPIClient
}

func (s MenuAPIGINRPCService) CreateMenu(context *gin.Context, request *pb.CreateMenuRequest) {
	response, err := s.client.CreateMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
}

func (s MenuAPIGINRPCService) DeleteMenu(context *gin.Context, request *pb.DeleteMenuRequest) {
	response, err := s.client.DeleteMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Empty,
	})
}

func (s MenuAPIGINRPCService) GetMenu(context *gin.Context, request *pb.GetMenuRequest) {
	response, err := s.client.GetMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
}

func (s MenuAPIGINRPCService) ListMenus(context *gin.Context, request *pb.ListMenusRequest) {
	response, err := s.client.ListMenus(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Total:   response.TotalSize,
		Data:    response.Menus,
	})
}

func (s MenuAPIGINRPCService) UpdateMenu(context *gin.Context, request *pb.UpdateMenuRequest) {
	response, err := s.client.UpdateMenu(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Menu,
	})
}

// NewMenuAPIGINRPCService new a menu service.
func NewMenuAPIGINRPCService(client pb.MenuAPIClient) *MenuAPIGINRPCService {
	return &MenuAPIGINRPCService{client: client}
}

// NewMenuAPIGINRPCAgent new a menu service.
func NewMenuAPIGINRPCAgent(client pb.MenuAPIClient) pb.MenuAPIGINRPCAgent {
	return &MenuAPIGINRPCService{client: client}
}
func NewMenuServerAgent(client *service.GRPCClient) pb.MenuAPIGINRPCAgent {
	cli := pb.NewMenuAPIClient(client)
	return NewMenuAPIGINRPCAgent(cli)
}

var _ pb.MenuAPIGINRPCAgent = (*MenuAPIGINRPCService)(nil)
