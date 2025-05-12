package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// auth := router.Group("/")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := router.Group("/api")
	// api.Use(middlewares.AuthMiddleware())
	{
		SetupUserRoutes(api)
		SetupEntityRoutes(api)
		SetupPeriodRoutes(api)
		SetupUserPermissionRoutes(api)
		SetupDeviceRoutes(api)
		SetupEventRoutes(api)
		SetupHistoryRoutes(api)
		SetupAdjustmentRoutes(api)
		SetupEncodeRoutes(api)
	}
}
