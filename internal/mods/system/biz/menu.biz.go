/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package biz is a biz layer for the system module of OrigAdmin.
package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/grpc"

	pb "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/mods/system/dto"
)

// MenusBiz is a Menu use case.
type MenusBiz struct {
	dao     dto.MenuRepo
	limiter pagination.PageLimiter
	log     *log.Helper
}

func (biz MenusBiz) ListMenus(ctx context.Context, in *pb.ListMenusRequest, opts ...grpc.CallOption) (*pb.ListMenusResponse, error) {
	var option dto.MenuQueryOption
	if err := option.FromListRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("ListMenus")
	result, total, err := biz.dao.List(ctx, in, option)
	if err != nil {
		return nil, err
	}
	return dto.ListMenuResponse(result, in, int32(total))
}

func (biz MenusBiz) GetMenu(ctx context.Context, in *pb.GetMenuRequest, opts ...grpc.CallOption) (*pb.GetMenuResponse, error) {
	var option dto.MenuQueryOption
	if err := option.FromGetRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("GetMenu")
	result, err := biz.dao.Get(ctx, in.GetId(), option)
	if err != nil {
		return nil, err
	}
	return &pb.GetMenuResponse{
		Menu: result,
	}, nil
}

func (biz MenusBiz) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest, opts ...grpc.CallOption) (*pb.CreateMenuResponse, error) {
	var option dto.MenuQueryOption
	if err := option.FromCreateRequest(in, biz.limiter); err != nil {
		return nil, err
	}
	log.Info("CreateMenu")
	result, err := biz.dao.Create(ctx, in.Menu, option)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMenuResponse{
		Menu: result,
	}, nil
}

func (biz MenusBiz) UpdateMenu(ctx context.Context, in *pb.UpdateMenuRequest, opts ...grpc.CallOption) (*pb.UpdateMenuResponse, error) {
	//var option dto.UpdateMenuOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	log.Info("UpdateMenu")
	result, err := biz.dao.Update(ctx, in.Menu)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateMenuResponse{
		Menu: result,
	}, nil
}

func (biz MenusBiz) DeleteMenu(ctx context.Context, in *pb.DeleteMenuRequest, opts ...grpc.CallOption) (*pb.DeleteMenuResponse, error) {
	//var option dto.DeleteMenuOption
	//if err := option.FromListRequest(in, biz.limiter); err != nil {
	//	return nil, err
	//}
	//_, err := biz.dao.Get(ctx, in.GetId())
	//if err != nil {
	//	return nil, err
	//}
	if err := biz.dao.Delete(ctx, in.GetId()); err != nil {
		return nil, err
	}
	return &pb.DeleteMenuResponse{}, nil
}

// NewMenusBiz new a Menu use case.
func NewMenusBiz(repo dto.MenuRepo, logger log.Logger) *MenusBiz {
	return &MenusBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

// NewMenusClient new a Menu use case.
func NewMenusClient(repo dto.MenuRepo, logger log.Logger) pb.MenuAPIClient {
	return &MenusBiz{dao: repo, limiter: defaultLimiter, log: log.NewHelper(logger)}
}

var _ pb.MenuAPIClient = (*MenusBiz)(nil)
