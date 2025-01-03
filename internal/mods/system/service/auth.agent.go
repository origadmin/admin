/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
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

func (s AuthAPIGINRPCService) Authenticate(context transhttp.Context, request *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	response, err := s.client.Authenticate(context, request)
	if err != nil {
		return nil, err
	}
	if !response.IsValid {
		return nil, ErrorInvalidToken
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) CreateToken(context transhttp.Context, request *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	response, err := s.client.CreateToken(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) DestroyToken(context transhttp.Context, request *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error) {
	response, err := s.client.DestroyToken(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) ValidateToken(context transhttp.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	response, err := s.client.ValidateToken(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Data:    response,
	})
	return nil, nil
}

func (s AuthAPIGINRPCService) ListAuthResources(context transhttp.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	response, err := s.client.ListAuthResources(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Page{
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
