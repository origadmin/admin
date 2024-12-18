// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: basis/login.proto

package basis

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

const OperationLoginAPICaptchaID = "/api.v1.services.basis.LoginAPI/CaptchaID"
const OperationLoginAPICaptchaImage = "/api.v1.services.basis.LoginAPI/CaptchaImage"
const OperationLoginAPICaptchaResource = "/api.v1.services.basis.LoginAPI/CaptchaResource"
const OperationLoginAPICaptchaResources = "/api.v1.services.basis.LoginAPI/CaptchaResources"
const OperationLoginAPICurrentMenus = "/api.v1.services.basis.LoginAPI/CurrentMenus"
const OperationLoginAPICurrentUser = "/api.v1.services.basis.LoginAPI/CurrentUser"
const OperationLoginAPILogin = "/api.v1.services.basis.LoginAPI/Login"
const OperationLoginAPILogout = "/api.v1.services.basis.LoginAPI/Logout"
const OperationLoginAPIRefresh = "/api.v1.services.basis.LoginAPI/Refresh"
const OperationLoginAPIRegister = "/api.v1.services.basis.LoginAPI/Register"

type LoginAPIHTTPServer interface {
	CaptchaID(context.Context, *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(context.Context, *CaptchaImageRequest) (*CaptchaImageResponse, error)
	CaptchaResource(context.Context, *CaptchaResourceRequest) (*CaptchaResourceResponse, error)
	CaptchaResources(context.Context, *CaptchaResourcesRequest) (*CaptchaResourcesResponse, error)
	CurrentMenus(context.Context, *CurrentMenusRequest) (*CurrentMenusResponse, error)
	CurrentUser(context.Context, *CurrentUserRequest) (*CurrentUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterLoginAPIHTTPServer(s *http.Server, srv LoginAPIHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/captcha/id", _LoginAPI_CaptchaID0_HTTP_Handler(srv))
	r.GET("/api/v1/captcha/image", _LoginAPI_CaptchaImage0_HTTP_Handler(srv))
	r.GET("/api/v1/captcha/id/{id}/{resource}", _LoginAPI_CaptchaResource0_HTTP_Handler(srv))
	r.GET("/api/v1/captcha/id/{id}", _LoginAPI_CaptchaResources0_HTTP_Handler(srv))
	r.POST("/api/v1/login", _LoginAPI_Login0_HTTP_Handler(srv))
	r.POST("/api/v1/register", _LoginAPI_Register0_HTTP_Handler(srv))
	r.POST("/api/v1/refresh_token", _LoginAPI_Refresh0_HTTP_Handler(srv))
	r.POST("/api/v1/current/logout", _LoginAPI_Logout0_HTTP_Handler(srv))
	r.POST("/api/v1/current/user", _LoginAPI_CurrentUser0_HTTP_Handler(srv))
	r.GET("/api/v1/current/menus", _LoginAPI_CurrentMenus0_HTTP_Handler(srv))
}

func _LoginAPI_CaptchaID0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaID)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaID(ctx, req.(*CaptchaIDRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaIDResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_CaptchaImage0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaImageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaImage)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaImage(ctx, req.(*CaptchaImageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaImageResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_CaptchaResource0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaResource)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaResource(ctx, req.(*CaptchaResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaResourceResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_CaptchaResources0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaResources)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaResources(ctx, req.(*CaptchaResourcesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaResourcesResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_Login0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPILogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_Register0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPIRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_Refresh0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPIRefresh)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Refresh(ctx, req.(*RefreshRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_Logout0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPILogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_CurrentUser0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrentUserRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentUser(ctx, req.(*CurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentUserResponse)
		return ctx.Result(200, reply)
	}
}

func _LoginAPI_CurrentMenus0_HTTP_Handler(srv LoginAPIHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CurrentMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICurrentMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentMenus(ctx, req.(*CurrentMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentMenusResponse)
		return ctx.Result(200, reply)
	}
}

type LoginAPIHTTPClient interface {
	CaptchaID(ctx context.Context, req *CaptchaIDRequest, opts ...http.CallOption) (rsp *CaptchaIDResponse, err error)
	CaptchaImage(ctx context.Context, req *CaptchaImageRequest, opts ...http.CallOption) (rsp *CaptchaImageResponse, err error)
	CaptchaResource(ctx context.Context, req *CaptchaResourceRequest, opts ...http.CallOption) (rsp *CaptchaResourceResponse, err error)
	CaptchaResources(ctx context.Context, req *CaptchaResourcesRequest, opts ...http.CallOption) (rsp *CaptchaResourcesResponse, err error)
	CurrentMenus(ctx context.Context, req *CurrentMenusRequest, opts ...http.CallOption) (rsp *CurrentMenusResponse, err error)
	CurrentUser(ctx context.Context, req *CurrentUserRequest, opts ...http.CallOption) (rsp *CurrentUserResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutResponse, err error)
	Refresh(ctx context.Context, req *RefreshRequest, opts ...http.CallOption) (rsp *RefreshResponse, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResponse, err error)
}

type LoginAPIHTTPClientImpl struct {
	cc *http.Client
}

func NewLoginAPIHTTPClient(client *http.Client) LoginAPIHTTPClient {
	return &LoginAPIHTTPClientImpl{client}
}

func (c *LoginAPIHTTPClientImpl) CaptchaID(ctx context.Context, in *CaptchaIDRequest, opts ...http.CallOption) (*CaptchaIDResponse, error) {
	var out CaptchaIDResponse
	pattern := "/api/v1/captcha/id"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginAPICaptchaID))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) CaptchaImage(ctx context.Context, in *CaptchaImageRequest, opts ...http.CallOption) (*CaptchaImageResponse, error) {
	var out CaptchaImageResponse
	pattern := "/api/v1/captcha/image"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginAPICaptchaImage))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) CaptchaResource(ctx context.Context, in *CaptchaResourceRequest, opts ...http.CallOption) (*CaptchaResourceResponse, error) {
	var out CaptchaResourceResponse
	pattern := "/api/v1/captcha/id/{id}/{resource}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginAPICaptchaResource))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) CaptchaResources(ctx context.Context, in *CaptchaResourcesRequest, opts ...http.CallOption) (*CaptchaResourcesResponse, error) {
	var out CaptchaResourcesResponse
	pattern := "/api/v1/captcha/id/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginAPICaptchaResources))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) CurrentMenus(ctx context.Context, in *CurrentMenusRequest, opts ...http.CallOption) (*CurrentMenusResponse, error) {
	var out CurrentMenusResponse
	pattern := "/api/v1/current/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLoginAPICurrentMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...http.CallOption) (*CurrentUserResponse, error) {
	var out CurrentUserResponse
	pattern := "/api/v1/current/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginAPICurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/api/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginAPILogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutResponse, error) {
	var out LogoutResponse
	pattern := "/api/v1/current/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginAPILogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) Refresh(ctx context.Context, in *RefreshRequest, opts ...http.CallOption) (*RefreshResponse, error) {
	var out RefreshResponse
	pattern := "/api/v1/refresh_token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginAPIRefresh))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *LoginAPIHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResponse, error) {
	var out RegisterResponse
	pattern := "/api/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLoginAPIRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Data, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
