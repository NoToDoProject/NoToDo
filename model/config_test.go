package model

import "testing"

func TestConfig(t *testing.T) {
	config := Config{}
	config.Server.Host = "0.0.0.0"
	config.Server.Port = "8888"
	config.Mongo.Uri = "mongodb://localhost:27017"
	config.Log.Level = "info"
	t.Log(config)
}
