/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
)

// policyProvider implements authz.PolicyReader interface using direct database access.
// It reads policy data directly from the 'casbin_rule' table.
type policyProvider struct {
	db  *ent.Database
	log *log.Helper
}

// NewPolicyProvider creates a new PolicyReader that uses direct database access.
func NewPolicyProvider(db *ent.Database, logger log.Logger) authz.PolicyReader {
	return &policyProvider{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "dal.policy_db")),
	}
}

// ListPolicies fetches all authorization policies directly from the casbin_rule table.
func (p *policyProvider) ListPolicies(ctx context.Context, base *authzv1.PolicySpec, opts ...authz.PolicyFilterOption) ([]*authzv1.PolicySpec, error) {
	p.log.WithContext(ctx).Info("Listing all policies from casbin_rule table")

	// 1. Fetch all rules from the casbin_rule table.
	rules, err := p.db.CasbinRule(ctx).Query().All(ctx)
	if err != nil {
		p.log.WithContext(ctx).Errorf("Failed to list policies from casbin_rule table: %v", err)
		return nil, err
	}

	// 2. Transform the casbin rules into the PolicySpec format.
	policies := make([]*authzv1.PolicySpec, len(rules))
	for i, rule := range rules {
		policies[i] = &authzv1.PolicySpec{
			Type:      rule.Ptype,
			Subject:   rule.V0,
			Domain:    &rule.V1,
			Resources: []string{rule.V2},
			Actions:   []string{rule.V3},
			// V4 and V5 are not mapped in this context.
		}
	}

	p.log.WithContext(ctx).Infof("Successfully retrieved and transformed %d policies", len(policies))
	return policies, nil
}

var _ authz.PolicyReader = (*policyProvider)(nil)
