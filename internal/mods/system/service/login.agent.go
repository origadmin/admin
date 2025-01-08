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

// LoginServiceAgent is a Login service.
type LoginServiceAgent struct {
	resp.Response

	client pb.LoginServiceClient
}

func (s LoginServiceAgent) PersonalLogout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

func (s LoginServiceAgent) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
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

func (s LoginServiceAgent) Captcha(ctx context.Context, request *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Captcha(ctx, request)
	if err != nil {
		log.Errorf("Captcha error: %v", err)
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Any2AnyPB(response),
	})
	return nil, nil
}

func (s LoginServiceAgent) CaptchaId(ctx context.Context, request *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	log.Debugf("CaptchaId: Request:%+v", request)
	response, err := s.client.CaptchaId(ctx, request)
	log.Debugf("CaptchaId: Response:%+v, Error:%+v", response, err)
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

func (s LoginServiceAgent) CaptchaAudio(ctx context.Context, request *pb.CaptchaAudioRequest) (*pb.CaptchaAudioResponse, error) {
	_, err := s.client.CaptchaAudio(ctx, request)
	if err != nil {
		log.Errorf("Logout error: %v", err)
		return nil, err
	}
	return nil, nil
}
func (s LoginServiceAgent) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
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

func (s LoginServiceAgent) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
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

func (s LoginServiceAgent) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
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

func (s LoginServiceAgent) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

// NewLoginServiceAgent new a Login service.
func NewLoginServiceAgent(client pb.LoginServiceClient) *LoginServiceAgent {
	return &LoginServiceAgent{client: client}
}

// NewLoginServiceAgentPB new a Login service.
func NewLoginServiceAgentPB(client pb.LoginServiceClient) pb.LoginServiceAgent {
	return &LoginServiceAgent{client: client}
}
func NewLoginServiceAgentClient(client *service.GRPCClient) pb.LoginServiceAgent {
	cli := pb.NewLoginServiceClient(client)
	return NewLoginServiceAgent(cli)
}

var _ pb.LoginServiceAgent = (*LoginServiceAgent)(nil)
