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

// UserServiceHookedBridge is a menu service.
type UserServiceHookedBridge struct {
	pb.UnimplementedUserServiceHooked
	log *log.KHelper
}

func (h UserServiceHookedBridge) CompleteCreateUser(ctx transhttp.Context, request *pb.CreateUserRequest, response *pb.CreateUserResponse) error {
	marshal, err := json.Marshal(response.User)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h UserServiceHookedBridge) CompleteDeleteUser(ctx transhttp.Context, request *pb.DeleteUserRequest, response *pb.DeleteUserResponse) error {
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    nil,
	})
}

func (h UserServiceHookedBridge) CompleteGetUser(ctx transhttp.Context, request *pb.GetUserRequest, response *pb.GetUserResponse) error {
	marshal, err := json.Marshal(response.User)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func (h UserServiceHookedBridge) CompleteListUsers(ctx transhttp.Context, request *pb.ListUsersRequest, response *pb.ListUsersResponse) error {
	marshal, err := json.Marshal(response.Users)
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

func (h UserServiceHookedBridge) CompleteUpdateUser(ctx transhttp.Context, request *pb.UpdateUserRequest, response *pb.UpdateUserResponse) error {
	marshal, err := json.Marshal(response.User)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.SourceData{
		Success: true,
		Data:    marshal,
	})
}

func NewUserServiceHookedBridge(r runtime.Runtime, client pb.UserServiceHTTPServer) pb.UserServiceHookedBridger {
	return pb.WithUserServiceHook(&UserServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/system")),
	})(client)
}

// NewUserServiceBridge new a menu service.
func NewUserServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.UserServiceServer {
	return pb.NewUserServiceBridge(client)
}

func NewUserServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.UserServiceServer {
	if c, ok := clients["system"]; ok {
		return pb.NewUserServiceBridge(c)
	} else {
		return pb.UnimplementedUserServiceServer{}
	}
}

// NewUserServiceHTTPBridge new a menu service.
func NewUserServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.UserServiceHTTPServer {
	return pb.NewUserServiceHTTPBridge(client)
}

var _ pb.UserServiceHooker = (*UserServiceHookedBridge)(nil)
