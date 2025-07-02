package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
)

func SetupDeviceRoutes(r *gin.RouterGroup) {
	service := services.NewDeviceService(database.DB)
	controller := controllers.NewDeviceController(service)
	{
		r.GET("/devices", controller.GetDevices)
		r.GET("/devices/:id", controller.GetDeviceByID)
		r.POST("/devices", controller.CreateDevice)
		r.PUT("/devices/:id", controller.UpdateDevice)
		r.DELETE("/devices/:id", controller.DeleteDevice)
		r.GET("/devices/stream/:id", controller.StreamDevice)
	}
}
