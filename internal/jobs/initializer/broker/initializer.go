package broker

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	natsio "github.com/nats-io/nats.go"

	"github.com/origadmin/runtime/log"
	confpb "origadmin/application/admin/internal/conf/pb"
	"origadmin/application/admin/internal/helpers/pubsub"
)

// Initializer handles pre-start initialization for message brokers.
// For now, it specifically handles NATS JetStream provisioning to ensure
// streams and consumers exist before the application starts subscribing.
type Initializer struct {
	bootstrap *confpb.Bootstrap
	logger    log.Logger
	log       *log.Helper
}

// NewInitializer creates a new broker Initializer.
func NewInitializer(b *confpb.Bootstrap, logger log.Logger) *Initializer {
	return &Initializer{
		bootstrap: b,
		logger:    logger,
		log:       log.NewHelper(log.With(logger, "module", "initializer.broker")),
	}
}

// Init triggers the provisioning for any brokers that require it.
func (i *Initializer) Init(_ context.Context) error {
	i.log.Info("Starting broker initialization...")

	allBrokers := i.bootstrap.GetBrokers().GetConfigs()
	if def := i.bootstrap.GetBrokers().GetDefault(); def != nil {
		isDup := false
		for _, b := range allBrokers {
			if b.GetName() == def.GetName() && def.GetName() != "" {
				isDup = true
				break
			}
		}
		if !isDup {
			allBrokers = append(allBrokers, def)
		}
	}

	if len(allBrokers) == 0 {
		i.log.Info("No brokers configured, skipping initialization.")
		return nil
	}

	for _, brokerConfig := range allBrokers {
		// Currently, only NATS JetStream requires pre-initialization.
		// This can be extended for other broker types in the future.
		if brokerConfig.GetType() == "nats" {
			if err := i.provisionNatsJetStream(brokerConfig.GetUrl()); err != nil {
				// We return the error to halt startup if provisioning fails,
				// as it will lead to fatal errors on subscription attempts.
				return fmt.Errorf("failed to provision NATS JetStream for URL %s: %w", brokerConfig.GetUrl(), err)
			}
		}
	}

	i.log.Info("Broker initialization completed successfully.")
	return nil
}

// provisionNatsJetStream uses the Watermill NATS driver's AutoProvision feature
// to ensure a JetStream stream and its consumers are correctly provisioned before use.
// It relies on the simple behavior where the stream name is the same as the topic name.
func (i *Initializer) provisionNatsJetStream(brokerURL string) error {
	parsedURL, err := url.Parse(brokerURL)
	if err != nil {
		i.log.Warnf("Skipping invalid NATS broker URL: %s, error: %v", brokerURL, err)
		// Don't treat as a fatal error, just skip.
		return nil
	}

	// We only care about URLs that are explicitly marked for JetStream.
	if parsedURL.Query().Get("jetstream") != "true" {
		return nil
	}

	topic := strings.Trim(parsedURL.Path, "/")
	if topic == "" {
		i.log.Warnf("Skipping JetStream NATS broker with missing topic in URL path: %s", brokerURL)
		return nil
	}

	queueGroup := parsedURL.Query().Get("queue_group")
	if queueGroup == "" {
		i.log.Warnf("Skipping JetStream NATS broker with missing 'queue_group' query parameter: %s", brokerURL)
		return nil
	}

	i.log.Infof("Provisioning JetStream for topic: '%s' with queue group: '%s'...", topic, queueGroup)

	// Parse deliver_policy from URL parameters
	deliverPolicy := parsedURL.Query().Get("deliver_policy")
	if deliverPolicy == "" {
		deliverPolicy = "last" // Default to 'last' to prevent consuming all historical messages
	}
	i.log.Infof("Using deliver_policy: '%s' for topic: '%s'", deliverPolicy, topic)

	// Configure JetStream to use AutoProvision. This will create a stream with the same name as the topic.
	// The StreamConfig field does not exist in this version of the library and must not be used.
	subOpts := []natsio.SubOpt{natsio.Durable(queueGroup)}

	// Add deliver policy option based on configuration
	switch deliverPolicy {
	case "new":
		subOpts = append(subOpts, natsio.DeliverNew())
	case "last":
		subOpts = append(subOpts, natsio.DeliverLast())
	case "last_per_subject":
		subOpts = append(subOpts, natsio.DeliverLastPerSubject())
	case "all":
		subOpts = append(subOpts, natsio.DeliverAll())
	default:
		// Default to 'last' for safety
		i.log.Warnf("Unknown deliver_policy '%s', defaulting to 'last'", deliverPolicy)
		subOpts = append(subOpts, natsio.DeliverLast())
	}

	jetStreamConfig := nats.JetStreamConfig{
		AutoProvision:    true,
		SubscribeOptions: subOpts,
	}

	// Use the Watermill driver to provision the stream.
	// We create a temporary subscriber; the act of creation with JetStream config
	// is what triggers the stream and consumer creation in Watermill.
	subscriber, err := nats.NewSubscriber(
		nats.SubscriberConfig{
			URL:              brokerURL,
			NatsOptions:      nil,
			Unmarshaler:      nats.GobMarshaler{},
			QueueGroupPrefix: queueGroup,
			JetStream:        jetStreamConfig,
		},
		pubsub.NewWatermillLogger(i.logger),
	)
	if err != nil {
		return fmt.Errorf("could not create watermill subscriber for provisioning: %w", err)
	}

	// We don't actually subscribe, just creating the client is enough to trigger provisioning.
	// Close it immediately to clean up resources.
	if err := subscriber.Close(); err != nil {
		// Log this error but don't fail, as the stream might have been created
		// successfully before the close error occurred.
		i.log.Warnf("Error closing temporary watermill subscriber for topic '%s': %v", topic, err)
	}

	i.log.Infof("Successfully provisioned JetStream for topic: '%s' with queue group: '%s'", topic, queueGroup)
	return nil
}
