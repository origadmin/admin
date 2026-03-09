/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package providers

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"

	"github.com/origadmin/runtime/contracts/component"
	"github.com/origadmin/runtime/helpers/comp"
	"github.com/origadmin/runtime/log"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/pubsub"
)

// NewPublisher is the engine provider for watermill.Publisher (NATS).
func NewPublisher(ctx context.Context, h component.Handle) (any, error) {
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

	// Retrieve Runtime logger and wrap it with Watermill adapter
	logger, _ := comp.GetDefault[log.Logger](ctx, h.Locator().In(component.CategoryLogger))
	wmLogger := pubsub.NewWatermillLogger(logger)

	// 1. Prepare NATS Publisher Config
	publisherConfig := nats.PublisherConfig{
		URL: brokerUrl,
	}

	// 2. Check for JetStream
	if brokerConfig.GetDefault().GetType() == "nats" {
		if strings.Contains(brokerUrl, "jetstream=true") {
			publisherConfig.JetStream = nats.JetStreamConfig{
				Disabled: false,
			}
		}
	}

	// 3. Create Publisher
	publisher, err := pubsub.NewPublisher(publisherConfig, wmLogger)
	if err != nil {
		return nil, fmt.Errorf("engine: failed to create NATS publisher: %w", err)
	}

	return publisher, nil
}
