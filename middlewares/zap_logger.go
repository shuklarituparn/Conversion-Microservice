package middlewares

//import (
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"gorm.io/gorm/logger"
//	"time"
//)
//
//func ZapLogger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//		c.Next() // Process request
//		duration := time.Since(start)
//		logger.Info("request",
//			zap.String("method", c.Request.Method),
//			zap.String("path", c.Request.URL.Path),
//			zap.Int("status", c.Writer.Status()),
//			zap.Duration("duration", duration),
//		)
//	}
//}
