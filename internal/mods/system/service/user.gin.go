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

// UserAPIGINRPCService is a menu service.
type UserAPIGINRPCService struct {
	resp.Response

	client pb.UserAPIClient
}

func (s UserAPIGINRPCService) CreateUser(context *gin.Context, request *pb.CreateUserRequest) {
	response, err := s.client.CreateUser(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s UserAPIGINRPCService) DeleteUser(context *gin.Context, request *pb.DeleteUserRequest) {
	response, err := s.client.DeleteUser(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.GetEmpty()
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s UserAPIGINRPCService) GetUser(context *gin.Context, request *pb.GetUserRequest) {
	response, err := s.client.GetUser(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s UserAPIGINRPCService) ListUsers(context *gin.Context, request *pb.ListUsersRequest) {
	response, err := s.client.ListUsers(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Total = response.TotalSize
		res.Data = response.Users
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s UserAPIGINRPCService) UpdateUser(context *gin.Context, request *pb.UpdateUserRequest) {
	response, err := s.client.UpdateUser(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.User
	}
	s.Any(context, http.StatusOK, &res, err)
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
