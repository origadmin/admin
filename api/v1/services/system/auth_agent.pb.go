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
	ListResources(http.Context, *ListResourcesRequest) (*ListResourcesResponse, error)
}

func _AuthAPI_ListResources0_HTTPAgent_Handler(srv AuthAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in ListResourcesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthAPIListResources)
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

func RegisterAuthAPIAgent(ag agent.HTTPAgent, srv AuthAPIAgent) {
	r := ag.Route()
	r.GET("/sys/resources", _AuthAPI_ListResources0_HTTPAgent_Handler(srv))
}
