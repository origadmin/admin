package nats

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/origadmin/runtime/log"
	"origadmin/application/admin/internal/conf"
)

// ProviderSet exports the NATS initializer and its dependencies.
var ProviderSet = wire.NewSet(
	ProvideConnection,
	NewInitializer,
)

// Initializer implements the Initializer interface for NATS JetStream.
type Initializer struct {
	nc  *nats.Conn
	log *log.Helper
}

// NewInitializer creates a new NATS Initializer.
func NewInitializer(nc *nats.Conn, logger log.Logger) *Initializer {
	return &Initializer{
		nc:  nc,
		log: log.NewHelper(log.With(logger, "module", "initializer.nats")),
	}
}

// ProvideConnection creates a new NATS connection for the initializer.
func ProvideConnection(c *conf.Config, logger log.Logger) (*nats.Conn, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "initializer/nats"))
	nc, err := nats.Connect(c.GetBootstrap().GetBrokers().GetDefault().GetUrl())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to nats: %w", err)
	}
	cleanup := func() {
		nc.Close()
		log.Info("NATS connection for initializer closed.")
	}
	return nc, cleanup, nil
}

// Init checks for and creates all required JetStream Streams.
func (n *Initializer) Init(ctx context.Context) error {
	n.log.Info("Starting NATS JetStream initialization...")
	js, err := n.nc.JetStream(nats.Context(ctx))
	if err != nil {
		return fmt.Errorf("failed to get JetStream context: %w", err)
	}

	// Stream for Business Events
	err = n.provisionStream(js, "EVENTS", "system.>")
	if err != nil {
		return err
	}

	// Stream for Casbin Watcher
	err = n.provisionStream(js, "CASBIN", "casbin_channel")
	if err != nil {
		return err
	}

	n.log.Info("NATS JetStream initialization completed successfully.")
	return nil
}

// provisionStream is a helper to create a stream idempotently.
func (n *Initializer) provisionStream(js nats.JetStreamContext, streamName, subject string) error {
	streamConfig := &nats.StreamConfig{
		Name:      streamName,
		Subjects:  []string{subject},
		Retention: nats.WorkQueuePolicy,
		Storage:   nats.FileStorage,
	}

	_, err := js.StreamInfo(streamName)
	if err == nil {
		n.log.Infof("Stream '%s' already exists, no action needed.", streamName)
		return nil
	}

	if err != nats.ErrStreamNotFound {
		return fmt.Errorf("failed to get info for stream '%s': %w", streamName, err)
	}

	n.log.Infof("Stream '%s' not found, creating it now for subject '%s'...", streamName, subject)
	_, err = js.AddStream(streamConfig)
	if err != nil {
		return fmt.Errorf("failed to create stream '%s': %w", streamName, err)
	}

	n.log.Infof("Successfully created NATS JetStream stream '%s'", streamName)
	return nil
}
