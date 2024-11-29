/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleAPIGINService is a menu service.
type RoleAPIGINService struct {
	client pb.RoleAPIClient
}

func (m RoleAPIGINService) CreateRole(context *gin.Context, request *pb.CreateRoleRequest) {
	response, err := m.client.CreateRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINService) DeleteRole(context *gin.Context, request *pb.DeleteRoleRequest) {
	response, err := m.client.DeleteRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINService) GetRole(context *gin.Context, request *pb.GetRoleRequest) {
	response, err := m.client.GetRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINService) ListRoles(context *gin.Context, request *pb.ListRolesRequest) {
	response, err := m.client.ListRoles(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = int64(response.TotalSize)
		res.Data = response.Roles
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINService) UpdateRole(context *gin.Context, request *pb.UpdateRoleRequest) {
	response, err := m.client.UpdateRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

// NewRoleAPIGINService new a menu service.
func NewRoleAPIGINService(client pb.RoleAPIClient) *RoleAPIGINService {
	return &RoleAPIGINService{client: client}
}

// NewRoleAPIGINServer new a menu service.
func NewRoleAPIGINServer(client pb.RoleAPIClient) pb.RoleAPIGINRPCServer {
	return &RoleAPIGINService{client: client}
}

var _ pb.RoleAPIGINRPCServer = (*RoleAPIGINService)(nil)
