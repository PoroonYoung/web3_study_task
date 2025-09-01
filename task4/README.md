# Web3 Study Task4 - RESTful API

基于 Go + Gin + GORM 的标准化 Web API 项目。

## 项目结构

```
task4/
├── cmd/server/          # 应用程序入口
│   └── main.go
├── config/              # 配置文件
│   ├── config.yaml      # yaml 配置文件
│   └── config.go        # 配置结构定义
├── internal/            # 内部包（不对外暴露）
│   ├── handler/         # HTTP 处理器（控制器层）
│   │   └── user.go
│   ├── service/         # 业务逻辑层
│   │   └── user.go
│   ├── repository/      # 数据访问层
│   │   └── user.go
│   └── middleware/      # 中间件
│       └── auth.go
├── pkg/                 # 可重用的包
│   ├── entity/          # 数据实体
│   │   ├── User.go
│   │   ├── Post.go
│   │   └── Comment.go
│   ├── jwt/             # JWT 工具
│   │   └── jwtutil.go
│   ├── response/        # 响应处理
│   │   └── response.go
│   └── utils/           # 工具函数
└── README.md
```

## 技术栈

- **语言**: Go 1.24+
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **配置**: YA=ML

## 功能特性

- ✅ 标准化的项目结构
- ✅ 分层架构（Handler -> Service -> Repository）
- ✅ JWT 认证中间件
- ✅ 配置文件管理
- ✅ 统一响应格式
- ✅ 数据库连接池
- ✅ GORM 关联关系（无外键约束）

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置数据库

修改 `config/config.yaml` 中的数据库配置：

```yaml
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "your_password"
  dbname: "your_database"
```

### 3. 运行应用

```bash
go run cmd/server/main.go
```

服务器将在 `:8080` 端口启动。

## API 接口

### 认证相关

#### 用户登录
```
POST /user/login?username=admin&password=admin
```

响应：
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 604800,
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com"
    }
  }
}
```

### 用户相关（需要认证）

#### 获取用户信息
```
GET /api/user/getUserInfo
Authorization: Bearer <token>
```

#### 测试接口
```
GET /api/hello
Authorization: Bearer <token>
```

## 配置说明

### 服务器配置
- `port`: 服务端口
- `mode`: 运行模式 (debug/release/test)

### 数据库配置
- 支持连接池配置
- 自动迁移数据表
- 禁用外键约束（保持数据灵活性）

### JWT 配置
- 可配置密钥和过期时间
- 支持自定义发行方
- 安全的 token 验证

## 开发规范

### 目录说明
- `cmd/`: 应用程序入口点
- `internal/`: 项目内部代码，不对外暴露
- `pkg/`: 可以被外部项目使用的库代码
- `config/`: 配置文件和配置相关代码

### 分层架构
1. **Handler层**: 处理HTTP请求，参数验证，调用Service
2. **Service层**: 业务逻辑处理，调用Repository
3. **Repository层**: 数据访问，数据库操作
4. **Entity层**: 数据模型定义

## 扩展功能

项目支持轻松扩展：
- 添加新的实体和关联关系
- 实现更多业务接口
- 集成更多中间件（日志、限流等）
- 添加单元测试
- 集成 Docker 部署
