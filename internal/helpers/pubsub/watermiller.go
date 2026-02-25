// Package pubsub implements the functions, types, and contracts for the module.
package pubsub

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
)

func NewPublisher(cfg nats.PublisherConfig, logger watermill.LoggerAdapter) (*nats.Publisher, error) {
	publisher, err := nats.NewPublisher(
		cfg,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS publisher: %w", err)
	}
	return publisher, nil
}

func NewSubscriber(cfg nats.SubscriberConfig, logger watermill.LoggerAdapter) (*nats.Subscriber, error) {
	subscriber, err := nats.NewSubscriber(
		cfg,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS subscriber: %w", err)
	}
	return subscriber, nil
}
