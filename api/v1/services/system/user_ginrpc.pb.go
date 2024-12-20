// Code generated by protoc-gen-go-ginrpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-ginrpc unknown
// - protoc             (unknown)
// source: system/user.proto

package system

import (
	context "context"
	gins "github.com/origadmin/contrib/transport/gins"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)

const _ = gins.SupportPackageIsVersion1

const UserAPI_CreateUser_FullOperation = "/api.v1.services.system.UserAPI/CreateUser"
const UserAPI_DeleteUser_FullOperation = "/api.v1.services.system.UserAPI/DeleteUser"
const UserAPI_GetUser_FullOperation = "/api.v1.services.system.UserAPI/GetUser"
const UserAPI_ListUsers_FullOperation = "/api.v1.services.system.UserAPI/ListUsers"
const UserAPI_UpdateUser_FullOperation = "/api.v1.services.system.UserAPI/UpdateUser"

type UserAPIGINRPCAgentResponder interface {
	// Error returns a error
	Error(*gins.Context, int, error)
	// JSON returns a json data
	JSON(*gins.Context, int, any)
	// Any returns errors or any data
	Any(*gins.Context, int, any, error)
}

type UserAPIGINRPCAgent interface {
	UserAPIGINRPCAgentResponder
	CreateUser(*gins.Context, *CreateUserRequest)
	DeleteUser(*gins.Context, *DeleteUserRequest)
	GetUser(*gins.Context, *GetUserRequest)
	ListUsers(*gins.Context, *ListUsersRequest)
	UpdateUser(*gins.Context, *UpdateUserRequest)
}

func RegisterUserAPIGINRPCAgent(router gins.IRouter, srv UserAPIGINRPCAgent) {
	router.GET("/api/v1/sys/users", _UserAPI_ListUsers0_GINRPC_Handler(srv))
	router.GET("/api/v1/sys/users/:id", _UserAPI_GetUser0_GINRPC_Handler(srv))
	router.POST("/api/v1/sys/users", _UserAPI_CreateUser0_GINRPC_Handler(srv))
	router.PUT("/api/v1/sys/users/:user.id", _UserAPI_UpdateUser0_GINRPC_Handler(srv))
	router.DELETE("/api/v1/sys/users/:id", _UserAPI_DeleteUser0_GINRPC_Handler(srv))
}

func _UserAPI_ListUsers0_GINRPC_Handler(srv UserAPIGINRPCAgent) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in ListUsersRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_ListUsers_OperationName)
		srv.ListUsers(ctx, &in)
	}
}

func _UserAPI_GetUser0_GINRPC_Handler(srv UserAPIGINRPCAgent) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in GetUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_GetUser_OperationName)
		srv.GetUser(ctx, &in)
	}
}

func _UserAPI_CreateUser0_GINRPC_Handler(srv UserAPIGINRPCAgent) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in CreateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_CreateUser_OperationName)
		srv.CreateUser(ctx, &in)
	}
}

func _UserAPI_UpdateUser0_GINRPC_Handler(srv UserAPIGINRPCAgent) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in UpdateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_UpdateUser_OperationName)
		srv.UpdateUser(ctx, &in)
	}
}

func _UserAPI_DeleteUser0_GINRPC_Handler(srv UserAPIGINRPCAgent) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in DeleteUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			srv.Error(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_DeleteUser_OperationName)
		srv.DeleteUser(ctx, &in)
	}
}
