/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"fmt"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
	"origadmin/application/admin/internal/features/system/dto"
)

// casbinPolicyModifier implements the authz.PolicyModifier interface for Casbin.
type casbinPolicyModifier struct {
	db  *ent.Database
	log *log.Helper
}

// NewCasbinModifier creates a new Casbin PolicyModifier.
func NewCasbinModifier(db *ent.Database, logger log.Logger) (authz.PolicyModifier, error) {
	return &casbinPolicyModifier{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "system.dal.casbin_policy_modifier")),
	}, nil
}

// encode converts a PolicySpec to a full Casbin rule row (ptype + v0-v5).
// Returns a slice of size 7: [ptype, v0, v1, v2, v3, v4, v5]
// Standard mapping:
// p: ptype, sub, dom, obj, act
// g: ptype, user, role, dom
func (m *casbinPolicyModifier) encode(policy *authzv1.PolicySpec) []string {
	rule := make([]string, 7) // Ptype + V0-V5
	rule[0] = policy.Type     // Ptype

	switch policy.Type {
	case "p":
		rule[1] = policy.Subject // V0
		if policy.Domain != nil {
			rule[2] = *policy.Domain // V1
		}
		if len(policy.Resources) > 0 {
			rule[3] = policy.Resources[0] // V2
		}
		if len(policy.Actions) > 0 {
			rule[4] = policy.Actions[0] // V3
		}
	case "g":
		rule[1] = policy.Subject // V0
		if len(policy.Roles) > 0 {
			rule[2] = policy.Roles[0] // V1
		}
		if policy.Domain != nil {
			rule[3] = *policy.Domain // V2
		}
	default:
		// Default fallback mapping
		rule[1] = policy.Subject // V0
	}

	return rule
}

// decode converts a Casbin rule to a PolicySpec.
func (m *casbinPolicyModifier) decode(rule *ent.CasbinRule) *authzv1.PolicySpec {
	policy := &authzv1.PolicySpec{
		Type: rule.Ptype,
	}

	switch rule.Ptype {
	case "p":
		policy.Subject = rule.V0
		if rule.V1 != "" {
			policy.Domain = &rule.V1
		}
		if rule.V2 != "" {
			policy.Resources = []string{rule.V2}
		}
		if rule.V3 != "" {
			policy.Actions = []string{rule.V3}
		}
	case "g":
		policy.Subject = rule.V0
		if rule.V1 != "" {
			policy.Roles = []string{rule.V1}
		}
		if rule.V2 != "" {
			policy.Domain = &rule.V2
		}
	default:
		policy.Subject = rule.V0
	}

	return policy
}

// getDomainIndex returns the domain index for the given ptype.
func (m *casbinPolicyModifier) getDomainIndex(ptype string) int {
	if ptype == "g" {
		return 2
	}
	return 1
}

// ListPolicies queries policies matching the filter criteria.
func (m *casbinPolicyModifier) ListPolicies(ctx context.Context, base *authzv1.PolicySpec, opts ...authz.PolicyFilterOption) ([]*authzv1.PolicySpec, error) {
	filter := authz.BuildPolicyFilter(base, opts...)
	m.log.WithContext(ctx).Debugf("Listing policies with filter: Type=%v, Subject=%v, Domain=%v",
		dto.NonNilStr(filter.Type), dto.NonNilStr(filter.Subject), dto.NonNilStr(filter.Domain))

	policies := make([]*authzv1.PolicySpec, 0)

	ptypes := []string{"p", "g"}
	for _, ptype := range ptypes {
		ptypePolicies, err := m.listPoliciesByPType(ctx, ptype, filter)
		if err != nil {
			return nil, err
		}
		policies = append(policies, ptypePolicies...)
	}

	return policies, nil
}

