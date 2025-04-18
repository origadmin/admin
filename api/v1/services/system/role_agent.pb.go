// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/role.proto

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

type RoleServiceAgent interface {
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*DeleteRoleResponse, error)
	GetRole(context.Context, *GetRoleRequest) (*GetRoleResponse, error)
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*UpdateRoleResponse, error)
}

func _RoleService_ListRoles0_HTTPAgent_Handler(srv RoleServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListRolesRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationRoleServiceListRoles)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListRoles(ctx, req.(*ListRolesRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListRolesResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _RoleService_GetRole0_HTTPAgent_Handler(srv RoleServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in GetRoleRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationRoleServiceGetRole)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.GetRole(ctx, req.(*GetRoleRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoleResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _RoleService_CreateRole0_HTTPAgent_Handler(srv RoleServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in CreateRoleRequest
		if err := cctx.Bind(&in.Role); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationRoleServiceCreateRole)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.CreateRole(ctx, req.(*CreateRoleRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateRoleResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _RoleService_UpdateRole0_HTTPAgent_Handler(srv RoleServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdateRoleRequest
		if err := cctx.Bind(&in.Role); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationRoleServiceUpdateRole)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdateRole(ctx, req.(*UpdateRoleRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateRoleResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _RoleService_DeleteRole0_HTTPAgent_Handler(srv RoleServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in DeleteRoleRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationRoleServiceDeleteRole)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.DeleteRole(ctx, req.(*DeleteRoleRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteRoleResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func RegisterRoleServiceAgent(ag agent.HTTPAgent, srv RoleServiceAgent) {
	r := ag.Route()
	r.GET("/sys/roles", _RoleService_ListRoles0_HTTPAgent_Handler(srv))
	r.GET("/sys/roles/{id}", _RoleService_GetRole0_HTTPAgent_Handler(srv))
	r.POST("/sys/roles", _RoleService_CreateRole0_HTTPAgent_Handler(srv))
	r.PUT("/sys/roles/{role.id}", _RoleService_UpdateRole0_HTTPAgent_Handler(srv))
	r.DELETE("/sys/roles/{id}", _RoleService_DeleteRole0_HTTPAgent_Handler(srv))
}
