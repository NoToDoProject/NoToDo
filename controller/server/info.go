package server

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	StartTimestamp = time.Now().Unix() // 服务器启动时间戳
)

// Information 服务器信息
func Information(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	nc.Success(gin.H{
		"appName":     "NoToDo",                           // 应用名称
		"currentTime": time.Now().Unix(),                  // 当前时间戳
		"startTime":   StartTimestamp,                     // 服务器启动时间戳
		"uptime":      time.Now().Unix() - StartTimestamp, // 服务器运行时间
	})
}
