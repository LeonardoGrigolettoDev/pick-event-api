package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"

	"github.com/gin-gonic/gin"
)

// Login de usuário
func Login(ctx *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar credenciais e gerar token
	token, user, err := services.LoginUser(request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
