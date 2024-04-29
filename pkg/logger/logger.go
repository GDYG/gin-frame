package logger

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func init() {
	// 初始化 Logrus
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
}

// RequestLogger 中间件用于记录请求日志
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置日志输出到文件
		log.AddHook(NewLfsHook("runtime/logs", "request"))
		// 在处理请求之前记录请求日志
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"ip":     c.ClientIP(),
		}).Info("Incoming request")

		// 继续处理请求
		c.Next()

		// 请求处理完成后记录请求状态
		log.WithFields(log.Fields{
			"status": c.Writer.Status(),
		}).Info("Request handled")
	}
}

// NewLfsHook 创建一个新的 Logrus 钩子，用于将日志输出到文件
func NewLfsHook(logDir, prefix string) log.Hook {
	// 确保目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logPath := logDir + "/" + prefix + "_" + time.Now().Format("2006-01-02") + ".log"
	_, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	hook := lfshook.NewHook(lfshook.PathMap{
		log.InfoLevel:  logPath,
		log.ErrorLevel: logPath,
		log.WarnLevel:  logPath,
		log.FatalLevel: logPath,
		log.PanicLevel: logPath,
		log.TraceLevel: logPath,
		log.DebugLevel: logPath,
	}, &log.JSONFormatter{})

	log.AddHook(hook)

	return hook
}

// RecoveryMiddleware 中间件用于从 panic 中恢复并返回 200 状态码
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("Recovered from panic: %v", r)
				c.JSON(http.StatusOK, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		c.Next()
	}
}
