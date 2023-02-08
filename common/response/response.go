package response

import "github.com/gin-gonic/gin"

// ContextEx 扩展 Context
type ContextEx struct {
	*gin.Context
}

// Response 响应
func (c *ContextEx) Response(status int, code Code, msg string, data interface{}) {
	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Success 成功响应
func (c *ContextEx) Success(data interface{}) {
	c.Response(200, Success, "Success", data)
}

// Failure 失败响应
func (c *ContextEx) Failure(code Code, msg string) {
	c.Response(200, code, msg, nil)
}
