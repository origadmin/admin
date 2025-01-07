// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: system/menu.proto

package system

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationMenuServiceCreateMenu = "/api.v1.services.system.MenuService/CreateMenu"
const OperationMenuServiceDeleteMenu = "/api.v1.services.system.MenuService/DeleteMenu"
const OperationMenuServiceGetMenu = "/api.v1.services.system.MenuService/GetMenu"
const OperationMenuServiceListMenus = "/api.v1.services.system.MenuService/ListMenus"
const OperationMenuServiceUpdateMenu = "/api.v1.services.system.MenuService/UpdateMenu"

type MenuServiceHTTPServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*CreateMenuResponse, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuResponse, error)
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuResponse, error)
	ListMenus(context.Context, *ListMenusRequest) (*ListMenusResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*UpdateMenuResponse, error)
}

func RegisterMenuServiceHTTPServer(s *http.Server, srv MenuServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/sys/menus", _MenuService_ListMenus0_HTTP_Handler(srv))
	r.GET("/sys/menus/{id}", _MenuService_GetMenu0_HTTP_Handler(srv))
	r.POST("/sys/menus", _MenuService_CreateMenu0_HTTP_Handler(srv))
	r.PUT("/sys/menus/{menu.id}", _MenuService_UpdateMenu0_HTTP_Handler(srv))
	r.DELETE("/sys/menus/{id}", _MenuService_DeleteMenu0_HTTP_Handler(srv))
}

func _MenuService_ListMenus0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceListMenus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenus(ctx, req.(*ListMenusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenusResponse)
		return ctx.Result(200, reply)
	}
}

func _MenuService_GetMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceGetMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenu(ctx, req.(*GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenuResponse)
		return ctx.Result(200, reply)
	}
}

func _MenuService_CreateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenuRequest
		if err := ctx.Bind(&in.Menu); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*CreateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateMenuResponse)
		return ctx.Result(200, reply)
	}
}

func _MenuService_UpdateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenuRequest
		if err := ctx.Bind(&in.Menu); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceUpdateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenu(ctx, req.(*UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateMenuResponse)
		return ctx.Result(200, reply)
	}
}

func _MenuService_DeleteMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuServiceDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*DeleteMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMenuResponse)
		return ctx.Result(200, reply)
	}
}

type MenuServiceHTTPClient interface {
	CreateMenu(ctx context.Context, req *CreateMenuRequest, opts ...http.CallOption) (rsp *CreateMenuResponse, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuRequest, opts ...http.CallOption) (rsp *DeleteMenuResponse, err error)
	GetMenu(ctx context.Context, req *GetMenuRequest, opts ...http.CallOption) (rsp *GetMenuResponse, err error)
	ListMenus(ctx context.Context, req *ListMenusRequest, opts ...http.CallOption) (rsp *ListMenusResponse, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuRequest, opts ...http.CallOption) (rsp *UpdateMenuResponse, err error)
}

type MenuServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuServiceHTTPClient(client *http.Client) MenuServiceHTTPClient {
	return &MenuServiceHTTPClientImpl{client}
}

func (c *MenuServiceHTTPClientImpl) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...http.CallOption) (*CreateMenuResponse, error) {
	var out CreateMenuResponse
	pattern := "/sys/menus"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuServiceCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Menu, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceHTTPClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...http.CallOption) (*DeleteMenuResponse, error) {
	var out DeleteMenuResponse
	pattern := "/sys/menus/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceHTTPClientImpl) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...http.CallOption) (*GetMenuResponse, error) {
	var out GetMenuResponse
	pattern := "/sys/menus/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceGetMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceHTTPClientImpl) ListMenus(ctx context.Context, in *ListMenusRequest, opts ...http.CallOption) (*ListMenusResponse, error) {
	var out ListMenusResponse
	pattern := "/sys/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuServiceListMenus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *MenuServiceHTTPClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...http.CallOption) (*UpdateMenuResponse, error) {
	var out UpdateMenuResponse
	pattern := "/sys/menus/{menu.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuServiceUpdateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Menu, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
