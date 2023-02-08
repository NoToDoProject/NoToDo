package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var (
	mongoInstance *mongo.Client   = nil
	database      *mongo.Database = nil
	ctx, cancel                   = context.WithTimeout(context.Background(), 10*time.Second)
)

func init() {
	log.Debugf("database init")

	//clientOpts := options.Client().ApplyURI(
	//	"mongodb://localhost:27017/?connect=direct")
	//client, err := mongo.Connect(context.TODO(), clientOpts)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//mongoInstance = client
}

func Connect(uri string) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("database create error: %s", err)
	}
	mongoInstance = client

	if err = mongoInstance.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("database connect error: %s", err)
	}
	log.Info("database connected")
	database = mongoInstance.Database("notodo")
}

func GetDatabase() *mongo.Database {
	return database
}
