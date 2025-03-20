/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/system"
)

type (
	ListPoliciesRequest   = pb.ListPoliciesRequest
	ListPoliciesResponse  = pb.ListPoliciesResponse
	ListGroupingsRequest  = pb.ListGroupingsRequest
	ListGroupingsResponse = pb.ListGroupingsResponse
	//WatchUpdateRequest    = pb.WatchUpdateRequest
	//WatchUpdateResponse   = pb.WatchUpdateResponse
)

type CasbinSourceRepo interface {
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	ListGroupings(context.Context, *ListGroupingsRequest) (*ListGroupingsResponse, error)
	//WatchUpdate(context.Context, *WatchUpdateRequest) (*WatchUpdateResponse, error)
}
