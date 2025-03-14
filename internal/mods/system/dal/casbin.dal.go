// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dal

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"

	"origadmin/application/admin/internal/mods/system/dal/entity/ent"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/casbinrule"
	"origadmin/application/admin/internal/mods/system/dal/entity/ent/predicate"
)

const (
	DefaultTableName = "casbin_rule"
	DefaultDatabase  = "casbin"
)

type casbinAdapterRepo struct {
	ctx      context.Context
	data     *Data
	filtered bool
}

type CasbinAdapterSetting struct {
	filtered bool
}

type Option = func(a *CasbinAdapterSetting) error

func WithFiltered(filtered bool) Option {
	return func(a *CasbinAdapterSetting) error {
		a.filtered = filtered
		return nil
	}
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

func open(driverName, dataSourceName string) (*ent.Client, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	var drv dialect.Driver
	if driverName == "pgx" {
		drv = entsql.OpenDB(dialect.Postgres, db)
	} else {
		drv = entsql.OpenDB(driverName, db)
	}
	return ent.NewClient(ent.Driver(drv)), nil
}

// NewAdapter returns an adapter by driver name and data source string.
func NewAdapter(data *Data, options ...Option) (persist.Adapter, error) {
	a := &casbinAdapterRepo{
		data:     data,
		filtered: false,
	}
	var setting CasbinAdapterSetting
	for _, option := range options {
		if err := option(&setting); err != nil {
			return nil, err
		}
	}
	a.filtered = setting.filtered
	return a, nil
}

// NewAdapterWithClient create an adapter with client passed in.
// This method does not ensure the existence of database, user should create database manually.
func NewAdapterWithClient(client *ent.Client, options ...Option) (persist.Adapter, error) {
	a := &casbinAdapterRepo{
		data: NewDataWithClient(client),
	}
	var setting CasbinAdapterSetting
	for _, option := range options {
		if err := option(&setting); err != nil {
			return nil, err
		}
	}
	a.filtered = setting.filtered
	return a, nil
}

// LoadPolicy loads all policy rules from the storage.
func (repo *casbinAdapterRepo) LoadPolicy(model model.Model) error {
	policies, err := repo.data.CasbinRule(repo.ctx).Query().Order(ent.Asc("id")).All(repo.ctx)
	if err != nil {
		return err
	}
	for _, policy := range policies {
		loadPolicyLine(policy, model)
	}
	return nil
}

// LoadFilteredPolicy loads only policy rules that match the filter.
// Filter parameter here is a Filter structure
func (repo *casbinAdapterRepo) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	filterValue, ok := filter.(Filter)
	if !ok {
		return fmt.Errorf("invalid filter type: %v", reflect.TypeOf(filter))
	}

	query := repo.data.CasbinRule(repo.ctx).Query()
	if len(filterValue.Ptype) != 0 {
		query.Where(casbinrule.PtypeIn(filterValue.Ptype...))
	}
	if len(filterValue.V0) != 0 {
		query.Where(casbinrule.V0In(filterValue.V0...))
	}
	if len(filterValue.V1) != 0 {
		query.Where(casbinrule.V1In(filterValue.V1...))
	}
	if len(filterValue.V2) != 0 {
		query.Where(casbinrule.V2In(filterValue.V2...))
	}
	if len(filterValue.V3) != 0 {
		query.Where(casbinrule.V3In(filterValue.V3...))
	}
	if len(filterValue.V4) != 0 {
		query.Where(casbinrule.V4In(filterValue.V4...))
	}
	if len(filterValue.V5) != 0 {
		query.Where(casbinrule.V5In(filterValue.V5...))
	}

	lines, err := query.All(repo.ctx)
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	repo.filtered = true

	return nil
}

// IsFiltered returns true if the loaded policy has been filtered.
func (repo *casbinAdapterRepo) IsFiltered() bool {
	return repo.filtered
}

// SavePolicy saves all policy rules to the storage.
func (repo *casbinAdapterRepo) SavePolicy(model model.Model) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		if _, err := repo.data.CasbinRule(ctx).Delete().Exec(repo.ctx); err != nil {
			return err
		}
		lines := make([]*ent.CasbinRuleCreate, 0)

		for ptype, ast := range model["p"] {
			for _, policy := range ast.Policy {
				line := repo.savePolicyLine(ctx, ptype, policy)
				lines = append(lines, line)
			}
		}

		for ptype, ast := range model["g"] {
			for _, policy := range ast.Policy {
				line := repo.savePolicyLine(ctx, ptype, policy)
				lines = append(lines, line)
			}
		}

		_, err := repo.data.CasbinRule(ctx).CreateBulk(lines...).Save(repo.ctx)
		return err
	})

}

