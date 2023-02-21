# NoToDo 后端程序

[![Go Report Card](https://goreportcard.com/badge/github.com/NoToDoProject/NoToDo)](https://goreportcard.com/report/github.com/NoToDoProject/NoToDo)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/4912f94720f24de6b5924062f89160bf)](https://www.codacy.com/gh/NoToDoProject/NoToDo/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=NoToDoProject/NoToDo&amp;utm_campaign=Badge_Grade)

不要做！！！，NoToDo，发音 /nɒt tʊ duː/，即 Not To Do，
是一个基于Web的待办事项管理应用，支持多用户，~~多设备同步(WIP)~~。

简体中文 | [English](./README.en.md)

在线示例 : [NoToDo.akagiyui.com(WIP)](https://notodo.akagiyui.com)

## 快速开始

你需要先修改配置文件 `config.yaml` 或者环境变量，然后运行程序。

```bash
export MONGO_URL=<YOUR_MONGO_URL>
```

你也可以使用环境变量文件 `.env`。

```dotenv
MONGO_URL=<YOUR_MONGO_URL>
```

支持的配置项与环境变量：

| 环境变量        | 配置项         | 默认值                       | 说明           |
|-------------|-------------|---------------------------|--------------|
| SERVER_HOST | server.host | 0.0.0.0                   | 监听地址         |
| SERVER_PORT | server.port | 8888                      | 监听端口         |
| MONGO_URI   | mongo.uri   | mongodb://localhost:27017 | MongoDB 连接地址 |
| LOG_LEVEL   | log.level   | info                      | 日志等级         |

如果你处于中国大陆地区，你还可能希望配置代理：

```bash
export GOPROXY=https://goproxy.cn
```

### 使用 Docker

你需要在 `docker-compose.yml` 中修改环境变量。

```bash
docker-compose up -d
```

### 从代码开始

```bash
go run main.go
```

## 开发

### 依赖

  - [gin](https://github.com/gin-gonic/gin) - Web框架
  - [viper](https://github.com/spf13/viper) - 配置管理
  - [logrus](https://github.com/sirupsen/logrus) - 日志处理
  - [gin-cors](https://github.com/gin-contrib/cors) - 跨域处理
  - [websocket](https://github.com/gorilla/websocket) - WebSocket
  - [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) - mongodb数据库驱动
  - [crypto](https://pkg.go.dev/golang.org/x/crypto) - 密码加密
  - [gin-jwt](https://github.com/appleboy/gin-jwt) - JWT认证
  - [nanoid](https://github.com/jaevor/go-nanoid) - NanoID 生成

### 运行测试

```bash
go test ./...
```
