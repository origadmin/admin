/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"encoding/json"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// RoleServiceHookedBridge is a menu service.
type RoleServiceHookedBridge struct {
	pb.UnimplementedRoleServiceHooked
	log *log.KHelper
}

func (h RoleServiceHookedBridge) CompleteCreateRole(ctx transhttp.Context, request *pb.CreateRoleRequest, response *pb.CreateRoleResponse) error {
	marshal, err := json.Marshal(response.Role)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h RoleServiceHookedBridge) CompleteDeleteRole(ctx transhttp.Context, request *pb.DeleteRoleRequest, response *pb.DeleteRoleResponse) error {
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    nil,
	})
}

func (h RoleServiceHookedBridge) CompleteGetRole(ctx transhttp.Context, request *pb.GetRoleRequest, response *pb.GetRoleResponse) error {
	marshal, err := json.Marshal(response.Role)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h RoleServiceHookedBridge) CompleteListRoles(ctx transhttp.Context, request *pb.ListRolesRequest, response *pb.ListRolesResponse) error {
	marshal, err := json.Marshal(response.Roles)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourcePage{
		Success: true,
		Total:   response.GetTotalSize(),
		Data:    marshal,
		//Current:  request.GetCurrent(),
		//PageSize: nil,
		//Extra: "",
	})
}

func (h RoleServiceHookedBridge) CompleteUpdateRole(ctx transhttp.Context, request *pb.UpdateRoleRequest, response *pb.UpdateRoleResponse) error {
	marshal, err := json.Marshal(response.Role)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func NewRoleServiceHookedBridge(r runtime.Runtime, client pb.RoleServiceHTTPServer) pb.RoleServiceHookedBridger {
	return pb.WithRoleServiceHook(&RoleServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/system")),
	})(client)
}

// NewRoleServiceBridge new a menu service.
func NewRoleServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.RoleServiceServer {
	return pb.NewRoleServiceBridge(client)
}

func NewRoleServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.RoleServiceServer {
	if v, ok := clients["system"]; ok {
		return NewRoleServiceBridge(r, v)
	} else {
		return pb.UnimplementedRoleServiceServer{}
	}
}

// NewRoleServiceHTTPBridge new a menu service.
func NewRoleServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.RoleServiceHTTPServer {
	return pb.NewRoleServiceHTTPBridge(client)
}

var _ pb.RoleServiceHooker = (*RoleServiceHookedBridge)(nil)
