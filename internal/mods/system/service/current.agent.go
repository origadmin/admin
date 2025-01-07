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

// CurrentAPIAgentService is a Current service.
type CurrentAPIAgentService struct {
	resp.Response

	client pb.CurrentAPIClient
}

func (s CurrentAPIAgentService) RefreshCurrentToken(ctx context.Context, request *pb.RefreshCurrentTokenRequest) (*pb.RefreshCurrentTokenResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s CurrentAPIAgentService) GetCurrentUser(ctx context.Context, request *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error) {
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

func (s CurrentAPIAgentService) CurrentLogout(ctx context.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CurrentLogout(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentAPIAgentService) ListCurrentMenus(ctx context.Context, request *pb.ListCurrentMenusRequest) (*pb.ListCurrentMenusResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListCurrentMenus(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentAPIAgentService) ListCurrentRoles(ctx context.Context, request *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListCurrentRoles(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Data:    resp.Proto2AnyPBArray(response.Roles...),
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentSetting(ctx context.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentSetting(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentUser(ctx context.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentUser(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentUserPassword(ctx context.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.UpdateCurrentUserPassword(ctx, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

//func (s CurrentAPIAgentService) CurrentMenus(ctx context.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
//	response, err := s.client.CurrentMenus(context, request)
//	if err != nil {
//		log.Errorf("CurrentMenus error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s CurrentAPIAgentService) CurrentUser(ctx context.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
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

//func (s CurrentAPIAgentService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

// NewCurrentAPIAgentService new a Current service.
func NewCurrentAPIAgentService(client pb.CurrentAPIClient) *CurrentAPIAgentService {
	return &CurrentAPIAgentService{client: client}
}

// NewCurrentAPIAgent new a Current service.
func NewCurrentAPIAgent(client pb.CurrentAPIClient) pb.CurrentAPIAgent {
	return &CurrentAPIAgentService{client: client}
}
func NewCurrentServerAgent(client *service.GRPCClient) pb.CurrentAPIAgent {
	cli := pb.NewCurrentAPIClient(client)
	return NewCurrentAPIAgent(cli)
}

var _ pb.CurrentAPIAgent = (*CurrentAPIAgentService)(nil)
