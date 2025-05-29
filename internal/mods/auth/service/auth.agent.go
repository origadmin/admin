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

// AuthServiceAgent is a menu service.
type AuthServiceAgent struct {
	resp.Response

	client pb.AuthServiceClient
}

func (s AuthServiceAgent) AuthLogout(ctx context.Context, request *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s AuthServiceAgent) Authenticate(ctx context.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
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

func (s AuthServiceAgent) CreateToken(ctx context.Context, request *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
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

func (s AuthServiceAgent) DestroyToken(ctx context.Context, request *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
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

func (s AuthServiceAgent) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
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

func (s AuthServiceAgent) ListAuthResources(ctx context.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
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

// NewAuthServiceAgent new a menu service.
func NewAuthServiceAgent(client pb.AuthServiceClient) *AuthServiceAgent {
	return &AuthServiceAgent{client: client}
}

// NewAuthServiceAgentPB new a menu service.
func NewAuthServiceAgentPB(client pb.AuthServiceClient) pb.AuthServiceAgent {
	return &AuthServiceAgent{client: client}
}
func NewAuthServiceAgentClient(client *service.GRPCClient) pb.AuthServiceAgent {
	cli := pb.NewAuthServiceClient(client)
	return NewAuthServiceAgent(cli)
}

var _ pb.AuthServiceAgent = (*AuthServiceAgent)(nil)
