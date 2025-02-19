// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/department.proto

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

type DepartmentServiceAgent interface {
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*CreateDepartmentResponse, error)
	DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentResponse, error)
	GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentResponse, error)
	ListDepartments(context.Context, *ListDepartmentsRequest) (*ListDepartmentsResponse, error)
	UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*UpdateDepartmentResponse, error)
}

func _DepartmentService_ListDepartments0_HTTPAgent_Handler(srv DepartmentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListDepartmentsRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationDepartmentServiceListDepartments)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListDepartments(ctx, req.(*ListDepartmentsRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDepartmentsResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _DepartmentService_GetDepartment0_HTTPAgent_Handler(srv DepartmentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in GetDepartmentRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationDepartmentServiceGetDepartment)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.GetDepartment(ctx, req.(*GetDepartmentRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDepartmentResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _DepartmentService_CreateDepartment0_HTTPAgent_Handler(srv DepartmentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in CreateDepartmentRequest
		if err := cctx.Bind(&in.Department); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationDepartmentServiceCreateDepartment)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.CreateDepartment(ctx, req.(*CreateDepartmentRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDepartmentResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _DepartmentService_UpdateDepartment0_HTTPAgent_Handler(srv DepartmentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdateDepartmentRequest
		if err := cctx.Bind(&in.Department); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationDepartmentServiceUpdateDepartment)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdateDepartment(ctx, req.(*UpdateDepartmentRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDepartmentResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _DepartmentService_DeleteDepartment0_HTTPAgent_Handler(srv DepartmentServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in DeleteDepartmentRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationDepartmentServiceDeleteDepartment)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.DeleteDepartment(ctx, req.(*DeleteDepartmentRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDepartmentResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func RegisterDepartmentServiceAgent(ag agent.HTTPAgent, srv DepartmentServiceAgent) {
	r := ag.Route()
	r.GET("/sys/departments", _DepartmentService_ListDepartments0_HTTPAgent_Handler(srv))
	r.GET("/sys/departments/{id}", _DepartmentService_GetDepartment0_HTTPAgent_Handler(srv))
	r.POST("/sys/departments", _DepartmentService_CreateDepartment0_HTTPAgent_Handler(srv))
	r.PUT("/sys/departments/{department.id}", _DepartmentService_UpdateDepartment0_HTTPAgent_Handler(srv))
	r.DELETE("/sys/departments/{id}", _DepartmentService_DeleteDepartment0_HTTPAgent_Handler(srv))
}
