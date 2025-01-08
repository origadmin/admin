/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package dto implements the functions, types, and interfaces for the module.
package dto

import (
	"github.com/casbin/casbin/v2/model"
)

type CasbinAdapterRepo interface {
	LoadPolicy(model model.Model) error
	LoadFilteredPolicy(model model.Model, filter interface{}) error
	IsFiltered() bool
	SavePolicy(model model.Model) error
	AddPolicy(sec string, ptype string, rule []string) error
	RemovePolicy(sec string, ptype string, rule []string) error
	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
	AddPolicies(sec string, ptype string, rules [][]string) error
	RemovePolicies(sec string, ptype string, rules [][]string) error
	UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error
	UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error
	UpdateFilteredPolicies(sec string, ptype string, newPolicies [][]string, fieldIndex int,
		fieldValues ...string) ([][]string, error)
}
