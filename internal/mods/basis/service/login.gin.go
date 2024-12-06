/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

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
	response, err := s.client.CaptchaID(context, request)
	log.Debugf("CaptchaID: Request:%+v, Response:%+v, Error:%+v", request, response, err)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response.Data
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s LoginAPIGINRPCService) CaptchaImage(context *gins.Context, request *pb.CaptchaImageRequest) {
	response, err := s.client.CaptchaImage(context, request)
	if err != nil {
		s.Error(context, 404, err)
		return
	}
	if _, err := context.Writer.Write(response.Image); err != nil {
		s.Error(context, 404, err)
		return
	}
	for k, v := range response.Headers {
		context.Writer.Header().Set(k, v)
	}
	context.Writer.Flush()
	context.Writer.WriteHeader(200)
	return
}

func (s LoginAPIGINRPCService) CurrentMenus(context *gins.Context, request *pb.CurrentMenusRequest) {
	response, err := s.client.CurrentMenus(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s LoginAPIGINRPCService) CurrentUser(context *gins.Context, request *pb.CurrentUserRequest) {
	response, err := s.client.CurrentUser(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s LoginAPIGINRPCService) Login(context *gins.Context, request *pb.LoginRequest) {
	response, err := s.client.Login(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	s.Any(context, http.StatusOK, &res, err)
}

func (s LoginAPIGINRPCService) Logout(context *gins.Context, request *pb.LogoutRequest) {
	response, err := s.client.Logout(context, request)
	var res resp.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	s.Any(context, http.StatusOK, &res, err)
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
