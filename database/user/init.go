package user

import (
	"context"
	"github.com/NoToDoProject/NoToDo/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection mongo collection
var Collection *mongo.Collection

// IsUserExist check user exist
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

// IsUserExistWithPassword check user exist with password
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

// AddUser add user
func AddUser(user model.User) bool {
	_, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Errorf("add user %+v error: %s", user, err)
		return false
	}
	log.Debugf("add user %+v", user)
	return true
}

// GetUser get user
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

// GetUserByUid get user by uid
func GetUserByUid(uid int) (user model.User, err error) {
	filter := bson.M{"uid": uid}
	r := Collection.FindOne(context.Background(), filter)
	if r.Err() == mongo.ErrNoDocuments {
		log.Errorf("user uid %d not found", uid)
		err = r.Err()
		return
	}
	err = r.Decode(&user)
	return
}

// GetUnusedUid get a new uid
func GetUnusedUid() (uid int) {
	// todo thread safe
	filter := bson.M{} // all
	opt := &options.FindOneOptions{
		Sort: bson.M{"uid": -1}, // desc
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

// IsEmailExist check if email exist
func IsEmailExist(email string) bool {
	filter := bson.M{"email": email}
	count, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("find email %s error: %s", email, err)
		return false
	}
	log.Debugf("email %s count: %d", email, count)
	return count > 0
}

// AllocateUid give an uid to user with -1 uid
func AllocateUid(user model.User) bool {
	if user.Uid != -1 {
		log.Panicf("user %s already has uid: %d", user.Username, user.Uid)
	}

	filter, _ := bson.Marshal(user)
	update := bson.M{"$set": bson.M{"uid": GetUnusedUid()}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Errorf("update user %+v error: %s", user, err)
		return false
	}
	if result.ModifiedCount == 0 {
		log.Errorf("update user %+v failed", user)
		return false
	}

	return true
}
