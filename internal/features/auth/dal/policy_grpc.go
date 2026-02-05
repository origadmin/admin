/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/auth/dto"
)

// policyGRPCProvider implements PolicyProvider using gRPC calls to the system service.
// This is used when auth and system services have separate databases.
type policyGRPCProvider struct {
	client systemv1.AuthorizationServiceClient
	log    *log.Helper
}

// NewPolicyGRPCProvider creates a PolicyProvider that uses gRPC calls to the system service.
func NewPolicyGRPCProvider(client systemv1.AuthorizationServiceClient, logger log.Logger) dto.PolicyProvider {
	return &policyGRPCProvider{
		client: client,
		log:    log.NewHelper(log.With(logger, "module", "dal.policy_grpc")),
	}
}

// ListAllPolicies retrieves all policies by calling the system service via gRPC.
func (p *policyGRPCProvider) ListAllPolicies(ctx context.Context) (*systemv1.ListAllPoliciesResponse, error) {
	p.log.WithContext(ctx).Info("Listing all policies via gRPC from system service")
	return p.client.ListAllPolicies(ctx, &systemv1.ListAllPoliciesRequest{})
}

// ListPoliciesForRoles retrieves policies for specific roles via gRPC.
func (p *policyGRPCProvider) ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systemv1.AccessRule, error) {
	p.log.WithContext(ctx).Infof("Listing policies for roles: %v via gRPC from system service", roleKeywords)
	resp, err := p.client.ListPoliciesForRoles(ctx, &systemv1.ListPoliciesForRolesRequest{
		RoleKeywords: roleKeywords,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetAccessRules(), nil
}

var _ dto.PolicyProvider = (*policyGRPCProvider)(nil)
