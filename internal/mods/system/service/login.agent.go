/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

// LoginAPIAgentService is a Login service.
type LoginAPIAgentService struct {
	resp.Response

	client pb.LoginAPIClient
}

func (s LoginAPIAgentService) CurrentLogout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Logout(ctx, request)
	if err != nil {
		log.Errorf("Logout error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s LoginAPIAgentService) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Register(ctx, request)
	if err != nil {
		log.Errorf("Register error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaResource(ctx context.Context, request *pb.CaptchaResourceRequest) (*pb.CaptchaResourceResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CaptchaResource(ctx, request)
	if err != nil {
		log.Errorf("CaptchaResource error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaResources(ctx context.Context, request *pb.CaptchaResourcesRequest) (*pb.CaptchaResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CaptchaResources(ctx, request)
	if err != nil {
		log.Errorf("CaptchaResources error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaID(ctx context.Context, request *pb.CaptchaIDRequest) (*pb.CaptchaIDResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	log.Debugf("CaptchaID: Request:%+v", request)
	response, err := s.client.CaptchaID(ctx, request)
	log.Debugf("CaptchaID: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		return nil, err
	}

	s.JSON(httpCtx, http.StatusOK, &resp.StringResult{
		Success: true,
		Data:    response.Data,
	})
	return nil, nil
}

func (s LoginAPIAgentService) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	log.Debugf("CaptchaImage: Request:%+v", request)
	response, err := s.client.CaptchaImage(ctx, request)
	log.Debugf("CaptchaImage: Response:%+v, Error:%+v", response, err)
	if err != nil {
		log.Errorf("CaptchaImage error: %v", err)
		return nil, err
	}
	log.Debugf("CaptchaImage: Setting headers: %+v", response.Headers)
	for k, v := range response.Headers {
		httpCtx.Response().Header().Set(k, v)
	}
	log.Debugf("CaptchaImage: Writing response headers")
	httpCtx.Response().WriteHeader(http.StatusOK)
	log.Debugf("CaptchaImage: Writing response image")
	if _, err := httpCtx.Response().Write(response.Image); err != nil {
		log.Errorf("CaptchaImage error writing response: %v", err)
		return nil, err
	}
	//log.Debugf("CaptchaImage: Flushing response writer")
	//context.Response().Flush()
	log.Debugf("CaptchaImage: Completed successfully")
	return nil, nil
}

func (s LoginAPIAgentService) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.TokenRefresh(ctx, request)
	if err != nil {
		log.Errorf("Refresh error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(resp.FromToken(response.Token)),
	})
	return nil, nil
}

func (s LoginAPIAgentService) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Login(ctx, request)
	if err != nil {
		log.Errorf("Login error: %v", err)
		return nil, err
	}
	token := resp.FromToken(response.Token)
	log.Debugf("Login: Token:%+v", token)
	s.JSON(httpCtx, http.StatusOK, &resp.Result{
		Success: true,
		Data:    token,
	})
	return nil, nil
}

func (s LoginAPIAgentService) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Logout(ctx, request)
	if err != nil {
		log.Errorf("Logout error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
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
		pb.RegisterLoginAPIAgent(agent.NewHTTP(server), cli)
	}
}

var _ pb.LoginAPIAgent = (*LoginAPIAgentService)(nil)
