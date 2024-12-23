// Code generated by protoc-gen-go-gins. DO NOT EDIT.
// versions:
// - protoc-gen-go-gins 0.0.11
// - protoc             (unknown)
// source: system/menu.proto

package system

import (
	context "context"
	gins "github.com/origadmin/contrib/transport/gins"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)

const _ = gins.SupportPackageIsVersion1

const MenuAPI_CreateMenu_OperationName = "/api.v1.services.system.MenuAPI/CreateMenu"
const MenuAPI_DeleteMenu_OperationName = "/api.v1.services.system.MenuAPI/DeleteMenu"
const MenuAPI_GetMenu_OperationName = "/api.v1.services.system.MenuAPI/GetMenu"
const MenuAPI_ListMenus_OperationName = "/api.v1.services.system.MenuAPI/ListMenus"
const MenuAPI_UpdateMenu_OperationName = "/api.v1.services.system.MenuAPI/UpdateMenu"

type MenuAPIGINSServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuResponse, error)
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	ListMenus(context.Context, *ListMenusRequest) (*ListMenusResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
}

func RegisterMenuAPIGINSServer(router gins.IRouter, srv MenuAPIGINSServer) {
	router.GET("/sys/menus", _MenuAPI_ListMenus0_GIN_Handler(srv))
	router.GET("/sys/menus/:id", _MenuAPI_GetMenu0_GIN_Handler(srv))
	router.POST("/sys/menus", _MenuAPI_CreateMenu0_GIN_Handler(srv))
	router.PUT("/sys/menus/:menu.id", _MenuAPI_UpdateMenu0_GIN_Handler(srv))
	router.DELETE("/sys/menus/:id", _MenuAPI_DeleteMenu0_GIN_Handler(srv))
}

func _MenuAPI_ListMenus0_GIN_Handler(srv MenuAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in ListMenusRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, MenuAPI_ListMenus_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.ListMenus(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _MenuAPI_GetMenu0_GIN_Handler(srv MenuAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in GetMenuRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, MenuAPI_GetMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.GetMenu(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _MenuAPI_CreateMenu0_GIN_Handler(srv MenuAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in CreateMenuRequest
		if err := gins.BindBody(ctx, &in.Menu); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, MenuAPI_CreateMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.CreateMenu(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _MenuAPI_UpdateMenu0_GIN_Handler(srv MenuAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in UpdateMenuRequest
		if err := gins.BindBody(ctx, &in.Menu); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, MenuAPI_UpdateMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.UpdateMenu(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

func _MenuAPI_DeleteMenu0_GIN_Handler(srv MenuAPIGINSServer) gins.HandlerFunc {
	return func(ctx *gins.Context) {
		var in DeleteMenuRequest
		if err := gins.BindQuery(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		if err := gins.BindURI(ctx, &in); err != nil {
			gins.JSON(ctx, 400, err)
			return
		}
		gins.SetOperation(ctx, MenuAPI_DeleteMenu_OperationName)
		newCtx := gins.NewContext(ctx)
		reply, err := srv.DeleteMenu(newCtx, &in)
		if err != nil {
			gins.JSON(ctx, 500, err)
			return
		}
		gins.JSON(ctx, 200, reply)
		return
	}
}

type MenuAPIGINSClient interface {
	CreateMenu(ctx context.Context, req *CreateMenuRequest, opts ...gins.CallOption) (rsp *CreateMenuResponse, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuRequest, opts ...gins.CallOption) (rsp *DeleteMenuResponse, err error)
	GetMenu(ctx context.Context, req *GetMenuRequest, opts ...gins.CallOption) (rsp *GetMenuResponse, err error)
	ListMenus(ctx context.Context, req *ListMenusRequest, opts ...gins.CallOption) (rsp *ListMenusResponse, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuRequest, opts ...gins.CallOption) (rsp *UpdateMenuResponse, err error)
}

type MenuAPIGINSClientImpl struct {
	cc *gins.Client
}

func NewMenuAPIGINSClient(client *gins.Client) MenuAPIGINSClient {
	return &MenuAPIGINSClientImpl{client}
}

func (c *MenuAPIGINSClientImpl) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...gins.CallOption) (*CreateMenuResponse, error) {
	var out CreateMenuResponse
	pattern := "/sys/menus"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(MenuAPI_CreateMenu_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Menu, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuAPIGINSClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...gins.CallOption) (*DeleteMenuResponse, error) {
	var out DeleteMenuResponse
	pattern := "/sys/menus/{id}"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(MenuAPI_DeleteMenu_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuAPIGINSClientImpl) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...gins.CallOption) (*GetMenuResponse, error) {
	var out GetMenuResponse
	pattern := "/sys/menus/{id}"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(MenuAPI_GetMenu_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuAPIGINSClientImpl) ListMenus(ctx context.Context, in *ListMenusRequest, opts ...gins.CallOption) (*ListMenusResponse, error) {
	var out ListMenusResponse
	pattern := "/sys/menus"
	path := gins.EncodeURL(pattern, in, true)
	opts = append(opts, gins.Operation(MenuAPI_ListMenus_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuAPIGINSClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...gins.CallOption) (*UpdateMenuResponse, error) {
	var out UpdateMenuResponse
	pattern := "/sys/menus/{menu.id}"
	path := gins.EncodeURL(pattern, in, false)
	opts = append(opts, gins.Operation(MenuAPI_UpdateMenu_OperationName))
	opts = append(opts, gins.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Menu, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
