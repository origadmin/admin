// Package pubsub implements the functions, types, and contracts for the module.
package pubsub

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"

	runtimebroker "github.com/origadmin/runtime/contracts/broker"
)

type pub struct {
	message.Publisher
}

func (p *pub) Publish(ctx context.Context, topic string, messages ...runtimebroker.Message) error {
	wmMessages := make([]*message.Message, len(messages))
	for i, msg := range messages {
		wmMessages[i] = message.NewMessage(msg.GetId(), msg.GetPayload())
		wmMessages[i].Metadata = msg.GetMetadata()
	}
	return p.Publisher.Publish(topic, wmMessages...)
}

func WatermillPublisher(publisher message.Publisher) runtimebroker.Publisher {
	return &pub{
		Publisher: publisher,
	}
}

func NewPublisher(cfg nats.PublisherConfig, logger watermill.LoggerAdapter) (message.Publisher, error) {
	publisher, err := nats.NewPublisher(
		cfg,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS publisher: %w", err)
	}
	return publisher, nil
}

type sub struct {
	message.Subscriber
	outputChan chan runtimebroker.Message
}

type brokerMessage struct {
	*message.Message
}

func (b brokerMessage) GetId() string {
	return b.UUID
}

func (b brokerMessage) GetMetadata() map[string]string {
	return b.Metadata
}

func (b brokerMessage) GetPayload() []byte {
	return b.Payload
}

func (s *sub) Subscribe(ctx context.Context, topic string) (<-chan runtimebroker.Message, error) {
	wmChan, err := s.Subscriber.Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}
	go func() {
		for msg := range wmChan {
			s.outputChan <- &brokerMessage{
				Message: msg,
			}
		}
	}()
	return s.outputChan, nil
}

func (s *sub) Close() error {
	if s.outputChan != nil {
		close(s.outputChan)
		s.outputChan = nil
	}
	return s.Subscriber.Close()
}

func WatermillSubscriber(subscriber message.Subscriber) runtimebroker.Subscriber {
	return &sub{
		Subscriber: subscriber,
		outputChan: make(chan runtimebroker.Message, 100),
	}
}

func NewSubscriber(cfg nats.SubscriberConfig, logger watermill.LoggerAdapter) (message.Subscriber, error) {
	subscriber, err := nats.NewSubscriber(
		cfg,
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create NATS subscriber: %w", err)
	}
	return subscriber, nil
}
