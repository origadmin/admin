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

// UserAPIGINRPCService is a menu service.
type UserAPIGINRPCService struct {
	client pb.UserAPIClient
}

func (m UserAPIGINRPCService) CreateUser(context *gin.Context, request *pb.CreateUserRequest) {
	response, err := m.client.CreateUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINRPCService) DeleteUser(context *gin.Context, request *pb.DeleteUserRequest) {
	response, err := m.client.DeleteUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINRPCService) GetUser(context *gin.Context, request *pb.GetUserRequest) {
	response, err := m.client.GetUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINRPCService) ListUsers(context *gin.Context, request *pb.ListUsersRequest) {
	response, err := m.client.ListUsers(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Users
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINRPCService) UpdateUser(context *gin.Context, request *pb.UpdateUserRequest) {
	response, err := m.client.UpdateUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

// NewUserAPIGINRPCService new a menu service.
func NewUserAPIGINRPCService(client pb.UserAPIClient) *UserAPIGINRPCService {
	return &UserAPIGINRPCService{client: client}
}

// NewUserAPIGINRPCAgent new a menu service.
func NewUserAPIGINRPCAgent(client pb.UserAPIClient) pb.UserAPIGINRPCAgent {
	return &UserAPIGINRPCService{client: client}
}
func NewUserServerAgent(client *service.GRPCClient) pb.UserAPIGINRPCAgent {
	c := pb.NewUserAPIClient(client)
	return NewUserAPIGINRPCAgent(c)
}

var _ pb.UserAPIGINRPCAgent = (*UserAPIGINRPCService)(nil)
