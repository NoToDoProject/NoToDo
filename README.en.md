# NoToDo Backend

[![Go Report Card](https://goreportcard.com/badge/github.com/NoToDoProject/NoToDo)](https://goreportcard.com/report/github.com/NoToDoProject/NoToDo)

NoToDo!!!，pronounced /nɒt tʊ duː/，Not To Do，
is a web-based todo list application that supports multiple users, ~~multi-device synchronization (WIP)~~.

[简体中文](./README.md) | English

Online Demo : [NoToDo.akagiyui.com(WIP)](https://notodo.akagiyui.com)

## Quick Start

You need to modify the configuration file `config.yaml` or environment variables before running.

```bash
export MONGO_URL=<YOUR_MONGO_URL>
```

You can also save the environment variables in a `.env` file.

```dotenv
MONGO_URL=<YOUR_MONGO_URL>
```

Supported environment variables and configuration items are as follows:

| Environment Variables | Configuration | Default                   | Commentary     |
|-----------------------|---------------|---------------------------|----------------|
| SERVER_HOST           | server.host   | 0.0.0.0                   | Listening Host |
| SERVER_PORT           | server.port   | 8888                      | Listening Port |
| MONGO_URI             | mongo.uri     | mongodb://localhost:27017 | MongoDB URI    |
| LOG_LEVEL             | log.level     | info                      | Log Level      |


### Docker

You must modify `docker-compose.yml` for configuration.

```bash
docker-compose up -d
```

### Source

```bash
go run main.go
```

## Development

### Dependencies

- [gin](https://github.com/gin-gonic/gin) - Web framework
- [viper](https://github.com/spf13/viper) - Configuration management
- [logrus](https://github.com/sirupsen/logrus) - Log
- [gin-cors](https://github.com/gin-contrib/cors) - CORS
- [websocket](https://github.com/gorilla/websocket) - WebSocket
- [mongo-driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) - MongoDB Driver
- [crypto](https://pkg.go.dev/golang.org/x/crypto) - Password encryption

### Test

```bash
go test ./...
```
