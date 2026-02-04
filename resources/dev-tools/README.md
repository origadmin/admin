# Dev Tools Workspace

This directory contains Docker Compose configurations and related files exclusively for local development, integration
testing, and evaluation purposes.

## ⚠️ Important Legal & Usage Notice

These configurations are **NOT** part of the official distribution, deployment, or runtime of the main project.

- **Purpose**: The files here are provided as a convenience to quickly spin up a complete environment for development
  and testing. They simulate the integration with various external services.
- **Licensing**: The Docker images and software defined here (e.g., Windmill, Consul) are governed by their own
  respective licenses (e.g., AGPLv3, RSALv2/SSPL, Apache 2.0). You are responsible for understanding and complying with
  these licenses if you run them.
- **Production Use**: **DO NOT** use this docker-compose file directly for production deployment. For production, you
  must obtain, install, and configure these third-party services independently according to their official documentation
  and licensing terms.

## 📁 Contents

- `docker-compose.yml`: The main Docker Compose file that defines all dependent services (database, message queue,
  service discovery, etc.).
- `nats/nats.conf`: NATS server configuration file with JetStream enabled.
- `config/`: (Optional) Directory for service-specific configuration files.
- `scripts/`: (Optional) Directory for helper scripts to initialize or seed test data.

## 🚀 Quick Start

To start the complete development environment, ensure you are in this directory and run:

```bash
docker-compose up
```

To run in detached mode, add the `-d` flag:

```bash
docker-compose up -d
```

To stop and remove all containers and volumes:

```bash
docker-compose down -v
```

## 🔧 Included Services & Configuration

This environment includes the following services. Check the docker-compose.yml file for the definitive list and
versions.

| Service     | Purpose (in this context)               | Image / Source            | Key License        | Note for Production                                                                                           |
|-------------|-----------------------------------------|---------------------------|--------------------|---------------------------------------------------------------------------------------------------------------|
| PostgreSQL  | Primary database                        | `postgres:18.1-alpine`    | PostgreSQL License | Deploy independently. Default database: `origadmin`, user: `user`, password: `password`.                      |
| Redis       | Cache and ephemeral data store          | `redis:8.4.0-alpine`      | BSD-3-Clause       | Use version ≤7.2 for BSD terms, or deploy Valkey/Redis newer versions independently under their new licenses. |
| NATS Server | Message broker with JetStream           | `nats:2.12.3-alpine`      | Apache-2.0         | Deploy independently. JetStream enabled for persistence.                                                      |
| Consul      | Service discovery & configuration store | `hashicorp/consul:1.22.2` | MPL-2.0            | Deploy independently as a cluster.                                                                            |

## 🔗 Connecting from the Main Project

When the services are running, your main application (outside this workspace) can connect to them using the hostnames
and ports defined in the docker-compose.yml file (usually via localhost or the service name as hostname within the
Docker network).

Example connection strings for configuration:

- **Database**: `postgresql://user:password@localhost:5432/origadmin`
- **Redis**: `redis://localhost:6379`
- **NATS**: `nats://localhost:4222`
- **Consul HTTP API**: `http://localhost:8500`
- **Consul DNS**: `localhost:8600`

## 📝 Maintenance

- Update the service images and configurations as needed for your development.
- Regularly review the official licenses of the included software, as they may change.
- Keep this README.md updated to reflect changes in the included services.

---

**By using the configurations in this directory, you acknowledge that you have read, understood, and agree to comply
with the licenses of all third-party software involved.**
