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

// AuthAPIGINRPCService is a menu service.
type AuthAPIGINRPCService struct {
	resp.Response

	client pb.AuthAPIClient
}

func (s AuthAPIGINRPCService) ListAuthResources(context transhttp.Context, request *pb.ListAuthResourcesRequest) (*pb.ListAuthResourcesResponse, error) {
	response, err := s.client.ListAuthResources(context, request)
	if err != nil {
		return nil, err
	}
	s.JSON(context, http.StatusOK, &resp.Result{
		Success: true,
		Total:   response.TotalSize,
		Data:    response.Resources,
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
