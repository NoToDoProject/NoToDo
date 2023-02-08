package model

// Config 配置结构体
type Config struct {
	Server struct {
		Host string
		Port string
	}
	Mongo struct {
		Host     string
		Port     string
		Db       string
		User     string
		Password string
	}
	Log struct {
		Level string
	}
}

// ConfigInfo 配置项关联信息结构体
type ConfigInfo struct {
	Path    string
	Env     string
	Default any
}
