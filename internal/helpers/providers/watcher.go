/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"fmt"

	"github.com/origadmin/casbin-watcher/v3"
	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/pubsub"
)

// NewWatcher is the engine provider for casbin.Watcher (NATS).
func NewWatcher(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Bootstrap](h)
	if err != nil {
		return nil, err
	}

	brokerConfig := cfg.GetBrokers()
	if brokerConfig == nil || brokerConfig.GetDefault() == nil {
		return nil, fmt.Errorf("engine: broker configuration not found for watcher")
	}

	url := brokerConfig.GetDefault().GetUrl()

	// Retrieve Runtime logger and wrap it with Watermill adapter
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
	wmLogger := pubsub.NewWatermillLogger(logger)

	// Create Watcher using the core package (driver is registered via side-effect import in register.go)
	w, err := watcher.NewWatcher(ctx, url, watcher.WithLogger(wmLogger))
	if err != nil {
		return nil, fmt.Errorf("engine: failed to create watcher: %w", err)
	}

	return w, nil
}
