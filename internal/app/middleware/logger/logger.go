package logger

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const (
	TimestampFormat = "2006/01/02 15:04:05.000"
	maxAge          = 30 * 24 * time.Hour
	rotationTime    = 24 * time.Hour
	appNameKey      = "RUNTIME_APP_NAME"
)

var Log *logrus.Logger

var (
	logFilePath = "/tmp"
)

func GetAppName() string {
	appName := os.Getenv(appNameKey)
	return appName
}

func Logger() (*logrus.Logger, error) {

	if Log != nil {
		return Log, nil
	}
	// now := time.Now()
	appName := GetAppName()
	appLogFilePath := path.Join(logFilePath, appName)

	appLogWriter, err := rotatelogs.New(
		appLogFilePath+".%Y-%m-%d.%H%M.log",
		rotatelogs.WithLocation(time.Local),
		rotatelogs.WithLinkName(appLogFilePath+".log"),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	logFormat := logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: TimestampFormat,
	}

	logrus.SetFormatter(&logFormat)
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  appLogWriter,
			logrus.FatalLevel: appLogWriter,
			logrus.WarnLevel:  appLogWriter,
			logrus.DebugLevel: appLogWriter,
			logrus.ErrorLevel: appLogWriter,
			logrus.PanicLevel: appLogWriter,
		}, &logrus.JSONFormatter{}))

	return logger, nil
}

func LogMiddleware() gin.HandlerFunc {
	logger, err := Logger()
	if err != nil {
		logrus.Panic(err)
	}

	return func(c *gin.Context) {
		//var buf bytes.Buffer
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		//endTime := time.Now()
		// 执行时间
		//latencyTime := endTime.Sub(startTime)
		latencyTime := time.Since(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//requests id
		requestId := c.Writer.Header().Get("X-Request-Id")
		//trace id
		traceId := c.Writer.Header().Get("TraceId")
		//request headers
		//requestHeader := c.Request.Header
		//request body
		//tee := io.TeeReader(c.Request.Body, &buf)
		//requestBody, _ := io.ReadAll(tee)
		//c.Request.Body = io.NopCloser(&buf)
		//requestBody, _ := io.ReadAll(c.Request.Body)
		//response

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"x-request-id": requestId,
			"trace_id":     traceId,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
			//"req_body":     string(requestBody),
			//"req_header":   requestHeader,
		}).Info()
	}
}
