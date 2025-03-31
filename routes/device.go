package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupDeviceRoutes(r *gin.RouterGroup) {
	{
		r.GET("/devices", controllers.GetDevices)
		r.GET("/devices/:id", controllers.GetDeviceByID)
		r.POST("/devices", controllers.CreateDevice)
		r.PUT("/devices/:id", controllers.UpdateDevice)
		r.DELETE("/devices/:id", controllers.DeleteDevice)
	}
}
