// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: system/user.proto

package system

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserAPICreateUser = "/api.v1.services.system.UserAPI/CreateUser"
const OperationUserAPIDeleteUser = "/api.v1.services.system.UserAPI/DeleteUser"
const OperationUserAPIGetUser = "/api.v1.services.system.UserAPI/GetUser"
const OperationUserAPIListUsers = "/api.v1.services.system.UserAPI/ListUsers"
const OperationUserAPIResetUserPassword = "/api.v1.services.system.UserAPI/ResetUserPassword"
const OperationUserAPIUpdateUser = "/api.v1.services.system.UserAPI/UpdateUser"
const OperationUserAPIUpdateUserRoles = "/api.v1.services.system.UserAPI/UpdateUserRoles"
const OperationUserAPIUpdateUserStatus = "/api.v1.services.system.UserAPI/UpdateUserStatus"

type UserAPIHTTPServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	// ResetUserPassword ResetUserPassword reset the user s password
	ResetUserPassword(context.Context, *ResetUserPasswordRequest) (*ResetUserPasswordResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// UpdateUserRoles UpdateUserRoles update the user roles
	UpdateUserRoles(context.Context, *UpdateUserRolesRequest) (*UpdateUserRolesResponse, error)
	// UpdateUserStatus UpdateUserStatus Update the status of the user information
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error)
}

func RegisterUserAPIHTTPServer(s *http.Server, srv UserAPIHTTPServer) {
	r := s.Route("/")
	r.GET("/sys/users", _UserAPI_ListUsers0_HTTP_Handler(srv))
	r.GET("/sys/users/{id}", _UserAPI_GetUser0_HTTP_Handler(srv))
	r.POST("/sys/users", _UserAPI_CreateUser0_HTTP_Handler(srv))
	r.PUT("/sys/users/{user.id}", _UserAPI_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("/sys/users/{user.id}", _UserAPI_DeleteUser0_HTTP_Handler(srv))
	r.PUT("/sys/users/{user.id}/status", _UserAPI_UpdateUserStatus0_HTTP_Handler(srv))
	r.PUT("/sys/users/{user.id}/roles", _UserAPI_UpdateUserRoles0_HTTP_Handler(srv))
	r.POST("/sys/users/password/reset", _UserAPI_ResetUserPassword0_HTTP_Handler(srv))
}

func _UserAPI_ListUsers0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListUsersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIListUsers)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListUsers(ctx, req.(*ListUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUsersResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_GetUser0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_CreateUser0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPICreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_UpdateUser0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_DeleteUser0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_UpdateUserStatus0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserStatusRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIUpdateUserStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserStatusResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_UpdateUserRoles0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRolesRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIUpdateUserRoles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserRoles(ctx, req.(*UpdateUserRolesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserRolesResponse)
		return ctx.Result(200, reply)
	}
}

func _UserAPI_ResetUserPassword0_HTTP_Handler(srv UserAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetUserPasswordRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIResetUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResetUserPassword(ctx, req.(*ResetUserPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResetUserPasswordResponse)
		return ctx.Result(200, reply)
	}
}

type UserAPIHTTPClient interface {
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserResponse, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserResponse, err error)
	ListUsers(ctx context.Context, req *ListUsersRequest, opts ...http.CallOption) (rsp *ListUsersResponse, err error)
	ResetUserPassword(ctx context.Context, req *ResetUserPasswordRequest, opts ...http.CallOption) (rsp *ResetUserPasswordResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserResponse, err error)
	UpdateUserRoles(ctx context.Context, req *UpdateUserRolesRequest, opts ...http.CallOption) (rsp *UpdateUserRolesResponse, err error)
	UpdateUserStatus(ctx context.Context, req *UpdateUserStatusRequest, opts ...http.CallOption) (rsp *UpdateUserStatusResponse, err error)
}

type UserAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewUserAPIHTTPClient(client *http.Client) UserAPIHTTPClient {
	return &UserAPIHTTPClientImpl{client}
}

func (c *UserAPIHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserResponse, error) {
	var out CreateUserResponse
	pattern := "/sys/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPICreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserResponse, error) {
	var out DeleteUserResponse
	pattern := "/sys/users/{user.id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserAPIDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserResponse, error) {
	var out GetUserResponse
	pattern := "/sys/users/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserAPIGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...http.CallOption) (*ListUsersResponse, error) {
	var out ListUsersResponse
	pattern := "/sys/users"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserAPIListUsers))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) ResetUserPassword(ctx context.Context, in *ResetUserPasswordRequest, opts ...http.CallOption) (*ResetUserPasswordResponse, error) {
	var out ResetUserPasswordResponse
	pattern := "/sys/users/password/reset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPIResetUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserResponse, error) {
	var out UpdateUserResponse
	pattern := "/sys/users/{user.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPIUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) UpdateUserRoles(ctx context.Context, in *UpdateUserRolesRequest, opts ...http.CallOption) (*UpdateUserRolesResponse, error) {
	var out UpdateUserRolesResponse
	pattern := "/sys/users/{user.id}/roles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPIUpdateUserRoles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserAPIHTTPClientImpl) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...http.CallOption) (*UpdateUserStatusResponse, error) {
	var out UpdateUserStatusResponse
	pattern := "/sys/users/{user.id}/status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserAPIUpdateUserStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.User, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
