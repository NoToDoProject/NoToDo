package user

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

// AddUser 添加用户
//func AddUser(user model.User) {
//	one, err := Collection.InsertOne(context.Background(), user)
//	log.Tracef("add user: %s", one)
//	if err != nil {
//		log.Errorf("add user error: %s", err)
//	}
//}

// IsUserExist 判断用户是否存在
func IsUserExist(name string) bool {
	filter := map[string]string{"name": name}
	count, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("find user error: %s", err)
	}
	return count > 0
}
