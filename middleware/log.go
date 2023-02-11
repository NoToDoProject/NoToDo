package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LogMiddleware 日志输出中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		reqUrl := c.Request.URL.Path // c.Request.RequestURI 区别：RequestURI 包含查询参数
		fields := log.Fields{
			"Method": method,
			"Url":    reqUrl,
			"Client": c.Request.RemoteAddr,
		}
		logText := "<- HTTP In"
		if c.IsWebsocket() {
			logText = "<<- WebSocket In"
		}
		log.WithFields(fields).Info(logText)

		c.Next()
		statusCode := c.Writer.Status()

		fields = log.Fields{
			"Status": statusCode,
		}
		spendTime := c.GetString("spend_time")
		if spendTime != "" {
			fields["Last"] = spendTime
		}

		logText = "-> HTTP Out"
		if c.IsWebsocket() {
			logText = "->> Websocket finished"
		}
		log.WithFields(fields).Info(logText)
	}
}
