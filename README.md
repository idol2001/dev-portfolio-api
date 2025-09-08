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
│   ├── handlers/
│   └── routes/
├── configs/
│   └── config.se.yml
├── initialize/
├── internal/
│   ├── dto/
│   └── services/
├── logs/
├── middleware/
├── models/
├── pkg/
│   └── utils/
├── main.go     # Application entry point
├── go.mod
├── go.sum
├── Dockerfile (optional)
└── README.md
```

## 快速开始

### 克隆项目

```bash
git clone https://github.com/idol2001/dev-portfolio-api.git
cd dev-portfolio-api
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


## 贡献

欢迎贡献代码！请提交 Pull Request 或报告问题。

## 许可证

本项目使用 [MIT 许可证](LICENSE)。  
