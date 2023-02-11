package config

import (
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var Config model.Config

func init() {
	// 读取 .env 文件并设置环境变量
	if common.IsFileExists(".env") {
		v := viper.New()
		v.SetConfigName(".env") // 读取 .env 文件
		v.SetConfigType("env")  // 使用 env 格式
		v.AddConfigPath(".")    // 读取当前目录
		err := v.ReadInConfig() // 读取配置文件
		if err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}
		// 循环设置环境变量
		for _, key := range v.AllKeys() {
			err := os.Setenv(key, v.GetString(key))
			if err != nil {
				log.Fatalf("Error setting environment variable, %s", err)
			}
		}
	}
}

// LoadConfig 读取配置
func LoadConfig() {
	// 读取 config.yaml 文件
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

	v.AutomaticEnv() // 读取环境变量
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Config not exists, using default config and environment variables.")
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	err = v.Unmarshal(&Config) // 反序列化
	if err != nil {
		log.Fatalf("Error unmarshaling config, %s", err)
	}
}
