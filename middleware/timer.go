// https://github.com/dtest11/reponse/blob/1f462b27a8930673eb85868e1bdc402a488693db/reponse_time.go

package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// XResponseTimeWriter Patch ResponseWriter
type XResponseTimeWriter struct {
	gin.ResponseWriter              // origin ResponseWriter
	context            *gin.Context // origin context
	startTime          time.Time    // response start time
}

// TimerMiddleware record response time
func TimerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		newWriter := &XResponseTimeWriter{ResponseWriter: c.Writer, context: c, startTime: time.Now()}
		c.Writer = newWriter
		c.Next()
		// WebSocket will not call WriteHeader, set spend time here
		c.Set("spend_time", time.Since(newWriter.startTime).String())
	}
}

// WriteHeader override WriteHeader
func (w *XResponseTimeWriter) WriteHeader(statusCode int) {
	duration := time.Since(w.startTime)
	if _, exist := w.context.Get("spend_time"); !exist {
		w.context.Set("spend_time", duration.String())
	}
	w.Header().Set("X-Response-Time", duration.String())
	w.ResponseWriter.WriteHeader(statusCode)
}
