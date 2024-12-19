/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package service implements the functions, types, and interfaces for the module.
package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// UserAPIGINRPCService is a menu service.
type UserAPIGINRPCService struct {
	resp.Response

	client pb.UserAPIClient
}

func (s UserAPIGINRPCService) CreateUser(context transhttp.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	response, err := s.client.CreateUser(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.User,
	})
	return response, nil
}

func (s UserAPIGINRPCService) DeleteUser(context transhttp.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	response, err := s.client.DeleteUser(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Empty,
	})
	return response, nil
}

func (s UserAPIGINRPCService) GetUser(context transhttp.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	response, err := s.client.GetUser(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.User,
	})
	return response, nil
}

func (s UserAPIGINRPCService) ListUsers(context transhttp.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	response, err := s.client.ListUsers(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Total:   response.TotalSize,
		Data:    response.Users,
	})
	return response, nil
}

func (s UserAPIGINRPCService) UpdateUser(context transhttp.Context, request *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	response, err := s.client.UpdateUser(context, request)
	if err != nil {
		s.Error(context, http.StatusNotFound, err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.User,
	})
	return response, nil
}

// NewUserAPIGINRPCService new a menu service.
func NewUserAPIGINRPCService(client pb.UserAPIClient) *UserAPIGINRPCService {
	return &UserAPIGINRPCService{client: client}
}

// NewUserAPIAgent new a menu service.
func NewUserAPIAgent(client pb.UserAPIClient) pb.UserAPIAgent {
	return &UserAPIGINRPCService{client: client}
}
func NewUserServerAgent(client *service.GRPCClient) pb.UserAPIAgent {
	c := pb.NewUserAPIClient(client)
	return NewUserAPIAgent(c)
}

var _ pb.UserAPIAgent = (*UserAPIGINRPCService)(nil)
