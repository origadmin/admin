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

// PermissionServiceHookedBridge is a menu service.
type PermissionServiceHookedBridge struct {
	pb.UnimplementedPermissionServiceHooked
	log *log.KHelper
}

func (h PermissionServiceHookedBridge) CompleteCreatePermission(ctx transhttp.Context, request *pb.CreatePermissionRequest, response *pb.CreatePermissionResponse) error {
	marshal, err := json.Marshal(response.Permission)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h PermissionServiceHookedBridge) CompleteDeletePermission(ctx transhttp.Context, request *pb.DeletePermissionRequest, response *pb.DeletePermissionResponse) error {
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    nil,
	})
}

func (h PermissionServiceHookedBridge) CompleteGetPermission(ctx transhttp.Context, request *pb.GetPermissionRequest, response *pb.GetPermissionResponse) error {
	marshal, err := json.Marshal(response.Permission)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h PermissionServiceHookedBridge) CompleteListPermissions(ctx transhttp.Context, request *pb.ListPermissionsRequest, response *pb.ListPermissionsResponse) error {
	marshal, err := json.Marshal(response.Permissions)
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

func (h PermissionServiceHookedBridge) CompleteUpdatePermission(ctx transhttp.Context, request *pb.UpdatePermissionRequest, response *pb.UpdatePermissionResponse) error {
	marshal, err := json.Marshal(response.Permission)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func NewPermissionServiceHookedBridge(r runtime.Runtime, client pb.PermissionServiceHTTPServer) pb.PermissionServiceHookedBridger {
	return pb.WithPermissionServiceHook(&PermissionServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/permission")),
	})(client)
}

// NewPermissionServiceBridge new a menu service.
func NewPermissionServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.PermissionServiceServer {
	return pb.NewPermissionServiceBridge(client)
}

func NewPermissionServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.PermissionServiceServer {
	if c, ok := clients["system"]; ok {
		return pb.NewPermissionServiceBridge(c)
	} else {
		return pb.UnimplementedPermissionServiceServer{}
	}
}

// NewPermissionServiceHTTPBridge new a menu service.
func NewPermissionServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.PermissionServiceHTTPServer {
	return pb.NewPermissionServiceHTTPBridge(client)
}

var _ pb.PermissionServiceHooker = (*PermissionServiceHookedBridge)(nil)
