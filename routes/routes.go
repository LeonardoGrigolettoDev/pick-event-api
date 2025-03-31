package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// auth := router.Group("/")
	api := router.Group("/api")
	{
		SetupUserRoutes(api)
		SetupEntityRoutes(api)
		SetupPeriodRoutes(api)
		SetupUserPermissionRoutes(api)
		SetupDeviceRoutes(api)
		SetupEventRoutes(api)
		SetupHistoryRoutes(api)
		SetupAdjustmentRoutes(api)
		api.Use(middlewares.AuthMiddleware())
	}
}
