/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
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

func (s CurrentAPIAgentService) CurrentLogout(context transhttp.Context, request *pb.CurrentLogoutRequest) (*pb.CurrentLogoutResponse, error) {
	response, err := s.client.CurrentLogout(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s CurrentAPIAgentService) ListCurrentMenus(context transhttp.Context, request *pb.ListCurrentMenusRequest) (*pb.ListCurrentMenusResponse, error) {
	response, err := s.client.ListCurrentMenus(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentRoles(context transhttp.Context, request *pb.UpdateCurrentRolesRequest) (*pb.UpdateCurrentRolesResponse, error) {
	response, err := s.client.UpdateCurrentRoles(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentSetting(context transhttp.Context, request *pb.UpdateCurrentSettingRequest) (*pb.UpdateCurrentSettingResponse, error) {
	response, err := s.client.UpdateCurrentSetting(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentUser(context transhttp.Context, request *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error) {
	response, err := s.client.UpdateCurrentUser(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s CurrentAPIAgentService) UpdateCurrentUserPassword(context transhttp.Context, request *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error) {
	response, err := s.client.UpdateCurrentUserPassword(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

//func (s CurrentAPIAgentService) CurrentMenus(context transhttp.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
//	response, err := s.client.CurrentMenus(context, request)
//	if err != nil {
//		log.Errorf("CurrentMenus error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s CurrentAPIAgentService) CurrentUser(context transhttp.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
//	response, err := s.client.CurrentUser(context, request)
//	if err != nil {
//		log.Errorf("CurrentUser error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}

//func (s CurrentAPIAgentService) Logout(context transhttp.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
//	response, err := s.client.Logout(context, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	s.JSON(context, http.StatusOK, &resp.Result{
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
