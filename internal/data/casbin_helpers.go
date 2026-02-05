/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// --- Helper Functions ---

// policyToRule converts a CasbinRule entity to a string slice representing the rule.
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
	// Pad the rule to 6 fields to ensure all V-fields are set.
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
	var preds []predicate.CasbinRule
	if ptype != "" {
		preds = append(preds, casbinrule.PtypeEQ(ptype))
	}
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

// loadPolicyLine correctly formats a policy rule from the database for the Casbin model.
func loadPolicyLine(line *ent.CasbinRule, model model.Model) {
	key := line.Ptype
	sec := key[:1]

	// Get the assertion from the model to determine the number of tokens.
	assertion, ok := model[sec][key]
	if !ok {
		return // Should not happen
	}

	tokens := assertion.Tokens

	// Build the policy rule with the correct number of tokens.
	rule := make([]string, len(tokens))
	rule[0] = line.V0
	if len(tokens) > 1 {
		rule[1] = line.V1
	}
	if len(tokens) > 2 {
		rule[2] = line.V2
	}
	if len(tokens) > 3 {
		rule[3] = line.V3
	}
	if len(tokens) > 4 {
		rule[4] = line.V4
	}
	if len(tokens) > 5 {
		rule[5] = line.V5
	}

	// Join the tokens to form the policy line text.
	lineText := line.Ptype + ", " + strings.Join(rule, ", ")
	_ = persist.LoadPolicyLine(lineText, model)
}
