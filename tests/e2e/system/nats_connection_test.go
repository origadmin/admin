package system

import (
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
)

func TestNATSConnection(t *testing.T) {
	// NATS server URL, assuming it's running on localhost as per docker-compose.tools.yml
	natsURL := "nats://localhost:4222"

	t.Logf("Attempting to connect to NATS at: %s", natsURL)

	// Connect to NATS
	nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))

	// Check for connection errors
	require.NoError(t, err, "Failed to connect to NATS server. Please ensure the NATS container is running and the port is correct.")

	// If connection is successful, check the status
	if nc.IsConnected() {
		t.Log("Successfully connected to NATS server!")
		nc.Close()
	} else {
		t.Fatal("NATS connection was not established, even though no error was returned.")
	}
}
