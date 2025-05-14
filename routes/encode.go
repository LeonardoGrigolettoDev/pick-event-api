package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
)

func SetupEncodeRoutes(r *gin.RouterGroup) {
	service := services.NewEncodeService(database.DB)
	controller := controllers.NewEncodeController(service)
	{
		r.GET("/encodes", controller.GetEncodes)
		r.GET("/encodes/:id", controller.GetEncodeByID)
		r.POST("/encodes", controller.RegisterEncode)
		r.PUT("/encodes/:id", controller.UpdateEncode)
		r.DELETE("/encodes/:id", controller.DeleteEncode)
	}
}
