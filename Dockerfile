# 使用官方的Golang镜像作为基础镜像
FROM golang:1.17-alpine as builder

# 设置工作目录
WORKDIR /app

# 将项目中的go.mod和go.sum文件复制到工作目录
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目源代码复制到工作目录
COPY . .

# 编译Golang程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 使用官方的Nginx镜像作为基础镜像
FROM nginx:1.21-alpine

# 将编译好的Golang程序复制到Nginx镜像中
COPY --from=builder /app/main /app/main

# 将Nginx配置文件复制到镜像中
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 启动Golang程序和Nginx
CMD ["/bin/sh", "-c", "/app/main & nginx -g 'daemon off;'"]
