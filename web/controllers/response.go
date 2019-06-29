// Created by Hisen at 2019-06-26.
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponseHandle(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
	})
}
func SuccessResponseHandle(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}
