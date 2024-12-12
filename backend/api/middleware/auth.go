package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/chaos-mesh/chaos-platform/pkg/errors"
	"github.com/chaos-mesh/chaos-platform/pkg/jwt"
	"github.com/chaos-mesh/chaos-platform/pkg/logger"
)

type AuthMiddleware struct {
	jwtService *jwt.Service
	logger     *logger.Logger
}

func NewAuthMiddleware(jwtService *jwt.Service, logger *logger.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
		logger:     logger,
	}
}

func (m *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			err := errors.NewAppError(
				errors.ErrUnauthorized,
				"No authorization token provided",
				"",
			)
			c.AbortWithStatusJSON(401, err)
			return
		}
		
		claims, err := m.jwtService.ValidateToken(token)
		if err != nil {
			m.logger.Error("Token validation failed", 
				"error", err,
				"token", token,
			)
			err := errors.NewAppError(
				errors.ErrUnauthorized,
				"Invalid token",
				err.Error(),
			)
			c.AbortWithStatusJSON(401, err)
			return
		}
		
		c.Set("userId", claims.UserId)
		c.Set("roles", claims.Roles)
		c.Next()
	}
} 