// AddPolicy adds a policy rule to the storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) AddPolicy(sec string, ptype string, rule []string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		_, err := repo.savePolicyLine(ctx, ptype, rule).Save(repo.ctx)
		return err
	})
}

// RemovePolicy removes a policy rule from the storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) RemovePolicy(sec string, ptype string, rule []string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		instance := repo.toInstance(ptype, rule)
		_, err := repo.data.CasbinRule(ctx).Delete().Where(
			casbinrule.PtypeEQ(instance.Ptype),
			casbinrule.V0EQ(instance.V0),
			casbinrule.V1EQ(instance.V1),
			casbinrule.V2EQ(instance.V2),
			casbinrule.V3EQ(instance.V3),
			casbinrule.V4EQ(instance.V4),
			casbinrule.V5EQ(instance.V5),
		).Exec(repo.ctx)
		return err
	})
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		cond := make([]predicate.CasbinRule, 0)
		cond = append(cond, casbinrule.PtypeEQ(ptype))
		if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) && len(fieldValues[0-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V0EQ(fieldValues[0-fieldIndex]))
		}
		if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) && len(fieldValues[1-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V1EQ(fieldValues[1-fieldIndex]))
		}
		if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) && len(fieldValues[2-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V2EQ(fieldValues[2-fieldIndex]))
		}
		if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) && len(fieldValues[3-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V3EQ(fieldValues[3-fieldIndex]))
		}
		if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) && len(fieldValues[4-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V4EQ(fieldValues[4-fieldIndex]))
		}
		if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) && len(fieldValues[5-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V5EQ(fieldValues[5-fieldIndex]))
		}
		_, err := repo.data.CasbinRule(ctx).Delete().Where(
			cond...,
		).Exec(repo.ctx)
		return err
	})
}

// AddPolicies adds policy rules to the storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) AddPolicies(sec string, ptype string, rules [][]string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		return repo.createPolicies(ctx, ptype, rules)
	})
}

// RemovePolicies removes policy rules from the storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) RemovePolicies(sec string, ptype string, rules [][]string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		for _, rule := range rules {
			instance := repo.toInstance(ptype, rule)
			if _, err := repo.data.CasbinRule(ctx).Delete().Where(
				casbinrule.PtypeEQ(instance.Ptype),
				casbinrule.V0EQ(instance.V0),
				casbinrule.V1EQ(instance.V1),
				casbinrule.V2EQ(instance.V2),
				casbinrule.V3EQ(instance.V3),
				casbinrule.V4EQ(instance.V4),
				casbinrule.V5EQ(instance.V5),
			).Exec(repo.ctx); err != nil {
				return err
			}
		}
		return nil
	})
}

func loadPolicyLine(line *ent.CasbinRule, model model.Model) {
	var p = []string{
		line.Ptype,
		line.V0, line.V1, line.V2, line.V3, line.V4, line.V5,
	}

	var lineText string
	if line.V5 != "" {
		lineText = strings.Join(p, ", ")
	} else if line.V4 != "" {
		lineText = strings.Join(p[:6], ", ")
	} else if line.V3 != "" {
		lineText = strings.Join(p[:5], ", ")
	} else if line.V2 != "" {
		lineText = strings.Join(p[:4], ", ")
	} else if line.V1 != "" {
		lineText = strings.Join(p[:3], ", ")
	} else if line.V0 != "" {
		lineText = strings.Join(p[:2], ", ")
	}

	persist.LoadPolicyLine(lineText, model)
}

func (repo *casbinAdapterRepo) toInstance(ptype string, rule []string) *ent.CasbinRule {
	instance := &ent.CasbinRule{}

	instance.Ptype = ptype

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

func (repo *casbinAdapterRepo) savePolicyLine(ctx context.Context, ptype string, rule []string) *ent.CasbinRuleCreate {
	line := repo.data.CasbinRule(ctx).Create()

	line.SetPtype(ptype)
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

// UpdatePolicy updates a policy rule from storage.
// This is part of the Auto-Save feature.
func (repo *casbinAdapterRepo) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		rule := repo.toInstance(ptype, oldRule)
		line := repo.data.CasbinRule(ctx).Update().Where(
			casbinrule.PtypeEQ(rule.Ptype),
			casbinrule.V0EQ(rule.V0),
			casbinrule.V1EQ(rule.V1),
			casbinrule.V2EQ(rule.V2),
			casbinrule.V3EQ(rule.V3),
			casbinrule.V4EQ(rule.V4),
			casbinrule.V5EQ(rule.V5),
		)
		rule = repo.toInstance(ptype, newPolicy)
		line.SetV0(rule.V0)
		line.SetV1(rule.V1)
		line.SetV2(rule.V2)
		line.SetV3(rule.V3)
		line.SetV4(rule.V4)
		line.SetV5(rule.V5)
		_, err := line.Save(repo.ctx)
		return err
	})
}

