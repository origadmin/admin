package e2e

import (
	"os"
	"testing"
)

// TestMain is the entry point for E2E tests.
func TestMain(m *testing.M) {
	// Here you can add global setup and teardown logic,
	// such as starting a server or connecting to a database.
	exitCode := m.Run()
	os.Exit(exitCode)
}
