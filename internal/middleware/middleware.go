package middleware

import (
	"simple-forum/internal/configs"
	"simple-forum/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.GetConfig().Service.SecretJWT

	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "Missing Token",
			})
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
