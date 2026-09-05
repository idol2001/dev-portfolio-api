# ============================================================
# 构建阶段
# ============================================================
FROM golang:1.23-alpine AS builder

# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# 安装 git
RUN apk add --no-cache git

# 下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制代码
COPY . .

# 【关键修复】先执行 go mod tidy 整理依赖关系，再编译
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# --- 运行阶段 ---
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata sqlite

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/init-sqlite.sql .

RUN mkdir -p uploads logs data

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/ || exit 1

CMD ["./main"]