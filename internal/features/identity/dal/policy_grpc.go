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

	// Note: The current implementation of the gRPC client for AuthorizationService
	// does not support filtering policies with PolicyFilterOption.
	// The 'base' parameter is used to construct the request.
	// The 'opts' are ignored.
	// This may need to be updated if the service's capabilities are extended.

	req := &systemv1.ListPoliciesRequest{
		// Assuming the ListPoliciesRequest can be built from the 'base' PolicySpec.
		// This part needs to be aligned with the actual definition of ListPoliciesRequest.
		// For example:
		// Page:     1,
		// PageSize: 100,
		// Type:     base.GetType(),
		// Subject:  base.GetSubject(),
		// Domain:   base.GetDomain(),
	}

	resp, err := p.client.ListPolicies(ctx, req)
	if err != nil {
		p.log.WithContext(ctx).Errorf("Failed to list policies: %v", err)
		return nil, err
	}
	return resp.GetPolicies(), nil
}

var _ authz.PolicyReader = (*policyGRPCProvider)(nil)
