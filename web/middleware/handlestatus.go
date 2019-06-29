// Created by Hisen at 2019-06-29.
package middleware

import (
	"code.hanx.xin/bms/logger"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"time"
)

func LogHandleStatus() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Next()
		requestUri := context.Request.RequestURI
		logger.Logger.Debugf("请求URL%s耗时%s", requestUri, time.Since(start).String())
	}

}
func MarkTraceID() gin.HandlerFunc {
	return func(context *gin.Context) {
		u4, err := uuid.NewV4()
		uuid4 := u4.String()
		if err != nil {
			uuid4 = "-----"
		}
		context.Set("trace_id", uuid4)
		context.Next()
	}

}
