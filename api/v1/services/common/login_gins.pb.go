// Code generated by protoc-gen-go-gins. DO NOT EDIT.
// versions:
// - protoc-gen-go-gins 0.0.6
// - protoc             (unknown)
// source: common/login.proto

package common

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

const LoginAPI_CaptchaID_OperationName = "/api.v1.services.common.LoginAPI/CaptchaID"
const LoginAPI_CaptchaImage_OperationName = "/api.v1.services.common.LoginAPI/CaptchaImage"
const LoginAPI_CurrentMenus_OperationName = "/api.v1.services.common.LoginAPI/CurrentMenus"
const LoginAPI_CurrentUser_OperationName = "/api.v1.services.common.LoginAPI/CurrentUser"
const LoginAPI_Login_OperationName = "/api.v1.services.common.LoginAPI/Login"
const LoginAPI_Logout_OperationName = "/api.v1.services.common.LoginAPI/Logout"

type LoginAPIGINSServer interface {
	CaptchaID(context.Context, *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(context.Context, *CaptchaImageRequest) (*CaptchaImageResponse, error)
	CurrentMenus(context.Context, *CurrentMenusRequest) (*CurrentMenusResponse, error)
	CurrentUser(context.Context, *CurrentUserRequest) (*CurrentUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

func RegisterLoginAPIGINSServer(router gins.IRouter, srv LoginAPIGINSServer) {
	router.GET("/api/v1/captcha/id", _LoginAPI_CaptchaID0_GIN_Handler(srv))
	router.GET("/api/v1/captcha/image", _LoginAPI_CaptchaImage0_GIN_Handler(srv))
	router.POST("/api/v1/login", _LoginAPI_Login0_GIN_Handler(srv))
	router.POST("/api/v1/current/logout", _LoginAPI_Logout0_GIN_Handler(srv))
	router.POST("/api/v1/current/user", _LoginAPI_CurrentUser0_GIN_Handler(srv))
	router.GET("/api/v1/current/menus", _LoginAPI_CurrentMenus0_GIN_Handler(srv))
}

func _LoginAPI_CaptchaID0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in CaptchaIDRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_CaptchaID_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CaptchaID(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply)
		return
	}
}

func _LoginAPI_CaptchaImage0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in CaptchaImageRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_CaptchaImage_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CaptchaImage(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply.Data.CaptchaImg)
		return
	}
}

func _LoginAPI_Login0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in LoginRequest
		if err := gins.BindBody(ctx, &in.Data); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_Login_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.Login(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply)
		return
	}
}

func _LoginAPI_Logout0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in LogoutRequest
		if err := gins.BindBody(ctx, &in.Data); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_Logout_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.Logout(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply)
		return
	}
}

func _LoginAPI_CurrentUser0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in CurrentUserRequest
		if err := gins.BindBody(ctx, &in.Data); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_CurrentUser_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CurrentUser(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply)
		return
	}
}

func _LoginAPI_CurrentMenus0_GIN_Handler(srv LoginAPIGINSServer) func(ctx *gins.Context) {
	return func(ctx *gins.Context) {
		var in CurrentMenusRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.SetOperation(ctx, LoginAPI_CurrentMenus_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CurrentMenus(newCtx, &in)
		if err != nil {
			gins.ResultError(ctx, err)
			return
		}
		gins.ResultJSON(ctx, 200, reply)
		return
	}
}

type LoginAPIGINSClient interface {
	CaptchaID(ctx context.Context, req *CaptchaIDRequest, opts ...gins.CallOption) (rsp *CaptchaIDResponse, err error)
	CaptchaImage(ctx context.Context, req *CaptchaImageRequest, opts ...gins.CallOption) (rsp *CaptchaImageResponse, err error)
	CurrentMenus(ctx context.Context, req *CurrentMenusRequest, opts ...gins.CallOption) (rsp *CurrentMenusResponse, err error)
	CurrentUser(ctx context.Context, req *CurrentUserRequest, opts ...gins.CallOption) (rsp *CurrentUserResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...gins.CallOption) (rsp *LoginResponse, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...gins.CallOption) (rsp *LogoutResponse, err error)
}

type LoginAPIGINSClientImpl struct {
	cc *gins.Client
}

func NewLoginAPIGINSClient(client *gins.Client) LoginAPIGINSClient {
	return &LoginAPIGINSClientImpl{client}
}

func (c *LoginAPIGINSClientImpl) CaptchaID(ctx context.Context, in *CaptchaIDRequest, opts ...gins.CallOption) (*CaptchaIDResponse, error) {
	var out CaptchaIDResponse
	pattern := "/api/v1/captcha/id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(LoginAPI_CaptchaID_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIGINSClientImpl) CaptchaImage(ctx context.Context, in *CaptchaImageRequest, opts ...gins.CallOption) (*CaptchaImageResponse, error) {
	var out CaptchaImageResponse
	pattern := "/api/v1/captcha/image"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(LoginAPI_CaptchaImage_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out.Data.CaptchaImg, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIGINSClientImpl) CurrentMenus(ctx context.Context, in *CurrentMenusRequest, opts ...gins.CallOption) (*CurrentMenusResponse, error) {
	var out CurrentMenusResponse
	pattern := "/api/v1/current/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(LoginAPI_CurrentMenus_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIGINSClientImpl) CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...gins.CallOption) (*CurrentUserResponse, error) {
	var out CurrentUserResponse
	pattern := "/api/v1/current/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(LoginAPI_CurrentUser_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIGINSClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...gins.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(LoginAPI_Login_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIGINSClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...gins.CallOption) (*LogoutResponse, error) {
	var out LogoutResponse
	pattern := "/api/v1/current/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(LoginAPI_Logout_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}