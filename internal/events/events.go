// Package events implements the functions, types, and interfaces for the module.
package events

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/origadmin/runtime/context"
)

type Publisher interface {
	Publish(topic string, messages ...*message.Message) (err error)
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (o <-chan *message.Message, err error)
	Close() error
}

type PubSub interface {
	Publisher
	Subscriber
}
