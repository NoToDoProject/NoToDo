package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
func (c *ContextEx) Success(data ...interface{}) {
	// if len == 0, data = nil
	if len(data) == 0 {
		c.Response(http.StatusOK, Success, "Success", nil)
		return
	}
	// if len == 1, data = data[0]
	if len(data) == 1 {
		c.Response(http.StatusOK, Success, "Success", data[0])
		return
	}
	c.Response(http.StatusOK, Success, "Success", data)
}

// Failure 失败响应
func (c *ContextEx) Failure(code Code, msg string) {
	c.Response(http.StatusOK, code, msg, nil)
}

// Unauthorized 未授权响应
func (c *ContextEx) Unauthorized() {
	c.Response(http.StatusUnauthorized, Error, "Unauthorized", nil)
}

// InternalServerError 服务器错误响应
func (c *ContextEx) InternalServerError() {
	c.Response(http.StatusInternalServerError, InternalServerError, "Internal Server Error", nil)
}

// NotFound 未找到响应
func (c *ContextEx) NotFound() {
	c.Response(http.StatusNotFound, NotFound, "Not Found", nil)
}

// NoContent 无内容响应
func (c *ContextEx) NoContent() {
	c.Status(http.StatusNoContent)
}
