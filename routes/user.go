package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/controllers"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	service := services.NewUserService(database.DB)
	controller := controllers.NewUserController(service)
	{
		r.GET("/users", controller.GetUsers)
		r.GET("/users/:id", controller.GetUserByID)
		r.POST("/users", controller.CreateUser)
		r.PUT("/users/:id", controller.UpdateUser)
		r.DELETE("/users/:id", controller.DeleteUser)
		// r.POST("/users/register", controllers.CreateUser)
		r.POST("/login", controllers.Login)
	}
}
