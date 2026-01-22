package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// Publisher defines the interface for publishing events.
type Publisher interface {
	Publish(ctx context.Context, topic string, messages ...*message.Message) error
	Close() error
}

// Subscriber defines the interface for subscribing to events.
type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
}

// PolicyUpdateEventTopic is the topic for policy update events.
const PolicyUpdateEventTopic = "events.policy.updated"

// PolicyUpdateEvent represents a generic event indicating that a Casbin policy has changed.
// We use a simple event and trigger a full policy reload on the subscriber side.
// This is simpler and more robust than trying to replicate granular changes.
type PolicyUpdateEvent struct {
	// Timestamp allows subscribers to ignore older events if needed.
	Timestamp int64 `json:"timestamp"`
	// Source identifies the service that triggered the update.
	Source string `json:"source"`
}

// NewPolicyUpdateMessage creates a new Watermill message for a policy update event.
func NewPolicyUpdateMessage(sourceService string) (*message.Message, error) {
	event := PolicyUpdateEvent{
		Timestamp: time.Now().Unix(),
		Source:    sourceService,
	}
	payload, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	return msg, nil
}
