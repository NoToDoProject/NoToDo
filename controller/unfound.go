package controller

import "github.com/gin-gonic/gin"
import "github.com/NoToDoProject/NoToDo/common/response"

// NotFoundRoute 404 route
func NotFoundRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		nc := response.ContextEx{Context: c}
		nc.NotFound()
	}
}
