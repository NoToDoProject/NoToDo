package user

import (
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/gin-gonic/gin"
)

// User Controller
type User struct {
}

// InitRouter add user router
func (_ User) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.Use()
	{
		userGroup.POST("/login", middleware.AuthMiddleware.LoginHandler)
		userGroup.POST("/register", register)
	}
	userGroup.Use(middleware.MiddleFunc)
	{
		userGroup.GET("/info", info)
		userGroup.GET("/exist", middleware.IsAdminMiddleware(), isUserExist)
		userGroup.GET("/refresh_token", middleware.IsAdminMiddleware(), middleware.AuthMiddleware.RefreshHandler)
	}
}
