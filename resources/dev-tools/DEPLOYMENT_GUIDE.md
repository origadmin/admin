# OrigAdmin Backend 部署指南

## 📋 部署配置说明

本项目提供多种部署选项以适应不同使用场景。

## 🚀 部署选项

### 选项 1：分离式部署（推荐 - 规避许可证风险）

```bash
# 1. 启动基础设施服务
docker-compose -f docker-compose.infra.yml up -d

# 2. 启动应用服务
docker-compose -f docker-compose.app.yml up -d
```

### 选项 2：一体化部署（便利但有许可证风险）

```bash
# 使用原始的 docker-compose.yml（包含所有服务）
docker-compose up -d
```

⚠️ 注意：此方式会同时分发 MIT 和 AGPLv3 许可的组件

## 📊 部署选项对比

| 部署方式                         | 包含服务  | 适用场景     | 特点          |
|------------------------------|-------|----------|-------------|
| `docker-compose.yml`         | 完整服务栈 | **默认推荐** | 开发测试，开箱即用   |
| `docker-compose.minimal.yml` | 仅应用服务 | 生产环境     | 轻量级，需外部基础设施 |
| 分离式部署                        | 分别部署  | 企业级部署    | 灵活管理，高可用    |

## 🛠 环境变量配置

应用服务通过环境变量连接外部基础设施：

```bash
# 可选的外部服务配置
# KRATOS_DISCOVERY_CONSUL_ADDRESS=your-consul-host:8500
# NATS_SERVER_URL=nats://your-nats-host:4222
```

## 🔧 故障排除

如果遇到连接问题，请检查：

1. 应用服务是否正常启动
2. 端口配置是否正确
3. 环境变量设置是否需要调整