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
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Role,
	})
}

func (s RoleAPIGINRPCService) DeleteRole(context *gin.Context, request *pb.DeleteRoleRequest) {
	response, err := s.client.DeleteRole(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Empty,
	})
}

func (s RoleAPIGINRPCService) GetRole(context *gin.Context, request *pb.GetRoleRequest) {
	response, err := s.client.GetRole(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Role,
	})
}

func (s RoleAPIGINRPCService) ListRoles(context *gin.Context, request *pb.ListRolesRequest) {
	response, err := s.client.ListRoles(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Total:   response.TotalSize,
		Data:    response.Roles,
	})
}

func (s RoleAPIGINRPCService) UpdateRole(context *gin.Context, request *pb.UpdateRoleRequest) {
	response, err := s.client.UpdateRole(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Role,
	})
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
