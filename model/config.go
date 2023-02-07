package model

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Mongo struct {
		Host string
		Port string
		User string
	}
}
