package model

import "time"

// ATodo one todo_item
type ATodo struct {
	Id          string    `json:"id" bson:"id"` // unique, auto create
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	UserId      int       `json:"user_id" bson:"user_id"` // belong to which user
	IsDone      bool      `json:"is_done" bson:"is_done"` // if done
	ListId      int       `json:"list_id" bson:"list_id"` // belong to which list
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

// TodoList one todo_list
type TodoList struct {
	Id          string    `bson:"id"` // unique, auto create
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	UserId      int       `bson:"user_id"`
	CreatedAt   time.Time `bson:"created_at"`
}

// CreateListBody body for create list
type CreateListBody struct {
	Title string `json:"title" binding:"required"`
}

// ListExist check if list exist query
type ListExist struct {
	UserID    int    `bson:"user_id"`
	ListTitle string `bson:"title"`
}

// ListsPageQuery query for lists page
type ListsPageQuery struct {
	PageSize int `form:"size" binding:"required"`
	PageNum  int `form:"page" binding:"required"`
}

// SetListTitleQuery query for set list title
type SetListTitleQuery struct {
	ListID   int    `json:"list_id" binding:"required"`
	NewTitle string `json:"new_title" binding:"required"`
}
