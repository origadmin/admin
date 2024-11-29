// Code generated by protoc-gen-go-ginrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-ginrpc 0.0.6
// - protoc             (unknown)
// source: system/user.proto

package system

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	gins "github.com/origadmin/contrib/transport/gins"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(gin.H)
var _ = binding.EncodeURL

const _ = gins.SupportPackageIsVersion1

const UserAPI_CreateUser_FullOperation = "/api.v1.services.system.UserAPI/CreateUser"
const UserAPI_DeleteUser_FullOperation = "/api.v1.services.system.UserAPI/DeleteUser"
const UserAPI_GetUser_FullOperation = "/api.v1.services.system.UserAPI/GetUser"
const UserAPI_ListUsers_FullOperation = "/api.v1.services.system.UserAPI/ListUsers"
const UserAPI_UpdateUser_FullOperation = "/api.v1.services.system.UserAPI/UpdateUser"

type UserAPIGINRPCServer interface {
	CreateUser(*gin.Context, *CreateUserRequest)
	DeleteUser(*gin.Context, *DeleteUserRequest)
	GetUser(*gin.Context, *GetUserRequest)
	ListUsers(*gin.Context, *ListUsersRequest)
	UpdateUser(*gin.Context, *UpdateUserRequest)
}

func RegisterUserAPIGINRPCServer(router gins.IRouter, srv UserAPIGINRPCServer) {
	router.GET("/api/v1/sys/users", _UserAPI_ListUsers0_GINRPC_Handler(srv))
	router.GET("/api/v1/sys/users/:id", _UserAPI_GetUser0_GINRPC_Handler(srv))
	router.POST("/api/v1/sys/users", _UserAPI_CreateUser0_GINRPC_Handler(srv))
	router.PUT("/api/v1/sys/users/:user.id", _UserAPI_UpdateUser0_GINRPC_Handler(srv))
	router.DELETE("/api/v1/sys/users/:id", _UserAPI_DeleteUser0_GINRPC_Handler(srv))
}

func _UserAPI_ListUsers0_GINRPC_Handler(srv UserAPIGINRPCServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in ListUsersRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_ListUsers_OperationName)
		srv.ListUsers(ctx, &in)
	}
}

func _UserAPI_GetUser0_GINRPC_Handler(srv UserAPIGINRPCServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in GetUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_GetUser_OperationName)
		srv.GetUser(ctx, &in)
	}
}

func _UserAPI_CreateUser0_GINRPC_Handler(srv UserAPIGINRPCServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in CreateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_CreateUser_OperationName)
		srv.CreateUser(ctx, &in)
	}
}

func _UserAPI_UpdateUser0_GINRPC_Handler(srv UserAPIGINRPCServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in UpdateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_UpdateUser_OperationName)
		srv.UpdateUser(ctx, &in)
	}
}

func _UserAPI_DeleteUser0_GINRPC_Handler(srv UserAPIGINRPCServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in DeleteUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_DeleteUser_OperationName)
		srv.DeleteUser(ctx, &in)
	}
}
