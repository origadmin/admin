/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package debounce

import (
	"context"
	"time"

	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/helpers/comp"
	confpb "origadmin/application/admin/internal/conf/pb"
)

const (
	// CategoryDebounce is the engine category for debounce components.
	CategoryDebounce component.Category = "infrastructure/debounce"
)

// Provider is the provider function for debounce components.
func Provider(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Bootstrap](h)
	if err != nil {
		return nil, err
	}
	var delay time.Duration
	policyDelay := cfg.GetAuth().GetPolicySyncDelay()
	if policyDelay != nil {
		delay = policyDelay.AsDuration()
	}
	return NewExecutor(delay), nil
}

// Resolver resolves debounce configuration.
func Resolver(ctx context.Context, root any, opts *component.LoadOptions) (*component.ModuleConfig, error) {
	return &component.ModuleConfig{
		Entries: []component.ConfigEntry{{Name: "default", Value: root}},
		Active:  "default",
	}, nil
}
