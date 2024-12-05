/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/common"
	"origadmin/application/admin/helpers/resp"
)

// LoginAPIGINRPCService is a Login service.
type LoginAPIGINRPCService struct {
	//pb.UnimplementedLoginAPIServer

	client pb.LoginAPIClient
}

func (l LoginAPIGINRPCService) CaptchaID(context *gins.Context, request *pb.CaptchaIDRequest) {
	log.Infof("CaptchaID Request %+v", request)
	response, err := l.client.CaptchaID(context, request)
	log.Infof("CaptchaID: %v", response)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response.Data
	}
	resp.Result(context, res, err)
}

func (l LoginAPIGINRPCService) CaptchaImage(context *gins.Context, request *pb.CaptchaImageRequest) {
	response, err := l.client.CaptchaImage(context, request)
	if err != nil {
		gins.ResultError(context, err)
		return
	}
	if _, err := context.Writer.Write(response.Image); err != nil {
		gins.ResultError(context, err)
		return
	}
	for k, v := range response.Headers {
		context.Writer.Header().Set(k, v)
	}
	context.Writer.Flush()
	context.Writer.WriteHeader(200)
	return
}

func (l LoginAPIGINRPCService) CurrentMenus(context *gins.Context, request *pb.CurrentMenusRequest) {
	response, err := l.client.CurrentMenus(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	resp.Result(context, res, err)
}

func (l LoginAPIGINRPCService) CurrentUser(context *gins.Context, request *pb.CurrentUserRequest) {
	response, err := l.client.CurrentUser(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	resp.Result(context, res, err)
}

func (l LoginAPIGINRPCService) Login(context *gins.Context, request *pb.LoginRequest) {
	response, err := l.client.Login(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	resp.Result(context, res, err)
}

func (l LoginAPIGINRPCService) Logout(context *gins.Context, request *pb.LogoutRequest) {
	response, err := l.client.Logout(context, request)
	var res gins.Result
	if err == nil {
		res.Success = true
		res.Data = response
	}
	resp.Result(context, res, err)
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

var _ pb.LoginAPIGINRPCAgent = (*LoginAPIGINRPCService)(nil)
