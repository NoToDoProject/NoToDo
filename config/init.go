// Package config loads the configuration file and environment variables.
package config

import (
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

// Config global config
var Config model.Config

func init() {
	// load .env file
	if common.IsFileExists(".env") {
		v := viper.New()
		v.SetConfigName(".env")
		v.SetConfigType("env")
		v.AddConfigPath(".")
		err := v.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		// set environment variables
		for _, key := range v.AllKeys() {
			err := os.Setenv(key, v.GetString(key))
			if err != nil {
				log.Fatalf("Error setting environment variable, %s", err)
			}
		}
	}
}

// LoadConfig load config from config.yaml and environment variables
func LoadConfig() {
	// load config.yaml
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	configInfos := []model.ConfigInfo{
		{Path: "server.host", Env: "SERVER_HOST", Default: "0.0.0.0"},
		{Path: "server.port", Env: "SERVER_PORT", Default: "8888"},

		{Path: "mongo.uri", Env: "MONGO_URI", Default: "mongodb://localhost:27017"},

		{Path: "log.level", Env: "LOG_LEVEL", Default: "info"},
	}

	for _, config := range configInfos {
		v.SetDefault(config.Path, config.Default)
		err := v.BindEnv(config.Path, config.Env)
		if err != nil {
			log.Fatalf("Error binding environment variable, %s", err)
		}
	}

	v.AutomaticEnv() // get environment variables
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Config not exists, using default config and environment variables.")
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	err = v.Unmarshal(&Config) // unmarshal config to struct
	if err != nil {
		log.Fatalf("Error unmarshaling config, %s", err)
	}
}
