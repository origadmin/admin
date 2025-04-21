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
    go run main.go start
    ```
