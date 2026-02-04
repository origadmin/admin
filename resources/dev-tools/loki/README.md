# Loki + Grafana Log Integration

## ⚠️ License Notice

**This configuration uses AGPL 3.0 licensed components:**

- Grafana Loki: [AGPL 3.0](https://github.com/grafana/loki/blob/main/LICENSE)
- Grafana Promtail: [AGPL 3.0](https://github.com/grafana/loki/blob/main/LICENSE)
- Grafana: [AGPL 3.0](https://github.com/grafana/grafana/blob/main/LICENSE)

**AGPL 3.0 Risk Warning:**

- If you provide services based on these components over the network
- You may be required to provide source code to users
- **Please evaluate legal risks before commercial deployment**

## Use Cases

✅ **Recommended:**

- Internal development and testing
- Open source projects
- Learning and research

⚠️ **Use with Caution:**

- Private deployment to a single client
- Internal enterprise use

❌ **Not Recommended:**

- Providing multi-tenant SaaS services
- As a core component of commercial products

## Quick Start

```bash
cd resources/dev-tools
docker compose -f docker-compose.loki.yml up -d
```

## Access URLs

| Service  | URL                   | Credentials |
|----------|-----------------------|-------------|
| Grafana  | http://localhost:3000 | admin/admin |
| Loki API | http://localhost:3100 | -           |

## Configuration Files

- `loki/loki-config.yaml` - Loki log storage configuration
- `loki/promtail-config.yaml` - Promtail log collection configuration
- `loki/grafana/provisioning/` - Grafana automatic data source configuration

## Log Query Examples

Use LogQL in Grafana Explore:

```logql
# Query all logs
{service="auth"}

# Query error logs
{level="error"}

# Search for specific content
{service="gateway"} |= "error"

# Regex matching
{service="auth"} |~ "trace_id=.*"

# Calculate error rate
rate({level="error"}[5m])
```

## Data Persistence

Loki and Grafana use Docker volumes for data persistence:

```bash
# List data volumes
docker volume ls | grep loki

# Delete data (⚠️ use with caution)
docker volume rm origadmin_loki_data origadmin_grafana_data
```

## Stop Services

```bash
docker compose -f docker-compose.loki.yml down
```

## Cleanup Data

```bash
docker compose -f docker-compose.loki.yml down -v
```

## Security Recommendations

1. **Change Passwords in Production**
   ```yaml
   # docker-compose.loki.yml
   environment:
     - GF_SECURITY_ADMIN_PASSWORD=your-secure-password
   ```

2. **Configure HTTPS** (production environment)

3. **Restrict Access**
   ```yaml
   # docker-compose.loki.yml
   ports:
     - "127.0.0.1:3000:3000"  # Local access only
   ```

## Alternative Solutions

If you need a fully open source and commercially friendly solution, use:

- [OpenSearch Solution](../docker-compose.opensearch.yml) - Apache 2.0

## More Information

- [Loki Documentation](https://grafana.com/docs/loki/latest/)
- [Promtail Documentation](https://grafana.com/docs/loki/latest/clients/promtail/)
- [Grafana Documentation](https://grafana.com/docs/grafana/latest/)
