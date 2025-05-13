package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
)

func SetupEventRoutes(r *gin.RouterGroup) {
	service := services.NewEventService(database.DB)
	controller := controllers.NewEventController(service)
	{
		r.GET("/events", controller.GetEvents)
		r.GET("/events/:id", controller.GetEventByID)
		r.PUT("/events/:id", controller.UpdateEvent)
		r.DELETE("/events/:id", controller.DeleteEvent)
		r.POST("/events", controller.CreateEvent)
	}
}
