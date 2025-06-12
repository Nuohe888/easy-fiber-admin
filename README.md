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
```

### 运行

```bash
# 安装依赖
go mod tidy

# 启动服务
go run main.go
```

服务默认在`http://localhost:18888`启动。

## 扩展开发

可以参考`module/system`模块的实现，添加自己的业务模块。新模块应在`module`目录下创建，并在`module/init.go`中初始化。

## 贡献

欢迎提交问题和功能需求，也欢迎提交Pull Request贡献代码。

## 许可

[Apache License 2.0](LICENSE) 