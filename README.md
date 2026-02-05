# Backend

This is the backend for the project `OrigAdmin`

## Introduction

The architecture diagram demonstrates a multi-layered directory structure where each directory serves specific
functionalities:

1. **api**: Houses API-related implementations with subdirectories:
    - `http`: HTTP interface definitions
    - `multiplatform`: Cross-platform interfaces
    - `proto`: Protocol Buffer definitions
    - `services`: Service implementations

2. **cmd**: Contains CLI tool implementations with subdirectories:
    - `internal`: Core CLI logic
    - `multiplatform`: Platform-agnostic commands
    - `root.go`: Root command definitions
    - `system`: System management commands

3. **data**: Stores data files (e.g., admin.db) for persistent storage.

4. **generate.go**: Code generation utilities.

5. **helpers**: Utility modules including:
    - `command`: CLI toolkit
    - `ent`: Entity management
    - `errors`: Error handling framework
    - `protobuf`: Protocol Buffer utilities
    - `resp`: Response handlers

6. **internal**: Core internal components:
    - `configs`: Configuration management
    - `generate.go`: Internal code generators
    - `loader`: Resource loading system
    - `mods`: Modular components

7. **main.go**: Application entry point with initialization logic.

8. **Makefile**: Build automation scripts.

9. **resources**: Resource files including:
    - `configs`: Configuration templates
    - `docs`: Documentation assets

10. **third_party**: External dependencies' integration:
    - Authentication systems
    - Code generation tools
    - Configuration management
    - Error handling libraries
    - Protocol Buffer extensions
    - Google API integrations
    - Pagination utilities
    - Token management
    - Validation frameworks

This project provides a comprehensive API service solution featuring HTTP interfaces, cross-platform support, protocol
definitions, and service implementations, complemented by configuration management, code generation, and modular
architecture.

## Getting Started

1. Clone the repository

    ```bash
    # git clone URL_ADDRESS 
    git clone https://github.com/OrigAdmin/backend.git
    ```

2. Install dependencies

    ```bash
    cd backend
    go mod tidy
    ```

3. Run the application

    ```bash
    # 默认安全部署（无法律风险）
    docker-compose up -d
    
    # 或使用其他配置选项
    docker-compose -f docker-compose.dev.yml up -d
    
    # 查看部署指南了解所有选项
    more DEPLOYMENT_GUIDE.md
    ```

## 📋 部署说明

**默认配置**：`docker-compose.yml`（安全配置，仅包含 MIT 许可证组件）

*其他配置：*

- `docker-compose.dev.yml` - 更精简的配置
- 分离式部署 - 基础设施与应用服务分开管理

# SourceTree

```plainText
.
├── api/                # API interface definitions
│   ├── http/           # HTTP interface
│   ├── multiplatform/  # Cross-platform interfaces
│   ├── proto/          # Protocol Buffer definitions
│   └── services/       # Service implementations
├── cmd/                # CLI tool implementations
│   ├── internal/       # Core CLI logic
│   ├── multiplatform/  # Platform-agnostic commands
│   ├── root.go         # Root command definitions
│   └── system/         # System management commands
├── data/               # Data storage
├── generate.go         # Code generation utilities
├── helpers/            # Utility modules
│   ├── command/        # CLI toolkit
│   ├── ent/            # Entity management
│   ├── errors/         # Error handling framework
│   ├── protobuf/       # Protocol Buffer utilities
│   └── resp/           # Response handlers
├── internal/           # Core internal components
│   ├── configs/        # Configuration management
│   ├── generate.go     # Internal code generators
│   ├── loader/         # Resource loading system
│   └── mods/           # Modular components
├── main.go             # Application entry point
├── Makefile            # Build automation scripts
├── resources/          # Resource files
│   ├── configs/        # Configuration templates
│   └── docs/           # Documentation assets
├── third_party/        # External dependencies
│   ├── auth/           # Authentication systems
│   ├── codegen/        # Code generation tools
│   ├── config/         # Configuration management
│   ├── errors/         # Error handling libraries
│   ├── proto/          # Protocol Buffer extensions
│   ├── google/         # Google API integrations
│   ├── pagination/     # Pagination utilities
│   ├── token/          # Token management
│   └── validation/     # Validation frameworks
└── go.mod              # Go module dependencies
```


