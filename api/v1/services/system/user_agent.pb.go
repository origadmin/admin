// Code generated by protoc-gen-go-agent. DO NOT EDIT.
// versions:
// - protoc-gen-go-agent unknown
// - protoc             (unknown)
// source: system/user.proto

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

type UserAPIAgent interface {
	CreateUser(http.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(http.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetUser(http.Context, *GetUserRequest) (*GetUserResponse, error)
	ListUsers(http.Context, *ListUsersRequest) (*ListUsersResponse, error)
	UpdateUser(http.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func RegisterUserAPIAgent(ag agent.Agent, srv UserAPIAgent) {
	r := ag.Route()
	r.GET("/sys/users", _UserAPI_ListUsers0_Agent_Handler(srv))
	r.GET("/sys/users/:id", _UserAPI_GetUser0_Agent_Handler(srv))
	r.POST("/sys/users", _UserAPI_CreateUser0_Agent_Handler(srv))
	r.PUT("/sys/users/:user.id", _UserAPI_UpdateUser0_Agent_Handler(srv))
	r.DELETE("/sys/users/:id", _UserAPI_DeleteUser0_Agent_Handler(srv))
}

func _UserAPI_ListUsers0_Agent_Handler(srv UserAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in ListUsersRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIListUsers)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.ListUsers(ctx, req.(*ListUsersRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListUsersResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _UserAPI_GetUser0_Agent_Handler(srv UserAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIGetUser)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _UserAPI_CreateUser0_Agent_Handler(srv UserAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPICreateUser)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _UserAPI_UpdateUser0_Agent_Handler(srv UserAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in.User); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIUpdateUser)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}

func _UserAPI_DeleteUser0_Agent_Handler(srv UserAPIAgent) http.HandlerFunc {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserAPIDeleteUser)
		h := ctx.Middleware(func(_ context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserResponse)
		if reply == nil {
			return nil
		}
		return ctx.Result(200, reply)
	}
}
