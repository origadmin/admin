/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package casbin implements the functions, types, and interfaces for the module.
package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type adapter struct {
	typedPolicies map[string][][]string
}

func (a *adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	for _, rule := range rules {
		err := a.AddPolicy(sec, ptype, rule)
		if err != nil {
			return fmt.Errorf("error adding policy: %v", err)
		}
	}
	return nil
}

func (a *adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	for _, rule := range rules {
		err := a.RemovePolicy(sec, ptype, rule)
		if err != nil {
			return fmt.Errorf("error removing policy: %v", err)
		}
	}
	return nil
}

func (a *adapter) LoadPolicy(model model.Model) error {
	var (
		idx      int
		policies [][]string
		ptype    string
	)
	for ptype, policies = range a.typedPolicies {
		for idx = range policies {
			err := persist.LoadPolicyArray(append([]string{ptype}, policies[idx]...), model)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *adapter) savePolicyLine(ptype string, rule []string) {
	rule = append([]string{ptype}, rule...)
	a.typedPolicies[ptype] = append(a.typedPolicies[ptype], rule)
}

func (a *adapter) SavePolicy(m model.Model) error {
	var (
		ptype string
		rule  []string
		ast   *model.Assertion
	)
	for ptype, ast = range m["p"] {
		for _, rule = range ast.Policy {
			a.savePolicyLine(ptype, rule)
		}
	}

	for ptype, ast = range m["g"] {
		for _, rule = range ast.Policy {
			a.savePolicyLine(ptype, rule)
		}
	}
	return nil
}

func (a *adapter) AddPolicy(sec string, ptype string, rule []string) error {
	a.savePolicyLine(ptype, rule)
	return nil
}

func (a *adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	var newPolices, polices [][]string
	polices, ok := a.typedPolicies[ptype]
	if !ok {
		return nil
	}
	var (
		idx int
	)
	for idx = range polices {
		if arrayEquals(rule, polices[idx][1:]) {
			continue
		}
		newPolices = append(newPolices, polices[idx])
	}
	a.typedPolicies[ptype] = newPolices
	return nil
}

func arrayEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (a *adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	var newPolices, polices [][]string
	polices, ok := a.typedPolicies[ptype]
	if !ok {
		return nil
	}
	var (
		idx int
	)
	for idx = range polices {
		if arrayEquals(fieldValues, polices[idx][fieldIndex+1:]) {
			continue
		}
		newPolices = append(newPolices, polices[idx])
	}
	a.typedPolicies[ptype] = newPolices
	return nil
}

func (a *adapter) SetPolicies(policies map[string][][]string) {
	a.typedPolicies = policies
}

func NewAdapter() persist.Adapter {
	return &adapter{
		typedPolicies: make(map[string][][]string),
	}
}
func NewAdapterWithPolicies(policies map[string][][]string) persist.Adapter {
	return &adapter{
		typedPolicies: policies,
	}
}

var _ persist.Adapter = (*adapter)(nil)
var _ persist.BatchAdapter = (*adapter)(nil)
