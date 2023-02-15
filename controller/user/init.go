package user

import (
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/gin-gonic/gin"
)

// User Controller
type User struct {
}

// InitRouter 初始化路由
func (_ User) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.Use()
	{
		userGroup.POST("/login", middleware.AuthMiddleware.LoginHandler) // 登录
		userGroup.POST("/register", Register)                            // 注册
	}
	userGroup.Use(middleware.MiddleFunc)
	{
		userGroup.GET("/", Info)                                                                                  // 获取自身信息
		userGroup.GET("/exist", middleware.IsAdminMiddleware(), IsUserExist)                                      // 判断用户是否存在
		userGroup.GET("/refresh_token", middleware.IsAdminMiddleware(), middleware.AuthMiddleware.RefreshHandler) // 刷新token
	}
}
