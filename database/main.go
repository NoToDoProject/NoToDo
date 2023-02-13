package database

import (
	"context"
	"github.com/NoToDoProject/NoToDo/config"
	"github.com/NoToDoProject/NoToDo/database/user"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	// mongoInstance mongo实例
	mongoInstance *mongo.Client = nil
	// Database      数据库实例
	Database *mongo.Database = nil
)

// Connect 连接数据库
func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("database create error: %s", err)
	}
	mongoInstance = client

	if err = mongoInstance.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("database connect error: %s", err)
	}
	log.Info("database connected")
	Database = mongoInstance.Database("notodo")

	ConfigCollection = Database.Collection("config")
	LoadConfig() // 读取配置

	// 初始化集合
	user.Collection = Database.Collection("user")
}
