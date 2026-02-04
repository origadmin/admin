# OpenSearch Log Integration ⭐ Recommended

## ✅ License Notice

**This configuration uses Apache 2.0 licensed components:**

- OpenSearch: [Apache 2.0](https://github.com/opensearch-project/OpenSearch/blob/main/LICENSE)
- OpenSearch Dashboards: [Apache 2.0](https://github.com/opensearch-project/OpenSearch-Dashboards/blob/main/LICENSE)
- Fluent Bit: [Apache 2.0](https://github.com/fluent/fluent-bit/blob/master/LICENSE)

**License Advantages:**

- ✅ Fully open source and commercially friendly
- ✅ No requirement to open source your application code
- ✅ Suitable for all scenarios (including commercial projects)

## Use Cases

✅ **Recommended For:**

- All scenarios (development, testing, production)
- Commercial projects
- Multi-tenant SaaS services
- Private deployment

## Quick Start

```bash
cd resources/dev-tools
docker compose -f docker-compose.opensearch.yml up -d
```

## Access URLs

| Service               | URL                   | Description             |
|-----------------------|-----------------------|-------------------------|
| OpenSearch Dashboards | http://localhost:5601 | Visualization Interface |
| OpenSearch API        | http://localhost:9200 | REST API                |

## Configuration Files

- `docker-compose.opensearch.yml` - Service orchestration configuration
- `fluent-bit/fluent-bit.conf` - Log collection configuration

## Querying Logs

### Via OpenSearch Dashboards

1. Visit http://localhost:5601
2. Navigate to the "Discover" panel
3. Select index pattern (e.g., `origadmin-logs*`)
4. Query logs

### Via API

```bash
# Search all logs
curl -X GET "localhost:9200/_search?q=*"

# Search by service
curl -X GET "localhost:9200/_search?q=service.name:auth"

# Search error logs
curl -X GET "localhost:9200/_search?q=level:error"
```

### Query Syntax Examples

```json
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "service.name": "auth"
          }
        },
        {
          "match": {
            "level": "error"
          }
        }
      ]
    }
  }
}
```

## Data Persistence

OpenSearch uses Docker volume for data persistence:

```bash
# List data volumes
docker volume ls | grep opensearch

# Backup data
docker run --rm -v origadmin_opensearch_data:/data -v $(pwd):/backup \
  ubuntu tar czf /backup/opensearch-backup.tar.gz /data
```

## Stop Services

```bash
docker compose -f docker-compose.opensearch.yml down
```

## Cleanup Data

```bash
docker compose -f docker-compose.opensearch.yml down -v
```

## Performance Optimization

### Adjust JVM Memory

```yaml
# docker-compose.opensearch.yml
opensearch:
  environment:
    - "OPENSEARCH_JAVA_OPTS=-Xms1g -Xmx1g"  # Adjust based on server memory
```

### Optimize Index Strategy

```json
PUT /_template/origadmin-template
{
  "index_patterns": ["origadmin-logs-*"],
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0,
    "index.lifecycle.name": "logs-policy"
  }
}
```

## Security Recommendations

### 1. Enable Security Authentication (Production)

```yaml
# docker-compose.opensearch.yml
opensearch:
  environment:
    - DISABLE_SECURITY_PLUGIN=false
    - discovery.type=single-node
    - plugins.security.ssl.http.enabled=false
```

### 2. Configure Firewall

```bash
# Allow only specific IP access
iptables -A INPUT -p tcp --dport 9200 -s 192.168.1.0/24 -j ACCEPT
iptables -A INPUT -p tcp --dport 9200 -j DROP
```

### 3. Use HTTPS

Configure reverse proxy (Nginx) to enable TLS

## Monitoring and Alerting

OpenSearch Dashboards supports creating alert rules:

1. Navigate to "Alerting" → "Create Alert"
2. Set conditions (e.g., error rate > 5%)
3. Configure notification methods (email, Webhook)

## Comparison with Other Solutions

| Feature             | OpenSearch   | Loki        | ELK     |
|---------------------|--------------|-------------|---------|
| License             | Apache 2.0 ✅ | AGPL 3.0 ⚠️ | SSPL ⚠️ |
| Performance         | ⭐⭐⭐⭐         | ⭐⭐⭐⭐⭐       | ⭐⭐⭐⭐    |
| Visualization       | ⭐⭐⭐⭐         | ⭐⭐⭐⭐⭐       | ⭐⭐⭐⭐    |
| Ease of Use         | ⭐⭐⭐⭐         | ⭐⭐⭐         | ⭐⭐⭐     |
| Commercial Friendly | ✅            | ⚠️          | ⚠️      |

## FAQ

### Q: What's the difference between OpenSearch and Elasticsearch?

A: OpenSearch is an open source fork of Elasticsearch, maintained by AWS, fully open source.

### Q: How's the performance?

A: OpenSearch has excellent performance, suitable for large-scale log collection. Recommend at least 4GB RAM.

### Q: How to upgrade?

A:

```bash
docker compose -f docker-compose.opensearch.yml pull
docker compose -f docker-compose.opensearch.yml up -d
```

## More Information

- [OpenSearch Documentation](https://opensearch.org/docs/)
- [Fluent Bit Documentation](https://docs.fluentbit.io/manual/)
- [OpenSearch Dashboards Guide](https://opensearch.org/docs/latest/dashboards/index/)
