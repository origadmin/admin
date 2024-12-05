/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleAPIGINRPCService is a menu service.
type RoleAPIGINRPCService struct {
	client pb.RoleAPIClient
}

func (m RoleAPIGINRPCService) CreateRole(context *gin.Context, request *pb.CreateRoleRequest) {
	response, err := m.client.CreateRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINRPCService) DeleteRole(context *gin.Context, request *pb.DeleteRoleRequest) {
	response, err := m.client.DeleteRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINRPCService) GetRole(context *gin.Context, request *pb.GetRoleRequest) {
	response, err := m.client.GetRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINRPCService) ListRoles(context *gin.Context, request *pb.ListRolesRequest) {
	response, err := m.client.ListRoles(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Roles
	}
	resp.Result(context, res, err)
}

func (m RoleAPIGINRPCService) UpdateRole(context *gin.Context, request *pb.UpdateRoleRequest) {
	response, err := m.client.UpdateRole(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	resp.Result(context, res, err)
}

// NewRoleAPIGINRPCService new a menu service.
func NewRoleAPIGINRPCService(client pb.RoleAPIClient) *RoleAPIGINRPCService {
	return &RoleAPIGINRPCService{client: client}
}

// NewRoleAPIGINRPCAgent new a menu service.
func NewRoleAPIGINRPCAgent(client pb.RoleAPIClient) pb.RoleAPIGINRPCAgent {
	return &RoleAPIGINRPCService{client: client}
}
func NewRoleServerAgent(client *service.GRPCClient) pb.RoleAPIGINRPCAgent {
	c := pb.NewRoleAPIClient(client)
	return NewRoleAPIGINRPCAgent(c)
}

var _ pb.RoleAPIGINRPCAgent = (*RoleAPIGINRPCService)(nil)
