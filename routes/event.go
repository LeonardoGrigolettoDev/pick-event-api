package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupEventRoutes(r *gin.RouterGroup) {
	{
		r.GET("/events", controllers.GetEvents)
		r.GET("/events/:id", controllers.GetEventByID)
		r.PUT("/events/:id", controllers.UpdateEvent)
		r.DELETE("/events/:id", controllers.DeleteEvent)
		r.POST("/events", controllers.CreateEvent)
	}
}
