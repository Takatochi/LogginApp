package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		statusCode := c.Writer.Status()
		log.Printf("[%d] %s %s | %v", statusCode, method, path, latency)
	}
}
