/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package casbin

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	watcher "github.com/origadmin/casbin-watcher/v3"
	_ "github.com/origadmin/casbin-watcher/v3/drivers/nats"
)

const (
	natsEnvVar     = "NATS_SERVER_URL"
	natsDefaultURL = "nats://localhost:4222"
	testTimeout    = 15 * time.Second
	receiveTimeout = 10 * time.Second
	initWait       = 500 * time.Millisecond
)

// setupWatcherTest is a helper function to create a new watcher for testing.
// It returns the watcher and a context cancellation function for cleanup.
func setupWatcherTest(t *testing.T, ctx context.Context) (*watcher.Watcher, func()) {
	natsURL := os.Getenv(natsEnvVar)
	if natsURL == "" {
		natsURL = natsDefaultURL
	}

	w, err := watcher.NewWatcher(ctx, natsURL)
	require.NoError(t, err, "Failed to create watcher")

	cleanup := func() {
		w.Close()
	}
	return w, cleanup
}

// TestWatcherOneToMany verifies that a single publisher's message is received by multiple subscribers.
// This is the core scenario for policy synchronization across multiple service replicas.
func TestWatcherOneToMany(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	// Create one publisher (like the 'auth' service)
	publisher, publisherCleanup := setupWatcherTest(t, ctx)
	defer publisherCleanup()

	// Create multiple subscribers (like replicas of the 'system' service)
	numSubscribers := 3
	var wg sync.WaitGroup
	wg.Add(numSubscribers)

	for i := 0; i < numSubscribers; i++ {
		go func(subscriberID int) {
			defer wg.Done()

			subscriber, subscriberCleanup := setupWatcherTest(t, ctx)
			defer subscriberCleanup()

			receivedSignal := make(chan string, 1)
			err := subscriber.SetUpdateCallback(func(msg string) {
				select {
				case receivedSignal <- msg:
				default:
				}
			})
			require.NoError(t, err)

			// Wait for the message
			select {
			case <-receivedSignal:
				t.Logf("Subscriber %d received the message.", subscriberID)
			case <-time.After(receiveTimeout):
				t.Errorf("FAILURE: Subscriber %d timed out waiting for message.", subscriberID)
			case <-ctx.Done():
				t.Errorf("FAILURE: Test context cancelled for subscriber %d.", subscriberID)
			}
		}(i)
	}

	// Give all subscribers a moment to connect and subscribe.
	time.Sleep(initWait)

	// Publish a single message.
	t.Logf("Publisher sending update notification")
	require.NoError(t, publisher.Update())

	// Wait for all subscriber goroutines to finish.
	wg.Wait()
}

// TestWatcherMultipleDispatches verifies that a subscriber receives multiple messages sent in quick succession.
func TestWatcherMultipleDispatches(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	publisher, publisherCleanup := setupWatcherTest(t, ctx)
	defer publisherCleanup()

	subscriber, subscriberCleanup := setupWatcherTest(t, ctx)
	defer subscriberCleanup()

	numMessages := 5
	receivedCount := 0
	var mu sync.Mutex
	receivedSignal := make(chan struct{}, numMessages)

	err := subscriber.SetUpdateCallback(func(msg string) {
		mu.Lock()
		receivedCount++
		mu.Unlock()
		receivedSignal <- struct{}{}
	})
	require.NoError(t, err)

	time.Sleep(initWait)

	// Publish multiple messages
	for i := 0; i < numMessages; i++ {
		require.NoError(t, publisher.Update())
	}

	// Wait for all messages to be received, with a timeout.
	for i := 0; i < numMessages; i++ {
		select {
		case <-receivedSignal:
		// Message received
		case <-time.After(receiveTimeout):
			t.Fatalf("FAILURE: Timed out waiting for all messages. Received %d of %d.", receivedCount, numMessages)
		}
	}

	mu.Lock()
	require.Equal(t, numMessages, receivedCount, "Subscriber should have received all dispatched messages.")
	mu.Unlock()
	t.Logf("SUCCESS: Subscriber received all %d messages.", numMessages)
}

// TestWatcherSelfReception verifies that a watcher can receive its own message.
// This is important for the 'auth' service to update its own policy engine.
func TestWatcherSelfReception(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	// A single watcher that is both a publisher and a subscriber.
	watcherNode, cleanup := setupWatcherTest(t, ctx)
	defer cleanup()

	receivedSignal := make(chan struct{}, 1)
	err := watcherNode.SetUpdateCallback(func(msg string) {
		select {
		case receivedSignal <- struct{}{}:
		default:
		}
	})
	require.NoError(t, err)

	time.Sleep(initWait)

	t.Logf("Watcher sending self-notification")
	require.NoError(t, watcherNode.Update())

	select {
	case <-receivedSignal:
		t.Log("SUCCESS: Watcher successfully received its own update.")
	case <-time.After(receiveTimeout):
		t.Fatal("FAILURE: Timed out waiting for self-reception.")
	}
}
