/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// CasbinAdapter implements the casbin persist.UpdatableAdapter for ent.
type CasbinAdapter struct {
	Ctx context.Context
	DB  *ent.Database
}

// NewAdapter creates a new casbin adapter.
func NewAdapter(rt *runtime.App, db *ent.Database) (*CasbinAdapter, error) {
	return &CasbinAdapter{Ctx: rt.Context(), DB: db}, nil
}

// Context returns the context of the adapter.
func (a *CasbinAdapter) Context() context.Context {
	return a.Ctx
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(m model.Model) error {
	client := a.DB.CasbinRule(a.Context())
	policies, err := client.Query().Order(ent.Asc("id")).All(a.Context())
	if err != nil {
		return err
	}

	for _, policy := range policies {
		key := policy.Ptype
		sec := key[:1]
		ast, ok := m[sec][key]
		if !ok {
			return fmt.Errorf("policy type '%s' not found in Casbin model", key)
		}
		expectedLen := len(ast.Tokens)

		allDBFields := []string{policy.V0, policy.V1, policy.V2, policy.V3, policy.V4, policy.V5}

		rule := make([]string, expectedLen)
		for i := 0; i < expectedLen; i++ {
			if i < len(allDBFields) {
				rule[i] = allDBFields[i]
			} else {
				rule[i] = ""
			}
		}

		ruleStr := strings.Join(rule, ", ")
		line := policy.Ptype + ", " + ruleStr

		if err := persist.LoadPolicyLine(line, m); err != nil {
			return err
		}
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model model.Model) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		if _, err := cr.Delete().Exec(ctx); err != nil {
			return err
		}
		lines := make([]*ent.CasbinRuleCreate, 0)
		for ptype, ast := range model["p"] {
			for _, policy := range ast.Policy {
				lines = append(lines, savePolicyLine(cr, ptype, policy))
			}
		}
		for ptype, ast := range model["g"] {
			for _, policy := range ast.Policy {
				lines = append(lines, savePolicyLine(cr, ptype, policy))
			}
		}
		_, err := cr.CreateBulk(lines...).Save(ctx)
		return err
	})
}

// AddPolicy adds a policy rule to the storage.
func (a *CasbinAdapter) AddPolicy(_ string, ptype string, rule []string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		_, err := savePolicyLine(cr, ptype, rule).Save(ctx)
		return err
	})
}

// AddPolicies adds multiple policy rules to the storage.
func (a *CasbinAdapter) AddPolicies(_ string, ptype string, rules [][]string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		lines := make([]*ent.CasbinRuleCreate, len(rules))
		for i, rule := range rules {
			lines[i] = savePolicyLine(cr, ptype, rule)
		}
		_, err := cr.CreateBulk(lines...).Save(ctx)
		return err
	})
}

// UpdatePolicy updates a policy rule from storage.
func (a *CasbinAdapter) UpdatePolicy(_ string, ptype string, oldRule []string, newRule []string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		return a.updatePolicyInTx(ctx, ptype, oldRule, newRule)
	})
}

// UpdatePolicies updates multiple policy rules from storage.
func (a *CasbinAdapter) UpdatePolicies(_ string, ptype string, oldRules [][]string, newRules [][]string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		if len(oldRules) != len(newRules) {
			return fmt.Errorf("old and new rules must have the same length")
		}
		for i, oldRule := range oldRules {
			newRule := newRules[i]
			err := a.updatePolicyInTx(ctx, ptype, oldRule, newRule)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// updatePolicyInTx is a helper for UpdatePolicies to run within a transaction.
func (a *CasbinAdapter) updatePolicyInTx(ctx context.Context, ptype string, oldRule []string, newRule []string) error {
	cr := a.DB.CasbinRule(ctx)
	oldFilter, err := buildInstanceFilter(ptype, oldRule)
	if err != nil {
		return err
	}
	newValues := instanceLine(ptype, newRule)

	count, err := cr.Update().
		Where(oldFilter...).
		SetV0(newValues.V0).
		SetV1(newValues.V1).
		SetV2(newValues.V2).
		SetV3(newValues.V3).
		SetV4(newValues.V4).
		SetV5(newValues.V5).
		Save(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("policy to update not found")
	}
	return nil
}

// UpdateFilteredPolicies updates policy rules that match the filter from storage.
func (a *CasbinAdapter) UpdateFilteredPolicies(_ string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	var oldPolicies [][]string
	err := a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		filter, err := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
		if err != nil {
			return err
		}

		policies, err := cr.Query().Where(filter...).All(ctx)
		if err != nil {
			return err
		}
		for _, p := range policies {
			oldPolicies = append(oldPolicies, policyToRule(p))
		}

		count, err := cr.Delete().Where(filter...).Exec(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("policies to update not found for filter")
		}

		for _, newRule := range newRules {
			if _, err := savePolicyLine(cr, ptype, newRule).Save(ctx); err != nil {
				return err
			}
		}
		return nil
	})
	return oldPolicies, err
}

// RemovePolicy removes a policy rule from the storage.
func (a *CasbinAdapter) RemovePolicy(_ string, ptype string, rule []string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		filter, err := buildInstanceFilter(ptype, rule)
		if err != nil {
			return err
		}
		count, err := cr.Delete().Where(filter...).Exec(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("policy to remove not found: %s, %v", ptype, rule)
		}
		return err
	})
}

