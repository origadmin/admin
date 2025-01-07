// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/position.proto

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

type PositionServiceAgent interface {
	CreatePosition(context.Context, *CreatePositionRequest) (*CreatePositionResponse, error)
	DeletePosition(context.Context, *DeletePositionRequest) (*DeletePositionResponse, error)
	GetPosition(context.Context, *GetPositionRequest) (*GetPositionResponse, error)
	ListPositions(context.Context, *ListPositionsRequest) (*ListPositionsResponse, error)
	UpdatePosition(context.Context, *UpdatePositionRequest) (*UpdatePositionResponse, error)
}

func _PositionService_ListPositions0_HTTPAgent_Handler(srv PositionServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in ListPositionsRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationPositionServiceListPositions)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.ListPositions(ctx, req.(*ListPositionsRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPositionsResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _PositionService_GetPosition0_HTTPAgent_Handler(srv PositionServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in GetPositionRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationPositionServiceGetPosition)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.GetPosition(ctx, req.(*GetPositionRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPositionResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _PositionService_CreatePosition0_HTTPAgent_Handler(srv PositionServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in CreatePositionRequest
		if err := cctx.Bind(&in.Position); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationPositionServiceCreatePosition)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.CreatePosition(ctx, req.(*CreatePositionRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePositionResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _PositionService_UpdatePosition0_HTTPAgent_Handler(srv PositionServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in UpdatePositionRequest
		if err := cctx.Bind(&in.Position); err != nil {
			return err
		}
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationPositionServiceUpdatePosition)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.UpdatePosition(ctx, req.(*UpdatePositionRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePositionResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func _PositionService_DeletePosition0_HTTPAgent_Handler(srv PositionServiceAgent) http.HandlerFunc {
	return func(cctx http.Context) error {
		var in DeletePositionRequest
		if err := cctx.BindQuery(&in); err != nil {
			return err
		}
		if err := cctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(cctx, OperationPositionServiceDeletePosition)
		h := cctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = agent.NewHTTPContext(ctx, cctx)
			return srv.DeletePosition(ctx, req.(*DeletePositionRequest))
		})
		out, err := h(cctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePositionResponse)
		if reply == nil {
			return nil
		}
		return cctx.Result(200, reply)
	}
}

func RegisterPositionServiceAgent(ag agent.HTTPAgent, srv PositionServiceAgent) {
	r := ag.Route()
	r.GET("/sys/positions", _PositionService_ListPositions0_HTTPAgent_Handler(srv))
	r.GET("/sys/positions/:id", _PositionService_GetPosition0_HTTPAgent_Handler(srv))
	r.POST("/sys/positions", _PositionService_CreatePosition0_HTTPAgent_Handler(srv))
	r.PUT("/sys/positions/:position.id", _PositionService_UpdatePosition0_HTTPAgent_Handler(srv))
	r.DELETE("/sys/positions/:id", _PositionService_DeletePosition0_HTTPAgent_Handler(srv))
}
