package model

// Config 配置结构体
type Config struct {
	Server struct { // 服务器配置
		Host string // 监听地址
		Port string // 监听端口
	}
	Mongo struct { // MongoDB 配置
		Uri string // 连接地址
	}
	Log struct { // 日志配置
		Level string // 日志级别
	}
}

// ConfigInfo 配置项关联信息结构体
type ConfigInfo struct {
	Path    string // 配置项路径
	Env     string // 配置项关联环境变量
	Default any    // 配置项默认值
}
