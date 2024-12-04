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

// UserAPIGINService is a menu service.
type UserAPIGINService struct {
	client pb.UserAPIClient
}

func (m UserAPIGINService) CreateUser(context *gin.Context, request *pb.CreateUserRequest) {
	response, err := m.client.CreateUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINService) DeleteUser(context *gin.Context, request *pb.DeleteUserRequest) {
	response, err := m.client.DeleteUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINService) GetUser(context *gin.Context, request *pb.GetUserRequest) {
	response, err := m.client.GetUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINService) ListUsers(context *gin.Context, request *pb.ListUsersRequest) {
	response, err := m.client.ListUsers(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Users
	}
	resp.Result(context, res, err)
}

func (m UserAPIGINService) UpdateUser(context *gin.Context, request *pb.UpdateUserRequest) {
	response, err := m.client.UpdateUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	resp.Result(context, res, err)
}

// NewUserAPIGINService new a menu service.
func NewUserAPIGINService(client pb.UserAPIClient) *UserAPIGINService {
	return &UserAPIGINService{client: client}
}

// NewUserAPIGINServer new a menu service.
func NewUserAPIGINServer(client pb.UserAPIClient) pb.UserAPIGINRPCAgent {
	return &UserAPIGINService{client: client}
}

var _ pb.UserAPIGINRPCAgent = (*UserAPIGINService)(nil)
