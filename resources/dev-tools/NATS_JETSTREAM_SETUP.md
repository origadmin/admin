# NATS JetStream Setup Guide

This document provides instructions on how to correctly configure NATS JetStream for this project to avoid common
startup errors.

## 1. The Problem: `no stream matches subject`

When starting the application, you may encounter a fatal error similar to this:

```
FATAL msg="failed to wire app: cannot subscribe: nats: no stream matches subject"
```

This error occurs because the application's services (like the Casbin `Watcher` or Watermill `Subscriber`) are trying to
create durable consumers on NATS JetStream `Streams` that do not exist on the server.

## 2. The Cause: Missing Stream Definition

NATS JetStream requires that `Streams` are explicitly defined on the server *before* a consumer can subscribe. A
`Stream` acts as a message log that captures and persists all messages published to a specific set of subjects.

## 3. The Solution: Automated Initialization via `initializer` Command

This project provides a dedicated `initializer` command to automatically and idempotently prepare all necessary
infrastructure and seed initial data. This command should be run **once** when setting up a new environment or whenever
infrastructure/data changes require it.

### Step 1: Ensure NATS Server is Running

Make sure your NATS server is running and accessible at the address specified in your configuration files (e.g.,
`localhost:4222`).

### Step 2: Run the `initializer` Command

From the root of the `projects/backend` directory, execute the following command:

```sh
go run ./cmd/initializer/main.go -conf ./resources/configs
```

This command will:

1. Read your application's configuration.
2. Connect to the NATS server.
3. Execute all registered `Initializer` components, which include:
    - **NATS JetStream Stream Provisioning**: It will check if the required `Streams` (`EVENTS` for business events and
      `CASBIN` for the watcher) exist. If a `Stream` does not exist, it will be created with the correct configuration.
      If it already exists, it will do nothing (idempotent).
    - **Data Seeding**: It will execute any configured data seeding logic (e.g., creating initial users, roles,
      permissions). This process is also designed to be idempotent.

You only need to run this command **once** for a new environment setup, or whenever new infrastructure components or
initial data are required by the application. It is safe to run multiple times.

### What the `initializer` does

The `initializer` command automatically prepares the following:

1. **NATS JetStream Streams**:
    - **`EVENTS` Stream**:
        - **Subjects**: `system.>`
        - **Purpose**: Captures all business events published by the `system` service.
    - **`CASBIN` Stream**:
        - **Subjects**: `casbin_channel` (or as configured for the watcher)
        - **Purpose**: Used by the Casbin `Watcher` to synchronize policy updates across multiple service instances.
2. **Initial Data Seeding**:
    - **Purpose**: Populates the database with essential initial data (e.g., default admin user, basic roles, system
      settings).

## 4. Application Configuration (`server.yaml`)

Ensure your application's `server.yaml` correctly points to the NATS server and has JetStream enabled for Watermill
consumers.

**File:** `resources/configs/server.yaml`

```yaml
# ... (other server configurations) ...

# Auth Service Watermill Server (for event consumption)
- name: "auth"
  protocol: "watermill"
  watermill:
    broker:
      type: "nats"
      nats:
        address: "localhost:4222" # Your NATS server address
        jetstream_enabled: true # This must be true
```

## 5. Summary

By using the dedicated `initializer` command, the setup of NATS JetStream and initial data seeding are automated and
integrated into the project's workflow, similar to database migrations. This eliminates the need for manual setup and
ensures a consistent and reliable environment.
