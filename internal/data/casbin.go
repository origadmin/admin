/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"context"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// casbinAdapter implements the casbin persist.Adapter for casbin.
type casbinAdapter struct {
	db *ent.Database
}

// NewAdapter creates a new casbin adapter.
func NewAdapter(db *ent.Database) (persist.Adapter, error) {
	return &casbinAdapter{db: db}, nil
}

// LoadPolicy loads all policy rules from the storage.
func (a *casbinAdapter) LoadPolicy(model model.Model) error {
	ctx := context.Background()
	client := a.db.Client(ctx)
	policies, err := client.CasbinRule.Query().Order(ent.Asc("id")).All(ctx)
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
	ctx := context.Background()
	tx, err := a.db.Client(ctx).Tx(ctx)
	if err != nil {
		return err
	}
	if _, err := tx.CasbinRule.Delete().Exec(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	lines := make([]*ent.CasbinRuleCreate, 0)
	for ptype, ast := range model["p"] {
		for _, policy := range ast.Policy {
			lines = append(lines, savePolicyLine(tx, ptype, policy))
		}
	}
	for ptype, ast := range model["g"] {
		for _, policy := range ast.Policy {
			lines = append(lines, savePolicyLine(tx, ptype, policy))
		}
	}
	if _, err := tx.CasbinRule.CreateBulk(lines...).Save(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// AddPolicy adds a policy rule to the storage.
func (a *casbinAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	ctx := context.Background()
	tx, err := a.db.Client(ctx).Tx(ctx)
	if err != nil {
		return err
	}
	if _, err := savePolicyLine(tx, ptype, rule).Save(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// RemovePolicy removes a policy rule from the storage.
func (a *casbinAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	ctx := context.Background()
	tx, err := a.db.Client(ctx).Tx(ctx)
	if err != nil {
		return err
	}
	instance := toInstance(ptype, rule)
	if _, err := tx.CasbinRule.Delete().Where(buildInstanceFilter(instance)...).Exec(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *casbinAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	ctx := context.Background()
	tx, err := a.db.Client(ctx).Tx(ctx)
	if err != nil {
		return err
	}
	cond := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
	if _, err := tx.CasbinRule.Delete().Where(cond...).Exec(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// --- Helper Functions ---

func loadPolicyLine(line *ent.CasbinRule, model model.Model) {
	key := line.Ptype
	sec := key[:1]
	model[sec][key].Policy = append(model[sec][key].Policy, []string{line.V0, line.V1, line.V2, line.V3, line.V4, line.V5})
}

func toInstance(ptype string, rule []string) *ent.CasbinRule {
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

func savePolicyLine(tx *ent.Tx, ptype string, rule []string) *ent.CasbinRuleCreate {
	line := tx.CasbinRule.Create().SetPtype(ptype)
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
