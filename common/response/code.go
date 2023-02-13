package response

// Code 响应码枚举类
type Code int

const (
	Success             Code = 1000 // 成功
	Error               Code = 1001 // 一般错误
	NotFound            Code = 1002 // 未找到
	InternalServerError Code = 1003 // 服务器错误
	ParameterError      Code = 1004 // 参数错误
	Unauthorized        Code = 1005 // 未授权
	RegisterDisabled    Code = 1006 // 注册已关闭
)
