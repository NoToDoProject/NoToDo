// NoToDo 后端服务
package main

import (
	"fmt"
	"github.com/NoToDoProject/NoToDo/common"
	"github.com/NoToDoProject/NoToDo/common/response"
	"github.com/NoToDoProject/NoToDo/controller"
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

// loadConfig 读取配置
func loadConfig() (config model.Config, err error) {
	// 读取 .env 文件并设置环境变量
	if common.IsFileExists(".env") {
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
	}

	// 读取 config.yaml 文件
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	configInfos := []model.ConfigInfo{
		{Path: "server.host", Env: "SERVER_HOST", Default: "0.0.0.0"},
		{Path: "server.port", Env: "SERVER_PORT", Default: "8080"},

		{Path: "mongo.host", Env: "MONGO_HOST", Default: "127.0.0.1"},
		{Path: "mongo.port", Env: "MONGO_PORT", Default: "27017"},
		{Path: "mongo.db", Env: "MONGO_DB", Default: "notodo"},
		{Path: "mongo.user", Env: "MONGO_USER", Default: "notodo"},
		{Path: "mongo.password", Env: "MONGO_PASSWORD", Default: "notodo"},

		{Path: "log.level", Env: "LOG_LEVEL", Default: "info"},
	}

	for _, config := range configInfos {
		v.SetDefault(config.Path, config.Default)
		err = v.BindEnv(config.Path, config.Env)
		if err != nil {
			return model.Config{}, err
		}
	}

	v.AutomaticEnv() // 读取环境变量
	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("Config not exists, using default config and environment variables.")
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	err = v.Unmarshal(&config) // 反序列化
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return
}

// init 初始化
func init() {
	//gin.SetMode(gin.ReleaseMode)   // 设置gin运行模式
	//gin.DefaultWriter = io.Discard // 设置gin日志输出到空

	// 设置日志格式为Text格式
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
		ForceColors:     true,                  // 强制颜色
		//FullTimestamp:   true,                  // 显示完整时间
	})

	// 设置将日志输出到标准输出（默认的输出为stderr）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为Trace
	log.SetLevel(log.TraceLevel)
}

// main 主函数
func main() {
	config, _ := loadConfig() // 加载配置文件
	log.Debug(fmt.Sprintf("config: %v", config))

	engine := gin.New()                        // 创建无中间件应用
	_ = engine.SetTrustedProxies(nil)          // 允许所有代理
	engine.NoRoute(controller.NotFoundRoute()) // 设置404路由

	// 设置全局中间件
	middlewares := []gin.HandlerFunc{
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		}), // 设置跨域中间件
		middleware.GetRemotePortMiddleware(), // 设置获取客户端端口中间件
		middleware.LogMiddleware(),           // 设置日志中间件
		middleware.TimerMiddleware(),         // 设置计时中间件
		middleware.Recovery(),                // 设置恢复中间件
	}
	engine.Use(middlewares...) // 使用中间件

	engine.GET("/ping", func(c *gin.Context) {
		nc := response.ContextEx{Context: c}
		nc.Failure(response.Error, "pong")
	})

	err := engine.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)) // 监听并启动服务
	if err != nil {
		log.Fatal(err)
	}
}
