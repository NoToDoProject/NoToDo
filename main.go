// NoToDo 后端服务
package main

import (
	"context"
	"fmt"
	"github.com/NoToDoProject/NoToDo/config"
	"github.com/NoToDoProject/NoToDo/controller"
	serverController "github.com/NoToDoProject/NoToDo/controller/server"
	userController "github.com/NoToDoProject/NoToDo/controller/user"
	db "github.com/NoToDoProject/NoToDo/database"
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// init 初始化
func init() {
	gin.SetMode(gin.ReleaseMode) // 设置gin运行模式
	//gin.DefaultWriter = io.Discard // 设置gin日志输出到空
}

var startTime = time.Now() // 启动时间
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// main 主函数
func main() {
	config.LoadConfig() // 加载配置文件

	// 设置日志等级
	switch config.Config.Log.Level {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.Debug(fmt.Sprintf("config: %v", config.Config))

	db.Connect() // 连接数据库

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
		middleware.LogMiddleware(),   // 设置日志中间件
		middleware.TimerMiddleware(), // 设置计时中间件
		middleware.Recovery(),        // 设置恢复中间件
	}
	engine.Use(middlewares...) // 使用中间件

	// 设置路由
	routers := []model.Controller{
		serverController.Server{}, // 服务器相关
		userController.User{},     // 用户相关
	}
	for _, router := range routers {
		router.InitRouter(engine)
	}

	// websocket 测试
	engine.GET("/ws", func(c *gin.Context) {
		if !c.IsWebsocket() {
			log.Errorf("Not websocket request")
			return
		}
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Errorf("Failed to set websocket upgrade: %+v", err)
			return
		}

		logger := log.WithFields(log.Fields{
			"remote_addr": c.Request.RemoteAddr,
		})

		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				logger.Errorf("Failed to close websocket: %+v", err)
			}
		}(ws)

		logger.Infof("<<- New websocket connection")
		for {
			messageType, message, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					logger.Infof("<<- Websocket connection closed")
					break
				}
				logger.Errorf("Failed to read message: %+v", err)
				break
			}

			logger.Infof("<<- WS Recv: %-10s", message)

			for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
				message[i], message[j] = message[j], message[i]
			}
			err = ws.WriteMessage(messageType, message)
			logger.Infof("->> WS Send: %-10s", message)
			if err != nil {
				logger.Errorf("Failed to write message: %+v", err)
				break
			}
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Config.Server.Host, config.Config.Server.Port),
		Handler: engine,
	}

	go func() {
		log.Infof("Server is running at %s:%s", config.Config.Server.Host, config.Config.Server.Port)
		log.Debugf("Start using: %s", time.Since(startTime).String())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
