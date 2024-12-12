package middleware

import (
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "No authorization token provided"
            })
            return
        }
        
        // 验证token
        claims, err := validateToken(token)
        if err != nil {
            c.AbortWithStatusJSON(401, gin.H{
                "error": "Invalid token"
            })
            return
        }
        
        // 设置用户信息
        c.Set("userId", claims.UserId)
        c.Set("roles", claims.Roles)
        
        c.Next()
    }
} 