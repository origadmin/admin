/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"context"

	pb "origadmin/application/admin/api/v1/services/auth"
)

type (
	ListPoliciesRequest   = pb.ListPoliciesRequest
	ListPoliciesResponse  = pb.ListPoliciesResponse
	ListGroupingsRequest  = pb.ListGroupingsRequest
	ListGroupingsResponse = pb.ListGroupingsResponse
)

type CasbinRepo interface {
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
	ListGroupings(context.Context, *ListGroupingsRequest) (*ListGroupingsResponse, error)
}
