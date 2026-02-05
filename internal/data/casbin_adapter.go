/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"fmt"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	kratosLog "github.com/go-kratos/kratos/v2/log"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/data/entity/ent"
)

// Adapter implements the casbin persist.UpdatableAdapter for ent.
// Its sole responsibility is to act as a persistence layer for a Casbin Enforcer.
// It should not contain business logic or custom data manipulation methods.
type Adapter struct {
	ctx context.Context
	db  *ent.Database
	log *kratosLog.Helper
}

// NewAdapter creates a new casbin adapter.
func NewAdapter(ctx context.Context, db *ent.Database, logger log.Logger) (*Adapter, error) {
	return &Adapter{
		ctx: ctx,
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "data.casbin_adapter")),
	}, nil
}

// NewAdapterFromApp creates a new casbin adapter.
func NewAdapterFromApp(app *runtime.App, db *ent.Database) (*Adapter, error) {
	return &Adapter{
		ctx: app.Context(),
		db:  db,
		log: log.NewHelper(log.With(app.Logger(), "module", "data.casbin_adapter")),
	}, nil
}

// LoadPolicy loads all policy rules from the storage.
func (a *Adapter) LoadPolicy(m model.Model) error {
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

// SavePolicy saves all policy rules to the storage.
func (a *Adapter) SavePolicy(model model.Model) error {
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
func (a *Adapter) AddPolicy(_ string, ptype string, rule []string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		_, err := savePolicyLine(cr, ptype, rule).Save(ctx)
		return err
	})
}

// AddPolicies adds multiple policy rules to the storage.
func (a *Adapter) AddPolicies(_ string, ptype string, rules [][]string) error {
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
func (a *Adapter) UpdatePolicy(_ string, ptype string, oldRule []string, newRule []string) error {
	return a.db.Tx(a.ctx, func(ctx context.Context) error {
		return a.updatePolicyInTx(ctx, ptype, oldRule, newRule)
	})
}

// UpdatePolicies updates multiple policy rules from storage.
func (a *Adapter) UpdatePolicies(_ string, ptype string, oldRules [][]string, newRules [][]string) error {
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
func (a *Adapter) updatePolicyInTx(ctx context.Context, ptype string, oldRule []string, newRule []string) error {
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
func (a *Adapter) UpdateFilteredPolicies(_ string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
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
func (a *Adapter) RemovePolicy(_ string, ptype string, rule []string) error {
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
func (a *Adapter) RemovePolicies(_ string, ptype string, rules [][]string) error {
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
func (a *Adapter) RemoveFilteredPolicy(_ string, ptype string, fieldIndex int, fieldValues ...string) error {
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

var _ persist.UpdatableAdapter = (*Adapter)(nil)
