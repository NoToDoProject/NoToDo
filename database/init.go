// Package database MongoDB operations
package database

import (
	"context"
	"github.com/NoToDoProject/NoToDo/config"
	"github.com/NoToDoProject/NoToDo/database/todo"
	"github.com/NoToDoProject/NoToDo/database/user"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	// mongoInstance mongoDB client
	mongoInstance *mongo.Client = nil
	// database      database instance
	database *mongo.Database = nil
)

// Connect connect to database
func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Config.Mongo.URI))
	if err != nil {
		log.Fatalf("database create error: %s", err)
	}
	mongoInstance = client

	if err = mongoInstance.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("database connect error: %s", err)
	}
	log.Info("database connected")
	database = mongoInstance.Database("notodo")

	ConfigCollection = database.Collection("config")
	LoadConfig() // get config from database

	// init collections
	user.Collection = database.Collection("user")
	todo.ListCollection = database.Collection("todo_list")
}
