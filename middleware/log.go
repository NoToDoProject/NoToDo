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
		clientIP := c.ClientIP()     // c.RemoteIP() 区别：ClientIP 会优先获取 X-Forwarded-For
		log.WithFields(log.Fields{
			"Method":     method,
			"Url":        reqUrl,
			"ClientIp":   clientIP,
			"ClientPort": c.GetInt("remote_port"),
		}).Info("<- HTTP In")
		c.Next()
		statusCode := c.Writer.Status()
		fields := log.Fields{
			"status": statusCode,
		}
		spendTime := c.GetString("spend_time")
		if spendTime != "" {
			fields["Last"] = spendTime
		}
		log.WithFields(fields).Info("-> HTTP Out")
	}
}
