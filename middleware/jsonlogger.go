package middleware

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/lyhgo/gin/tool/logging"
)

type RequestLog struct {
	Level        string
	ClientIP     string
	TS           string
	Method       string
	Path         string
	Proto        string
	StatusCode   int
	Latency      float64
	UserAgent    string
	ErrorMessage string
}

func JsonLogger() gin.HandlerFunc {
	return jsonLoggerHandler()
}

// jsonLoggerHandler returns a middleware to record json format request log.
func jsonLoggerHandler() gin.HandlerFunc {
	logger := logging.GetLogger()
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		r := RequestLog{
			Level:        "info",
			ClientIP:     param.ClientIP,
			TS:           param.TimeStamp.Format(time.RFC1123),
			Method:       param.Method,
			Path:         param.Path,
			Proto:        param.Request.Proto,
			StatusCode:   param.StatusCode,
			Latency:      param.Latency.Seconds(),
			UserAgent:    param.Request.UserAgent(),
			ErrorMessage: param.ErrorMessage,
		}
		if r.ErrorMessage != "" {
			r.Level = "error"
		}
		logBytes, err := sonic.Marshal(&r)
		if err != nil {
			logger.Error(err.Error())
		}
		logBytes = append(logBytes, '\n')
		return string(logBytes)
	})
}
