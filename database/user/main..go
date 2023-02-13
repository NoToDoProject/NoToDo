package user

import (
	"context"
	"github.com/NoToDoProject/NoToDo/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection mongo集合
var Collection *mongo.Collection

// IsUserExist 判断用户是否存在
func IsUserExist(user model.IsUserExist) bool {
	filter, _ := bson.Marshal(user)
	count, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("find user %+v error: %s", user, err)
		return false
	}
	log.Debugf("user %+v count: %d", user, count)
	return count > 0
}

// IsUserExistWithPassword 判断密码是否正确
func IsUserExistWithPassword(user model.UserWithPassword) bool {
	filter, _ := bson.Marshal(user)
	count, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("find user %+v error: %s", user, err)
		return false
	}
	log.Debugf("user %+v count: %d", user, count)
	return count > 0
}

// AddUser 添加用户
func AddUser(user model.User) bool {
	_, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Errorf("add user %+v error: %s", user, err)
		return false
	}
	log.Debugf("add user %+v", user)
	return true
}

// GetUser 获取用户
func GetUser(_user model.IsUserExist) (user model.User, err error) {
	filter, err := bson.Marshal(_user)
	if err != nil {
		log.Errorf("marshal user %+v error: %s", _user, err)
		return
	}
	r := Collection.FindOne(context.Background(), filter)
	if r.Err() == mongo.ErrNoDocuments {
		log.Errorf("user %+v not found", _user)
		err = r.Err()
		return
	}
	err = r.Decode(&user)
	return
}

// GetUnusedUid 获取未使用的uid，从1开始
func GetUnusedUid() (uid int) {
	filter := bson.M{} // 查询所有
	opt := &options.FindOneOptions{
		Sort: bson.M{"uid": -1}, // 降序
	}
	r := Collection.FindOne(context.Background(), filter, opt)
	if r.Err() == mongo.ErrNoDocuments {
		// not found
		return 1
	}
	var user model.User
	err := r.Decode(&user)
	if err != nil {
		log.Panicf("decode user error: %s", err)
	}
	return user.Uid + 1
}
