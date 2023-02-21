package server

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/database"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	startTimestamp = time.Now().Unix()
	config         = &database.Config
)

// Information 服务器信息
func Information(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	nc.Success(gin.H{
		"appName":     "NoToDo",                           // app name
		"currentTime": time.Now().Unix(),                  // current time
		"startTime":   startTimestamp,                     // server start time
		"uptime":      time.Now().Unix() - startTimestamp, // server uptime
	})
}

// CanRegister check register enable
func CanRegister(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	nc.Success(config.CanRegister)
}

// RefreshConfig reload config from database
func RefreshConfig(c *gin.Context) {
	nc := response.ContextEx{Context: c}
	database.LoadConfig()
	nc.Success()
}
