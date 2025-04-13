package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupEncodeRoutes(r *gin.RouterGroup) {
	{
		r.GET("/encodes", controllers.GetEncodes)
		r.GET("/encodes/:id", controllers.GetEncodeByID)
		r.POST("/encodes", controllers.CreateEncode)
		r.POST("/encodes/register", controllers.RegisterEncode)
		r.PUT("/encodes/:id", controllers.UpdateEncode)
		r.DELETE("/encodes/:id", controllers.DeleteEncode)
	}
}
