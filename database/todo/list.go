package todo

import (
	"context"
	"github.com/NoToDoProject/NoToDo/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ListCollection mongo collection
var ListCollection *mongo.Collection

// CreateList create list
func CreateList(list model.TodoList) bool {
	_, err := ListCollection.InsertOne(context.Background(), list)
	if err != nil {
		log.Errorf("create list %+v error: %s", list, err)
		return false
	}
	log.Debugf("create list %+v", list)
	return true
}

// IsListExist check list exist
func IsListExist(list model.ListExist) bool {
	count, err := ListCollection.CountDocuments(context.Background(), list)
	if err != nil {
		log.Errorf("list exist error: %s", err)
		return false
	}
	return count > 0
}

// GetList get lists by uid
func GetList(uid int, page model.ListsPageQuery) (lists []model.TodoList, total int64, pageCount int64) {
	var size = int64(page.PageSize)
	var skip = int64((page.PageNum - 1) * page.PageSize)
	var filter = bson.M{"user_id": uid}
	var opts = options.FindOptions{
		Skip:  &skip,
		Limit: &size,
	}
	cur, err := ListCollection.Find(context.Background(), filter, &opts)
	if err != nil {
		log.Errorf("get list error: %s", err)
		return
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			log.Errorf("close cursor error: %s", err)
		}
	}(cur, context.Background())
	for cur.Next(context.Background()) {
		var list model.TodoList
		err := cur.Decode(&list)
		if err != nil {
			log.Errorf("decode list error: %s", err)
			continue
		}
		lists = append(lists, list)
	}
	total, err = ListCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Errorf("get list count error: %s", err)
		return
	}
	pageCount = total / size
	if total%size != 0 {
		pageCount++
	}
	return
}
