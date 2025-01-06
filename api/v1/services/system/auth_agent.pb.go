// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/auth.proto

package system

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	agent "github.com/origadmin/runtime/agent"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1
const _ = agent.ApiVersionV1

type AuthAPIAgent interface {
	Authenticate(http.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// CreateToken CreateToken generates a new JWT token for the given user.
	CreateToken(http.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	// DestroyToken DestroyToken invalidates a JWT token.
	DestroyToken(http.Context, *DestroyTokenRequest) (*DestroyTokenResponse, error)
	ListAuthResources(http.Context, *ListAuthResourcesRequest) (*ListAuthResourcesResponse, error)
	// ValidateToken ValidateToken verifies the validity of a JWT token.
	ValidateToken(http.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

func _AuthAPI_ListAuthResources0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in ListAuthResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPIListAuthResources)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.ListAuthResources(ctx, req.(*ListAuthResourcesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAuthResourcesResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _AuthAPI_CreateToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CreateTokenRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPICreateToken)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CreateToken(ctx, req.(*CreateTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTokenResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _AuthAPI_ValidateToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in ValidateTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPIValidateToken)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.ValidateToken(ctx, req.(*ValidateTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValidateTokenResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _AuthAPI_DestroyToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in DestroyTokenRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPIDestroyToken)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.DestroyToken(ctx, req.(*DestroyTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DestroyTokenResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _AuthAPI_Authenticate0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in AuthenticateRequest
		if err := ctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPIAuthenticate)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.Authenticate(ctx, req.(*AuthenticateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthenticateResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func RegisterAuthAPIAgent(ag agent.HTTPAgent, srv AuthAPIAgent) {
	r := ag.Route()
	r.GET("/sys/auth/resources", _AuthAPI_ListAuthResources0_HTTPAgent_Handler(srv))
	r.POST("/sys/auth/token", _AuthAPI_CreateToken0_HTTPAgent_Handler(srv))
	r.GET("/sys/auth/validate", _AuthAPI_ValidateToken0_HTTPAgent_Handler(srv))
	r.POST("/sys/auth/destroy", _AuthAPI_DestroyToken0_HTTPAgent_Handler(srv))
	r.POST("/sys/auth/authenticate", _AuthAPI_Authenticate0_HTTPAgent_Handler(srv))
}
