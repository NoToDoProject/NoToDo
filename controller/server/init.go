// Package server API to control the server.
package server

import (
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/gin-gonic/gin"
)

// Server Controller
type Server struct {
}

// InitRouter add server router
func (_ Server) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/server")
	userGroup.Use()
	{
		userGroup.GET("/info", Information)
		userGroup.GET("/can_register", CanRegister)
	}
	userGroup.Use(middleware.MiddleFunc)
	{
		userGroup.POST("/refresh_config", middleware.IsAdminMiddleware(), RefreshConfig)
	}
}
