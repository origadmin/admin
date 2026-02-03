/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	kratosLog "github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// CasbinAdapter implements the casbin persist.UpdatableAdapter for ent.
// Its sole responsibility is to act as a persistence layer for a Casbin Enforcer.
// It should not contain business logic or custom data manipulation methods.
type CasbinAdapter struct {
	ctx      context.Context
	db       *ent.Database
	log      *kratosLog.Helper
	filtered bool
}

type Filter struct {
	Ptype []string
	V0    []string
	V1    []string
	V2    []string
	V3    []string
	V4    []string
	V5    []string
}

// NewAdapter creates a new casbin adapter.
func NewAdapter(ctx context.Context, db *ent.Database, logger log.Logger) (*CasbinAdapter, error) {
	return &CasbinAdapter{
		ctx: ctx,
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "data.casbin_adapter")),
	}, nil
}

// NewAdapterFromApp creates a new casbin adapter.
func NewAdapterFromApp(app *runtime.App, db *ent.Database) (*CasbinAdapter, error) {
	return &CasbinAdapter{
		ctx: app.Context(),
		db:  db,
		log: log.NewHelper(log.With(app.Logger(), "module", "data.casbin_adapter")),
	}, nil
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(m model.Model) error {
	client := a.db.CasbinRule(a.ctx)
	policies, err := client.Query().Order(ent.Asc("id")).All(a.ctx)
	if err != nil {
		return err
	}
	for _, policy := range policies {
		loadPolicyLine(policy, m)
	}
	return nil
}

// LoadFilteredPolicy loads only policy rules that match the filter.
func (a *CasbinAdapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	filterValue, ok := filter.(Filter)
	if !ok {
		return fmt.Errorf("invalid filter type: %v", reflect.TypeOf(filter))
	}

	session := a.db.CasbinRule(a.ctx).Query()
	if len(filterValue.Ptype) != 0 {
		session.Where(casbinrule.PtypeIn(filterValue.Ptype...))
	}
	if len(filterValue.V0) != 0 {
		session.Where(casbinrule.V0In(filterValue.V0...))
	}
	if len(filterValue.V1) != 0 {
		session.Where(casbinrule.V1In(filterValue.V1...))
	}
	if len(filterValue.V2) != 0 {
		session.Where(casbinrule.V2In(filterValue.V2...))
	}
	if len(filterValue.V3) != 0 {
		session.Where(casbinrule.V3In(filterValue.V3...))
	}
	if len(filterValue.V4) != 0 {
		session.Where(casbinrule.V4In(filterValue.V4...))
	}
	if len(filterValue.V5) != 0 {
		session.Where(casbinrule.V5In(filterValue.V5...))
	}

	lines, err := session.All(a.ctx)
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	a.filtered = true

	return nil
}

// IsFiltered returns true if the loaded policy has been filtered.
func (a *CasbinAdapter) IsFiltered() bool {
	return a.filtered
}

// SavePolicy saves all policy rules to the storage.
func (a *CasbinAdapter) SavePolicy(model model.Model) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
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
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		_, err := savePolicyLine(cr, ptype, rule).Save(ctx)
		return err
	})
}

// AddPolicies adds multiple policy rules to the storage.
func (a *CasbinAdapter) AddPolicies(_ string, ptype string, rules [][]string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
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
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		return a.updatePolicyInTx(ctx, ptype, oldRule, newRule)
	})
}

// UpdatePolicies updates multiple policy rules from storage.
func (a *CasbinAdapter) UpdatePolicies(_ string, ptype string, oldRules [][]string, newRules [][]string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
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
	cr := a.db.CasbinRule(ctx)
	oldFilter, err := buildInstanceFilter(ptype, oldRule)
	if err != nil {
		return err
	}
	newValues := instanceLine(ptype, newRule)

	_, err = cr.Update().
		Where(oldFilter...).
		SetV0(newValues.V0).
		SetV1(newValues.V1).
		SetV2(newValues.V2).
		SetV3(newValues.V3).
		SetV4(newValues.V4).
		SetV5(newValues.V5).
		Save(ctx)
	return err
}

// UpdateFilteredPolicies updates policy rules that match the filter from storage.
func (a *CasbinAdapter) UpdateFilteredPolicies(_ string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	var oldPolicies [][]string
	err := a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
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

		if _, err := cr.Delete().Where(filter...).Exec(ctx); err != nil {
			return err
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
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		filter, err := buildInstanceFilter(ptype, rule)
		if err != nil {
			return err
		}
		if _, err := cr.Delete().Where(filter...).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

// RemovePolicies removes multiple policy rules from the storage.
func (a *CasbinAdapter) RemovePolicies(_ string, ptype string, rules [][]string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		for _, rule := range rules {
			filter, err := buildInstanceFilter(ptype, rule)
			if err != nil {
				return err
			}
			if _, err := cr.Delete().Where(filter...).Exec(ctx); err != nil {
				return err
			}
		}
		return nil
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *CasbinAdapter) RemoveFilteredPolicy(_ string, ptype string, fieldIndex int, fieldValues ...string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		cond, err := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
		if err != nil {
			return err
		}
		if _, err := cr.Delete().Where(cond...).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

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

var _ persist.UpdatableAdapter = (*CasbinAdapter)(nil)
