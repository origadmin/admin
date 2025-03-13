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

// PersonalServiceAgent is a Personal service.
type PersonalServiceAgent struct {
	resp.Response

	client pb.PersonalServiceClient
}

func (s PersonalServiceAgent) RefreshPersonalToken(ctx context.Context, request *pb.RefreshPersonalTokenRequest) (*pb.RefreshPersonalTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s PersonalServiceAgent) GetPersonalProfile(ctx context.Context, request *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error) {
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

func (s PersonalServiceAgent) PersonalLogout(ctx context.Context, request *pb.PersonalLogoutRequest) (*pb.PersonalLogoutResponse, error) {
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

func (s PersonalServiceAgent) ListPersonalResources(ctx context.Context, request *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error) {
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

func (s PersonalServiceAgent) ListPersonalRoles(ctx context.Context, request *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error) {
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

func (s PersonalServiceAgent) UpdatePersonalSetting(ctx context.Context, request *pb.UpdatePersonalSettingRequest) (*pb.UpdatePersonalSettingResponse, error) {
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

func (s PersonalServiceAgent) UpdatePersonalProfile(ctx context.Context, request *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error) {
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

func (s PersonalServiceAgent) UpdatePersonalPassword(ctx context.Context, request *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error) {
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

//func (s PersonalServiceAgent) PersonalResources(ctx context.Context, request *pb.PersonalResourcesRequest) (*pb.PersonalResourcesResponse, error) {
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

//func (s PersonalServiceAgent) PersonalProfile(ctx context.Context, request *pb.PersonalProfileRequest) (*pb.PersonalProfileResponse, error) {
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

//func (s PersonalServiceAgent) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

// NewPersonalServiceAgent new a Personal service.
func NewPersonalServiceAgent(client pb.PersonalServiceClient) *PersonalServiceAgent {
	return &PersonalServiceAgent{client: client}
}

// NewPersonalServiceAgentPB new a Personal service.
func NewPersonalServiceAgentPB(client pb.PersonalServiceClient) pb.PersonalServiceAgent {
	return &PersonalServiceAgent{client: client}
}
func NewPersonalServiceAgentClient(client *service.GRPCClient) pb.PersonalServiceAgent {
	cli := pb.NewPersonalServiceClient(client)
	return NewPersonalServiceAgent(cli)
}

var _ pb.PersonalServiceAgent = (*PersonalServiceAgent)(nil)
