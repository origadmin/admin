// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/current.proto

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

type CurrentServiceAgent interface {
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

func _CurrentService_CurrentLogout0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in CurrentLogoutRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceCurrentLogout)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.CurrentLogout(ctx, req.(*CurrentLogoutRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CurrentLogoutResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentUserPassword0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdateCurrentUserPasswordRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceUpdateCurrentUserPassword)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdateCurrentUserPassword(ctx, req.(*UpdateCurrentUserPasswordRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentUserPasswordResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentUser0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdateCurrentUserRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceUpdateCurrentUser)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentUserResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_GetCurrentUser0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in GetCurrentUserRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceGetCurrentUser)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCurrentUserResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_ListCurrentResources0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListCurrentResourcesRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceListCurrentResources)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListCurrentResources(ctx, req.(*ListCurrentResourcesRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCurrentResourcesResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_ListCurrentRoles0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListCurrentRolesRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceListCurrentRoles)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListCurrentRoles(ctx, req.(*ListCurrentRolesRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListCurrentRolesResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_UpdateCurrentSetting0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdateCurrentSettingRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceUpdateCurrentSetting)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdateCurrentSetting(ctx, req.(*UpdateCurrentSettingRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateCurrentSettingResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _CurrentService_RefreshCurrentToken0_HTTPAgent_Handler(srv CurrentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in RefreshCurrentTokenRequest
		if err := cctx.Bind(&in.Data); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationCurrentServiceRefreshCurrentToken)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.RefreshCurrentToken(ctx, req.(*RefreshCurrentTokenRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshCurrentTokenResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func RegisterCurrentServiceAgent(ag agent.HTTPAgent, srv CurrentServiceAgent) {
	r := ag.Route()
	r.POST("/sys/current/logout", _CurrentService_CurrentLogout0_HTTPAgent_Handler(srv))
	r.PUT("/sys/current/user/password", _CurrentService_UpdateCurrentUserPassword0_HTTPAgent_Handler(srv))
	r.PUT("/sys/current/user", _CurrentService_UpdateCurrentUser0_HTTPAgent_Handler(srv))
	r.GET("/sys/current/user", _CurrentService_GetCurrentUser0_HTTPAgent_Handler(srv))
	r.GET("/sys/current/menus", _CurrentService_ListCurrentResources0_HTTPAgent_Handler(srv))
	r.GET("/sys/current/roles", _CurrentService_ListCurrentRoles0_HTTPAgent_Handler(srv))
	r.PUT("/sys/current/setting", _CurrentService_UpdateCurrentSetting0_HTTPAgent_Handler(srv))
	r.POST("/sys/current/token/refresh", _CurrentService_RefreshCurrentToken0_HTTPAgent_Handler(srv))
}
