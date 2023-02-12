# NoToDo Backend

[![Go Report Card](https://goreportcard.com/badge/github.com/NoToDoProject/NoToDo)](https://goreportcard.com/report/github.com/NoToDoProject/NoToDo)

不要做！！！

NoToDo (pronounced /nɒt tʊ duː/, like Not To Do)

Demo: [NoToDo(WIP)](https://notodo.akagiyui.com)

## Quick Start

### Docker

```bash
export MONGO_URL=<YOUR_MONGO_URL>
docker-compose up -d
```

### Source

```bash
export MONGO_URL=<YOUR_MONGO_URL>
go run main.go
```


## Packages

- [gin](https://github.com/gin-gonic/gin) - web框架
- [viper](https://github.com/spf13/viper) - 配置文件
- [logrus](https://github.com/sirupsen/logrus) - 日志处理
- [gin-cors](https://github.com/gin-contrib/cors) - 跨域处理
- [websocket](https://github.com/gorilla/websocket) - websocket
- [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) - mongodb数据库驱动
