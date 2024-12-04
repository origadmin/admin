/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"github.com/origadmin/contrib/transport/gins"

	pb "origadmin/application/admin/api/v1/services/common"
)

// LoginAPIGINService is a Login service.
type LoginAPIGINService struct {
	//pb.UnimplementedLoginAPIServer

	client pb.LoginAPIClient
}

func (l LoginAPIGINService) CaptchaID(context *gins.Context, request *pb.CaptchaIDRequest) {
	//TODO implement me
	panic("implement me")
}

func (l LoginAPIGINService) CaptchaImage(context *gins.Context, request *pb.CaptchaImageRequest) {
	//TODO implement me
	panic("implement me")
}

func (l LoginAPIGINService) CurrentMenus(context *gins.Context, request *pb.CurrentMenusRequest) {
	//TODO implement me
	panic("implement me")
}

func (l LoginAPIGINService) CurrentUser(context *gins.Context, request *pb.CurrentUserRequest) {
	//TODO implement me
	panic("implement me")
}

func (l LoginAPIGINService) Login(context *gins.Context, request *pb.LoginRequest) {
	//TODO implement me
	panic("implement me")
}

func (l LoginAPIGINService) Logout(context *gins.Context, request *pb.LogoutRequest) {
	//TODO implement me
	panic("implement me")
}

// NewLoginAPIGINService new a Login service.
func NewLoginAPIGINService(client pb.LoginAPIClient) *LoginAPIGINService {
	return &LoginAPIGINService{client: client}
}

// NewLoginAPIGINServer new a Login service.
func NewLoginAPIGINServer(client pb.LoginAPIClient) pb.LoginAPIGINRPCAgent {
	return &LoginAPIGINService{client: client}
}

var _ pb.LoginAPIGINRPCAgent = (*LoginAPIGINService)(nil)
