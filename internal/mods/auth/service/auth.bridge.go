/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	context2 "context"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/auth"
	typespb "origadmin/application/admin/api/v1/services/types"
	"origadmin/application/admin/helpers/resp"
)

var (
	ErrorInvalidToken = typespb.ErrorSystemErrorReasonInvalidToken("invalid token")
)

// AuthServiceHookedBridge is a menu service.
type AuthServiceHookedBridge struct {
	pb.UnimplementedAuthServiceHooked
	log *log.KHelper
}

func (s AuthServiceHookedBridge) PrepareAuthLogout(h transhttp.Context, request *pb.AuthLogoutRequest) (context2.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteAuthLogout(h transhttp.Context, request *pb.AuthLogoutRequest, response *pb.AuthLogoutResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteAuthenticate(h transhttp.Context, request *pb.AuthenticateRequest, response *pb.AuthenticateResponse) error {
	if !response.IsValid {
		return ErrorInvalidToken
	}
	return h.JSON(http.StatusOK, &resp.Result{
		Success: true,
	})
}

func (s AuthServiceHookedBridge) PrepareCreateToken(h transhttp.Context, request *pb.CreateTokenRequest) (context2.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteCreateToken(h transhttp.Context, request *pb.CreateTokenRequest, response *pb.CreateTokenResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) PrepareDestroyToken(h transhttp.Context, request *pb.DestroyTokenRequest) (context2.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteDestroyToken(h transhttp.Context, request *pb.DestroyTokenRequest, response *pb.DestroyTokenResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) PrepareListAuthResources(h transhttp.Context, request *pb.ListAuthResourcesRequest) (context2.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteListAuthResources(h transhttp.Context, request *pb.ListAuthResourcesRequest, response *pb.ListAuthResourcesResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) PrepareValidateToken(h transhttp.Context, request *pb.ValidateTokenRequest) (context2.Context, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) CompleteValidateToken(h transhttp.Context, request *pb.ValidateTokenRequest, response *pb.ValidateTokenResponse) error {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceHookedBridge) AuthLogout(ctx context.Context, request *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

//func (s AuthServiceHookedBridge) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.Authenticate(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	if !response.IsValid {
//		return nil, ErrorInvalidToken
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//	})
//	return nil, nil
//}
//
//func (s AuthServiceHookedBridge) CreateToken(ctx context.Context, request *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.CreateToken(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}
//
//func (s AuthServiceHookedBridge) DestroyToken(ctx context.Context, request *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.DestroyToken(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}
//
//func (s AuthServiceHookedBridge) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.ValidateToken(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Result{
//		Success: true,
//		Data:    response,
//	})
//	return nil, nil
//}
//
//func (s AuthServiceHookedBridge) ListAuthResources(ctx context.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
//	httpCtx := agent.FromHTTPContext(ctx)
//	response, err := s.client.ListAuthResources(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	s.JSON(httpCtx, http.StatusOK, &resp.Page{
//		Success: true,
//		Total:   response.TotalSize,
//		Data:    resp.Proto2AnyPBArray(response.Resources...),
//	})
//	return nil, nil
//}

func NewAuthServiceHookedBridge(r runtime.Runtime, client pb.AuthServiceHTTPServer) pb.AuthServiceHookedBridger {
	return pb.WithAuthServiceHook(&AuthServiceHookedBridge{
		log: log.NewHelper(r.WithLogger("module", "service/permission")),
	})(client)
}

// NewAuthServiceBridge new a menu service.
func NewAuthServiceBridge(r runtime.Runtime, client *service.GRPCClient) pb.AuthServiceServer {
	return pb.NewAuthServiceBridge(client)
}

// NewAuthServiceHTTPBridge new a menu service.
func NewAuthServiceHTTPBridge(r runtime.Runtime, client *service.HTTPClient) pb.AuthServiceHTTPServer {
	return pb.NewAuthServiceHTTPBridge(client)
}
