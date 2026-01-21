/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// casbinAdapter implements the casbin persist.Adapter for casbin.
type casbinAdapter struct {
	ctx context.Context
	db  *ent.Database
}

// NewAdapter creates a new casbin adapter.
func NewAdapter(rt *runtime.App, db *ent.Database) (persist.Adapter, error) {
	return &casbinAdapter{ctx: rt.Context(), db: db}, nil
}

// Context returns the context of the adapter.
func (a *casbinAdapter) Context() context.Context {
	return a.ctx
}

// LoadPolicy loads all policy rules from the storage.
func (a *casbinAdapter) LoadPolicy(model model.Model) error {
	client := a.db.CasbinRule(a.Context())
	policies, err := client.Query().Order(ent.Asc("id")).All(a.Context())
	if err != nil {
		return err
	}
	for _, policy := range policies {
		loadPolicyLine(policy, model)
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *casbinAdapter) SavePolicy(model model.Model) error {
	return a.db.Tx(a.Context(), func(ctx context.Context) error {
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
func (a *casbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	return a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		if _, err := savePolicyLine(cr, ptype, rule).Save(ctx); err != nil {
			return err
		}
		return nil
	})
}

// RemovePolicy removes a policy rule from the storage.
func (a *casbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		instance := instanceLine(ptype, rule)
		if _, err := cr.Delete().Where(buildInstanceFilter(instance)...).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *casbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		cond := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
		if _, err := cr.Delete().Where(cond...).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}

// --- Helper Functions ---

func loadPolicyLine(line *ent.CasbinRule, model model.Model) {
	key := line.Ptype
	sec := key[:1]
	model[sec][key].Policy = append(model[sec][key].Policy, []string{line.V0, line.V1, line.V2, line.V3, line.V4, line.V5})
}

func instanceLine(ptype string, rule []string) *ent.CasbinRule {
	instance := &ent.CasbinRule{Ptype: ptype}
	if len(rule) > 0 {
		instance.V0 = rule[0]
	}
	if len(rule) > 1 {
		instance.V1 = rule[1]
	}
	if len(rule) > 2 {
		instance.V2 = rule[2]
	}
	if len(rule) > 3 {
		instance.V3 = rule[3]
	}
	if len(rule) > 4 {
		instance.V4 = rule[4]
	}
	if len(rule) > 5 {
		instance.V5 = rule[5]
	}
	return instance
}

func savePolicyLine(cr *ent.CasbinRuleClient, ptype string, rule []string) *ent.CasbinRuleCreate {
	line := cr.Create().SetPtype(ptype)
	if len(rule) > 0 {
		line.SetV0(rule[0])
	}
	if len(rule) > 1 {
		line.SetV1(rule[1])
	}
	if len(rule) > 2 {
		line.SetV2(rule[2])
	}
	if len(rule) > 3 {
		line.SetV3(rule[3])
	}
	if len(rule) > 4 {
		line.SetV4(rule[4])
	}
	if len(rule) > 5 {
		line.SetV5(rule[5])
	}
	return line
}

func buildInstanceFilter(instance *ent.CasbinRule) []predicate.CasbinRule {
	return []predicate.CasbinRule{
		casbinrule.PtypeEQ(instance.Ptype),
		casbinrule.V0EQ(instance.V0),
		casbinrule.V1EQ(instance.V1),
		casbinrule.V2EQ(instance.V2),
		casbinrule.V3EQ(instance.V3),
		casbinrule.V4EQ(instance.V4),
		casbinrule.V5EQ(instance.V5),
	}
}

func buildFilteredFilter(ptype string, fieldIndex int, fieldValues ...string) []predicate.CasbinRule {
	cond := []predicate.CasbinRule{casbinrule.PtypeEQ(ptype)}
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V0EQ(fieldValues[0-fieldIndex]))
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V1EQ(fieldValues[1-fieldIndex]))
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V2EQ(fieldValues[2-fieldIndex]))
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V3EQ(fieldValues[3-fieldIndex]))
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V4EQ(fieldValues[4-fieldIndex]))
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		cond = append(cond, casbinrule.V5EQ(fieldValues[5-fieldIndex]))
	}
	return cond
}

var _ persist.Adapter = (*casbinAdapter)(nil)
