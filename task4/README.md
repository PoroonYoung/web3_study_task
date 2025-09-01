# Web3 Study Task4 - 博客系统 API

基于 Go + Gin + GORM 实现的博客系统后端 API，支持用户认证、文章管理和评论功能。

## 项目功能

### 用户系统
- 用户登录认证（JWT Token）
- 获取用户信息
- JWT 中间件保护需要认证的接口

### 文章管理
- 创建文章（需认证）
- 查看文章详情（公开）
- 获取所有文章列表（公开，支持分页）
- 获取用户的文章列表（需认证，支持分页）
- 更新文章（需认证，只能修改自己的文章）
- 删除文章（需认证，只能删除自己的文章）

### 评论系统
- 发布评论（需认证）
- 查看文章的所有评论（公开）

### 系统功能
- 全局异常处理中间件
- 统一的 JSON 响应格式
- YAML 配置文件管理
- MySQL 数据库连接池
- GORM 自动建表和关联关系

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

### 用户认证

#### 用户登录
```
POST /api/user/login?username=admin&password=admin
```

#### 获取用户信息（需认证）
```
GET /api/user/getUserInfo
Authorization: Bearer <token>
```

### 文章管理

#### 查看文章详情（公开）
```
GET /api/post/:id
```

#### 获取所有文章列表（公开）
```
GET /api/post/list?page=1&pageSize=10
```

#### 创建文章（需认证）
```
POST /api/post/create
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "文章标题",
  "content": "文章内容"
}
```

#### 获取我的文章列表（需认证）
```
GET /api/post/my?page=1&pageSize=10
Authorization: Bearer <token>
```

#### 更新文章（需认证）
```
POST /api/post/update
Authorization: Bearer <token>
Content-Type: application/json

{
  "id": 1,
  "title": "更新的标题",
  "content": "更新的内容"
}
```

#### 删除文章（需认证）
```
POST /api/post/delete/:id
Authorization: Bearer <token>
```

### 评论管理

#### 查看文章评论（公开）
```
GET /api/comment/all/:postId
```

#### 发布评论（需认证）
```
POST /api/comment/publish
Authorization: Bearer <token>
Content-Type: application/json

{
  "postId": 1,
  "content": "评论内容"
}
```

## 数据库表结构

### 用户表 (users)
- id: 主键
- username: 用户名
- password: 密码（加密存储）
- email: 邮箱
- created_at, updated_at: 时间戳

### 文章表 (posts)
- id: 主键
- title: 文章标题
- content: 文章内容
- user_id: 作者ID（关联用户表）
- created_at, updated_at: 时间戳

### 评论表 (comments)
- id: 主键
- content: 评论内容
- post_id: 文章ID（关联文章表）
- user_id: 评论者ID（关联用户表）
- created_at, updated_at: 时间戳

## 技术栈

- **语言**: Go
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **配置**: YAML

## 项目架构

采用分层架构设计：
- **Handler层**: 处理HTTP请求和响应
- **Service层**: 业务逻辑处理
- **Repository层**: 数据库操作
- **Entity层**: 数据模型定义
