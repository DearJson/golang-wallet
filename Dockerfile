# 多阶段构建，减小镜像体积
# 第一阶段：构建阶段
FROM golang:1.19-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的构建工具
RUN apk add --no-cache git ca-certificates tzdata

# 复制 go mod 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用 (静态链接，减少依赖)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o wallet-app \
    main.go

# 第二阶段：运行阶段 (极简安全镜像)
FROM alpine:3.18

# 安装必要的运行时依赖并创建用户
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -g 1000 appgroup \
    && adduser -D -s /bin/sh -u 1000 -G appgroup appuser

# 设置时区
ENV TZ=Asia/Shanghai

# 创建应用目录
RUN mkdir -p /app/data /app/logs \
    && chown -R appuser:appgroup /app

# 复制编译好的应用
COPY --from=builder /app/wallet-app /app/wallet-app

# 复制配置文件和必要的资源
COPY --from=builder /app/config /app/config
COPY --from=builder /app/public /app/public
COPY --from=builder /app/template /app/template

# 复制Docker专用配置文件
COPY --from=builder /app/config/docker.toml /app/config/config.toml

# 安装envsubst用于环境变量替换
RUN apk add --no-cache gettext

# 设置文件权限
RUN chmod +x /app/wallet-app \
    && chown -R appuser:appgroup /app

# 设置工作目录
WORKDIR /app

# 切换到非特权用户
USER appuser

# 暴露端口
EXPOSE 9090

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=30s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:9090/ || exit 1

# 启动应用
ENTRYPOINT ["/app/wallet-app"]
