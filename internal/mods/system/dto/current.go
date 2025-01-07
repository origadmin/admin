/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type CurrentRepo interface {
	GetCurrentUser(ctx context.Context, in *pb.GetCurrentUserRequest) (*pb.GetCurrentUserResponse, error)
	ListCurrentRoles(ctx context.Context, in *pb.ListCurrentRolesRequest) (*pb.ListCurrentRolesResponse, error)
	UpdateCurrentUserPassword(ctx context.Context, in *pb.UpdateCurrentUserPasswordRequest) (*pb.UpdateCurrentUserPasswordResponse, error)
	UpdateCurrentUser(ctx context.Context, in *pb.UpdateCurrentUserRequest) (*pb.UpdateCurrentUserResponse, error)
	ListCurrentResources(ctx context.Context, in *pb.ListCurrentResourcesRequest) (*pb.ListCurrentResourcesResponse, error)
}
