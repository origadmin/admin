/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto is the data transfer object package for the system module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type PersonalRepo interface {
	GetPersonalProfile(ctx context.Context, in *pb.GetPersonalProfileRequest) (*pb.GetPersonalProfileResponse, error)
	ListPersonalRoles(ctx context.Context, in *pb.ListPersonalRolesRequest) (*pb.ListPersonalRolesResponse, error)
	UpdatePersonalPassword(ctx context.Context, in *pb.UpdatePersonalPasswordRequest) (*pb.UpdatePersonalPasswordResponse, error)
	UpdatePersonalProfile(ctx context.Context, in *pb.UpdatePersonalProfileRequest) (*pb.UpdatePersonalProfileResponse, error)
	ListPersonalResources(ctx context.Context, in *pb.ListPersonalResourcesRequest) (*pb.ListPersonalResourcesResponse, error)
}
