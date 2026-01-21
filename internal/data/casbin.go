/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package data

import (
	"fmt"
	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	"github.com/origadmin/casbin-watcher/v3"

	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/context"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/casbinrule"
	"origadmin/application/admin/internal/data/entity/ent/predicate"
)

// CasbinAdapter implements the casbin persist.UpdatableAdapter for ent.
// It now includes a watcher to send notifications upon policy changes.
type CasbinAdapter struct {
	ctx     context.Context
	db      *ent.Database
	watcher *watcher.Watcher
}

// NewAdapter creates a new casbin adapter with an integrated watcher.
// The watcherURL is used to configure the message queue backend (e.g., "redis://localhost:6379/0").
func NewAdapter(rt *runtime.App, db *ent.Database, watcherURL string) (*CasbinAdapter, error) {
	w, err := watcher.NewWatcher(rt.Context(), watcherURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create integrated watcher for adapter: %w", err)
	}

	return &CasbinAdapter{
		ctx:     rt.Context(),
		db:      db,
		watcher: w,
	}, nil
}

// Close closes the integrated watcher. This should be called on service shutdown.
func (a *CasbinAdapter) Close() {
	if a.watcher != nil {
		a.watcher.Close()
	}
}

// Context returns the context of the adapter.
func (a *CasbinAdapter) Context() context.Context {
	return a.ctx
}

// LoadPolicy loads all policy rules from the storage.
func (a *CasbinAdapter) LoadPolicy(model model.Model) error {
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

// SavePolicy saves all policy rules to the storage and sends a notification.
func (a *CasbinAdapter) SavePolicy(model model.Model) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
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
	if err != nil {
		return err
	}
	// Notify after transaction is successful
	return a.watcher.Update()
}

// AddPolicy adds a policy rule to the storage and sends a notification.
func (a *CasbinAdapter) AddPolicy(_ string, ptype string, rule []string) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		_, err := savePolicyLine(cr, ptype, rule).Save(ctx)
		return err
	})
	if err != nil {
		return err
	}
	return a.watcher.Update()
}

// UpdatePolicy updates a policy rule from storage and sends a notification.
func (a *CasbinAdapter) UpdatePolicy(_ string, ptype string, oldRule []string, newRule []string) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		oldFilter := buildInstanceFilter(ptype, oldRule)
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
	})
	if err != nil {
		return err
	}
	return a.watcher.Update()
}

// UpdatePolicies updates multiple policy rules from storage and sends a notification.
func (a *CasbinAdapter) UpdatePolicies(_ string, ptype string, oldRules [][]string, newRules [][]string) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
		if len(oldRules) != len(newRules) {
			return fmt.Errorf("old and new rules must have the same length")
		}
		for i, oldRule := range oldRules {
			newRule := newRules[i]
			// We call the internal UpdatePolicy logic within the same transaction
			err := a.updatePolicyInTx(ctx, ptype, oldRule, newRule)
			if err != nil {
				return err // This will cause a rollback
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return a.watcher.Update()
}

// updatePolicyInTx is a helper for UpdatePolicies to run within a transaction.
func (a *CasbinAdapter) updatePolicyInTx(ctx context.Context, ptype string, oldRule []string, newRule []string) error {
	cr := a.db.CasbinRule(ctx)
	oldFilter := buildInstanceFilter(ptype, oldRule)
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

// UpdateFilteredPolicies updates policy rules that match the filter from storage and sends a notification.
func (a *CasbinAdapter) UpdateFilteredPolicies(_ string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	var oldPolicies [][]string
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
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
	if err != nil {
		return nil, err
	}
	return oldPolicies, a.watcher.Update()
}

// RemovePolicy removes a policy rule from the storage and sends a notification.
func (a *CasbinAdapter) RemovePolicy(_ string, ptype string, rule []string) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		filter := buildInstanceFilter(ptype, rule)
		_, err := cr.Delete().Where(filter...).Exec(ctx)
		return err
	})
	if err != nil {
		return err
	}
	return a.watcher.Update()
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage and sends a notification.
func (a *CasbinAdapter) RemoveFilteredPolicy(_ string, ptype string, fieldIndex int, fieldValues ...string) error {
	err := a.db.Tx(a.Context(), func(ctx context.Context) error {
		cr := a.db.CasbinRule(ctx)
		cond, err := buildFilteredFilter(ptype, fieldIndex, fieldValues...)
		if err != nil {
			return err
		}
		_, err = cr.Delete().Where(cond...).Exec(ctx)
		return err
	})
	if err != nil {
		return err
	}
	return a.watcher.Update()
}

// --- Helper Functions ---

func policyToRule(p *ent.CasbinRule) []string {
	return []string{p.V0, p.V1, p.V2, p.V3, p.V4, p.V5}
}

func loadPolicyLine(line *ent.CasbinRule, model model.Model) {
	key := line.Ptype
	sec := key[:1]
	rule := policyToRule(line)
	model[sec][key].Policy = append(model[sec][key].Policy, rule)
}

func instanceLine(ptype string, rule []string) *ent.CasbinRule {
	instance := &ent.CasbinRule{Ptype: ptype}
	for i, v := range rule {
		switch i {
		case 0:
			instance.V0 = v
		case 1:
			instance.V1 = v
		case 2:
			instance.V2 = v
		case 3:
			instance.V3 = v
		case 4:
			instance.V4 = v
		case 5:
			instance.V5 = v
		}
	}
	return instance
}

func savePolicyLine(cr *ent.CasbinRuleClient, ptype string, rule []string) *ent.CasbinRuleCreate {
	line := cr.Create().SetPtype(ptype)
	for i, v := range rule {
		switch i {
		case 0:
			line.SetV0(v)
		case 1:
			line.SetV1(v)
		case 2:
			line.SetV2(v)
		case 3:
			line.SetV3(v)
		case 4:
			line.SetV4(v)
		case 5:
			line.SetV5(v)
		}
	}
	return line
}

func buildInstanceFilter(ptype string, rule []string) []predicate.CasbinRule {
	paddedRule := make([]string, 6)
	copy(paddedRule, rule)

	return []predicate.CasbinRule{
		casbinrule.PtypeEQ(ptype),
		casbinrule.V0EQ(paddedRule[0]),
		casbinrule.V1EQ(paddedRule[1]),
		casbinrule.V2EQ(paddedRule[2]),
		casbinrule.V3EQ(paddedRule[3]),
		casbinrule.V4EQ(paddedRule[4]),
		casbinrule.V5EQ(paddedRule[5]),
	}
}

func buildFilteredFilter(ptype string, fieldIndex int, fieldValues ...string) ([]predicate.CasbinRule, error) {
	if fieldIndex < 0 || fieldIndex+len(fieldValues) > 6 {
		return nil, fmt.Errorf("invalid field index or field values for filtered policy")
	}

	cond := []predicate.CasbinRule{casbinrule.PtypeEQ(ptype)}
	for i, fieldValue := range fieldValues {
		fieldNum := fieldIndex + i
		switch fieldNum {
		case 0:
			cond = append(cond, casbinrule.V0EQ(fieldValue))
		case 1:
			cond = append(cond, casbinrule.V1EQ(fieldValue))
		case 2:
			cond = append(cond, casbinrule.V2EQ(fieldValue))
		case 3:
			cond = append(cond, casbinrule.V3EQ(fieldValue))
		case 4:
			cond = append(cond, casbinrule.V4EQ(fieldValue))
		case 5:
			cond = append(cond, casbinrule.V5EQ(fieldValue))
		}
	}
	return cond, nil
}

var _ persist.UpdatableAdapter = (*CasbinAdapter)(nil)
