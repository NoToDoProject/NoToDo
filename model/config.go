package model

// Config configuration
type Config struct {
	Server struct {
		Host string // listen host
		Port string // listen port
	}
	Mongo struct { // MongoDB
		URI string // connection uri
	}
	Log struct {
		Level string
	}
}

// ConfigInfo configuration related information
type ConfigInfo struct {
	Path    string // configuration path
	Env     string // related environment variable
	Default any    // default value
}
