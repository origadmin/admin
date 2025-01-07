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

// CurrentServiceAgent is a Current service.
type CurrentServiceAgent struct {
	resp.Response

	client pb.CurrentServiceClient
}

func (s CurrentServiceAgent) RefreshCurrentToken(ctx context.Context, request *pb.RefreshCurrentTokenRequest) (*pb.RefreshCurrentTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CurrentServiceAgent) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.GetCurrentUser(ctx, request)
	if err != nil {
		log.Errorf("GetCurrentUser error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentServiceAgent) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CurrentLogout(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentServiceAgent) ListCurrentResources(ctx context.Context, request *pb.ListCurrentResourcesRequest) (*pb.ListCurrentResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListCurrentResources(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentServiceAgent) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListCurrentRoles(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Data:    resp.Proto2AnyPBArray(response.Roles...),
	})
	return nil, nil
}

func (s CurrentServiceAgent) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentSetting(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentServiceAgent) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentUser(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentServiceAgent) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentUserPassword(ctx, request)
	if err != nil {
		log.Errorf("CurrentResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

//func (s CurrentServiceAgent) CurrentResources(ctx context.Context, request *pb.CurrentResourcesRequest) (*pb.CurrentResourcesResponse, error) {
//	response, err := s.client.CurrentResources(context, request)
//	if err != nil {
//		log.Errorf("CurrentResources error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s CurrentServiceAgent) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
//	response, err := s.client.CurrentUser(context, request)
//	if err != nil {
//		log.Errorf("CurrentUser error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s CurrentServiceAgent) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

// NewCurrentServiceAgent new a Current service.
func NewCurrentServiceAgent(client pb.CurrentServiceClient) *CurrentServiceAgent {
	return &CurrentServiceAgent{client: client}
}

// NewCurrentServiceAgentPB new a Current service.
func NewCurrentServiceAgentPB(client pb.CurrentServiceClient) pb.CurrentServiceAgent {
	return &CurrentServiceAgent{client: client}
}
func NewCurrentServiceAgentClient(client *service.GRPCClient) pb.CurrentServiceAgent {
	cli := pb.NewCurrentServiceClient(client)
	return NewCurrentServiceAgent(cli)
}

var _ pb.CurrentServiceAgent = (*CurrentServiceAgent)(nil)
