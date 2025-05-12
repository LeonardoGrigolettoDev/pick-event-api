package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	{
		r.GET("/users", controllers.GetUsers)
		r.GET("/users/:id", controllers.GetUserByID)
		r.POST("/users", controllers.CreateUser)
		r.PUT("/users/:id", controllers.UpdateUser)
		r.DELETE("/users/:id", controllers.DeleteUser)
		// r.POST("/users/register", controllers.CreateUser)
		r.POST("/login", controllers.Login)
	}
}
