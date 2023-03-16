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
func (Server) InitRouter(r *gin.Engine) {
	serverGroup := r.Group("/server")
	serverGroup.Use()
	{
		serverGroup.GET("/info", Information)
		serverGroup.GET("/can_register", CanRegister)
		serverGroup.GET("/ping", Ping)
	}
	serverGroup.Use(middleware.MiddleFunc)
	{
		serverGroup.POST("/refresh_config", middleware.IsAdminMiddleware(), RefreshConfig)
	}
}
