# 构建阶段
FROM golang:1.24-alpine AS builder

# 替换 apk 源为阿里云
RUN sed -i 's|https://dl-cdn.alpinelinux.org|https://mirrors.aliyun.com|g' /etc/apk/repositories

# 设置工作目录
WORKDIR /app

# 安装git（用于获取依赖）
RUN apk add --no-cache git

# 复制go.mod和go.sum文件
COPY go.mod go.sum ./

# 使用国内Go代理加速
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 安装 swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 生成 Swagger 文档
RUN swag init -g src/main.go -o src/docs

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main src/main.go

# 运行阶段
FROM alpine:latest

# 替换 apk 源为阿里云
RUN sed -i 's|https://dl-cdn.alpinelinux.org|https://mirrors.aliyun.com|g' /etc/apk/repositories

# 安装ca证书（用于HTTPS请求）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /data/

# 从构建阶段复制二进制文件
COPY --from=builder /app/main ./server

# 暴露端口
EXPOSE 8000

# 运行应用
CMD ["./server"]