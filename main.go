// NoToDo 后端服务
package main

import (
	"context"
	"fmt"
	"github.com/NoToDoProject/NoToDo/config"
	"github.com/NoToDoProject/NoToDo/controller"
	serverController "github.com/NoToDoProject/NoToDo/controller/server"
	todoController "github.com/NoToDoProject/NoToDo/controller/todo"
	userController "github.com/NoToDoProject/NoToDo/controller/user"
	db "github.com/NoToDoProject/NoToDo/database"
	"github.com/NoToDoProject/NoToDo/middleware"
	"github.com/NoToDoProject/NoToDo/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func init() {
	gin.SetMode(gin.ReleaseMode)   // show less info
	gin.DefaultWriter = io.Discard // disable gin log
}

// create gin engine
func setupEngine() *gin.Engine {
	engine := gin.New()                        // get gin engine without any middleware
	_ = engine.SetTrustedProxies(nil)          // allow all proxies
	engine.NoRoute(controller.NotFoundRoute()) // add 404 route

	// add global middlewares
	middlewares := []gin.HandlerFunc{
		cors.New(cors.Config{ // CORS
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			MaxAge: 12 * time.Hour,
		}),
		middleware.LogMiddleware(),   // log
		middleware.TimerMiddleware(), // calc time cost
		middleware.Recovery(),        // recover from panic
	}
	engine.Use(middlewares...)

	// add routers
	routers := []model.Controller{
		serverController.Server{},
		userController.User{},
		todoController.Todos{},
	}
	for _, router := range routers {
		router.InitRouter(engine)
	}

	return engine
}

// main entrance of the application
func main() {
	config.LoadConfig() // load local config

	// set log level
	switch strings.ToLower(config.Config.Log.Level) {
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

	db.Connect() // connect to database

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Config.Server.Host, config.Config.Server.Port),
		Handler: setupEngine(),
	}

	// start server
	go func() {
		log.Infof("Server is running at %s:%s", config.Config.Server.Host, config.Config.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	// force close in 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Panic("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}