// listPoliciesByPType queries policies for a specific policy type (p or g).
func (m *casbinPolicyModifier) listPoliciesByPType(ctx context.Context, ptype string, filter *authz.PolicyFilter) ([]*authzv1.PolicySpec, error) {
	var preds []predicate.CasbinRule
	preds = append(preds, casbinrule.PtypeEQ(ptype))

	if filter.Type != nil {
		if *filter.Type != ptype {
			return []*authzv1.PolicySpec{}, nil
		}
	}

	if filter.Subject != nil {
		preds = append(preds, casbinrule.V0EQ(*filter.Subject))
	}

	if filter.Domain != nil {
		domainIndex := m.getDomainIndex(ptype)
		switch domainIndex {
		case 1:
			preds = append(preds, casbinrule.V1EQ(*filter.Domain))
		case 2:
			preds = append(preds, casbinrule.V2EQ(*filter.Domain))
		}
	}

	rules, err := m.db.CasbinRule(ctx).Query().Where(preds...).All(ctx)
	if err != nil {
		return nil, err
	}

	policies := make([]*authzv1.PolicySpec, 0, len(rules))
	for _, rule := range rules {
		policy := m.decode(rule)
		if m.policyMatchesFilter(policy, filter) {
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

// AddPolicies adds one or more policies using a "query then insert" strategy to handle duplicates.
func (m *casbinPolicyModifier) AddPolicies(ctx context.Context, policies ...*authzv1.PolicySpec) (bool, error) {
	if len(policies) == 0 {
		return false, nil
	}
	m.log.WithContext(ctx).Debugf("Adding %d policies", len(policies))

	var anyPolicyAdded bool
	err := m.db.Tx(ctx, func(ctx context.Context) error {
		for _, policy := range policies {
			rule := m.encode(policy)
			// Check if the policy already exists.
			exists, err := m.db.CasbinRule(ctx).Query().Where(
				casbinrule.PtypeEQ(rule[0]),
				casbinrule.V0EQ(rule[1]),
				casbinrule.V1EQ(rule[2]),
				casbinrule.V2EQ(rule[3]),
				casbinrule.V3EQ(rule[4]),
				casbinrule.V4EQ(rule[5]),
				casbinrule.V5EQ(rule[6]),
			).Exist(ctx)
			if err != nil {
				return err
			}
			if exists {
				m.log.WithContext(ctx).Debugf("Skipping duplicate policy (already exists): %v", policy)
				continue
			}

			// Policy does not exist, create it.
			create := m.db.CasbinRule(ctx).Create().
				SetPtype(rule[0]).
				SetV0(rule[1]).
				SetV1(rule[2]).
				SetV2(rule[3]).
				SetV3(rule[4]).
				SetV4(rule[5]).
				SetV5(rule[6])

			err = create.Exec(ctx)
			if err != nil {
				return err
			}
			anyPolicyAdded = true
		}
		return nil
	})

	if err != nil {
		m.log.WithContext(ctx).Errorf("Failed to add policies: %v", err)
		return anyPolicyAdded, err
	}
	return anyPolicyAdded, nil
}

// UpdatePolicies updates policies in batch.
func (m *casbinPolicyModifier) UpdatePolicies(ctx context.Context, oldPolicies []*authzv1.PolicySpec, newPolicies []*authzv1.PolicySpec) (bool, error) {
	if len(oldPolicies) != len(newPolicies) {
		return false, fmt.Errorf("oldPolicies and newPolicies must have the same length")
	}

	m.log.WithContext(ctx).Debugf("Updating %d policies", len(oldPolicies))

	changed := false
	var firstErr error

	err := m.db.Tx(ctx, func(ctx context.Context) error {
		for i := 0; i < len(oldPolicies); i++ {
			oldRule := m.encode(oldPolicies[i])
			newRule := m.encode(newPolicies[i])

			if oldRule[0] != newRule[0] {
				err := fmt.Errorf("cannot change policy type from %s to %s", oldRule[0], newRule[0])
				if firstErr == nil {
					firstErr = err
				}
				continue
			}

			// Find the specific rule to update based on its old values
			rule, err := m.db.CasbinRule(ctx).Query().
				Where(
					casbinrule.PtypeEQ(oldRule[0]),
					casbinrule.V0EQ(oldRule[1]),
					casbinrule.V1EQ(oldRule[2]),
					casbinrule.V2EQ(oldRule[3]),
					casbinrule.V3EQ(oldRule[4]),
					casbinrule.V4EQ(oldRule[5]),
					casbinrule.V5EQ(oldRule[6]),
				).Only(ctx)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue
			}

			// Update the found rule with the new values
			_, err = m.db.CasbinRule(ctx).UpdateOneID(rule.ID).
				SetV0(newRule[1]).
				SetV1(newRule[2]).
				SetV2(newRule[3]).
				SetV3(newRule[4]).
				SetV4(newRule[5]).
				SetV5(newRule[6]).
				Save(ctx)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue
			}
			changed = true
		}
		return nil
	})

	if err != nil && firstErr == nil {
		return changed, err
	}
	if firstErr != nil {
		return changed, firstErr
	}
	return changed, nil
}

// RemovePolicies removes policies matching the filter criteria.
func (m *casbinPolicyModifier) RemovePolicies(ctx context.Context, base *authzv1.PolicySpec, opts ...authz.PolicyFilterOption) (bool, error) {
	filter := authz.BuildPolicyFilter(base, opts...)
	m.log.WithContext(ctx).Debugf("Removing policies with filter: Type=%v, Subject=%v, Domain=%v",
		dto.NonNilStr(filter.Type), dto.NonNilStr(filter.Subject), dto.NonNilStr(filter.Domain))

	changed := false
	var firstErr error

	ptypes := []string{"p", "g"}
	for _, ptype := range ptypes {
		ptypeChanged, err := m.removePoliciesByPType(ctx, ptype, filter)
		if err != nil && firstErr == nil {
			firstErr = err
		}
		if ptypeChanged {
			changed = true
		}
	}

	if firstErr != nil {
		return changed, firstErr
	}
	return changed, nil
}

// removePoliciesByPType removes policies for a specific policy type (p or g).
func (m *casbinPolicyModifier) removePoliciesByPType(ctx context.Context, ptype string, filter *authz.PolicyFilter) (bool, error) {
	var preds []predicate.CasbinRule
	preds = append(preds, casbinrule.PtypeEQ(ptype))

	if filter.Type != nil {
		if *filter.Type != ptype {
			return false, nil
		}
	}

	if filter.Subject != nil {
		preds = append(preds, casbinrule.V0EQ(*filter.Subject))
	}

	if filter.Domain != nil {
		domainIndex := m.getDomainIndex(ptype)
		switch domainIndex {
		case 1:
			preds = append(preds, casbinrule.V1EQ(*filter.Domain))
		case 2:
			preds = append(preds, casbinrule.V2EQ(*filter.Domain))
		}
	}

	rules, err := m.db.CasbinRule(ctx).Query().Where(preds...).All(ctx)
	if err != nil {
		return false, err
	}

	if len(rules) == 0 {
		return false, nil
	}

	changed := false
	for _, rule := range rules {
		policy := m.decode(rule)
		if !m.policyMatchesFilter(policy, filter) {
			continue
		}

		_, err := m.db.CasbinRule(ctx).Delete().
			Where(casbinrule.IDEQ(rule.ID)).
			Exec(ctx)
		if err != nil {
			m.log.WithContext(ctx).Errorf("Failed to delete policy %d: %v", rule.ID, err)
			return changed, err
		}
		changed = true
	}

	return changed, nil
}

// ClearPolicies removes all policies.
func (m *casbinPolicyModifier) ClearPolicies(ctx context.Context) (bool, error) {
	m.log.WithContext(ctx).Info("Clearing all policies from storage")
	_, err := m.db.CasbinRule(ctx).Delete().Exec(ctx)
	if err != nil {
		m.log.WithContext(ctx).Errorf("Failed to clear all policies: %v", err)
		return false, err
	}
	m.log.WithContext(ctx).Info("Successfully cleared all policies")
	return true, nil
}

// policyMatchesFilter checks if a policy matches the filter criteria.
func (m *casbinPolicyModifier) policyMatchesFilter(policy *authzv1.PolicySpec, filter *authz.PolicyFilter) bool {
	if filter.Type != nil && policy.Type != *filter.Type {
		return false
	}
	if filter.Subject != nil && policy.Subject != *filter.Subject {
		return false
	}
	if filter.Domain != nil {
		policyDomain := ""
		if policy.Domain != nil {
			policyDomain = *policy.Domain
		}
		if policyDomain != *filter.Domain {
			return false
		}
	}
	if len(filter.Actions) > 0 {
		actionMatch := false
		for _, filterAction := range filter.Actions {
			for _, policyAction := range policy.Actions {
				if policyAction == filterAction {
					actionMatch = true
					break
				}
			}
		}
		if !actionMatch {
			return false
		}
	}
	if len(filter.Resources) > 0 {
		resourceMatch := false
		for _, filterResource := range filter.Resources {
			for _, policyResource := range policy.Resources {
				if policyResource == filterResource {
					resourceMatch = true
					break
				}
			}
		}
		if !resourceMatch {
			return false
		}
	}
	if filter.Effect != nil {
		policyEffect := ""
		if policy.Effect != nil {
			policyEffect = *policy.Effect
		}
		if policyEffect != *filter.Effect {
			return false
		}
	}
	if filter.Disabled != nil {
		policyDisabled := false
		if policy.Disabled != nil {
			policyDisabled = *policy.Disabled
		}
		if policyDisabled != *filter.Disabled {
			return false
		}
	}
	return true
}