// RemovePolicies removes multiple policy rules from the storage.
func (a *CasbinAdapter) RemovePolicies(_ string, ptype string, rules [][]string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		for _, rule := range rules {
			filter, err := buildInstanceFilter(ptype, rule)
			if err != nil {
				return err
			}
			count, err := cr.Delete().Where(filter...).Exec(ctx)
			if err != nil {
				return err
			}
			if count == 0 {
				return fmt.Errorf("policy to remove not found: %s, %v", ptype, rule)
			}
		}
		return nil
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *CasbinAdapter) RemoveFilteredPolicy(_ string, ptype string, fieldIndex int, fieldValues ...string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		cond, err := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
		if err != nil {
			return err
		}
		count, err := cr.Delete().Where(cond...).Exec(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("policy to remove not found for filter")
		}
		return err
	})
}

// RemovePoliciesByFields removes policies that match a given set of field filters.
// This is a custom, generic method for the adapter.
func (a *CasbinAdapter) RemovePoliciesByFields(ptype string, filters map[string]string) error {
	return a.DB.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.DB.CasbinRule(ctx)
		preds, err := buildPredicatesFromFilters(ptype, filters)
		if err != nil {
			return err
		}
		count, err := cr.Delete().Where(preds...).Exec(ctx)
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("policies to remove not found for filter: %v", filters)
		}
		return err
	})
}

// --- Helper Functions ---

// policyToRule converts a CasbinRule entity to a string slice representing the rule.
// It returns all V0-V5 fields, even if they are empty, to maintain consistency with the database schema.
func policyToRule(p *ent.CasbinRule) []string {
	return []string{p.V0, p.V1, p.V2, p.V3, p.V4, p.V5}
}

// instanceLine creates a CasbinRule entity from a policy type and rule string slice.
func instanceLine(ptype string, rule []string) *ent.CasbinRule {
	instance := &ent.CasbinRule{Ptype: ptype}
	paddedRule := make([]string, 6)
	copy(paddedRule, rule)
	instance.V0 = paddedRule[0]
	instance.V1 = paddedRule[1]
	instance.V2 = paddedRule[2]
	instance.V3 = paddedRule[3]
	instance.V4 = paddedRule[4]
	instance.V5 = paddedRule[5]
	return instance
}

// savePolicyLine creates a CasbinRuleCreate builder from a policy type and rule string slice.
func savePolicyLine(cr *ent.CasbinRuleClient, ptype string, rule []string) *ent.CasbinRuleCreate {
	line := cr.Create().SetPtype(ptype)
	paddedRule := make([]string, 6)
	copy(paddedRule, rule)
	line.SetV0(paddedRule[0])
	line.SetV1(paddedRule[1])
	line.SetV2(paddedRule[2])
	line.SetV3(paddedRule[3])
	line.SetV4(paddedRule[4])
	line.SetV5(paddedRule[5])
	return line
}

// buildPredicatesFromFilters constructs a slice of predicates from a map of field filters.
func buildPredicatesFromFilters(ptype string, filters map[string]string) ([]predicate.CasbinRule, error) {
	preds := []predicate.CasbinRule{casbinrule.PtypeEQ(ptype)}
	for field, value := range filters {
		pred, err := appendVnEQPredicate(field, value)
		if err != nil {
			return nil, err
		}
		preds = append(preds, pred)
	}
	return preds, nil
}

// buildInstanceFilter creates a slice of predicates to filter by a specific policy instance.
func buildInstanceFilter(ptype string, rule []string) ([]predicate.CasbinRule, error) {
	preds := []predicate.CasbinRule{casbinrule.PtypeEQ(ptype)}
	for i, v := range rule {
		field := fmt.Sprintf("v%d", i)
		pred, err := appendVnEQPredicate(field, v)
		if err != nil {
			return nil, err
		}
		preds = append(preds, pred)
	}
	return preds, nil
}

// buildFilteredFilter creates a slice of predicates for filtered policy operations.
func buildFilteredFilter(ptype string, fieldIndex int, fieldValues ...string) ([]predicate.CasbinRule, error) {
	if fieldIndex < 0 || fieldIndex+len(fieldValues) > 6 {
		return nil, fmt.Errorf("invalid field index or field values for filtered policy")
	}
	filters := make(map[string]string)
	for i, v := range fieldValues {
		filters[fmt.Sprintf("v%d", fieldIndex+i)] = v
	}
	return buildPredicatesFromFilters(ptype, filters)
}

// appendVnEQPredicate returns a predicate for a given field (v0-v5) and value.
func appendVnEQPredicate(field, value string) (predicate.CasbinRule, error) {
	switch field {
	case "v0":
		return casbinrule.V0EQ(value), nil
	case "v1":
		return casbinrule.V1EQ(value), nil
	case "v2":
		return casbinrule.V2EQ(value), nil
	case "v3":
		return casbinrule.V3EQ(value), nil
	case "v4":
		return casbinrule.V4EQ(value), nil
	case "v5":
		return casbinrule.V5EQ(value), nil
	default:
		return nil, fmt.Errorf("invalid filter field: %s", field)
	}
}

var _ persist.UpdatableAdapter = (*CasbinAdapter)(nil)
