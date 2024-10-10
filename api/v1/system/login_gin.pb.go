// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.0
// - protoc             v5.27.2
// source: system/login.proto

package system

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	gins "github.com/origadmin/toolkits/runtime/kratos/transport/gins"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(gin.H)
var _ = gins.NewContext
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const LoginService_Login_OperationName = "/api.v1.system.LoginService/Login"
const LoginService_Logout_OperationName = "/api.v1.system.LoginService/Logout"
const LoginService_Register_OperationName = "/api.v1.system.LoginService/Register"

type LoginServiceGINServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterLoginServiceGINServer(router gin.IRouter, srv LoginServiceGINServer) {
	router.POST("/register", _LoginService_Register0_GIN_Handler(srv))
	router.POST("/login", _LoginService_Login0_GIN_Handler(srv))
	router.POST("/logout", _LoginService_Logout0_GIN_Handler(srv))
}

func _LoginService_Register0_GIN_Handler(srv LoginServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in RegisterRequest
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, LoginService_Register_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.Register(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _LoginService_Login0_GIN_Handler(srv LoginServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in LoginRequest
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, LoginService_Login_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.Login(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _LoginService_Logout0_GIN_Handler(srv LoginServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in LogoutRequest
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, LoginService_Logout_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.Logout(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

type LoginServiceGINClient interface {
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type LoginServiceGINClientImpl struct {
	cc *http.Client
}

func NewLoginServiceGINClient(client *http.Client) LoginServiceGINClient {
	return &LoginServiceGINClientImpl{client}
}

func (c *LoginServiceGINClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(LoginService_Login_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginServiceGINClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(LoginService_Logout_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginServiceGINClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(LoginService_Register_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
