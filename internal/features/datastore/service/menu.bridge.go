/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"encoding/json"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/cmp"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// MenuServiceHookedBridge is a menu service.
type MenuServiceHookedBridge struct {
	pb.UnimplementedMenuServiceHooked
	log *log.KHelper
}

func (h MenuServiceHookedBridge) CompleteCreateMenu(ctx transhttp.Context, request *pb.CreateMenuRequest, response *pb.CreateMenuResponse) error {
	marshal, err := json.Marshal(response.Menu)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h MenuServiceHookedBridge) CompleteDeleteMenu(ctx transhttp.Context, request *pb.DeleteMenuRequest, response *pb.DeleteMenuResponse) error {
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    nil,
	})
}

func (h MenuServiceHookedBridge) CompleteGetMenu(ctx transhttp.Context, request *pb.GetMenuRequest, response *pb.GetMenuResponse) error {
	marshal, err := json.Marshal(response.Menu)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h MenuServiceHookedBridge) CompleteListMenus(ctx transhttp.Context, request *pb.ListMenusRequest, response *pb.ListMenusResponse) error {
	if response == nil {
		return ctx.JSON(http.StatusOK, &resp.Result{
			Success: false,
			Data:    nil,
		})
	}
	marshal, err := resp.Proto2JSON(response.Menus...)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.Result{
		Success:       true,
		Data:          marshal,
		Total:         response.TotalSize,
		NextPageToken: cmp.If(response.NextPageToken != "", &response.NextPageToken, nil),
	})
}

func (h MenuServiceHookedBridge) CompleteUpdateMenu(ctx transhttp.Context, request *pb.UpdateMenuRequest, response *pb.UpdateMenuResponse) error {
	marshal, err := json.Marshal(response.Menu)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func NewMenuServiceHookedBridge(r runtime.Runtime, client pb.MenuServiceHTTPServer) pb.MenuServiceHookedBridger {
	return pb.WithMenuServiceHook(&MenuServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/system")),
	})(client)
}

// NewMenuServiceBridge new a menu service.
func NewMenuServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.MenuServiceServer {
	return pb.NewMenuServiceBridge(client)
}

func NewMenuServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.MenuServiceServer {
	if c, ok := clients["system"]; ok {
		return pb.NewMenuServiceBridge(c)
	} else {
		return pb.UnimplementedMenuServiceServer{}
	}
}

// NewMenuServiceHTTPBridge new a menu service.
func NewMenuServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.MenuServiceHTTPServer {
	return pb.NewMenuServiceHTTPBridge(client)
}

var _ pb.MenuServiceHooker = (*MenuServiceHookedBridge)(nil)
