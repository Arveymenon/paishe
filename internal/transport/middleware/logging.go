package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)
        log.Printf("%s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
    }
}
