package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"op-wx-api/internal/app/middleware/header"
	"op-wx-api/internal/app/middleware/logger"
	"op-wx-api/internal/app/weichat"
	"op-wx-api/internal/pkg/conf"
	"time"
)

// 初始化 config
func initConfig() {
	var err error
	err = conf.NacosReadRemoteConfig()
	//fmt.Println("init conf: ", conf.NacosConfig)
	if nil != err {
		log.Fatal(err)
	}
	log.Info("init conf success: ======")
}

// initLog 初始化log配置
func (s *server) initLog() *gin.Engine {
	logs := logger.LogMiddleware()
	s.App.Use(logs)
	log.Info("init log success: =======")
	return s.App
}

// initRequestHeader 初始化request header
func (s *server) initRequestHeader() *gin.Engine {
	requestHeader := header.RequestIDMiddleware()
	s.App.Use(requestHeader)
	log.Info("init request id header success: ========")
	return s.App
}

// initTraceId 初始化trace id
func (s *server) initTraceId() *gin.Engine {
	traceId := header.TraceIDMiddleware()
	s.App.Use(traceId)
	log.Info("init trace id success: ========")
	return s.App
}

// InitRouter 加载gin 路由配置
func (s *server) InitRouter(crossDomain bool) *gin.Engine {
	if true == crossDomain {
		s.App.Use(cors.New(cors.Config{
			//AllowAllOrigins: true,
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour,
		}))
		return s.App
	}
	return s.App
}

// LoadUrl 加载路由
func (s *server) LoadUrl() *gin.Engine {
	weichat.Url(s.App)

	log.Info("Load router Url success: =======")
	return s.App
}
