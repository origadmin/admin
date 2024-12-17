# Backend

This is the backend for the project `OrigAdmin`

## Introduction


这个项目的架构图展示了一个多层次的目录结构，每个目录都有特定的功能和用途。以下是对每个目录的详细解释：

1. **.github**：这个目录包含了GitHub相关的配置文件，如CODE_OF_CONDUCT.md、CONTRIBUTING.md、ISSUE_TEMPLATE等，用于管理和指导项目的开发和使用。

2. **api**：这个目录包含了API相关的代码，包括HTTP、multiplatform、proto和services等子目录。这些子目录分别用于定义HTTP接口、多平台接口、协议文件和服务代码。

3. **cmd**：这个目录包含了命令行工具的代码，包括internal、multiplatform、root.go和system等子目录。这些子目录分别用于定义命令行工具的内部逻辑、多平台逻辑、根命令和系统命令。

4. **data**：这个目录包含了数据相关的文件，如admin.db，用于存储项目所需的数据。

5. **generate.go**：这个文件包含了代码生成相关的代码，用于生成项目所需的代码文件。

6. **go.sum**：这个文件包含了项目依赖的校验和，用于确保依赖的一致性和安全性。

7. **helpers**：这个目录包含了辅助工具的代码，如command、ent、errors、protobuf和resp等子目录。这些子目录分别用于定义命令行工具、实体、错误处理、协议缓冲区和响应处理等辅助功能。

8. **internal**：这个目录包含了项目内部使用的代码，如configs、generate.go、loader和mods等子目录。这些子目录分别用于定义配置、代码生成、加载器和模块等内部功能。

9. **LICENSE**：这个文件包含了项目的许可证信息。

10. **main.go**：这个文件是项目的入口点，包含了项目的初始化和启动逻辑。

11. **Makefile**：这个文件包含了项目的构建和打包命令。

12. **README.md**：这个文件包含了项目的介绍和文档。

13. **resources**：这个目录包含了项目的资源文件，如configs和docs等子目录。这些子目录分别用于定义配置文件和文档文件。

14. **third_party**：这个目录包含了第三方库的代码，如auth、buf、config、errors、gnostic、google、options、pagination、pwt、README.md和validate等子目录。这些子目录分别用于定义认证、代码生成、配置、错误处理、协议缓冲区、谷歌API、选项、分页、令牌、README.md和验证等第三方功能。

总的来说，这个项目是一个多层次的目录结构，每个目录都有特定的功能和用途。这个项目的目的是提供一套完整的API服务，包括HTTP接口、多平台接口、协议文件和服务代码，以及相关的配置、代码生成、加载器和模块等内部功能。

1. Clone the repository

```bash
git clone https://github.com/origadmin/admin.git
```

2. Install dependencies
```bash
cd admin
go mod tidy
```

3. Run the server
```bash
go run main.go start
```