// UpdatePolicies updates some policy rules to storage, like db, redis.
func (repo *casbinAdapterRepo) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	return repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		for _, policy := range oldRules {
			rule := repo.toInstance(ptype, policy)
			if _, err := repo.data.CasbinRule(ctx).Delete().Where(
				casbinrule.PtypeEQ(rule.Ptype),
				casbinrule.V0EQ(rule.V0),
				casbinrule.V1EQ(rule.V1),
				casbinrule.V2EQ(rule.V2),
				casbinrule.V3EQ(rule.V3),
				casbinrule.V4EQ(rule.V4),
				casbinrule.V5EQ(rule.V5),
			).Exec(repo.ctx); err != nil {
				return err
			}
		}
		lines := make([]*ent.CasbinRuleCreate, 0)
		for _, policy := range newRules {
			lines = append(lines, repo.savePolicyLine(ctx, ptype, policy))
		}
		if _, err := repo.data.CasbinRule(ctx).CreateBulk(lines...).Save(repo.ctx); err != nil {
			return err
		}
		return nil
	})
}

// UpdateFilteredPolicies deletes old rules and adds new rules.
func (repo *casbinAdapterRepo) UpdateFilteredPolicies(sec string, ptype string, newPolicies [][]string, fieldIndex int,
	fieldValues ...string) ([][]string, error) {
	oldPolicies := make([][]string, 0)
	err := repo.data.Tx(repo.ctx, func(ctx context.Context) error {
		cond := make([]predicate.CasbinRule, 0)
		cond = append(cond, casbinrule.PtypeEQ(ptype))
		if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) && len(fieldValues[0-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V0EQ(fieldValues[0-fieldIndex]))
		}
		if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) && len(fieldValues[1-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V1EQ(fieldValues[1-fieldIndex]))
		}
		if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) && len(fieldValues[2-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V2EQ(fieldValues[2-fieldIndex]))
		}
		if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) && len(fieldValues[3-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V3EQ(fieldValues[3-fieldIndex]))
		}
		if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) && len(fieldValues[4-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V4EQ(fieldValues[4-fieldIndex]))
		}
		if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) && len(fieldValues[5-fieldIndex]) > 0 {
			cond = append(cond, casbinrule.V5EQ(fieldValues[5-fieldIndex]))
		}
		rules, err := repo.data.CasbinRule(ctx).Query().Where(cond...).All(repo.ctx)
		if err != nil {
			return err
		}
		ruleIDs := make([]int, 0, len(rules))
		for _, r := range rules {
			ruleIDs = append(ruleIDs, r.ID)
		}

		_, err = repo.data.CasbinRule(ctx).Delete().
			Where(casbinrule.IDIn(ruleIDs...)).
			Exec(repo.ctx)
		if err != nil {
			return err
		}

		if err := repo.createPolicies(ctx, ptype, newPolicies); err != nil {
			return err
		}
		for _, rule := range rules {
			oldPolicies = append(oldPolicies, CasbinRuleToStringArray(rule))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return oldPolicies, nil
}

func (repo *casbinAdapterRepo) createPolicies(ctx context.Context, ptype string, policies [][]string) error {
	lines := make([]*ent.CasbinRuleCreate, 0)
	for _, policy := range policies {
		lines = append(lines, repo.savePolicyLine(ctx, ptype, policy))
	}
	if _, err := repo.data.CasbinRule(ctx).CreateBulk(lines...).Save(repo.ctx); err != nil {
		return err
	}
	return nil
}

func CasbinRuleToStringArray(rule *ent.CasbinRule) []string {
	arr := make([]string, 0)
	if rule.V0 != "" {
		arr = append(arr, rule.V0)
	}
	if rule.V1 != "" {
		arr = append(arr, rule.V1)
	}
	if rule.V2 != "" {
		arr = append(arr, rule.V2)
	}
	if rule.V3 != "" {
		arr = append(arr, rule.V3)
	}
	if rule.V4 != "" {
		arr = append(arr, rule.V4)
	}
	if rule.V5 != "" {
		arr = append(arr, rule.V5)
	}
	return arr
}
