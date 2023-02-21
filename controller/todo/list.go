package todo

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/database/todo"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-gonic/gin"
	"github.com/jaevor/go-nanoid"
	"time"
)

var canonidID, _ = nanoid.Standard(21)

// list get todo_lists
func list(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	user := nc.GetUser()

	var page model.ListsPageQuery
	_ = nc.BindQuery(&page)

	list, total, pageCount := todo.GetList(user.Uid, page)
	nc.Success(gin.H{
		"lists":     list,
		"todoCount": total,
		"pageCount": pageCount,
	})
}

// createList create todo list
func createList(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	user := nc.GetUser()

	// get params
	var createList model.CreateListBody
	_ = nc.BindJSON(&createList)

	// check if list exists
	if todo.IsListExist(model.ListExist{
		UserID:    user.Uid,
		ListTitle: createList.Title,
	}) {
		nc.TodoListAlreadyExist()
		return
	}

	// create list
	var todoList model.TodoList
	todoList.Title = createList.Title
	todoList.Description = ""
	todoList.Id = canonidID()
	todoList.UserId = user.Uid
	todoList.CreatedAt = time.Now()
	if !todo.CreateList(todoList) {
		nc.InternalServerError()
		return
	}

	nc.Success(todoList)
}
