/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dal

import (
	"context"
	"fmt"
	"strings"

	authzv1 "github.com/origadmin/contrib/api/gen/go/security/authz/v1"
	"github.com/origadmin/contrib/security/authz"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
	"origadmin/application/admin/internal/features/system/dto"
)

// CasbinRuleConfig defines the configuration for mapping PolicySpec fields to Casbin rule columns (V0-V5).
type CasbinRuleConfig struct {
	SubjectIndex  int
	DomainIndex   int
	ResourceIndex int
	ActionIndex   int
	RoleIndex     int
	EffectIndex   int
}

// DefaultAccessRuleConfig defines the default mapping for access policies (p-rules).
var DefaultAccessRuleConfig = CasbinRuleConfig{
	SubjectIndex:  0,
	DomainIndex:   1,
	ResourceIndex: 2,
	ActionIndex:   3,
	EffectIndex:   4,
	RoleIndex:     -1,
}

// DefaultGroupingRuleConfig defines the default mapping for grouping policies (g-rules).
var DefaultGroupingRuleConfig = CasbinRuleConfig{
	SubjectIndex:  0,
	RoleIndex:     1,
	DomainIndex:   2,
	ResourceIndex: -1,
	ActionIndex:   -1,
	EffectIndex:   -1,
}

// CasbinRuleMapper handles encoding and decoding between PolicySpec and Casbin rules.
type CasbinRuleMapper struct {
	accessConfig   CasbinRuleConfig
	groupingConfig CasbinRuleConfig
}

// NewCasbinRuleMapper creates a new CasbinRuleMapper.
func NewCasbinRuleMapper(access, grouping CasbinRuleConfig) *CasbinRuleMapper {
	return &CasbinRuleMapper{
		accessConfig:   access,
		groupingConfig: grouping,
	}
}

// GetConfig returns the configuration for the given ptype.
func (m *CasbinRuleMapper) GetConfig(ptype string) CasbinRuleConfig {
	if ptype == "g" {
		return m.groupingConfig
	}
	return m.accessConfig
}

// Encode converts a PolicySpec to a Casbin rule (ptype, rule slice).
func (m *CasbinRuleMapper) Encode(policy *authzv1.PolicySpec) (string, []string) {
	ptype := policy.Type
	config := m.GetConfig(ptype)
	rule := make([]string, 6) // V0-V5

	// Map fields based on configuration
	if config.SubjectIndex >= 0 {
		rule[config.SubjectIndex] = policy.Subject
	}
	if config.DomainIndex >= 0 && policy.Domain != nil {
		rule[config.DomainIndex] = *policy.Domain
	}
	if config.ResourceIndex >= 0 && len(policy.Resources) > 0 {
		rule[config.ResourceIndex] = policy.Resources[0]
	}
	if config.ActionIndex >= 0 && len(policy.Actions) > 0 {
		rule[config.ActionIndex] = policy.Actions[0]
	}
	if config.RoleIndex >= 0 && len(policy.Roles) > 0 {
		rule[config.RoleIndex] = policy.Roles[0]
	}
	//if config.EffectIndex >= 0 {
	//	if policy.Effect != nil {
	//		rule[config.EffectIndex] = *policy.Effect
	//	} else {
	//		rule[config.EffectIndex] = "allow"
	//	}
	//}

	return ptype, rule
}

// Decode converts a Casbin rule to a PolicySpec.
func (m *CasbinRuleMapper) Decode(rule *ent.CasbinRule) *authzv1.PolicySpec {
	ptype := rule.Ptype
	config := m.GetConfig(ptype)

	policy := &authzv1.PolicySpec{
		Type: ptype,
	}

	// Map fields based on configuration
	if config.SubjectIndex >= 0 {
		policy.Subject = m.getRuleValue(rule, config.SubjectIndex)
	}
	if config.DomainIndex >= 0 {
		val := m.getRuleValue(rule, config.DomainIndex)
		if val != "" {
			policy.Domain = &val
		}
	}
	if config.ResourceIndex >= 0 {
		val := m.getRuleValue(rule, config.ResourceIndex)
		if val != "" {
			policy.Resources = []string{val}
		}
	}
	if config.ActionIndex >= 0 {
		val := m.getRuleValue(rule, config.ActionIndex)
		if val != "" {
			policy.Actions = []string{val}
		}
	}
	if config.RoleIndex >= 0 {
		val := m.getRuleValue(rule, config.RoleIndex)
		if val != "" {
			policy.Roles = []string{val}
		}
	}
	//if config.EffectIndex >= 0 {
	//	val := m.getRuleValue(rule, config.EffectIndex)
	//	if val != "" {
	//		policy.Effect = &val
	//	}
	//} else {
	//	// Default effect if not mapped
	//	policy.Effect = dto.StrPtr("allow")
	//}

	return policy
}

