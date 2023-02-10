# 构建二进制文件
FROM golang:1.20 as builder
COPY . /work
WORKDIR /work
ARG GOPROXY
ARG GOOS=linux
ARG GOARCH=amd64
RUN export GOOS=${GOOS} && \
    export GOARCH=${GOARCH} && \
    export CGO_ENABLED=0 && \
    export GOPROXY=${GOPROXY} && \
    go build -o build/notodo -x -v -trimpath -ldflags="-s -w" .


# 构建镜像
FROM scratch
COPY --from=builder /work/build/notodo /usr/local/bin/notodo
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/notodo"]