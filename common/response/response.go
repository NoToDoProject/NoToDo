// Package response implements a response package for NoToDo.
package response

import (
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ContextEx extend gin.Context
type ContextEx struct {
	*gin.Context
}

// BindWith override gin bind
func (c *ContextEx) BindWith(obj any, binding binding.Binding) error {
	err := c.ShouldBindWith(obj, binding)
	// when body is empty, err is EOF
	if err != nil {
		c.ParameterError()
		panic(validator.ValidationErrors{})
	}
	return err
}

// Bind general bind
func (c *ContextEx) Bind(obj any) error {
	return c.BindWith(obj, binding.Default(c.Request.Method, c.ContentType()))
}

// BindJSON bind json
func (c *ContextEx) BindJSON(obj any) error {
	return c.BindWith(obj, binding.JSON)
}

// BindQuery bind query
func (c *ContextEx) BindQuery(obj any) error {
	return c.BindWith(obj, binding.Query)
}

// GetUser get user in context
func (c *ContextEx) GetUser() model.User {
	if user, ok := c.Get(common.IdentityKeyInContext); ok {
		if u, ok := user.(model.User); ok {
			return u
		}
	}
	panic("Get user from context failed")
}

// Response general response
func (c *ContextEx) Response(status int, code Code, msg string, data any) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(status, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// Success general success
func (c *ContextEx) Success(data ...any) {
	if len(data) == 0 {
		c.Response(http.StatusOK, Success, "", nil)
		return
	}
	if len(data) == 1 {
		c.Response(http.StatusOK, Success, "", data[0])
		return
	}
	c.Response(http.StatusOK, Success, "", data)
}

// Failure general failure
func (c *ContextEx) Failure(code Code, msg string) {
	c.Response(http.StatusOK, code, msg, nil)
}

// Unauthorized unauthorized
func (c *ContextEx) Unauthorized() {
	c.Response(http.StatusUnauthorized, Error, "Unauthorized", nil)
}

// InternalServerError server error
func (c *ContextEx) InternalServerError() {
	c.Response(http.StatusInternalServerError, InternalServerError, "Internal Server Error", nil)
}

// NotFound not found
func (c *ContextEx) NotFound() {
	c.Response(http.StatusNotFound, NotFound, "Not Found", nil)
}

// NoContent no content
func (c *ContextEx) NoContent() {
	c.Status(http.StatusNoContent)
}

// ParameterError parameter error
func (c *ContextEx) ParameterError() {
	c.Response(http.StatusBadRequest, ParameterError, "Parameter error", nil)
}

// LoginError login error
func (c *ContextEx) LoginError() {
	c.Response(http.StatusOK, Unauthorized, "User not exist or password error", nil)
}

// RegisterDisabled register disabled
func (c *ContextEx) RegisterDisabled() {
	c.Response(http.StatusOK, RegisterDisabled, "Register disabled", nil)
}

// TodoListAlreadyExist todolist already exist
func (c *ContextEx) TodoListAlreadyExist() {
	c.Response(http.StatusOK, TodoListAlreadyExist, "Todo list already exist", nil)
}
