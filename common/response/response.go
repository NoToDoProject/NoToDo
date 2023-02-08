package response

import "github.com/gin-gonic/gin"

// ContextEx 扩展 Context
type ContextEx struct {
	*gin.Context
}

// Success 成功响应
func (c *ContextEx) Success(data interface{}) {
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  "",
		"data": data,
	})
}

// Failure 失败响应
func (c *ContextEx) Failure(code Code, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
