package model

import "github.com/gin-gonic/gin"

// Controller 控制器接口
type Controller interface {
	InitRouter(*gin.Engine) // 添加路由
}
