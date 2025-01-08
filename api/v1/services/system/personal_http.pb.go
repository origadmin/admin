// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: system/personal.proto

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

const OperationPersonalServiceGetPersonalProfile = "/api.v1.services.system.PersonalService/GetPersonalProfile"
const OperationPersonalServiceListPersonalResources = "/api.v1.services.system.PersonalService/ListPersonalResources"
const OperationPersonalServiceListPersonalRoles = "/api.v1.services.system.PersonalService/ListPersonalRoles"
const OperationPersonalServicePersonalLogout = "/api.v1.services.system.PersonalService/PersonalLogout"
const OperationPersonalServiceRefreshPersonalToken = "/api.v1.services.system.PersonalService/RefreshPersonalToken"
const OperationPersonalServiceUpdatePersonalPassword = "/api.v1.services.system.PersonalService/UpdatePersonalPassword"
const OperationPersonalServiceUpdatePersonalProfile = "/api.v1.services.system.PersonalService/UpdatePersonalProfile"
const OperationPersonalServiceUpdatePersonalSetting = "/api.v1.services.system.PersonalService/UpdatePersonalSetting"

type PersonalServiceHTTPServer interface {
	// GetPersonalProfile GetPersonalProfile Update the personal user information
	GetPersonalProfile(context.Context, *GetPersonalProfileRequest) (*GetPersonalProfileResponse, error)
	// ListPersonalResources ListPersonalResources List the personal user's menu
	ListPersonalResources(context.Context, *ListPersonalResourcesRequest) (*ListPersonalResourcesResponse, error)
	// ListPersonalRoles ListPersonalResources List the personal user's menu
	ListPersonalRoles(context.Context, *ListPersonalRolesRequest) (*ListPersonalRolesResponse, error)
	PersonalLogout(context.Context, *PersonalLogoutRequest) (*PersonalLogoutResponse, error)
	// RefreshPersonalToken RefreshPersonalToken Refresh the personal user's token
	RefreshPersonalToken(context.Context, *RefreshPersonalTokenRequest) (*RefreshPersonalTokenResponse, error)
	// UpdatePersonalPassword UpdatePersonalProfilePassword The user changes the password
	UpdatePersonalPassword(context.Context, *UpdatePersonalPasswordRequest) (*UpdatePersonalPasswordResponse, error)
	// UpdatePersonalProfile UpdatePersonalProfile Update the personal user information
	UpdatePersonalProfile(context.Context, *UpdatePersonalProfileRequest) (*UpdatePersonalProfileResponse, error)
	// UpdatePersonalSetting UpdatePersonalSetting User settings are saved
	UpdatePersonalSetting(context.Context, *UpdatePersonalSettingRequest) (*UpdatePersonalSettingResponse, error)
}

func RegisterPersonalServiceHTTPServer(s *http.Server, srv PersonalServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/sys/personal/logout", _PersonalService_PersonalLogout0_HTTP_Handler(srv))
	r.PUT("/sys/personal/password", _PersonalService_UpdatePersonalPassword0_HTTP_Handler(srv))
	r.PUT("/sys/personal/profile", _PersonalService_UpdatePersonalProfile0_HTTP_Handler(srv))
	r.GET("/sys/personal/profile", _PersonalService_GetPersonalProfile0_HTTP_Handler(srv))
	r.GET("/sys/personal/menus", _PersonalService_ListPersonalResources0_HTTP_Handler(srv))
	r.GET("/sys/personal/roles", _PersonalService_ListPersonalRoles0_HTTP_Handler(srv))
	r.PUT("/sys/personal/setting", _PersonalService_UpdatePersonalSetting0_HTTP_Handler(srv))
	r.POST("/sys/personal/token/refresh", _PersonalService_RefreshPersonalToken0_HTTP_Handler(srv))
}

