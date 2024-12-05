/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuAPIGINRPCService is a menu service.
type MenuAPIGINRPCService struct {
	//pb.UnimplementedMenuAPIServer

	client pb.MenuAPIClient
}

func (m MenuAPIGINRPCService) CreateMenu(context *gin.Context, request *pb.CreateMenuRequest) {
	response, err := m.client.CreateMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Menu
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINRPCService) DeleteMenu(context *gin.Context, request *pb.DeleteMenuRequest) {
	response, err := m.client.DeleteMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Empty
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINRPCService) GetMenu(context *gin.Context, request *pb.GetMenuRequest) {
	response, err := m.client.GetMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Menu
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINRPCService) ListMenus(context *gin.Context, request *pb.ListMenusRequest) {
	response, err := m.client.ListMenus(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Menus
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINRPCService) UpdateMenu(context *gin.Context, request *pb.UpdateMenuRequest) {
	response, err := m.client.UpdateMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Menu
	}
	resp.Result(context, res, err)
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
