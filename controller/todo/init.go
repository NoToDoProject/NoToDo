package todo

import (
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/gin-gonic/gin"
)

// Todos Controller
type Todos struct {
}

// InitRouter add todo_router
func (_ Todos) InitRouter(r *gin.Engine) {
	todoGroup := r.Group("/todo")
	todoGroup.Use(middleware.MiddleFunc)
	{
		listGroup := todoGroup.Group("/list")
		{
			listGroup.GET("/", list)
			listGroup.POST("/", createList)
		}
	}
}
