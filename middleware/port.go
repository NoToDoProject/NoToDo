package middleware

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"strings"
)

// GetRemotePortMiddleware 获取远程端口中间件
func GetRemotePortMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteAddr := c.Request.RemoteAddr
		_, port, _ := net.SplitHostPort(strings.TrimSpace(remoteAddr))
		portInt, _ := strconv.Atoi(port)
		c.Set("remote_port", portInt)
		c.Next()
	}
}
