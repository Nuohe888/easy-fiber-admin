# EasyFiberAdmin 项目

## 项目概述

EasyFiberAdmin 是一个基于 Go 语言开发的后端管理系统框架，使用了现代化的技术栈和架构模式。该项目采用模块化设计，提供了完整的用户认证、权限管理、数据字典等基础功能，可作为后端管理系统的基础框架快速开发各类应用。

## 项目前端

- **前端项目**: [Web](https://github.com/Nuohe888/geeker-admin) - 基于geeker-admin二次开发开发

## 技术栈

- **Web框架**: Gofiber (高性能、低内存占用的Web框架)
- **ORM框架**: GORM (Go语言优秀的ORM库)
- **数据库**: MySQL (默认配置，支持多种数据库)
- **认证**: JWT (JSON Web Token)
- **权限管理**: Casbin (灵活的权限控制框架)
- **日志**: Zap (高性能日志库)
- **配置管理**: TOML (简洁的配置文件格式)
- **缓存**: Redis (高性能键值数据库)
- **API文档**: Swagger (via swaggo, API自动文档生成)
- **错误报告**: Sentry (实时错误追踪与报告)

## 项目结构

```
easy-fiber-admin/
├── boot/           # 启动相关代码
├── model/          # 数据模型定义
├── module/         # 业务模块
│   └── system/     # 系统管理模块
├── pkg/            # 公共工具包
│   ├── cache/      # 缓存相关
│   ├── config/     # 配置管理
│   ├── generator/  # 代码生成器
│   ├── jwt/        # JWT认证
│   ├── logger/     # 日志工具
│   ├── server/     # 服务器相关
│   ├── sql/        # 数据库操作
│   └── uuid/       # UUID生成工具
├── main.go         # 程序入口
└── config.toml     # 配置文件
```

## 功能特性

- **用户认证**: 支持登录、登出、token刷新等功能
- **用户管理**: 提供用户的增删改查API
- **角色管理**: 完整的RBAC权限控制系统
- **数据字典**: 支持系统参数和数据字典的管理
- **权限控制**: 基于Casbin的灵活权限管理
- **模块化设计**: 清晰的项目结构，便于扩展

## 快速开始

### 环境要求

- Go 1.20+
- MySQL 5.7+ (或其他GORM支持的数据库)

### 配置

编辑`config.toml`文件，配置数据库连接信息和服务端口等参数：

```toml
[server]
port = 18888
domain = "http://127.0.0.1:18888"
storage = "local"

[sql]
user = "root"
pass = "password"
host = "127.0.0.1"
dbName = "easy-fiber-admin"
port = 3306

[redis]
host = "127.0.0.1"
port = 6379
password = ""
db = 0

[sentry]
dsn = "YOUR_SENTRY_DSN_HERE" # 请替换为您的Sentry DSN
```

### 运行

```bash
# 安装依赖
go mod tidy

# 启动服务
go run main.go
```

服务默认在`http://localhost:18888`启动。

## API 文档

项目使用 Swagger (通过 swaggo 工具) 自动生成 API 文档。

- **访问文档**: 服务启动后，可以通过浏览器访问 `http://localhost:18888/swagger/index.html` 来查看和测试 API。
- **更新文档**: 如果您修改了 API 的代码注释（例如：控制器方法、路由或数据模型），需要重新生成 Swagger 文档。在项目根目录下运行以下命令：
  ```bash
  go run github.com/swaggo/swag/cmd/swag init --parseDependency --parseInternal --parseDepth 2
  ```
  这会更新 `docs/` 目录下的文档文件。`docs/docs.go` 文件应被提交到版本控制中。

## 缓存

为了提升性能和减少数据库负载，项目集成了 Redis 进行数据缓存。

- **用途**: 缓存常用的查询结果或计算数据。
- **配置**: Redis 的连接信息在 `config.toml` 文件中的 `[redis]` 部分进行配置。

## 错误报告

项目集成了 Sentry 用于实时的错误追踪和报告。

- **用途**: 自动捕获应用运行时发生的错误，并将其发送到 Sentry 平台，帮助开发人员快速定位和解决问题。
- **配置**: Sentry DSN 在 `config.toml` 文件中的 `[sentry]` 部分进行配置。如果 DSN 为空，Sentry 将不会初始化。
- **手动上报**: 除了自动捕获未处理的 panic，您也可以在代码中手动上报特定的错误到 Sentry，例如：
  ```go
  import "github.com/getsentry/sentry-go"
  // ...
  if err != nil {
      sentry.CaptureException(fmt.Errorf("这是一个特定的错误: %w", err))
      // ... 正常处理错误
  }
  ```

## 扩展开发

可以参考`module/system`模块的实现，添加自己的业务模块。新模块应在`module`目录下创建，并在`module/init.go`中初始化。

## 贡献

欢迎提交问题和功能需求，也欢迎提交Pull Request贡献代码。

## 许可

[Apache License 2.0](LICENSE) 