/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package service

import (
	"context"

	"github.com/origadmin/runtime/service/transport"
	systemv1 "origadmin/application/admin/api/v1/services/system"
)

type SystemService struct {
	Resource    *ResourceService
	Role        *RoleService
	User        *UserService
	Permission  *PermissionService
	View        *ViewService
	PolicyQuery *PolicyQueryService
}

func (s *SystemService) RegisterHTTP(_ context.Context, srv *transport.HTTPServer) error {
	// Register HTTP handlers
	systemv1.RegisterUserServiceHTTPServer(srv, s.User)
	systemv1.RegisterRoleServiceHTTPServer(srv, s.Role)
	systemv1.RegisterPermissionServiceHTTPServer(srv, s.Permission)
	systemv1.RegisterResourceServiceHTTPServer(srv, s.Resource)
	systemv1.RegisterViewServiceHTTPServer(srv, s.View)
	return nil
}

func (s *SystemService) RegisterGRPC(ctx context.Context, srv *transport.GRPCServer) error {
	// Register gRPC handlers
	systemv1.RegisterUserServiceServer(srv, s.User)
	systemv1.RegisterRoleServiceServer(srv, s.Role)
	systemv1.RegisterPermissionServiceServer(srv, s.Permission)
	systemv1.RegisterResourceServiceServer(srv, s.Resource)
	systemv1.RegisterViewServiceServer(srv, s.View)
	systemv1.RegisterPolicyQueryServiceServer(srv, s.PolicyQuery)
	return nil
}

func NewSystemService(
	resource *ResourceService,
	role *RoleService,
	user *UserService,
	permission *PermissionService,
	view *ViewService,
	policyQuery *PolicyQueryService,
) *SystemService {
	return &SystemService{
		Resource:    resource,
		Role:        role,
		User:        user,
		Permission:  permission,
		View:        view,
		PolicyQuery: policyQuery,
	}
}
