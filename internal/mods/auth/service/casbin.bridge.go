/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/auth"
)

// CasbinServiceHookedBridge is a Casbin service.
type CasbinServiceHookedBridge struct {
	pb.UnimplementedCasbinSourceServiceHooked
	log *log.KHelper
}

func (c CasbinServiceHookedBridge) CompleteListGroupings(context http.Context, request *pb.ListGroupingsRequest, response *pb.ListGroupingsResponse) error {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHookedBridge) PrepareListPolicies(context http.Context, request *pb.ListPoliciesRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHookedBridge) CompleteListPolicies(context http.Context, request *pb.ListPoliciesRequest, response *pb.ListPoliciesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHookedBridge) PrepareWatchUpdate(context http.Context, request *pb.WatchUpdateRequest) (context.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (c CasbinServiceHookedBridge) CompleteWatchUpdate(context http.Context, request *pb.WatchUpdateRequest, response *pb.WatchUpdateResponse) error {
	//TODO implement me
	panic("implement me")
}

//func (s CasbinServiceHookedBridge) PersonalLogout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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
//func (s CasbinServiceHookedBridge) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
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
//func (s CasbinServiceHookedBridge) Captcha(ctx context.Context, request *pb.CaptchaRequest) (*pb.CaptchaResponse, error) {
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
//func (s CasbinServiceHookedBridge) CaptchaId(ctx context.Context, request *pb.CaptchaIdRequest) (*pb.CaptchaIdResponse, error) {
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
//func (s CasbinServiceHookedBridge) CaptchaAudio(ctx context.Context, request *pb.CaptchaAudioRequest) (*pb.CaptchaAudioResponse, error) {
//	_, err := s.client.CaptchaAudio(ctx, request)
//	if err != nil {
//		log.Errorf("Logout error: %v", err)
//		return nil, err
//	}
//	return nil, nil
//}
//func (s CasbinServiceHookedBridge) CaptchaImage(ctx context.Context, request *pb.CaptchaImageRequest) (*pb.CaptchaImageResponse, error) {
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
//func (s CasbinServiceHookedBridge) TokenRefresh(ctx context.Context, request *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
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
//func (s CasbinServiceHookedBridge) Casbin(ctx context.Context, request *pb.CasbinRequest) (*pb.CasbinResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Casbin(ctx, request)
//	if err != nil {
//		log.Errorf("Casbin error: %v", err)
//		return nil, err
//	}
//	token := resp.FromToken(response.Token)
//	log.Debugf("Casbin: Token:%+v", token)
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    token,
//	})
//	return nil, nil
//}
//
//func (s CasbinServiceHookedBridge) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.LogoutResponse, error) {
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

func NewCasbinServiceHookedBridge(r runtime.Runtime, client pb.CasbinSourceServiceHTTPServer) pb.
CasbinSourceServiceHookedBridger {
	return pb.WithCasbinSourceServiceHook(&CasbinServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/permission")),
	})(client)
}

// NewCasbinServiceBridge new a menu service.
func NewCasbinServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.CasbinSourceServiceServer {
	return pb.NewCasbinSourceServiceBridge(client)
}

func NewCasbinServiceBridgeClient(r runtime.Runtime, clients map[string]*service.GRPCClient) pb.CasbinSourceServiceServer {
	if v, ok := clients["auth"]; ok {
		return NewCasbinServiceBridge(r, v)
	} else {
		return pb.UnimplementedCasbinSourceServiceServer{}
	}
}

// NewCasbinServiceHTTPBridge new a menu service.
func NewCasbinServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.CasbinSourceServiceHTTPServer {
	return pb.NewCasbinSourceServiceHTTPBridge(client)
}
