# dev-portfolio-api

dev-portfolio 是一个用 Go 语言开发的开发者个人网站和博客的后端 API。

## 功能

- 用户认证与授权
- 博客文章的创建、读取、更新和删除 (CRUD)
- 分类与标签管理
- 评论系统
- 文件上传 (如头像和文章图片)
- API 文档支持

## 技术栈

- **语言**: Go
- **框架**: Gin
- **数据库**: MySQL
- **ORM**: GORM
- **认证**: JWT
- **其他**: Swagger, Docker

## 项目结构
```
dev-portfolio-api/
├── api/
│   ├── handlers/
│   └── routes/
├── configs/
│   └── config.yaml
├── internal/
│   ├── services/
│   ├── dto/
│   └── dao/
├── models/
│   └── album.go
├── pkg/
│   └── utils/
├── go.mod
├── go.sum
├── Dockerfile (optional)
└── main.go  # Application entry point

```

## 快速开始

### 克隆项目

```bash
git clone https://github.com/idol2001/dev-portfolio-api.git
cd dev-portfolio-api
```

### 配置环境变量

创建 `.env` 文件并配置以下内容：

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=dev_portfolio
JWT_SECRET=your_secret
```

### 安装依赖

```bash
go mod tidy
```

### 运行项目

```bash
go run main.go
```

### 运行测试

```bash
go test ./...
```

## API 文档

启动项目后，访问 [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) 查看 Swagger API 文档。

## 贡献

欢迎贡献代码！请提交 Pull Request 或报告问题。

## 许可证

本项目使用 [MIT 许可证](LICENSE)。  