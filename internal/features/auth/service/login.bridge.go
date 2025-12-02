/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"
	"google.golang.org/protobuf/encoding/protojson"

	pb "origadmin/application/admin/api/v1/services/auth"
	"origadmin/application/admin/helpers/resp"
)

// LoginServiceHookedBridge is a Login service.
type LoginServiceHookedBridge struct {
	pb.UnimplementedLoginServiceHooked
	log *log.KHelper
}

func (s LoginServiceHookedBridge) CompleteCaptcha(ctx transhttp.Context, request *pb.CaptchaRequest, response *pb.CaptchaResponse) error {
	marshal, err := protojson.Marshal(response)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.Result{
		Success: true,
		Data:    marshal,
	})
}

func (s LoginServiceHookedBridge) CompleteCaptchaAudio(ctx transhttp.Context, request *pb.CaptchaAudioRequest, response *pb.CaptchaAudioResponse) error {
	return ctx.JSON(http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
}

func (s LoginServiceHookedBridge) CompleteCaptchaId(ctx transhttp.Context, request *pb.CaptchaIdRequest, response *pb.CaptchaIdResponse) error {
	return ctx.JSON(http.StatusOK, &resp.Data{
		Success: true,
		Data:    resp.Proto2Any(response),
	})
}

func (s LoginServiceHookedBridge) CompleteCaptchaImage(ctx transhttp.Context, request *pb.CaptchaImageRequest, response *pb.CaptchaImageResponse) error {
	s.log.Debugf("CaptchaImage: Setting headers: %+v", response.Headers)
	for k, v := range response.Headers {
		ctx.Response().Header().Set(k, v)
	}
	s.log.Debugf("CaptchaImage: Writing response headers")
	ctx.Response().WriteHeader(http.StatusOK)
	s.log.Debugf("CaptchaImage: Writing response image")
	if _, err := ctx.Response().Write(response.Image); err != nil {
		log.Errorf("CaptchaImage error writing response: %v", err)
		return err
	}
	s.log.Debugf("CaptchaImage: Completed successfully")
	return nil
}

func (s LoginServiceHookedBridge) CompleteLogin(ctx transhttp.Context, request *pb.LoginRequest, response *pb.LoginResponse) error {
	marshal, err := protojson.Marshal(resp.FromToken(response.Token))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.Result{
		Success: true,
		Data:    marshal,
	})
}

func (s LoginServiceHookedBridge) PrepareLogout(ctx transhttp.Context, request *pb.LogoutRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s LoginServiceHookedBridge) CompleteLogout(ctx transhttp.Context, request *pb.LogoutRequest, response *pb.LogoutResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s LoginServiceHookedBridge) PrepareRegister(ctx transhttp.Context, request *pb.RegisterRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s LoginServiceHookedBridge) CompleteRegister(ctx transhttp.Context, request *pb.RegisterRequest, response *pb.RegisterResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s LoginServiceHookedBridge) CompleteTokenRefresh(ctx transhttp.Context, request *pb.TokenRefreshRequest, response *pb.TokenRefreshResponse) error {
	marshal, err := protojson.Marshal(resp.FromToken(response.Token))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &resp.Result{
		Success: true,
		Data:    marshal,
	})
}

//func (s LoginServiceHookedBridge) PersonalLogout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Logout(ctx, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    resp.Proto2Any(response),
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Register(ctx, request)
//	if err != nil {
//		log.Errorf("Register error: %v", err)
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    resp.Proto2Any(response),
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) Captcha(ctx context.Context, request *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Captcha(ctx, request)
//	if err != nil {
//		log.Errorf("Captcha error: %v", err)
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    resp.Proto2Any(response),
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) CaptchaId(ctx context.Context, request *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	log.Debugf("CaptchaId: Request:%+v", request)
//	response, err := s.client.CaptchaId(ctx, request)
//	log.Debugf("CaptchaId: Response:%+v, Error:%+v", response, err)
//	if err != nil {
//		log.Errorf("CaptchaImage error: %v", err)
//		return nil, err
//	}
//
//	s.JSON(httpCtx, http.StatusOK, &resp.StringResult{
//		Success: true,
//		Data:    response.Data,
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) CaptchaAudio(ctx context.Context, request *pb.CaptchaAudioRequest) (*pb.CaptchaAudioResponse, error) {
//	_, err := s.client.CaptchaAudio(ctx, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	return nil, nil
//}
//func (s LoginServiceHookedBridge) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	log.Debugf("CaptchaImage: Request:%+v", request)
//	response, err := s.client.CaptchaImage(ctx, request)
//	log.Debugf("CaptchaImage: Response:%+v, Error:%+v", response, err)
//	if err != nil {
//		log.Errorf("CaptchaImage error: %v", err)
//		return nil, err
//	}
//	log.Debugf("CaptchaImage: Setting headers: %+v", response.Headers)
//	for k, v := range response.Headers {
//		httpCtx.Response().Header().Set(k, v)
//	}
//	log.Debugf("CaptchaImage: Writing response headers")
//	httpCtx.Response().WriteHeader(http.StatusOK)
//	log.Debugf("CaptchaImage: Writing response image")
//	if _, err := httpCtx.Response().Write(response.Image); err != nil {
//		log.Errorf("CaptchaImage error writing response: %v", err)
//		return nil, err
//	}
//	//log.Debugf("CaptchaImage: Flushing response writer")
//	//context.Response().Flush()
//	log.Debugf("CaptchaImage: Completed successfully")
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.TokenRefresh(ctx, request)
//	if err != nil {
//		log.Errorf("Refresh error: %v", err)
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    resp.Proto2Any(resp.FromToken(response.Token)),
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Login(ctx, request)
//	if err != nil {
//		log.Errorf("Login error: %v", err)
//		return nil, err
//	}
//	token := resp.FromToken(response.Token)
//	log.Debugf("Login: Token:%+v", token)
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    token,
//	})
//	return nil, nil
//}
//
//func (s LoginServiceHookedBridge) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Logout(ctx, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Data{
//		Success: true,
//		Data:    resp.Proto2Any(response),
//	})
//	return nil, nil
//}

func NewLoginServiceHookedBridge(r runtime.Runtime, client pb.LoginServiceHTTPServer) pb.LoginServiceHookedBridger {
	return pb.WithLoginServiceHook(&LoginServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/auth")),
	})(client)
}

// NewLoginServiceBridge new a menu service.
func NewLoginServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.LoginServiceServer {
	return pb.NewLoginServiceBridge(client)
}

func NewLoginServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.LoginServiceServer {
	if v, ok := clients["auth"]; ok {
		return NewLoginServiceBridge(r, v)
	} else {
		return pb.UnimplementedLoginServiceServer{}
	}
}

// NewLoginServiceHTTPBridge new a menu service.
func NewLoginServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.LoginServiceHTTPServer {
	return pb.NewLoginServiceHTTPBridge(client)
}
