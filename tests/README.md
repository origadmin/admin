# Test Directory Structure and Usage Guide

## Directory Structure

```
tests/
├── tools/                  # Shared test utilities - USE THESE!
│   ├── casbin.go          # Casbin enforcer setup and helpers
│   ├── database.go        # Database setup (SQLite, PostgreSQL)
│   ├── http.go            # HTTP client for API testing
│   ├── server.go          # Test server setup
│   └── assertions.go      # Common assertion helpers
│
├── fixtures/              # Test data and fixtures
│   ├── data/              # Static test data files
│   │   ├── base_permissions.yaml
│   │   ├── base_roles.yaml
│   │   ├── rbac_model.conf
│   │   └── rbac_policy.csv
│   └── paths.go           # Helper functions for test file paths
│
├── unit/                  # Unit tests (fast, isolated)
│   ├── dal/               # Data Access Layer tests
│   │   └── user_repo_test.go
│   ├── casbin/
│   │   ├── adapter_test.go
│   │   ├── casbin_modifier_test.go
│   │   ├── model_test.go
│   │   └── watcher_test.go
│   └── authorization/
│
├── integration/           # Integration tests (multiple components)
│   ├── api/
│   │   ├── role_api_test.go
│   │   ├── user_api_test.go
│   │   └── permission_api_test.go
│   └── database/
│
└── e2e/                   # End-to-end tests (requires running server)
    ├── helpers.go          # E2E-specific helper functions
    ├── concurrent_access_test.go
    ├── rbac_flow_test.go
    ├── login_test.go
    ├── nats_connection_test.go
    └── permission_boundary_test.go
```

## Usage Examples

### Unit Tests with Database

```go
package dal

import (
    "testing"
    "context"
    "origadmin/application/admin/tests/tools"
    "origadmin/application/admin/internal/data/entity/ent"
)

func TestUserRepo(t *testing.T) {
    ctx := context.Background()
    
    // Setup test database
    client := tools.SetupTestDatabase(t)
    defer client.Close()
    
    db := ent.NewDatabaseWithClient(client)
    userRepo := dal.NewUserRepo(db)
    
    // Write your test...
}
```

### Unit Tests with Casbin

```go
package casbin

import (
    "testing"
    "context"
    "origadmin/application/admin/tests/tools"
    "origadmin/application/admin/tests/fixtures"
)

func TestCasbinEnforcer(t *testing.T) {
    ctx := context.Background()
    
    // Setup test database and enforcer
    client := tools.SetupTestDatabase(t)
    defer client.Close()
    
    modelPath := fixtures.CasbinModelPath()
    enforcer := tools.NewTestEnforcer(t, client, modelPath)
    
    // Write your test...
}
```

### Integration Tests with Server

```go
package api

import (
    "testing"
    "origadmin/application/admin/tests/tools"
)

func TestRoleAPI(t *testing.T) {
    // Setup test server
    components := tools.SetupTestServer(t)
    defer components.DBClient.Close()
    defer components.Server.Close()
    
    // Use components.Server.URL for API calls
    httpClient := tools.NewTestHTTPClient(components.Server.URL)
    token := httpClient.Login(t, "admin", "admin123")
    
    // Write your test...
}
```

### HTTP Client Usage

```go
// Create HTTP client
httpClient := tools.NewTestHTTPClient("http://localhost:8000")

// Login
token := httpClient.Login(t, "admin", "admin123")

// Make requests
resp := httpClient.Get(t, "/api/v1/users", token)
resp = httpClient.Post(t, "/api/v1/users", userPayload, token)
resp = httpClient.Put(t, "/api/v1/users/1", updatePayload, token)
resp = httpClient.Delete(t, "/api/v1/users/1", token)

// Custom request
resp := httpClient.Request(t, "PATCH", "/api/v1/users/1", body, token)

// Assertions
tools.AssertHTTPStatusCode(t, resp, http.StatusOK)
var result map[string]interface{}
tools.AssertJSONBody(t, resp, &result)
```

### Casbin Helper Usage

```go
enforcer := tools.NewTestEnforcer(t, client, modelPath)

// Add policies
tools.AddTestPolicies(t, enforcer, [][]string{
    {"user1", "resource1", "read"},
    {"user2", "resource2", "write"},
})

// Add grouping policies
tools.AddTestGroupingPolicies(t, enforcer, [][]string{
    {"user1", "admin"},
    {"user2", "editor"},
})

// Assert permissions
tools.AssertPermission(t, enforcer, "user1", "resource1", "read", "", true)

// Clear policies
tools.ClearTestPolicies(t, enforcer)
```

## Running Tests

### Run all unit tests (fast)

```bash
go test -v ./tests/unit/...
```

### Run specific unit test

```bash
go test -v ./tests/unit/casbin/...
go test -v ./tests/unit/dal/...
```

### Run integration tests

```bash
go test -v ./tests/integration/...
```

### Run e2e tests (requires running server)

```bash
# First, start the server
go run cmd/server/main.go

# Then run e2e tests in another terminal
go test -v ./tests/e2e/...
```

### Run with coverage

```bash
go test -coverprofile=coverage.out ./tests/unit/...
go tool cover -html=coverage.out
```

## Best Practices

1. **Always use shared tools**: Don't reinvent the wheel. Use functions from `tests/tools/`
2. **Use t.Helper()**: Mark helper functions with `t.Helper()` for better error reporting
3. **Cleanup resources**: Always use `defer` for cleanup (database close, server close, etc.)
4. **Use fixtures**: Reference test data files using `fixtures.*Path()` functions
5. **Keep tests isolated**: Each test should be independent and not rely on other tests
6. **Use table-driven tests**: For testing multiple scenarios with the same logic

## Migration Notes

If you're migrating old tests to use the new tools:

1. Replace `enttest.Open()` with `tools.SetupTestDatabase()`
2. Replace custom HTTP helpers with `tools.TestHTTPClient`
3. Replace custom Casbin setup with `tools.NewTestEnforcer()`
4. Replace custom server setup with `tools.SetupTestServer()`
5. Use `fixtures.CasbinModelPath()` instead of hardcoded paths

## Documentation Reference

For detailed architecture analysis and troubleshooting, see:

- `docs/RBAC_FAILURE_ARCHITECTURE_ANALYSIS.md` - RBAC architecture analysis
- `docs/REFACTORING_PLAN.md` - Test directory refactoring plan
