package response

// Code 响应码枚举类
type Code int

const (
	// Success 成功
	Success Code = 1000
	// Error 一般错误
	Error Code = 1001
	// NotFound 未找到
	NotFound Code = 1002
	// InternalServerError 服务器错误
	InternalServerError Code = 1003
	// ParameterError 参数错误
	ParameterError Code = 1004
	// Unauthorized 未授权
	Unauthorized Code = 1005
)
