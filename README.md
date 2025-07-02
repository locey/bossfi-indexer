# bossfi-indexer 项目 README

## 项目概述


## 根据官网命名规范建议

1. 文件名：全小写，单词间用下划线 如：config_loader.go
2. 函数/变量：驼峰式命名（CamelCase） 如：getUserInfo()、GetUserInfo()
3. 包名：推荐使用简洁、小写的单个单词命名，确实需要多个单词，可以如：package userutils

参考：https://go.dev/doc/effective_go#file_names

## 目录结构说明
```text
目录结构说明

├── sql/                      # SQL脚本目录
│   └── bossfi.sql
├── src/                      # 源代码目录
│   ├── app/                  # 应用程序目录（日常业务需求在此层开发）
│   │   ├── model/            # 数据模型目录（结构体 + 基础CRUD）
│   │   │   └── demo.go
│   │   ├── router/           # 路由目录
│   │   │   └── router_v1.go
│   │   ├── service/          # 服务层目录（对Model层的编排）
│   │   │   └── demo.go
│   │   └── api/              # 控制器目录（路由 + 参数校验 + 业务处理：调用Service）
│   │       ├── evm.go
│   │       └── demo.go
│   ├── core/                 # 核心功能目录（业务层面开发不动此包）
│   │   ├── db/               # 数据库相关目录
│   │   │   ├── init.go
│   │   │   ├── pgsql.go
│   │   │   └── redis.go
│   │   ├── ctx/              # 上下文相关目录
│   │   │   └── context.go
│   │   ├── gin/              # Gin相关目录
│   │   │   ├── router/       # 路由相关目录
│   │   │   │   └── router.go
│   │   │   └── middleware/   # 中间件目录
│   │   │       ├── recover.go  # 异常处理中间件
│   │   │       ├── http_log.go # HTTP日志中间件
│   │   │       └── language.go # 多语言处理中间件
│   │   ├── log/              # 日志相关目录
│   │   │   └── log.go
│   │   ├── app.go            # 应用程序入口相关文件
│   │   ├── config/           # 配置相关目录
│   │   │   └── config.go
│   │   ├── result/           # 结果处理相关目录
│   │   │   └── result.go
│   │   └── chainclient/      # 区块链客户端相关目录
│   │       ├── evm/          # EVM相关目录
│   │       │   └── evm.go
│   │       ├── domain/       # 领域模型相关目录
│   │       │   └── block.go
│   │       └── service.go    # 服务相关文件
│   ├── common/               # 公共代码目录
│   │   ├── chain/            # 区块链相关公共代码目录
│   │   │   └── constants.go
│   │   └── utils.go          # 工具函数文件
│   ├── main.go               # 主程序入口文件
├── config/                   # 配置文件目录
│   ├── config.toml           # 配置文件
│   └── config.toml.example   # 配置文件示例
├── go.mod                    # Go模块文件
├── go.sum                    # Go依赖校验文件
└── README.md                 # 项目说明文件

```

### go.mod & go.sum
- Go 模块依赖管理文件

## 核心功能

1. **多语言支持**:
    - 通过 `middleware/language.go` 实现语言中间件
    - 支持从请求头或 URL 参数获取语言标识

2. **统一响应格式**:
    - 定义在 `src/core/result/result.go` 中
    - 包含状态码、消息和数据字段

3. **错误处理**:
    - 预定义了多种错误码和对应的多语言消息

4. **数据库访问**:
    - 支持 PostgreSQL 和 Redis

## 快速开始

1. 克隆项目
2. 复制 `config.toml.example` 为 `config.toml` 并修改配置
3. 运行 `go mod tidy` 安装依赖
4. 运行 `go run main.go` 启动服务
5. 安装 swag 命令 `go install github.com/swaggo/swag/cmd/swag@latest`
6. 生成swagger文档 `swag init -g src/main.go -o src/docs`

## Swagger 文档
- http://localhost:8000/swagger/index.html

示例 获取用户余额

- GET http://localhost:8000/api/v1/user_balance/0x69b821F23bc4E537d82a65593b032B8ad13B6c0c


