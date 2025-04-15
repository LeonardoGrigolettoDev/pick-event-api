package middlewares

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/utils"

	"github.com/gin-gonic/gin"
)

// Middleware para verificar token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Could not find token in Authorization."})
			c.Abort()
			return
		}

		userID, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			c.Abort()
			return
		}

		// Define o user_id no contexto para outras funções usarem
		c.Set("user_id", userID)
		c.Next()
	}
}
