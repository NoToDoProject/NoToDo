package server

import (
	"github.com/gin-gonic/gin"
)

// Server Controller
type Server struct {
}

// InitRouter 初始化路由
func (_ Server) InitRouter(r *gin.Engine) {
	userGroup := r.Group("/server")
	{
		userGroup.GET("/info", Information)
	}
}
