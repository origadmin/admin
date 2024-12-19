// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
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

type LoginAPIAgent interface {
	CaptchaID(http.Context, *CaptchaIDRequest) (*CaptchaIDResponse, error)
	CaptchaImage(http.Context, *CaptchaImageRequest) (*CaptchaImageResponse, error)
	CaptchaResource(http.Context, *CaptchaResourceRequest) (*CaptchaResourceResponse, error)
	CaptchaResources(http.Context, *CaptchaResourcesRequest) (*CaptchaResourcesResponse, error)
	CurrentMenus(http.Context, *CurrentMenusRequest) (*CurrentMenusResponse, error)
	CurrentUser(http.Context, *CurrentUserRequest) (*CurrentUserResponse, error)
	Login(http.Context, *LoginRequest) (*LoginResponse, error)
	Logout(http.Context, *LogoutRequest) (*LogoutResponse, error)
	Refresh(http.Context, *RefreshRequest) (*RefreshResponse, error)
	Register(http.Context, *RegisterRequest) (*RegisterResponse, error)
}

func RegisterLoginAPIAgent(s *http.Server, srv LoginAPIAgent) {
	r := s.Route("/")
	r.GET("/api/v1/captcha/id", _LoginAPI_CaptchaID0_Agent_Handler(srv))
	r.GET("/api/v1/captcha/image", _LoginAPI_CaptchaImage0_Agent_Handler(srv))
	r.GET("/api/v1/captcha/id/:id/:resource", _LoginAPI_CaptchaResource0_Agent_Handler(srv))
	r.GET("/api/v1/captcha/id/:id", _LoginAPI_CaptchaResources0_Agent_Handler(srv))
	r.POST("/api/v1/login", _LoginAPI_Login0_Agent_Handler(srv))
	r.POST("/api/v1/register", _LoginAPI_Register0_Agent_Handler(srv))
	r.POST("/api/v1/refresh_token", _LoginAPI_Refresh0_Agent_Handler(srv))
	r.POST("/api/v1/current/logout", _LoginAPI_Logout0_Agent_Handler(srv))
	r.POST("/api/v1/current/user", _LoginAPI_CurrentUser0_Agent_Handler(srv))
	r.GET("/api/v1/current/menus", _LoginAPI_CurrentMenus0_Agent_Handler(srv))
}

func _LoginAPI_CaptchaID0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CaptchaIDRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaID)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaID(ctx, req.(*CaptchaIDRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_CaptchaImage0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CaptchaImageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaImage)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaImage(ctx, req.(*CaptchaImageRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_CaptchaResource0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CaptchaResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaResource)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaResource(ctx, req.(*CaptchaResourceRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_CaptchaResources0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CaptchaResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICaptchaResources)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CaptchaResources(ctx, req.(*CaptchaResourcesRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_Login0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPILogin)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_Register0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPIRegister)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_Refresh0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in RefreshRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPIRefresh)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.Refresh(ctx, req.(*RefreshRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_Logout0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPILogout)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_CurrentUser0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CurrentUserRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICurrentUser)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentUser(ctx, req.(*CurrentUserRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}

func _LoginAPI_CurrentMenus0_Agent_Handler(srv LoginAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CurrentMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLoginAPICurrentMenus)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CurrentMenus(ctx, req.(*CurrentMenusRequest))
		})
		_, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return nil
	}
}
