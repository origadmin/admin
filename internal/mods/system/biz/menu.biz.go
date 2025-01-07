/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// MenuServiceClientBiz is a Menu use case.
type MenuServiceClientBiz struct {
	dao     dto.MenuRepo
	limiter pagination.PageLimiter
	log     *log.KHelper
}

func (m MenuServiceClientBiz) ListMenus(ctx context.Context, in *pb.ListMenusRequest, opts ...grpc.CallOption) (*pb.ListMenusResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuServiceClientBiz) GetMenu(ctx context.Context, in *pb.GetMenuRequest, opts ...grpc.CallOption) (*pb.GetMenuResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuServiceClientBiz) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest, opts ...grpc.CallOption) (*pb.CreateMenuResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuServiceClientBiz) UpdateMenu(ctx context.Context, in *pb.UpdateMenuRequest, opts ...grpc.CallOption) (*pb.UpdateMenuResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MenuServiceClientBiz) DeleteMenu(ctx context.Context, in *pb.DeleteMenuRequest, opts ...grpc.CallOption) (*pb.DeleteMenuResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewMenuServiceClientBiz new a Menu use case.
func NewMenuServiceClientBiz(repo dto.MenuRepo, logger log.KLogger) *MenuServiceClientBiz {
	return &MenuServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewMenuServiceClient new a Menu use case.
func NewMenuServiceClient(repo dto.MenuRepo, logger log.KLogger) pb.MenuServiceClient {
	return &MenuServiceClientBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.MenuServiceClient = (*MenuServiceClientBiz)(nil)
