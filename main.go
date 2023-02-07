package main

import (
	"fmt"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// 读取配置
func loadConfig() (config model.Config, err error) {
	// 读取 .env 文件并设置环境变量
	v := viper.New()
	v.SetConfigName(".env") // 读取 .env 文件
	v.SetConfigType("env")  // 使用 env 格式
	v.AddConfigPath(".")    // 读取当前目录
	err = v.ReadInConfig()  // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 循环设置环境变量
	for _, key := range v.AllKeys() {
		err := os.Setenv(key, v.GetString(key))
		if err != nil {
			return model.Config{}, err
		}
	}

	// 读取 config.yaml 文件
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err = v.BindEnv("server.host", "SERVER_HOST")
	err = v.BindEnv("server.port", "SERVER_PORT")
	err = v.BindEnv("mongo.host", "MONGO_HOST")
	err = v.BindEnv("mongo.port", "MONGO_PORT")
	err = v.BindEnv("mongo.user", "MONGO_USER")
	err = v.BindEnv("mongo.password", "MONGO_PASSWORD")
	err = v.BindEnv("mongo.db", "MONGO_DB")
	if err != nil {
		return model.Config{}, err
	}

	v.AutomaticEnv()
	err = v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = v.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return
}

func main() {
	config, _ := loadConfig() // 加载配置文件

	// 启动 gin
	engine := gin.Default() // 创建应用

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := engine.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)) // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		return
	}
}
