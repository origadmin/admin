/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/cmp"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/helpers/resp"
)

// PersonalServiceHookedBridge is a menu service.
type PersonalServiceHookedBridge struct {
	pb.UnimplementedPersonalServiceHooked
	client pb.PersonalServiceHTTPServer
	log    *log.KHelper
}

//func (p PersonalServiceHookedBridge) PrepareListPersonalResources(ctx transhttp.Context, request *pb.ListPersonalResourcesRequest) (context.Context, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (p PersonalServiceHookedBridge) CompleteListPersonalResources(ctx transhttp.Context, request *pb.ListPersonalResourcesRequest, response *pb.ListPersonalResourcesResponse) error {
	marshal, err := resp.Proto2JSON(response.Resources...)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.Result{
		Success:       true,
		Data:          marshal,
		Total:         int32(response.TotalSize),
		NextPageToken: cmp.If(response.NextPageToken != "", &response.NextPageToken, nil),
	})
}

func (p PersonalServiceHookedBridge) PrepareGetPersonalProfile(ctx transhttp.Context, request *pb.GetPersonalProfileRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteGetPersonalProfile(ctx transhttp.Context, request *pb.GetPersonalProfileRequest, response *pb.GetPersonalProfileResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareListPersonalRoles(ctx transhttp.Context, request *pb.ListPersonalRolesRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteListPersonalRoles(ctx transhttp.Context, request *pb.ListPersonalRolesRequest, response *pb.ListPersonalRolesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PreparePersonalLogout(ctx transhttp.Context, request *pb.PersonalLogoutRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompletePersonalLogout(ctx transhttp.Context, request *pb.PersonalLogoutRequest, response *pb.PersonalLogoutResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareRefreshPersonalToken(ctx transhttp.Context, request *pb.RefreshPersonalTokenRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteRefreshPersonalToken(ctx transhttp.Context, request *pb.RefreshPersonalTokenRequest, response *pb.RefreshPersonalTokenResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalPassword(ctx transhttp.Context, request *pb.UpdatePersonalPasswordRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalPassword(ctx transhttp.Context, request *pb.UpdatePersonalPasswordRequest, response *pb.UpdatePersonalPasswordResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalProfile(ctx transhttp.Context, request *pb.UpdatePersonalProfileRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalProfile(ctx transhttp.Context, request *pb.UpdatePersonalProfileRequest, response *pb.UpdatePersonalProfileResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalSetting(ctx transhttp.Context, request *pb.UpdatePersonalSettingRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalSetting(ctx transhttp.Context, request *pb.UpdatePersonalSettingRequest, response *pb.UpdatePersonalSettingResponse) error {
	//TODO implement me
	panic("implement me")
}

func NewPersonalServiceHookedBridge(r runtime.Runtime, client pb.PersonalServiceHTTPServer) pb.PersonalServiceHookedBridger {
	return pb.WithPersonalServiceHook(&PersonalServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/auth")),
	})(client)
}

// NewPersonalServiceBridge new a menu service.
func NewPersonalServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.PersonalServiceServer {
	return pb.NewPersonalServiceBridge(client)
}

func NewPersonalServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.PersonalServiceServer {
	if c, ok := clients["auth"]; ok {
		return pb.NewPersonalServiceBridge(c)
	} else {
		return pb.UnimplementedPersonalServiceServer{}
	}
}

// NewPersonalServiceHTTPBridge new a menu service.
func NewPersonalServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.PersonalServiceHTTPServer {
	return pb.NewPersonalServiceHTTPBridge(client)
}

var _ pb.PersonalServiceHooker = (*PersonalServiceHookedBridge)(nil)
