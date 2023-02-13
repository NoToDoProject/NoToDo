package middleware

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// Recovery 恢复中间件，捕获500错误
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Error from ParseForm or Bind
				if _, ok := err.(validator.ValidationErrors); ok {
					return
				}
				log.WithFields(log.Fields{
					"error": err,
				}).Error("Recovery")
				nc := response.ContextEx{Context: c}
				nc.InternalServerError()
			}
		}()
		c.Next()
	}
}
