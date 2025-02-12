package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	auth := router.Group("/")
	api := router.Group("/api")
	{
		SetupUserRoutes(api)
		auth.Use(middlewares.AuthMiddleware())
	}
}
