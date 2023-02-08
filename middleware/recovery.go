package middleware

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Recovery 恢复中间件，捕获500错误
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Recovery")
				nc := response.ContextEx{Context: c}
				nc.Response(500, response.InternalServerError, "Internal Server Error", nil)
			}
		}()
		c.Next()
	}
}
