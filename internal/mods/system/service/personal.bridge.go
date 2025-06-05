/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
)

// PersonalServiceHookedBridge is a menu service.
type PersonalServiceHookedBridge struct {
	pb.UnimplementedPersonalServiceHooked
	client pb.PersonalServiceHTTPServer
	log    *log.KHelper
}

func (p PersonalServiceHookedBridge) PrepareGetPersonalProfile(context transhttp.Context, request *pb.GetPersonalProfileRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteGetPersonalProfile(context transhttp.Context, request *pb.GetPersonalProfileRequest, response *pb.GetPersonalProfileResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareListPersonalResources(context transhttp.Context, request *pb.ListPersonalResourcesRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteListPersonalResources(context transhttp.Context, request *pb.ListPersonalResourcesRequest, response *pb.ListPersonalResourcesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareListPersonalRoles(context transhttp.Context, request *pb.ListPersonalRolesRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteListPersonalRoles(context transhttp.Context, request *pb.ListPersonalRolesRequest, response *pb.ListPersonalRolesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PreparePersonalLogout(context transhttp.Context, request *pb.PersonalLogoutRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompletePersonalLogout(context transhttp.Context, request *pb.PersonalLogoutRequest, response *pb.PersonalLogoutResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareRefreshPersonalToken(context transhttp.Context, request *pb.RefreshPersonalTokenRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteRefreshPersonalToken(context transhttp.Context, request *pb.RefreshPersonalTokenRequest, response *pb.RefreshPersonalTokenResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalPassword(context transhttp.Context, request *pb.UpdatePersonalPasswordRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalPassword(context transhttp.Context, request *pb.UpdatePersonalPasswordRequest, response *pb.UpdatePersonalPasswordResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalProfile(context transhttp.Context, request *pb.UpdatePersonalProfileRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalProfile(context transhttp.Context, request *pb.UpdatePersonalProfileRequest, response *pb.UpdatePersonalProfileResponse) error {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) PrepareUpdatePersonalSetting(context transhttp.Context, request *pb.UpdatePersonalSettingRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (p PersonalServiceHookedBridge) CompleteUpdatePersonalSetting(context transhttp.Context, request *pb.UpdatePersonalSettingRequest, response *pb.UpdatePersonalSettingResponse) error {
	//TODO implement me
	panic("implement me")
}

func NewPersonalServiceHookedBridge(r runtime.Runtime, client pb.PersonalServiceHTTPServer) pb.PersonalServiceHooker {
	return &PersonalServiceHookedBridge{
		log:    log.NewHelper(r.WithLogger("module", "service/permission")),
		client: client,
	}
}

// NewPersonalServiceBridge new a menu service.
func NewPersonalServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.PersonalServiceServer {
	return pb.NewPersonalServiceBridge(client)
}

// NewPersonalServiceHTTPBridge new a menu service.
func NewPersonalServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.PersonalServiceHTTPServer {
	return pb.NewPersonalServiceHTTPBridge(client)
}

var _ pb.PersonalServiceHooker = (*PersonalServiceHookedBridge)(nil)
