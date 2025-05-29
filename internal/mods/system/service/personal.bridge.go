/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// PersonalServiceBridge is a Personal service.
type PersonalServiceBridge struct {
	resp.Response

	client pb.PersonalServiceClient
}

func (s PersonalServiceBridge) RefreshPersonalToken(ctx context.Context, request *pb.RefreshPersonalTokenRequest) (*pb.RefreshPersonalTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s PersonalServiceBridge) GetPersonalProfile(ctx context.Context, request *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetPersonalProfile(ctx, request)
	if err != nil {
		log.Errorf("GetPersonalProfile error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
	return nil, nil
}

func (s PersonalServiceBridge) PersonalLogout(ctx context.Context, request *pb.PersonalLogoutRequest) (*pb.PersonalLogoutResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.PersonalLogout(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
	return nil, nil
}

func (s PersonalServiceBridge) ListPersonalResources(ctx context.Context, request *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListPersonalResources(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   int32(response.TotalSize),
		Data:    resp.Proto2AnyPBArray(response.Resources...),
	})
	return nil, nil
}

func (s PersonalServiceBridge) ListPersonalRoles(ctx context.Context, request *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListPersonalRoles(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		//Total:   int32(response.TotalSize),
		Data: resp.Proto2AnyPBArray(response.Roles...),
	})
	return nil, nil
}

func (s PersonalServiceBridge) UpdatePersonalSetting(ctx context.Context, request *pb.UpdatePersonalSettingRequest) (*pb.UpdatePersonalSettingResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdatePersonalSetting(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
	return nil, nil
}

func (s PersonalServiceBridge) UpdatePersonalProfile(ctx context.Context, request *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdatePersonalProfile(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
	return nil, nil
}

func (s PersonalServiceBridge) UpdatePersonalPassword(ctx context.Context, request *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdatePersonalPassword(ctx, request)
	if err != nil {
		log.Errorf("PersonalResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
	return nil, nil
}

//func (s PersonalServiceBridge) PersonalResources(ctx context.Context, request *pb.PersonalResourcesRequest) (*pb.PersonalResourcesResponse, error) {
//	response, err := s.client.PersonalResources(context, request)
//	if err != nil {
//		log.Errorf("PersonalResources error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s PersonalServiceBridge) PersonalProfile(ctx context.Context, request *pb.PersonalProfileRequest) (*pb.PersonalProfileResponse, error) {
//	response, err := s.client.PersonalProfile(context, request)
//	if err != nil {
//		log.Errorf("PersonalProfile error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s PersonalServiceBridge) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
//	response, err := s.client.Logout(context, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

// NewPersonalServiceBridge new a Personal service.
func NewPersonalServiceBridge(client pb.PersonalServiceClient) *PersonalServiceBridge {
	return &PersonalServiceBridge{client: client}
}

// NewPersonalServiceBridgePB new a Personal service.
func NewPersonalServiceBridgePB(client pb.PersonalServiceClient) pb.PersonalServiceBridge {
	return &PersonalServiceBridge{client: client}
}
func NewPersonalServiceBridgeClient(client *service.GRPCClient) pb.PersonalServiceServer {
	cli := pb.NewPersonalServiceClient(client)
	return NewPersonalServiceBridge(cli)
}

var _ pb.PersonalServiceBridge = (*PersonalServiceBridge)(nil)
