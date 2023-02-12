package user

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/database/user"
	"github.com/gin-gonic/gin"
)

// User Controller
type User struct {
}

// InitRouter 初始化路由
func (_ User) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/exist", IsUserExist)
	}
}

// IsUserExist 判断用户是否存在
func IsUserExist(c *gin.Context) {
	userName := c.Query("name")
	nc := response.ContextEx{Context: c}
	nc.Success(user.IsUserExist(userName))
}
