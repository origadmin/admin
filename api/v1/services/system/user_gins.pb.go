// Code generated by protoc-gen-go-gins. DO NOT EDIT.
// versions:
// - protoc-gen-go-gins 0.0.11
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

const UserAPI_CreateUser_OperationName = "/api.v1.services.system.UserAPI/CreateUser"
const UserAPI_DeleteUser_OperationName = "/api.v1.services.system.UserAPI/DeleteUser"
const UserAPI_GetUser_OperationName = "/api.v1.services.system.UserAPI/GetUser"
const UserAPI_ListUsers_OperationName = "/api.v1.services.system.UserAPI/ListUsers"
const UserAPI_ResetUserPassword_OperationName = "/api.v1.services.system.UserAPI/ResetUserPassword"
const UserAPI_UpdateUser_OperationName = "/api.v1.services.system.UserAPI/UpdateUser"
const UserAPI_UpdateUserStatus_OperationName = "/api.v1.services.system.UserAPI/UpdateUserStatus"

type UserAPIGINSServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// ResetUserPassword ResetUserPassword reset the user s password
	ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*ResetUserPasswordResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// UpdateUserStatus UpdateUserStatus Update the status of the user information
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error)
}

func RegisterUserAPIGINSServer(router gins.IRouter, srv UserAPIGINSServer) {
	router.GET("/sys/users", _UserAPI_ListUsers0_GIN_Handler(srv))
	router.GET("/sys/users/:id", _UserAPI_GetUser0_GIN_Handler(srv))
	router.POST("/sys/users", _UserAPI_CreateUser0_GIN_Handler(srv))
	router.PUT("/sys/users/:user.id", _UserAPI_UpdateUser0_GIN_Handler(srv))
	router.DELETE("/sys/users/:id", _UserAPI_DeleteUser0_GIN_Handler(srv))
	router.PUT("/sys/users/:user.id/status", _UserAPI_UpdateUserStatus0_GIN_Handler(srv))
	router.POST("/sys/user/password/reset", _UserAPI_ResetUserPassword0_GIN_Handler(srv))
}

func _UserAPI_ListUsers0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in ListUsersRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_ListUsers_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.ListUsers(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_GetUser0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in GetUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_GetUser_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.GetUser(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_CreateUser0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in CreateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_CreateUser_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CreateUser(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_UpdateUser0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in UpdateUserRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_UpdateUser_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.UpdateUser(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_DeleteUser0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in DeleteUserRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_DeleteUser_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.DeleteUser(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_UpdateUserStatus0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in UpdateUserStatusRequest
		if err := gins.BindBody(ctx, &in.User); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_UpdateUserStatus_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.UpdateUserStatus(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _UserAPI_ResetUserPassword0_GIN_Handler(srv UserAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in ResetUserPasswordRequest
		if err := gins.BindBody(ctx, &in.Data); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, UserAPI_ResetUserPassword_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.ResetUserPassword(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

type UserAPIGINSClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...gins.CallOption) (rsp *CreateUserResponse, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...gins.CallOption) (rsp *DeleteUserResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...gins.CallOption) (rsp *GetUserResponse, err error)
	ListUsers(ctx context.Context, req *ListUsersRequest, opts ...gins.CallOption) (rsp *ListUsersResponse, err error)
	ResetUserPassword(ctx context.Context, req *ResetUserPasswordRequest, opts ...gins.CallOption) (rsp *ResetUserPasswordResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...gins.CallOption) (rsp *UpdateUserResponse, err error)
	UpdateUserStatus(ctx context.Context, req *UpdateUserStatusRequest, opts ...gins.CallOption) (rsp *UpdateUserStatusResponse, err error)
}

type UserAPIGINSClientImpl struct {
	cc *gins.Client
}

func NewUserAPIGINSClient(client *gins.Client) UserAPIGINSClient {
	return &UserAPIGINSClientImpl{client}
}

func (c *UserAPIGINSClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...gins.CallOption) (*CreateUserResponse, error) {
	var out CreateUserResponse
	pattern := "/sys/users"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(UserAPI_CreateUser_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...gins.CallOption) (*DeleteUserResponse, error) {
	var out DeleteUserResponse
	pattern := "/sys/users/{id}"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(UserAPI_DeleteUser_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...gins.CallOption) (*GetUserResponse, error) {
	var out GetUserResponse
	pattern := "/sys/users/{id}"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(UserAPI_GetUser_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...gins.CallOption) (*ListUsersResponse, error) {
	var out ListUsersResponse
	pattern := "/sys/users"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(UserAPI_ListUsers_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...gins.CallOption) (*ResetUserPasswordResponse, error) {
	var out ResetUserPasswordResponse
	pattern := "/sys/user/password/reset"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(UserAPI_ResetUserPassword_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...gins.CallOption) (*UpdateUserResponse, error) {
	var out UpdateUserResponse
	pattern := "/sys/users/{user.id}"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(UserAPI_UpdateUser_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIGINSClientImpl) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...gins.CallOption) (*UpdateUserStatusResponse, error) {
	var out UpdateUserStatusResponse
	pattern := "/sys/users/{user.id}/status"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(UserAPI_UpdateUserStatus_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
