package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/services"

	"github.com/gin-gonic/gin"
)

// Login de usuário
func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar credenciais e gerar token
	token, err := services.LoginUser(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
