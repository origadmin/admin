/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/runtime/service"

	"origadmin/application/admin/internal/mods/system/biz"
)

type UserService struct {
	grpcClient *service.GRPCClient
	httpClient *service.HTTPClient
	biz        *biz.UserServiceBiz    `wire:"-"`
	grpc       *UserServiceServer     `wire:"-"`
	http       *UserServiceHTTPServer `wire:"-"`
}

//func NewUserService(r runtime.Runtime, bootstrap *configs.Bootstrap, service *UserService) pb.UserServiceServer {
//	if r.IsClient() {
//		return NewUserServiceBridge(r, service.grpcClient)
//	}
//}
