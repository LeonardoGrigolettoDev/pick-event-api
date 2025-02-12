package middlewares

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/utils"

	"github.com/gin-gonic/gin"
)

// Middleware para verificar token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Define o user_id no contexto para outras funções usarem
		c.Set("user_id", userID)
		c.Next()
	}
}
