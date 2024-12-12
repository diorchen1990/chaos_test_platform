package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"chaos-platform/models"
)

func AuditMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		c.Next()
		
		// 记录审计日志
		audit := &models.AuditLog{
			UserId:    c.GetString("userId"),
			Action:    c.Request.Method + " " + c.Request.URL.Path,
			Resource:  c.Param("id"),
			Status:    c.Writer.Status(),
			Duration:  time.Since(start),
			ClientIP:  c.ClientIP(),
		}
		
		go auditLogger.Log(audit)
	}
} 