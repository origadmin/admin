/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pubsub

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"

	"github.com/origadmin/casbin-watcher/v3"
	"github.com/origadmin/runtime"
	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	confpb "origadmin/application/admin/internal/conf/pb"
)

const (
	CategoryPublisher component.Category = "publisher"
	CategoryWatcher   component.Category = "watcher"
)

const (
	NameWatcher   = "watcher"
	NamePublisher = "publisher"
)

// NewPublisherHandle is the engine provider for watermill.Publisher (NATS).
func NewPublisherHandle(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Bootstrap](h)
	if err != nil {
		return nil, err
	}

	brokerConfig := cfg.GetBrokers()
	if brokerConfig == nil {
		return nil, errors.New("engine: broker configuration not found")
	}

	brokerUrl := brokerConfig.GetDefault().GetUrl()
	if brokerUrl == "" {
		return nil, errors.New("engine: broker url not found")
	}

	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))
	wmLogger := NewWatermillLogger(logger)

	publisherConfig := nats.PublisherConfig{URL: brokerUrl}
	if brokerConfig.GetDefault().GetType() == "nats" && strings.Contains(brokerUrl, "jetstream=true") {
		publisherConfig.JetStream = nats.JetStreamConfig{Disabled: false}
	}

	return NewPublisher(publisherConfig, wmLogger)
}

// NewWatcherHandle is the engine provider for casbin.Watcher (NATS).
func NewWatcherHandle(ctx context.Context, h component.Handle) (any, error) {
	cfg, err := comp.AsConfig[confpb.Bootstrap](h)
	if err != nil {
		return nil, err
	}

	brokerConfig := cfg.GetBrokers()
	if brokerConfig == nil || brokerConfig.GetDefault() == nil {
		return nil, fmt.Errorf("engine: broker configuration not found for watcher")
	}

	url := brokerConfig.GetDefault().GetUrl()
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(runtime.CategoryLogger))
	wmLogger := NewWatermillLogger(logger)

	return watcher.NewWatcher(ctx, url, watcher.WithLogger(wmLogger))
}

func Resolver(ctx context.Context, source any, opts *component.LoadOptions) (*component.ModuleConfig, error) {
	switch opts.Category {
	case CategoryPublisher:
		return &component.ModuleConfig{Active: NamePublisher, Entries: []component.ConfigEntry{
			{Name: NamePublisher, Value: source}},
		}, nil
	case CategoryWatcher:
		return &component.ModuleConfig{Active: NameWatcher, Entries: []component.ConfigEntry{
			{Name: NameWatcher, Value: source}},
		}, nil
	}
	return nil, nil
}
