/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/helpers/resp"
)

// LoginAPIAgentService is a Login service.
type LoginAPIAgentService struct {
	resp.Response

	client pb.LoginAPIClient
}

func (s LoginAPIAgentService) Register(context transhttp.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	response, err := s.client.Register(context, request)
	if err != nil {
		log.Errorf("Register error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaResource(context transhttp.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	response, err := s.client.CaptchaResource(context, request)
	if err != nil {
		log.Errorf("CaptchaResource error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaResources(context transhttp.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	response, err := s.client.CaptchaResources(context, request)
	if err != nil {
		log.Errorf("CaptchaResources error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaID(context transhttp.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	log.Debugf("CaptchaID: Request:%+v", request)
	response, err := s.client.CaptchaID(context, request)
	log.Debugf("CaptchaID: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Data,
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaImage(context transhttp.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	log.Debugf("CaptchaImage: Request:%+v", request)
	response, err := s.client.CaptchaImage(context, request)
	log.Debugf("CaptchaImage: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		return nil, err
	}
	log.Debugf("CaptchaImage: Setting headers: %+v", response.Headers)
	for k, v := range response.Headers {
		context.Response().Header().Set(k, v)
	}
	log.Debugf("CaptchaImage: Writing response headers")
	context.Response().WriteHeader(http.StatusOK)
	log.Debugf("CaptchaImage: Writing response image")
	if _, err := context.Response().Write(response.Image); err != nil {
		log.Errorf("CaptchaImage error writing response: %v", err)
		return nil, err
	}
	//log.Debugf("CaptchaImage: Flushing response writer")
	//context.Response().Flush()
	log.Debugf("CaptchaImage: Completed successfully")
	return nil, nil
}

func (s LoginAPIAgentService) CurrentMenus(context transhttp.Context, request *pb.CurrentMenusRequest) (*pb.CurrentMenusResponse, error) {
	response, err := s.client.CurrentMenus(context, request)
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

func (s LoginAPIAgentService) CurrentUser(context transhttp.Context, request *pb.CurrentUserRequest) (*pb.CurrentUserResponse, error) {
	response, err := s.client.CurrentUser(context, request)
	if err != nil {
		log.Errorf("CurrentUser error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}
func (s LoginAPIAgentService) Refresh(context transhttp.Context, request *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	response, err := s.client.Refresh(context, request)
	if err != nil {
		log.Errorf("Refresh error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data: gin.H{
			"user_id":       response.Token.GetUserId(),
			"access_token":  response.Token.GetAccessToken(),
			"refresh_token": response.Token.GetRefreshToken(),
			"expires_at":    response.Token.GetExpirationTime(),
		},
	})
	return nil, nil
}
func (s LoginAPIAgentService) Login(context transhttp.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	response, err := s.client.Login(context, request)
	if err != nil {
		log.Errorf("Login error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data: gin.H{
			"user_id":       response.Token.GetUserId(),
			"access_token":  response.Token.GetAccessToken(),
			"refresh_token": response.Token.GetRefreshToken(),
			"expires_at":    response.Token.GetExpirationTime(),
		},
	})
	return nil, nil
}

func (s LoginAPIAgentService) Logout(context transhttp.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	response, err := s.client.Logout(context, request)
	if err != nil {
		log.Errorf("Logout error: %v", err)
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

// NewLoginAPIAgentService new a Login service.
func NewLoginAPIAgentService(client pb.LoginAPIClient) *LoginAPIAgentService {
	return &LoginAPIAgentService{client: client}
}

// NewLoginAPIAgent new a Login service.
func NewLoginAPIAgent(client pb.LoginAPIClient) pb.LoginAPIAgent {
	return &LoginAPIAgentService{client: client}
}
func NewLoginServerAgent(client *service.GRPCClient) pb.LoginAPIAgent {
	cli := pb.NewLoginAPIClient(client)
	return NewLoginAPIAgent(cli)
}

func NewLoginServerAgentGINRegister(client *service.GRPCClient) func(server *transhttp.Server) {
	cli := NewLoginServerAgent(client)
	return func(server *transhttp.Server) {
		pb.RegisterLoginAPIAgent(agent.New(server), cli)
	}
}

var _ pb.LoginAPIAgent = (*LoginAPIAgentService)(nil)
