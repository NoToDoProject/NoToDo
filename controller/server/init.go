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
		userGroup.GET("/info", Information)              // 获取服务器信息
		userGroup.GET("/can-register", CanRegister)      // 获取是否允许注册
		userGroup.POST("/refresh-config", RefreshConfig) // 刷新配置
	}
}
