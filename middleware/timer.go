// https://github.com/dtest11/reponse/blob/1f462b27a8930673eb85868e1bdc402a488693db/reponse_time.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// XResponseTimeWriter Patch ResponseWriter
type XResponseTimeWriter struct {
	gin.ResponseWriter              // 原始的 ResponseWriter
	context            *gin.Context // 上下文
	startTime          time.Time    // 请求开始时间
}

// TimerMiddleware 请求处理时间记录中间件
func TimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newWriter := &XResponseTimeWriter{ResponseWriter: c.Writer, context: c, startTime: time.Now()}
		c.Writer = newWriter
		c.Next()
	}
}

// WriteHeader 重写 WriteHeader 方法
func (w *XResponseTimeWriter) WriteHeader(statusCode int) {
	duration := time.Since(w.startTime)
	w.context.Set("spend_time", duration.String())
	w.Header().Set("X-Response-Time", duration.String())
	w.ResponseWriter.WriteHeader(statusCode)
}
