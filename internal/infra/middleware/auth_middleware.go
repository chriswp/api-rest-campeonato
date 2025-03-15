package middleware

import (
	config "github.com/chriswp/api-rest-campeonato/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := config.Envs.TokenAuth.ParseToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
