// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: system/current.proto

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

const OperationCurrentServiceCurrentLogout = "/api.v1.services.system.CurrentService/CurrentLogout"
const OperationCurrentServiceGetCurrentUser = "/api.v1.services.system.CurrentService/GetCurrentUser"
const OperationCurrentServiceListCurrentResources = "/api.v1.services.system.CurrentService/ListCurrentResources"
const OperationCurrentServiceListCurrentRoles = "/api.v1.services.system.CurrentService/ListCurrentRoles"
const OperationCurrentServiceRefreshCurrentToken = "/api.v1.services.system.CurrentService/RefreshCurrentToken"
const OperationCurrentServiceUpdateCurrentSetting = "/api.v1.services.system.CurrentService/UpdateCurrentSetting"
const OperationCurrentServiceUpdateCurrentUser = "/api.v1.services.system.CurrentService/UpdateCurrentUser"
const OperationCurrentServiceUpdateCurrentUserPassword = "/api.v1.services.system.CurrentService/UpdateCurrentUserPassword"

type CurrentServiceHTTPServer interface {
	CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error)
	// GetCurrentUser GetCurrentUser Update the current user information
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error)
	// ListCurrentResources ListCurrentResources List the current user's menu
	ListCurrentResources(context.Context, *ListCurrentResourcesRequest) (*ListCurrentResourcesResponse, error)
	// ListCurrentRoles ListCurrentResources List the current user's menu
	ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error)
	// RefreshCurrentToken RefreshCurrentToken Refresh the current user's token
	RefreshCurrentToken(context.Context, *RefreshCurrentTokenRequest) (*RefreshCurrentTokenResponse, error)
	// UpdateCurrentSetting UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error)
	// UpdateCurrentUser UpdateCurrentUser Update the current user information
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error)
	// UpdateCurrentUserPassword UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error)
}

func RegisterCurrentServiceHTTPServer(s *http.Server, srv CurrentServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/sys/current/logout", _CurrentService_CurrentLogout0_HTTP_Handler(srv))
	r.PUT("/sys/current/user/password", _CurrentService_UpdateCurrentUserPassword0_HTTP_Handler(srv))
	r.PUT("/sys/current/user", _CurrentService_UpdateCurrentUser0_HTTP_Handler(srv))
	r.GET("/sys/current/user", _CurrentService_GetCurrentUser0_HTTP_Handler(srv))
	r.GET("/sys/current/menus", _CurrentService_ListCurrentResources0_HTTP_Handler(srv))
	r.GET("/sys/current/roles", _CurrentService_ListCurrentRoles0_HTTP_Handler(srv))
	r.PUT("/sys/current/setting", _CurrentService_UpdateCurrentSetting0_HTTP_Handler(srv))
	r.POST("/sys/current/token/refresh", _CurrentService_RefreshCurrentToken0_HTTP_Handler(srv))
}

func _CurrentService_CurrentLogout0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrentLogoutRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceCurrentLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentLogout(ctx, req.(*CurrentLogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentLogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentUserPassword0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentUserPasswordRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceUpdateCurrentUserPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCurrentUserPassword(ctx, req.(*UpdateCurrentUserPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentUserPasswordResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentUser0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentUserRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceUpdateCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentUserResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_GetCurrentUser0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceGetCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentUserResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_ListCurrentResources0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCurrentResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceListCurrentResources)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCurrentResources(ctx, req.(*ListCurrentResourcesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCurrentResourcesResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_ListCurrentRoles0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCurrentRolesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceListCurrentRoles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCurrentRoles(ctx, req.(*ListCurrentRolesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCurrentRolesResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentSetting0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentSettingRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceUpdateCurrentSetting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCurrentSetting(ctx, req.(*UpdateCurrentSettingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentSettingResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentService_RefreshCurrentToken0_HTTP_Handler(srv CurrentServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshCurrentTokenRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentServiceRefreshCurrentToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshCurrentToken(ctx, req.(*RefreshCurrentTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshCurrentTokenResponse)
		return ctx.Result(200, reply)
	}
}

type CurrentServiceHTTPClient interface {
	CurrentLogout(ctx context.Context, req *CurrentLogoutRequest, opts ...http.CallOption) (rsp *CurrentLogoutResponse, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *GetCurrentUserResponse, err error)
	ListCurrentResources(ctx context.Context, req *ListCurrentResourcesRequest, opts ...http.CallOption) (rsp *ListCurrentResourcesResponse, err error)
	ListCurrentRoles(ctx context.Context, req *ListCurrentRolesRequest, opts ...http.CallOption) (rsp *ListCurrentRolesResponse, err error)
	RefreshCurrentToken(ctx context.Context, req *RefreshCurrentTokenRequest, opts ...http.CallOption) (rsp *RefreshCurrentTokenResponse, err error)
	UpdateCurrentSetting(ctx context.Context, req *UpdateCurrentSettingRequest, opts ...http.CallOption) (rsp *UpdateCurrentSettingResponse, err error)
	UpdateCurrentUser(ctx context.Context, req *UpdateCurrentUserRequest, opts ...http.CallOption) (rsp *UpdateCurrentUserResponse, err error)
	UpdateCurrentUserPassword(ctx context.Context, req *UpdateCurrentUserPasswordRequest, opts ...http.CallOption) (rsp *UpdateCurrentUserPasswordResponse, err error)
}

type CurrentServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCurrentServiceHTTPClient(client *http.Client) CurrentServiceHTTPClient {
	return &CurrentServiceHTTPClientImpl{client}
}

func (c *CurrentServiceHTTPClientImpl) CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...http.CallOption) (*CurrentLogoutResponse, error) {
	var out CurrentLogoutResponse
	pattern := "/sys/current/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentServiceCurrentLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*GetCurrentUserResponse, error) {
	var out GetCurrentUserResponse
	pattern := "/sys/current/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentServiceGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) ListCurrentResources(ctx context.Context, in *ListCurrentResourcesRequest, opts ...http.CallOption) (*ListCurrentResourcesResponse, error) {
	var out ListCurrentResourcesResponse
	pattern := "/sys/current/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentServiceListCurrentResources))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...http.CallOption) (*ListCurrentRolesResponse, error) {
	var out ListCurrentRolesResponse
	pattern := "/sys/current/roles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentServiceListCurrentRoles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) RefreshCurrentToken(ctx context.Context, in *RefreshCurrentTokenRequest, opts ...http.CallOption) (*RefreshCurrentTokenResponse, error) {
	var out RefreshCurrentTokenResponse
	pattern := "/sys/current/token/refresh"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentServiceRefreshCurrentToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...http.CallOption) (*UpdateCurrentSettingResponse, error) {
	var out UpdateCurrentSettingResponse
	pattern := "/sys/current/setting"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentServiceUpdateCurrentSetting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...http.CallOption) (*UpdateCurrentUserResponse, error) {
	var out UpdateCurrentUserResponse
	pattern := "/sys/current/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentServiceUpdateCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentServiceHTTPClientImpl) UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...http.CallOption) (*UpdateCurrentUserPasswordResponse, error) {
	var out UpdateCurrentUserPasswordResponse
	pattern := "/sys/current/user/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentServiceUpdateCurrentUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
