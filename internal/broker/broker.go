// Package broker defines the interfaces for message broker interactions.
package broker

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/origadmin/runtime/context"
)

// Publisher defines the interface for publishing messages.
type Publisher interface {
	Publish(topic string, messages ...*message.Message) (err error)
	Close() error
}

// Subscriber defines the interface for subscribing to messages.
type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (o <-chan *message.Message, err error)
	Close() error
}

// PubSub combines Publisher and Subscriber interfaces.
type PubSub interface {
	Publisher
	Subscriber
}