func (m *CasbinRuleMapper) getRuleValue(rule *ent.CasbinRule, index int) string {
	switch index {
	case 0:
		return rule.V0
	case 1:
		return rule.V1
	case 2:
		return rule.V2
	case 3:
		return rule.V3
	case 4:
		return rule.V4
	case 5:
		return rule.V5
	default:
		return ""
	}
}

// GetDomainIndex returns the domain index for the given ptype.
func (m *CasbinRuleMapper) GetDomainIndex(ptype string) int {
	return m.GetConfig(ptype).DomainIndex
}

// casbinPolicyModifier implements the authz.PolicyModifier interface for Casbin.
type casbinPolicyModifier struct {
	db     *ent.Database
	log    *log.Helper
	mapper *CasbinRuleMapper
}

// NewCasbinModifier creates a new Casbin PolicyModifier.
func NewCasbinModifier(db *ent.Database, logger log.Logger) (authz.PolicyModifier, error) {
	return &casbinPolicyModifier{
		db:     db,
		log:    log.NewHelper(log.With(logger, "module", "system.dal.casbin_policy_modifier")),
		mapper: NewCasbinRuleMapper(DefaultAccessRuleConfig, DefaultGroupingRuleConfig),
	}, nil
}

// WithAccessRuleConfig sets the configuration for access rules.
func (m *casbinPolicyModifier) WithAccessRuleConfig(config CasbinRuleConfig) *casbinPolicyModifier {
	m.mapper.accessConfig = config
	return m
}

// WithGroupingRuleConfig sets the configuration for grouping rules.
func (m *casbinPolicyModifier) WithGroupingRuleConfig(config CasbinRuleConfig) *casbinPolicyModifier {
	m.mapper.groupingConfig = config
	return m
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
		domainIndex := m.mapper.GetDomainIndex(ptype)
		switch domainIndex {
		case 1:
			preds = append(preds, casbinrule.V1EQ(*filter.Domain))
		case 2:
			preds = append(preds, casbinrule.V2EQ(*filter.Domain))
		case 3:
			preds = append(preds, casbinrule.V3EQ(*filter.Domain))
		case 4:
			preds = append(preds, casbinrule.V4EQ(*filter.Domain))
		case 5:
			preds = append(preds, casbinrule.V5EQ(*filter.Domain))
		}
	}

	rules, err := m.db.CasbinRule(ctx).Query().Where(preds...).All(ctx)
	if err != nil {
		return nil, err
	}

	policies := make([]*authzv1.PolicySpec, 0, len(rules))
	for _, rule := range rules {
		policy := m.mapper.Decode(rule)
		if m.policyMatchesFilter(policy, filter) {
			policies = append(policies, policy)
		}
	}

	return policies, nil
}

