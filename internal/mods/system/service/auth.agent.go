/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	"github.com/origadmin/runtime/agent"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/service"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/helpers/resp"
)

var (
	ErrorInvalidToken = pb.ErrorSystemErrorReasonInvalidToken("invalid token")
)

// AuthAPIGINRPCService is a menu service.
type AuthAPIGINRPCService struct {
	resp.Response

	client pb.AuthAPIClient
}

func (s AuthAPIGINRPCService) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.Authenticate(ctx, request)
	if err != nil {
		return nil, err
	}
	if !response.IsValid {
		return nil, ErrorInvalidToken
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Result{
		Success: true,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) CreateToken(ctx context.Context, request *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.CreateToken(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) DestroyToken(ctx context.Context, request *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.DestroyToken(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ValidateToken(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) ListAuthResources(ctx context.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	httpCtx := agent.FromHTTPContext(ctx)
	response, err := s.client.ListAuthResources(ctx, request)
	if err != nil {
		return nil, err
	}
	s.JSON(httpCtx, http.StatusOK, &resp.Page{
		Success: true,
		Total:   response.TotalSize,
		Data:    resp.Proto2AnyPBArray(response.Resources...),
	})
	return nil, nil
}

// NewAuthAPIGINRPCService new a menu service.
func NewAuthAPIGINRPCService(client pb.AuthAPIClient) *AuthAPIGINRPCService {
	return &AuthAPIGINRPCService{client: client}
}

// NewAuthAPIAgent new a menu service.
func NewAuthAPIAgent(client pb.AuthAPIClient) pb.AuthAPIAgent {
	return &AuthAPIGINRPCService{client: client}
}
func NewAuthServerAgent(client *service.GRPCClient) pb.AuthAPIAgent {
	cli := pb.NewAuthAPIClient(client)
	return NewAuthAPIAgent(cli)
}

var _ pb.AuthAPIAgent = (*AuthAPIGINRPCService)(nil)
