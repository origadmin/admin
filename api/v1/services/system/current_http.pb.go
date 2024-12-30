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

const OperationCurrentAPICurrentLogout = "/api.v1.services.system.CurrentAPI/CurrentLogout"
const OperationCurrentAPIGetCurrentUser = "/api.v1.services.system.CurrentAPI/GetCurrentUser"
const OperationCurrentAPIListCurrentMenus = "/api.v1.services.system.CurrentAPI/ListCurrentMenus"
const OperationCurrentAPIListCurrentRoles = "/api.v1.services.system.CurrentAPI/ListCurrentRoles"
const OperationCurrentAPIUpdateCurrentSetting = "/api.v1.services.system.CurrentAPI/UpdateCurrentSetting"
const OperationCurrentAPIUpdateCurrentUser = "/api.v1.services.system.CurrentAPI/UpdateCurrentUser"
const OperationCurrentAPIUpdateCurrentUserPassword = "/api.v1.services.system.CurrentAPI/UpdateCurrentUserPassword"

type CurrentAPIHTTPServer interface {
	CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error)
	// GetCurrentUser GetCurrentUser Update the current user information
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error)
	// ListCurrentMenus ListCurrentMenus List the current user's menu
	ListCurrentMenus(context.Context, *ListCurrentMenusRequest) (*ListCurrentMenusResponse, error)
	// ListCurrentRoles ListCurrentMenus List the current user's menu
	ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error)
	// UpdateCurrentSetting UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error)
	// UpdateCurrentUser UpdateCurrentUser Update the current user information
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error)
	// UpdateCurrentUserPassword UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error)
}

func RegisterCurrentAPIHTTPServer(s *http.Server, srv CurrentAPIHTTPServer) {
	r := s.Route("/")
	r.POST("/sys/current/logout", _CurrentAPI_CurrentLogout0_HTTP_Handler(srv))
	r.PUT("/sys/current/user/password", _CurrentAPI_UpdateCurrentUserPassword0_HTTP_Handler(srv))
	r.PUT("/sys/current/user", _CurrentAPI_UpdateCurrentUser0_HTTP_Handler(srv))
	r.GET("/sys/current/user", _CurrentAPI_GetCurrentUser0_HTTP_Handler(srv))
	r.GET("/sys/current/menus", _CurrentAPI_ListCurrentMenus0_HTTP_Handler(srv))
	r.GET("/sys/current/roles", _CurrentAPI_ListCurrentRoles0_HTTP_Handler(srv))
	r.PUT("/sys/current/setting", _CurrentAPI_UpdateCurrentSetting0_HTTP_Handler(srv))
}

func _CurrentAPI_CurrentLogout0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrentLogoutRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPICurrentLogout)
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

func _CurrentAPI_UpdateCurrentUserPassword0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentUserPasswordRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIUpdateCurrentUserPassword)
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

func _CurrentAPI_UpdateCurrentUser0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentUserRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIUpdateCurrentUser)
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

func _CurrentAPI_GetCurrentUser0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIGetCurrentUser)
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

func _CurrentAPI_ListCurrentMenus0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCurrentMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIListCurrentMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCurrentMenus(ctx, req.(*ListCurrentMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCurrentMenusResponse)
		return ctx.Result(200, reply)
	}
}

func _CurrentAPI_ListCurrentRoles0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCurrentRolesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIListCurrentRoles)
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

func _CurrentAPI_UpdateCurrentSetting0_HTTP_Handler(srv CurrentAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCurrentSettingRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCurrentAPIUpdateCurrentSetting)
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

type CurrentAPIHTTPClient interface {
	CurrentLogout(ctx context.Context, req *CurrentLogoutRequest, opts ...http.CallOption) (rsp *CurrentLogoutResponse, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *GetCurrentUserResponse, err error)
	ListCurrentMenus(ctx context.Context, req *ListCurrentMenusRequest, opts ...http.CallOption) (rsp *ListCurrentMenusResponse, err error)
	ListCurrentRoles(ctx context.Context, req *ListCurrentRolesRequest, opts ...http.CallOption) (rsp *ListCurrentRolesResponse, err error)
	UpdateCurrentSetting(ctx context.Context, req *UpdateCurrentSettingRequest, opts ...http.CallOption) (rsp *UpdateCurrentSettingResponse, err error)
	UpdateCurrentUser(ctx context.Context, req *UpdateCurrentUserRequest, opts ...http.CallOption) (rsp *UpdateCurrentUserResponse, err error)
	UpdateCurrentUserPassword(ctx context.Context, req *UpdateCurrentUserPasswordRequest, opts ...http.CallOption) (rsp *UpdateCurrentUserPasswordResponse, err error)
}

type CurrentAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewCurrentAPIHTTPClient(client *http.Client) CurrentAPIHTTPClient {
	return &CurrentAPIHTTPClientImpl{client}
}

func (c *CurrentAPIHTTPClientImpl) CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...http.CallOption) (*CurrentLogoutResponse, error) {
	var out CurrentLogoutResponse
	pattern := "/sys/current/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentAPICurrentLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*GetCurrentUserResponse, error) {
	var out GetCurrentUserResponse
	pattern := "/sys/current/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentAPIGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) ListCurrentMenus(ctx context.Context, in *ListCurrentMenusRequest, opts ...http.CallOption) (*ListCurrentMenusResponse, error) {
	var out ListCurrentMenusResponse
	pattern := "/sys/current/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentAPIListCurrentMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...http.CallOption) (*ListCurrentRolesResponse, error) {
	var out ListCurrentRolesResponse
	pattern := "/sys/current/roles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCurrentAPIListCurrentRoles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...http.CallOption) (*UpdateCurrentSettingResponse, error) {
	var out UpdateCurrentSettingResponse
	pattern := "/sys/current/setting"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentAPIUpdateCurrentSetting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...http.CallOption) (*UpdateCurrentUserResponse, error) {
	var out UpdateCurrentUserResponse
	pattern := "/sys/current/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentAPIUpdateCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CurrentAPIHTTPClientImpl) UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...http.CallOption) (*UpdateCurrentUserPasswordResponse, error) {
	var out UpdateCurrentUserPasswordResponse
	pattern := "/sys/current/user/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCurrentAPIUpdateCurrentUserPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}