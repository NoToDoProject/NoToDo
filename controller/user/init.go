package user

import "github.com/gin-gonic/gin"

// User Controller
type User struct {
}

// InitRouter 初始化路由
func (_ User) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/exist", IsUserExist)  // 判断用户是否存在
		userGroup.POST("/login", Login)       // 登录
		userGroup.POST("/register", Register) // 注册
	}
}
