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
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	// CreateToken CreateToken generates a new JWT token for the given user.
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	// DestroyToken DestroyToken invalidates a JWT token.
	DestroyToken(context.Context, *DestroyTokenRequest) (*DestroyTokenResponse, error)
	ListAuthResources(context.Context, *ListAuthResourcesRequest) (*ListAuthResourcesResponse, error)
	// ValidateToken ValidateToken verifies the validity of a JWT token.
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
}

func _AuthAPI_ListAuthResources0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListAuthResourcesRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationAuthAPIListAuthResources)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListAuthResources(ctx, req.(*ListAuthResourcesRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListAuthResourcesResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _AuthAPI_CreateToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in CreateTokenRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationAuthAPICreateToken)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.CreateToken(ctx, req.(*CreateTokenRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTokenResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _AuthAPI_ValidateToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ValidateTokenRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationAuthAPIValidateToken)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ValidateToken(ctx, req.(*ValidateTokenRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValidateTokenResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _AuthAPI_DestroyToken0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in DestroyTokenRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationAuthAPIDestroyToken)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.DestroyToken(ctx, req.(*DestroyTokenRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DestroyTokenResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _AuthAPI_Authenticate0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in AuthenticateRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationAuthAPIAuthenticate)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.Authenticate(ctx, req.(*AuthenticateRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthenticateResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
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
