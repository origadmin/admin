// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v1.0.0
// - protoc             v5.27.2
// source: system/admin.proto

package admin

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	gins "github.com/origadmin/contrib/transport/gins"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(gin.H)
var _ = gins.NewContext
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const MenuService_BatchCreate_OperationName = "/api.system.v1.MenuService/BatchCreate"
const MenuService_CreateMenu_OperationName = "/api.system.v1.MenuService/CreateMenu"
const MenuService_DeleteMenu_OperationName = "/api.system.v1.MenuService/DeleteMenu"
const MenuService_ListMenu_OperationName = "/api.system.v1.MenuService/ListMenu"
const MenuService_UpdateMenu_OperationName = "/api.system.v1.MenuService/UpdateMenu"

type MenuServiceGINServer interface {
	BatchCreate(context.Context, *BatchCreateMenusRequest) (*BatchCreateMenusResponse, error)
	CreateMenu(context.Context, *CreateMenuReq) (*CreateMenuReply, error)
	DeleteMenu(context.Context, *DeleteMenuReq) (*DeleteMenuReply, error)
	ListMenu(context.Context, *ListMenuReq) (*ListMenuReply, error)
	UpdateMenu(context.Context, *UpdateMenuReq) (*UpdateMenuReply, error)
}

func RegisterMenuServiceGINServer(router gin.IRouter, srv MenuServiceGINServer) {
	router.GET("/api/v1/system/menus", _MenuService_ListMenu0_GIN_Handler(srv))
	router.POST("/api/v1/system/menus", _MenuService_CreateMenu0_GIN_Handler(srv))
	router.PUT("/api/v1/system/menus/:id", _MenuService_UpdateMenu0_GIN_Handler(srv))
	router.DELETE("/api/v1/system/menus/:id", _MenuService_DeleteMenu0_GIN_Handler(srv))
	router.POST("/api/v1/system/menus/batch", _MenuService_BatchCreate0_GIN_Handler(srv))
}

func _MenuService_ListMenu0_GIN_Handler(srv MenuServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in ListMenuReq
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, MenuService_ListMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.ListMenu(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _MenuService_CreateMenu0_GIN_Handler(srv MenuServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in CreateMenuReq
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, MenuService_CreateMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CreateMenu(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _MenuService_UpdateMenu0_GIN_Handler(srv MenuServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in UpdateMenuReq
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, MenuService_UpdateMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.UpdateMenu(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _MenuService_DeleteMenu0_GIN_Handler(srv MenuServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in DeleteMenuReq
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, MenuService_DeleteMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.DeleteMenu(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

func _MenuService_BatchCreate0_GIN_Handler(srv MenuServiceGINServer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var in BatchCreateMenusRequest
		if err := gins.BindBody(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			ctx.Error(err)
			return
		}
		http.SetOperation(ctx, MenuService_BatchCreate_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.BatchCreate(newCtx, &in)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(200, reply)
		return
	}
}

type MenuServiceGINClient interface {
	BatchCreate(ctx context.Context, req *BatchCreateMenusRequest, opts ...http.CallOption) (rsp *BatchCreateMenusResponse, err error)
	CreateMenu(ctx context.Context, req *CreateMenuReq, opts ...http.CallOption) (rsp *CreateMenuReply, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuReq, opts ...http.CallOption) (rsp *DeleteMenuReply, err error)
	ListMenu(ctx context.Context, req *ListMenuReq, opts ...http.CallOption) (rsp *ListMenuReply, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuReq, opts ...http.CallOption) (rsp *UpdateMenuReply, err error)
}

type MenuServiceGINClientImpl struct {
	cc *http.Client
}

func NewMenuServiceGINClient(client *http.Client) MenuServiceGINClient {
	return &MenuServiceGINClientImpl{client}
}

func (c *MenuServiceGINClientImpl) BatchCreate(ctx context.Context, in *BatchCreateMenusRequest, opts ...http.CallOption) (*BatchCreateMenusResponse, error) {
	var out BatchCreateMenusResponse
	pattern := "/api/v1/system/menus/batch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(MenuService_BatchCreate_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceGINClientImpl) CreateMenu(ctx context.Context, in *CreateMenuReq, opts ...http.CallOption) (*CreateMenuReply, error) {
	var out CreateMenuReply
	pattern := "/api/v1/system/menus"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(MenuService_CreateMenu_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceGINClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuReq, opts ...http.CallOption) (*DeleteMenuReply, error) {
	var out DeleteMenuReply
	pattern := "/api/v1/system/menus/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(MenuService_DeleteMenu_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceGINClientImpl) ListMenu(ctx context.Context, in *ListMenuReq, opts ...http.CallOption) (*ListMenuReply, error) {
	var out ListMenuReply
	pattern := "/api/v1/system/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(MenuService_ListMenu_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceGINClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuReq, opts ...http.CallOption) (*UpdateMenuReply, error) {
	var out UpdateMenuReply
	pattern := "/api/v1/system/menus/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(MenuService_UpdateMenu_OperationName))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
