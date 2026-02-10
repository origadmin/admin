/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	systemv1 "origadmin/application/admin/api/v1/services/system"
)

// policyGRPCProvider implements authz.PolicyReader using gRPC calls to the system service.
// This is used when identity and system services have separate databases.
type policyGRPCProvider struct {
	client systemv1.PolicyQueryServiceClient
	log    *log.Helper
}

// NewPolicyGRPCProvider creates a PolicyReader that uses gRPC calls to the system service.
func NewPolicyGRPCProvider(client systemv1.PolicyQueryServiceClient, logger log.Logger) authz.PolicyReader {
	return &policyGRPCProvider{
		client: client,
		log:    log.NewHelper(log.With(logger, "module", "dal.policy_grpc")),
	}
}

// ListPolicies queries policies matching the filter criteria.
func (p *policyGRPCProvider) ListPolicies(ctx context.Context, base *authzv1.PolicySpec, opts ...authz.PolicyFilterOption) ([]*authzv1.PolicySpec, error) {
	p.log.WithContext(ctx).Info("Listing policies via gRPC from system service")

	// Apply filter options to build a complete filter
	filter := authz.BuildPolicyFilter(base, opts...)

	// Build the gRPC request from the filter
	req := &systemv1.ListPoliciesRequest{}
	if filter.Type != nil {
		req.Type = filter.Type
	}
	if filter.Subject != nil {
		req.Subject = filter.Subject
	}
	if len(filter.Actions) > 0 {
		req.Actions = filter.Actions
	}
	if len(filter.Resources) > 0 {
		req.Resources = filter.Resources
	}
	if filter.Effect != nil {
		req.Effect = filter.Effect
	}
	if filter.Domain != nil {
		req.Domain = filter.Domain
	}
	if filter.Disabled != nil {
		req.Disabled = filter.Disabled
	}

	// Default pagination values (can be overridden by options if needed)
	// Note: PolicyFilter does not have Page/PageSize fields, so we use defaults
	req.Page = 1
	req.PageSize = 1000 // Use a large default to get all results

	resp, err := p.client.ListPolicies(ctx, req)
	if err != nil {
		p.log.WithContext(ctx).Errorf("Failed to list policies: %v", err)
		return nil, err
	}
	return resp.GetPolicies(), nil
}

var _ authz.PolicyReader = (*policyGRPCProvider)(nil)