func _PersonalService_PersonalLogout0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PersonalLogoutRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServicePersonalLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PersonalLogout(ctx, req.(*PersonalLogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PersonalLogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_UpdatePersonalPassword0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePersonalPasswordRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceUpdatePersonalPassword)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePersonalPassword(ctx, req.(*UpdatePersonalPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePersonalPasswordResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_UpdatePersonalProfile0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePersonalProfileRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceUpdatePersonalProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePersonalProfile(ctx, req.(*UpdatePersonalProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePersonalProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_GetPersonalProfile0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPersonalProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceGetPersonalProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPersonalProfile(ctx, req.(*GetPersonalProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPersonalProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_ListPersonalResources0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPersonalResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceListPersonalResources)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPersonalResources(ctx, req.(*ListPersonalResourcesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPersonalResourcesResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_ListPersonalRoles0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPersonalRolesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceListPersonalRoles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPersonalRoles(ctx, req.(*ListPersonalRolesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPersonalRolesResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_UpdatePersonalSetting0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePersonalSettingRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceUpdatePersonalSetting)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePersonalSetting(ctx, req.(*UpdatePersonalSettingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePersonalSettingResponse)
		return ctx.Result(200, reply)
	}
}

func _PersonalService_RefreshPersonalToken0_HTTP_Handler(srv PersonalServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshPersonalTokenRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPersonalServiceRefreshPersonalToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshPersonalToken(ctx, req.(*RefreshPersonalTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshPersonalTokenResponse)
		return ctx.Result(200, reply)
	}
}

type PersonalServiceHTTPClient interface {
	GetPersonalProfile(ctx context.Context, req *GetPersonalProfileRequest, opts ...http.CallOption) (rsp *GetPersonalProfileResponse, err error)
	ListPersonalResources(ctx context.Context, req *ListPersonalResourcesRequest, opts ...http.CallOption) (rsp *ListPersonalResourcesResponse, err error)
	ListPersonalRoles(ctx context.Context, req *ListPersonalRolesRequest, opts ...http.CallOption) (rsp *ListPersonalRolesResponse, err error)
	PersonalLogout(ctx context.Context, req *PersonalLogoutRequest, opts ...http.CallOption) (rsp *PersonalLogoutResponse, err error)
	RefreshPersonalToken(ctx context.Context, req *RefreshPersonalTokenRequest, opts ...http.CallOption) (rsp *RefreshPersonalTokenResponse, err error)
	UpdatePersonalPassword(ctx context.Context, req *UpdatePersonalPasswordRequest, opts ...http.CallOption) (rsp *UpdatePersonalPasswordResponse, err error)
	UpdatePersonalProfile(ctx context.Context, req *UpdatePersonalProfileRequest, opts ...http.CallOption) (rsp *UpdatePersonalProfileResponse, err error)
	UpdatePersonalSetting(ctx context.Context, req *UpdatePersonalSettingRequest, opts ...http.CallOption) (rsp *UpdatePersonalSettingResponse, err error)
}

type PersonalServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPersonalServiceHTTPClient(client *http.Client) PersonalServiceHTTPClient {
	return &PersonalServiceHTTPClientImpl{client}
}

func (c *PersonalServiceHTTPClientImpl) GetPersonalProfile(ctx context.Context, in *GetPersonalProfileRequest, opts ...http.CallOption) (*GetPersonalProfileResponse, error) {
	var out GetPersonalProfileResponse
	pattern := "/sys/personal/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonalServiceGetPersonalProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) ListPersonalResources(ctx context.Context, in *ListPersonalResourcesRequest, opts ...http.CallOption) (*ListPersonalResourcesResponse, error) {
	var out ListPersonalResourcesResponse
	pattern := "/sys/personal/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonalServiceListPersonalResources))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) ListPersonalRoles(ctx context.Context, in *ListPersonalRolesRequest, opts ...http.CallOption) (*ListPersonalRolesResponse, error) {
	var out ListPersonalRolesResponse
	pattern := "/sys/personal/roles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPersonalServiceListPersonalRoles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) PersonalLogout(ctx context.Context, in *PersonalLogoutRequest, opts ...http.CallOption) (*PersonalLogoutResponse, error) {
	var out PersonalLogoutResponse
	pattern := "/sys/personal/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonalServicePersonalLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) RefreshPersonalToken(ctx context.Context, in *RefreshPersonalTokenRequest, opts ...http.CallOption) (*RefreshPersonalTokenResponse, error) {
	var out RefreshPersonalTokenResponse
	pattern := "/sys/personal/token/refresh"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonalServiceRefreshPersonalToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) UpdatePersonalPassword(ctx context.Context, in *UpdatePersonalPasswordRequest, opts ...http.CallOption) (*UpdatePersonalPasswordResponse, error) {
	var out UpdatePersonalPasswordResponse
	pattern := "/sys/personal/password"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonalServiceUpdatePersonalPassword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) UpdatePersonalProfile(ctx context.Context, in *UpdatePersonalProfileRequest, opts ...http.CallOption) (*UpdatePersonalProfileResponse, error) {
	var out UpdatePersonalProfileResponse
	pattern := "/sys/personal/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonalServiceUpdatePersonalProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PersonalServiceHTTPClientImpl) UpdatePersonalSetting(ctx context.Context, in *UpdatePersonalSettingRequest, opts ...http.CallOption) (*UpdatePersonalSettingResponse, error) {
	var out UpdatePersonalSettingResponse
	pattern := "/sys/personal/setting"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPersonalServiceUpdatePersonalSetting))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
