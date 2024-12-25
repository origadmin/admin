// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/resource.proto

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

type ResourceAPIAgent interface {
	CreateResource(http.Context, *CreateResourceRequest) (*CreateResourceResponse, error)
	DeleteResource(http.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)
	GetResource(http.Context, *GetResourceRequest) (*GetResourceResponse, error)
	ListResources(http.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
	UpdateResource(http.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
}

func _ResourceAPI_ListResources0_HTTPAgent_Handler(srv ResourceAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in ListResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceAPIListResources)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.ListResources(ctx, req.(*ListResourcesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResourcesResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _ResourceAPI_GetResource0_HTTPAgent_Handler(srv ResourceAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in GetResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceAPIGetResource)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.GetResource(ctx, req.(*GetResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetResourceResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _ResourceAPI_CreateResource0_HTTPAgent_Handler(srv ResourceAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CreateResourceRequest
		if err := ctx.Bind(&in.Resource); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceAPICreateResource)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CreateResource(ctx, req.(*CreateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResourceResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _ResourceAPI_UpdateResource0_HTTPAgent_Handler(srv ResourceAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in UpdateResourceRequest
		if err := ctx.Bind(&in.Resource); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceAPIUpdateResource)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateResource(ctx, req.(*UpdateResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResourceResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _ResourceAPI_DeleteResource0_HTTPAgent_Handler(srv ResourceAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in DeleteResourceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationResourceAPIDeleteResource)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteResource(ctx, req.(*DeleteResourceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResourceResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func RegisterResourceAPIAgent(ag agent.HTTPAgent, srv ResourceAPIAgent) {
	r := ag.Route()
	r.GET("/sys/resources", _ResourceAPI_ListResources0_HTTPAgent_Handler(srv))
	r.GET("/sys/resources/:id", _ResourceAPI_GetResource0_HTTPAgent_Handler(srv))
	r.POST("/sys/resources", _ResourceAPI_CreateResource0_HTTPAgent_Handler(srv))
	r.PUT("/sys/resources/:resource.id", _ResourceAPI_UpdateResource0_HTTPAgent_Handler(srv))
	r.DELETE("/sys/resources/:id", _ResourceAPI_DeleteResource0_HTTPAgent_Handler(srv))
}
