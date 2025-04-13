package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// auth := router.Group("/")
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
