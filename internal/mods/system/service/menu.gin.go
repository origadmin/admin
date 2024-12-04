/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuAPIGINService is a menu service.
type MenuAPIGINService struct {
	//pb.UnimplementedMenuAPIServer

	client pb.MenuAPIClient
}

func (m MenuAPIGINService) CreateMenu(context *gin.Context, request *pb.CreateMenuRequest) {
	response, err := m.client.CreateMenu(context, request)
	if err != nil {
		gins.ResultError(context, err)
	}
	gins.ResultSuccess(context, gins.Result{
		Success: true,
		Data:    response.Menu,
	})
}

func (m MenuAPIGINService) DeleteMenu(context *gin.Context, request *pb.DeleteMenuRequest) {
	response, err := m.client.DeleteMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Empty
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINService) GetMenu(context *gin.Context, request *pb.GetMenuRequest) {
	response, err := m.client.GetMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Menu
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINService) ListMenus(context *gin.Context, request *pb.ListMenusRequest) {
	response, err := m.client.ListMenus(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Menus
	}
	resp.Result(context, res, err)
}

func (m MenuAPIGINService) UpdateMenu(context *gin.Context, request *pb.UpdateMenuRequest) {
	response, err := m.client.UpdateMenu(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Menu
	}
	resp.Result(context, res, err)
}

// NewMenuAPIGINService new a menu service.
func NewMenuAPIGINService(client pb.MenuAPIClient) *MenuAPIGINService {
	return &MenuAPIGINService{client: client}
}

// NewMenuAPIGINServer new a menu service.
func NewMenuAPIGINServer(client pb.MenuAPIClient) pb.MenuAPIGINRPCAgent {
	return &MenuAPIGINService{client: client}
}

var _ pb.MenuAPIGINRPCAgent = (*MenuAPIGINService)(nil)