// AddPolicies adds one or more policies.
func (m *casbinPolicyModifier) AddPolicies(ctx context.Context, policies ...*authzv1.PolicySpec) (bool, error) {
	if len(policies) == 0 {
		return false, nil
	}
	m.log.WithContext(ctx).Debugf("Adding %d policies", len(policies))

	creates := make([]*ent.CasbinRuleCreate, 0, len(policies))
	for _, policy := range policies {
		ptype, rule := m.mapper.Encode(policy)
		cr := m.db.CasbinRule(ctx).Create().
			SetPtype(ptype).
			SetV0(rule[0])
		if len(rule) > 1 {
			cr.SetV1(rule[1])
		}
		if len(rule) > 2 {
			cr.SetV2(rule[2])
		}
		if len(rule) > 3 {
			cr.SetV3(rule[3])
		}
		if len(rule) > 4 {
			cr.SetV4(rule[4])
		}
		if len(rule) > 5 {
			cr.SetV5(rule[5])
		}
		creates = append(creates, cr)
	}

	var changed bool
	err := m.db.Tx(ctx, func(ctx context.Context) error {
		for _, create := range creates {
			_, err := create.Save(ctx)
			if err != nil {
				if strings.Contains(err.Error(), "duplicate") ||
					strings.Contains(err.Error(), "UNIQUE") {
					continue
				}
				return err
			}
			changed = true
		}
		return nil
	})

	if err != nil {
		m.log.WithContext(ctx).Errorf("Failed to add policies: %v", err)
		return changed, err
	}
	return changed, nil
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
			oldPtype, oldRule := m.mapper.Encode(oldPolicies[i])
			newPtype, newRule := m.mapper.Encode(newPolicies[i])

			if oldPtype != newPtype {
				err := fmt.Errorf("cannot change policy type from %s to %s", oldPtype, newPtype)
				if firstErr == nil {
					firstErr = err
				}
				continue
			}

			rule, err := m.db.CasbinRule(ctx).Query().
				Where(casbinrule.PtypeEQ(oldPtype)).
				Where(casbinrule.V0EQ(oldRule[0])).
				Only(ctx)
			if err != nil {
				if firstErr == nil {
					firstErr = err
				}
				continue
			}

			_, err = m.db.CasbinRule(ctx).UpdateOneID(rule.ID).
				SetV0(newRule[0]).
				SetV1(newRule[1]).
				SetV2(newRule[2]).
				SetV3(newRule[3]).
				SetV4(newRule[4]).
				SetV5(newRule[5]).
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
		domainIndex := m.mapper.GetDomainIndex(ptype)
		switch domainIndex {
		case 1:
			preds = append(preds, casbinrule.V1EQ(*filter.Domain))
		case 2:
			preds = append(preds, casbinrule.V2EQ(*filter.Domain))
		case 3:
			preds = append(preds, casbinrule.V3EQ(*filter.Domain))
		case 4:
			preds = append(preds, casbinrule.V4EQ(*filter.Domain))
		case 5:
			preds = append(preds, casbinrule.V5EQ(*filter.Domain))
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
		policy := m.mapper.Decode(rule)
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

// AddRoles adds roles for a subject.
// Deprecated: Use AddPolicies instead.
func (m *casbinPolicyModifier) AddRoles(ctx context.Context, subject string, roles ...authz.RoleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("AddRoles is deprecated, use AddPolicies with PolicySpec instead")
	if len(roles) == 0 {
		return false, nil
	}
	policies := make([]*authzv1.PolicySpec, len(roles))
	for i, r := range roles {
		policies[i] = &authzv1.PolicySpec{
			Type:    "g",
			Subject: subject,
			Domain:  dto.StrPtr(r.Domain),
			Roles:   []string{r.Role},
			Effect:  dto.StrPtr("allow"),
		}
	}
	return m.AddPolicies(ctx, policies...)
}

// RemoveRoles removes roles for a subject.
// Deprecated: Use RemovePolicies instead.
func (m *casbinPolicyModifier) RemoveRoles(ctx context.Context, subject string, roles ...authz.RoleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("RemoveRoles is deprecated, use RemovePolicies with PolicySpec instead")
	if len(roles) == 0 {
		return m.RemovePolicies(ctx, &authzv1.PolicySpec{
			Type:    "g",
			Subject: subject,
		})
	}
	changed := false
	var firstErr error
	for _, r := range roles {
		c, err := m.RemovePolicies(ctx, &authzv1.PolicySpec{
			Type:    "g",
			Subject: subject,
			Domain:  dto.StrPtr(r.Domain),
			Roles:   []string{r.Role},
		})
		if err != nil && firstErr == nil {
			firstErr = err
		}
		if c {
			changed = true
		}
	}
	if firstErr != nil {
		return changed, firstErr
	}
	return changed, nil
}

// UpdateRole updates a role for a subject.
// Deprecated: Use UpdatePolicies instead.
func (m *casbinPolicyModifier) UpdateRole(ctx context.Context, subject string, oldRole authz.RoleSpec, newRole authz.RoleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("UpdateRole is deprecated, use UpdatePolicies with PolicySpec instead")
	oldPolicy := &authzv1.PolicySpec{
		Type:    "g",
		Subject: subject,
		Domain:  dto.StrPtr(oldRole.Domain),
		Roles:   []string{oldRole.Role},
		Effect:  dto.StrPtr("allow"),
	}
	newPolicy := &authzv1.PolicySpec{
		Type:    "g",
		Subject: subject,
		Domain:  dto.StrPtr(newRole.Domain),
		Roles:   []string{newRole.Role},
		Effect:  dto.StrPtr("allow"),
	}
	return m.UpdatePolicies(ctx, []*authzv1.PolicySpec{oldPolicy}, []*authzv1.PolicySpec{newPolicy})
}

// AddPermissions adds permissions for a subject.
// Deprecated: Use AddPolicies instead.
func (m *casbinPolicyModifier) AddPermissions(ctx context.Context, subject string, permissions ...authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("AddPermissions is deprecated, use AddPolicies with PolicySpec instead")
	if len(permissions) == 0 {
		return false, nil
	}
	policies := make([]*authzv1.PolicySpec, len(permissions))
	for i, p := range permissions {
		policies[i] = &authzv1.PolicySpec{
			Type:      "p",
			Subject:   subject,
			Domain:    dto.StrPtr(p.Domain),
			Resources: []string{p.Resource},
			Actions:   []string{p.Action},
			Effect:    dto.StrPtr("allow"),
		}
	}
	return m.AddPolicies(ctx, policies...)
}

// RemovePermissions removes permissions for a subject.
// Deprecated: Use RemovePolicies instead.
func (m *casbinPolicyModifier) RemovePermissions(ctx context.Context, subject string, permissions ...authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("RemovePermissions is deprecated, use RemovePolicies with PolicySpec instead")
	if len(permissions) == 0 {
		return m.RemovePolicies(ctx, &authzv1.PolicySpec{
			Type:    "p",
			Subject: subject,
		})
	}
	changed := false
	var firstErr error
	for _, p := range permissions {
		c, err := m.RemovePolicies(ctx, &authzv1.PolicySpec{
			Type:      "p",
			Subject:   subject,
			Domain:    dto.StrPtr(p.Domain),
			Resources: []string{p.Resource},
			Actions:   []string{p.Action},
		})
		if err != nil && firstErr == nil {
			firstErr = err
		}
		if c {
			changed = true
		}
	}
	if firstErr != nil {
		return changed, firstErr
	}
	return changed, nil
}

// UpdatePermission updates a permission for a subject.
// Deprecated: Use UpdatePolicies instead.
func (m *casbinPolicyModifier) UpdatePermission(ctx context.Context, subject string, oldPerm authz.RuleSpec, newPerm authz.RuleSpec) (bool, error) {
	m.log.WithContext(ctx).Warn("UpdatePermission is deprecated, use UpdatePolicies with PolicySpec instead")
	oldPolicy := &authzv1.PolicySpec{
		Type:      "p",
		Subject:   subject,
		Domain:    dto.StrPtr(oldPerm.Domain),
		Resources: []string{oldPerm.Resource},
		Actions:   []string{oldPerm.Action},
		Effect:    dto.StrPtr("allow"),
	}
	newPolicy := &authzv1.PolicySpec{
		Type:      "p",
		Subject:   subject,
		Domain:    dto.StrPtr(newPerm.Domain),
		Resources: []string{newPerm.Resource},
		Actions:   []string{newPerm.Action},
		Effect:    dto.StrPtr("allow"),
	}
	return m.UpdatePolicies(ctx, []*authzv1.PolicySpec{oldPolicy}, []*authzv1.PolicySpec{newPolicy})
}
