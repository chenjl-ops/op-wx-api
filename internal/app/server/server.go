package server

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	conf2 "op-wx-api/internal/pkg/conf"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	Config *conf2.Specification
	App    *gin.Engine
	//Validate *validator.Validate
}

func NewServer() (*server, error) {
	//gin.ForceConsoleColor()
	return &server{
		Config: conf2.NacosConfig,
		App:    gin.New(),
	}, nil
}

func StartServer() error {
	//initApolloConfig()
	initConfig()

	server, err1 := NewServer()
	if err1 != nil {
		return err1
	}

	// 初始化header
	server.initRequestHeader()
	// 初始化trace id
	server.initTraceId()
	// 初始化日志
	server.initLog()

	// 初始化swagger
	//server.InitSwagger()

	// 初始化路由
	server.InitRouter(true)
	server.LoadUrl()

	//启动服务
	//err := server.Run()

	//优雅启动
	err := server.GracefulRun()
	if err != nil {
		return err
	}
	return nil
}

// Run 启动服务
func (s *server) Run() error {
	port := fmt.Sprintf(":%d", cmp.Or(s.Config.ServerRunPort, 8080))
	return s.App.Run(port)
}

// GracefulRun 优雅启动服务
func (s *server) GracefulRun() error {
	port := fmt.Sprintf(":%d", cmp.Or(s.Config.ServerRunPort, 8080))

	srv := &http.Server{
		Addr:    port,
		Handler: s.App.Handler(),
	}
	log.Info("Start Service Successful Listen Port", port)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		} else {
			log.Info("Start Service Error: ", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	return nil
}
