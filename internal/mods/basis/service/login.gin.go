/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/basis"
	"origadmin/application/admin/helpers/resp"
)

// LoginAPIGINRPCService is a Login service.
type LoginAPIGINRPCService struct {
	resp.Response

	client pb.LoginAPIClient
}

func (s LoginAPIGINRPCService) CaptchaResource(context *gins.Context, request *pb.CaptchaResourceRequest) {
	//TODO implement me
	panic("implement me")
}

func (s LoginAPIGINRPCService) CaptchaResources(context *gins.Context, request *pb.CaptchaResourcesRequest) {
	//TODO implement me
	panic("implement me")
}

func (s LoginAPIGINRPCService) CaptchaID(context *gins.Context, request *pb.CaptchaIDRequest) {
	log.Debugf("CaptchaID: Request:%+v", request)
	response, err := s.client.CaptchaID(context, request)
	log.Debugf("CaptchaID: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response.Data,
	})
}

func (s LoginAPIGINRPCService) CaptchaImage(context *gins.Context, request *pb.CaptchaImageRequest) {
	log.Debugf("CaptchaImage: Request:%+v", request)
	response, err := s.client.CaptchaImage(context, request)
	log.Debugf("CaptchaImage: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	log.Debugf("CaptchaImage: Setting headers: %+v", response.Headers)
	for k, v := range response.Headers {
		context.Writer.Header().Set(k, v)
	}
	log.Debugf("CaptchaImage: Writing response headers")
	context.Writer.WriteHeader(http.StatusOK)
	log.Debugf("CaptchaImage: Writing response image")
	if _, err := context.Writer.Write(response.Image); err != nil {
		log.Errorf("CaptchaImage error writing response: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	log.Debugf("CaptchaImage: Flushing response writer")
	context.Writer.Flush()
	log.Debugf("CaptchaImage: Completed successfully")
	return
}

func (s LoginAPIGINRPCService) CurrentMenus(context *gins.Context, request *pb.CurrentMenusRequest) {
	response, err := s.client.CurrentMenus(context, request)
	if err != nil {
		log.Errorf("CurrentMenus error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
}

func (s LoginAPIGINRPCService) CurrentUser(context *gins.Context, request *pb.CurrentUserRequest) {
	response, err := s.client.CurrentUser(context, request)
	if err != nil {
		log.Errorf("CurrentUser error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
}
func (s LoginAPIGINRPCService) Refresh(context *gins.Context, request *pb.RefreshRequest) {
	response, err := s.client.Refresh(context, request)
	if err != nil {
		log.Errorf("Refresh error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data: gin.H{
			"user_id":       response.Token.GetUserId(),
			"access_token":  response.Token.GetAccessToken(),
			"refresh_token": response.Token.GetRefreshToken(),
			"expires_at":    response.Token.GetExpirationTime().GetSeconds(),
		},
	})
}
func (s LoginAPIGINRPCService) Login(context *gins.Context, request *pb.LoginRequest) {
	response, err := s.client.Login(context, request)
	if err != nil {
		log.Errorf("Login error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data: gin.H{
			"user_id":       response.Token.GetUserId(),
			"access_token":  response.Token.GetAccessToken(),
			"refresh_token": response.Token.GetRefreshToken(),
			"expires_at":    response.Token.GetExpirationTime().GetSeconds(),
		},
	})
}

func (s LoginAPIGINRPCService) Logout(context *gins.Context, request *pb.LogoutRequest) {
	response, err := s.client.Logout(context, request)
	if err != nil {
		log.Errorf("Logout error: %v", err)
		s.Error(context, http.StatusNotFound, err)
		return
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
}

// NewLoginAPIGINRPCService new a Login service.
func NewLoginAPIGINRPCService(client pb.LoginAPIClient) *LoginAPIGINRPCService {
	return &LoginAPIGINRPCService{client: client}
}

// NewLoginAPIGINRPCAgent new a Login service.
func NewLoginAPIGINRPCAgent(client pb.LoginAPIClient) pb.LoginAPIGINRPCAgent {
	return &LoginAPIGINRPCService{client: client}
}
func NewLoginServerAgent(client *service.GRPCClient) pb.LoginAPIGINRPCAgent {
	cli := pb.NewLoginAPIClient(client)
	return NewLoginAPIGINRPCAgent(cli)
}

func NewLoginServerAgentGINRegister(client *service.GRPCClient) func(server gins.IRouter) {
	cli := NewLoginServerAgent(client)
	return func(server gins.IRouter) {
		pb.RegisterLoginAPIGINRPCAgent(server, cli)
	}
}

var _ pb.LoginAPIGINRPCAgent = (*LoginAPIGINRPCService)(nil)
