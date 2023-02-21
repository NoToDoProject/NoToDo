// Package response implements a response package for NoToDo.
package response

// Code 响应码枚举类
type Code int

const (
	Success              Code = 1000 // Success success
	Error                Code = 1001 // Error general error
	NotFound             Code = 1002 // NotFound not found
	InternalServerError  Code = 1003 // InternalServerError internal server error
	ParameterError       Code = 1004 // ParameterError parameter error
	Unauthorized         Code = 1005 // Unauthorized unauthorized
	RegisterDisabled     Code = 1006 // RegisterDisabled register disabled
	TodoListAlreadyExist Code = 1007 // TodoListAlreadyExist todolist already exist
)
