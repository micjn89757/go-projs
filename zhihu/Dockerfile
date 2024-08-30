# 模板
FROM golang:alpine AS builder

# 环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0   \
    GOPROXY=https://goproxy.cn,direct  \
    GOOS=linux \
    GOARCH=amd64 


# 移动到工作目录
WORKDIR /build

# # 复制项目中的go.mod和go.sum文件并下载依赖
# COPY go.mod .
# COPY go.sum .
# RUN go mod download   下载依赖到本地cache，默认为$GOPATH/pkg/mod目录


# 将代码复制到容器中
COPY . .
RUN go mod tidy


# 将代码编译成二进制可执行文件
RUN go build -o main .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 创建一个小镜像
FROM ubuntu:latest 
# 拷贝配置文件目录
COPY ./config /config

# 从builder镜像中把 /build/main  拷贝到当前目录
COPY --from=builder /build/dist/main /

# 声明服务端口
EXPOSE 8888

# 容器启动时需要运行的命令和参数
ENTRYPOINT [ "/main", "config/xx.yaml" ]