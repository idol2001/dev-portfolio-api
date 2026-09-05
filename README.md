# Go Gin RESTful API Template

一个基于 Go + Gin + GORM 的 RESTful API 脚手架项目，开箱即用。

## 特性

- **双数据库支持**：SQLite（默认，零配置）/ MySQL，通过配置文件一键切换
- **JWT 认证**：完整的登录、Token 签发与验证流程
- **RSA 密码加密**：前端公钥加密传输，后端私钥解密，防止密码明文传输
- **文件上传**：支持阿里云 OSS 和本地存储双模式
- **结构化日志**：基于 `slog` 的 JSON 日志输出，支持日志轮转
- **热加载配置**：基于 Viper，配置文件修改后自动重载
- **软删除**：基于 GORM 的软删除支持
- **Docker 部署**：多阶段构建，生产就绪的 Dockerfile
- **统一响应格式**：标准化的 API 响应结构

## 技术栈

| 组件 | 技术 |
|------|------|
| 语言 | Go 1.23 |
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io) (SQLite / MySQL) |
| 认证 | JWT ([golang-jwt](https://github.com/golang-jwt/jwt)) |
| 加密 | RSA (Go 标准库) + bcrypt |
| 配置 | [Viper](https://github.com/spf13/viper) + .env |
| 日志 | [slog](https://pkg.go.dev/log/slog) + [lumberjack](https://github.com/natefinch/lumberjack) |
| 跨域 | [gin-contrib/cors](https://github.com/gin-contrib/cors) |
| 压缩 | [gin-contrib/gzip](https://github.com/gin-contrib/gzip) |

## 项目结构

```
.
├── api/
│   ├── handlers/          # 请求处理器（Controller 层）
│   └── routes/            # 路由定义
├── configs/               # 配置文件（YAML）
│   ├── config.se.yml      # 开发环境
│   └── config.prd.yml     # 生产环境
├── initialize/            # 初始化逻辑（配置、日志、数据库、路由）
├── internal/
│   ├── dto/               # 数据传输对象
│   └── services/          # 业务逻辑层（Service 层）
├── middleware/             # 中间件（JWT、异常处理、访问日志）
├── models/                # 数据模型（Model 层）
├── pkg/
│   ├── crypto/            # RSA 加密工具
│   └── global/            # 全局变量与配置结构
├── main.go                # 入口文件
├── Dockerfile             # 多阶段 Docker 构建
├── go.mod
└── go.sum
```

## 快速开始

### 前置要求

- Go 1.23+
- SQLite（默认，无需安装）或 MySQL

### 1. 克隆项目

```bash
git clone https://github.com/idol2001/dev-portfolio-api.git
cd dev-portfolio-api
```

### 2. 配置环境变量

```bash
cp .env.example .env
```

默认使用 SQLite，无需修改。如需 MySQL，编辑 `.env` 设置数据库连接信息。

### 3. 安装依赖并运行

```bash
go mod tidy
go run main.go
```

服务默认运行在 `http://localhost:8080`。

### 4. Docker 部署

```bash
docker build -t go-api-template .
docker run -p 8080:8080 go-api-template
```

## API 概览

所有 API 前缀：`/dev-portfolio/v1`

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/profile/info` | 获取个人信息 |
| GET | `/profile/skills` | 获取技能列表 |
| GET | `/profile/socials` | 获取社交链接 |
| GET | `/projects` | 获取项目列表 |
| GET | `/blogs` | 获取博客列表（分页） |
| GET | `/blogs/slug/:slug` | 根据 slug 获取文章 |
| POST | `/auth/login` | 用户登录 |
| GET | `/auth/public-key` | 获取 RSA 公钥 |

### 需要认证的接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/user/me` | 获取当前用户信息 |
| PUT | `/profile/info` | 更新个人信息 |
| POST/PUT/DELETE | `/profile/socials/*` | 社交链接 CRUD |
| POST/PUT/DELETE | `/projects/*` | 项目 CRUD |
| POST/PUT/DELETE | `/blogs/*` | 博客 CRUD |
| POST | `/upload/:type` | 文件上传 |

## 配置说明

### 数据库切换

编辑 `configs/config.prd.yml`：

```yaml
system:
  db-type: sqlite   # 改为 mysql 使用 MySQL

mysql:
  username: "root"
  password: "your-password"
  database: "my_database"
  host: "127.0.0.1"
  port: 3306
```

### 启用阿里云 OSS

编辑 `.env`：

```env
ALIYUNOSS_ENABLE=true
ALIYUNOSS_ENDPOINT=oss-cn-hangzhou.aliyuncs.com
ALIYUNOSS_ACCESS_KEY_ID=your-key-id
ALIYUNOSS_ACCESS_KEY_SECRET=your-key-secret
ALIYUNOSS_BUCKET_NAME=your-bucket
```

## 响应格式

```json
{
  "request_id": "uuid",
  "code": 0,
  "data": {},
  "msg": "操作成功",
  "total": 0
}
```

## 许可证

MIT
