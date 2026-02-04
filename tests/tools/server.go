// Copyright 2024 OrigAdmin. All rights reserved.

package tools

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"entgo.io/ent/dialect"
	"origadmin/application/admin/internal/data/entity/ent"
	"origadmin/application/admin/internal/data/entity/ent/enttest"
)

// TestServerComponents holds all components for running integration tests.
type TestServerComponents struct {
	Ctx      context.Context
	DBClient *ent.Client
	Server   *httptest.Server
	Router   *gin.Engine
}

// SetupTestServer creates a test server with in-memory database and returns components.
// Note: This is a basic setup. You need to register routes manually for your specific use case.
func SetupTestServer(t *testing.T) *TestServerComponents {
	t.Helper()

	// Setup in-memory SQLite database
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")

	// Initialize router and server
	gin.SetMode(gin.TestMode)
	router := gin.New()
	httpServer := httptest.NewServer(router)
	t.Cleanup(httpServer.Close)

	// TODO: Register routes here based on your test requirements
	// Example: system.RegisterRoutes(router, client)

	return &TestServerComponents{
		Ctx:      context.Background(),
		DBClient: client,
		Server:   httpServer,
		Router:   router,
	}
}
