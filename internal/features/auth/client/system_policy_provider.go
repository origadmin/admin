/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package client

import (
	"context"

	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
	"origadmin/application/admin/internal/features/auth/biz"
)

// systemPolicyProvider implements the biz.PolicyProvider interface by calling
// the system service via gRPC.
type systemPolicyProvider struct {
	client systemv1.AuthorizationServiceClient
	log    *log.Helper
}

// NewSystemPolicyProvider creates a new systemPolicyProvider.
func NewSystemPolicyProvider(client systemv1.AuthorizationServiceClient, logger log.Logger) biz.PolicyProvider {
	return &systemPolicyProvider{
		client: client,
		log:    log.NewHelper(log.With(logger, "module", "client.system_policy_provider")),
	}
}

// ListAllPolicies calls the remote ListAllPolicies method on the system service.
func (p *systemPolicyProvider) ListAllPolicies(ctx context.Context) (*systemv1.ListAllPoliciesResponse, error) {
	p.log.WithContext(ctx).Info("Fetching all policies from remote system service...")
	// The actual gRPC call.
	return p.client.ListAllPolicies(ctx, &systemv1.ListAllPoliciesRequest{})
}

// ListPoliciesForRoles calls the remote ListPoliciesForRoles method on the system service.
func (p *systemPolicyProvider) ListPoliciesForRoles(ctx context.Context, roleKeywords ...string) ([]*systemv1.AccessRule, error) {
	p.log.WithContext(ctx).Infof("Fetching policies for roles %v from remote system service...", roleKeywords)
	resp, err := p.client.ListPoliciesForRoles(ctx, &systemv1.ListPoliciesForRolesRequest{
		RoleKeywords: roleKeywords,
	})
	if err != nil {
		return nil, err
	}
	return resp.GetAccessRules(), nil
}
