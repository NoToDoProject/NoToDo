package middleware

import (
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
)

// Recovery catch panic
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Error from ParseForm or Bind
				if _, ok := err.(validator.ValidationErrors); ok {
					return
				}
				log.Errorf("Recovery from panic: %v", err)
				debug.PrintStack()
				nc := response.ContextEx{Context: c}
				nc.InternalServerError()
				nc.Abort()
			}
		}()
		c.Next()
	}
}
