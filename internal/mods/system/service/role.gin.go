/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleAPIGINRPCService is a menu service.
type RoleAPIGINRPCService struct {
	resp.Response

	client pb.RoleAPIClient
}

func (s RoleAPIGINRPCService) CreateRole(context *gin.Context, request *pb.CreateRoleRequest) {
	response, err := s.client.CreateRole(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s RoleAPIGINRPCService) DeleteRole(context *gin.Context, request *pb.DeleteRoleRequest) {
	response, err := s.client.DeleteRole(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s RoleAPIGINRPCService) GetRole(context *gin.Context, request *pb.GetRoleRequest) {
	response, err := s.client.GetRole(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s RoleAPIGINRPCService) ListRoles(context *gin.Context, request *pb.ListRolesRequest) {
	response, err := s.client.ListRoles(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Roles
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s RoleAPIGINRPCService) UpdateRole(context *gin.Context, request *pb.UpdateRoleRequest) {
	response, err := s.client.UpdateRole(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.Role
	}
	s.Any(context, http.StatusOK, &res, err)
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
