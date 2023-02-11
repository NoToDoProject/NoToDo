// NoToDo 后端服务
package main

import (
	"fmt"
	"github.com/NoToDoProject/NoToDo/config"
	"github.com/NoToDoProject/NoToDo/controller"
	"github.com/NoToDoProject/NoToDo/controller/user"
	db "github.com/NoToDoProject/NoToDo/database"
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// init 初始化
func init() {
	gin.SetMode(gin.ReleaseMode) // 设置gin运行模式
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
	user.InitRouter(engine) // 用户操作

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

	log.Infof("Server is running at %s:%s", config.Config.Server.Host, config.Config.Server.Port)
	err := engine.Run(fmt.Sprintf("%s:%s", config.Config.Server.Host, config.Config.Server.Port)) // 监听并启动服务
	if err != nil {
		log.Fatal(err)
	}
}
