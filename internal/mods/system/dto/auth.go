/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	"github.com/origadmin/toolkits/net/pagination"

	pb "origadmin/application/admin/api/v1/services/system"
)

type (
	ListAuthResourcesRequest  = pb.ListAuthResourcesRequest
	ListAuthResourcesResponse = pb.ListAuthResourcesResponse
)

// AuthRepo is a Auth repository interface.
type AuthRepo interface {
	ListAuthResources(context.Context, *ListAuthResourcesRequest, ...AuthResourceQueryOption) ([]*ResourcePB, int32, error)
	CreateToken(context.Context, *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
	ValidateToken(context.Context, *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error)
	DestroyToken(context.Context, *pb.DestroyTokenRequest) (*pb.DestroyTokenResponse, error)
	Authenticate(context.Context, *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error)
	AuthLogout(context.Context, *pb.AuthLogoutRequest) (*pb.AuthLogoutResponse, error)
}

type AuthResourceQueryOption struct {
	SelectFields []string
	OmitFields   []string
	OrderFields  []string
	Fields       []string
}

func (o AuthResourceQueryOption) FromListRequest(in *ListAuthResourcesRequest, limiter pagination.PageLimiter) error {
	in.Current = limiter.Current(in.Current)
	in.PageSize = limiter.PerPage(in.PageSize)
	return nil
}

func ToListAuthResourcesResponse(result []*ResourcePB, in *ListAuthResourcesRequest, total int32, args ...any) (*ListAuthResourcesResponse, error) {
	response := &ListAuthResourcesResponse{
		Resources: result,
		TotalSize: total,
	}
	return response, nil
}
