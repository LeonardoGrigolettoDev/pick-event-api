package routes

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	{
		r.GET("/users", controllers.GetUsers)
		r.GET("/users/:id", controllers.GetUserByID)
		r.POST("/users", controllers.CreateUser)
		r.PUT("/users/:id", controllers.UpdateUser)
		r.DELETE("/users/:id", controllers.DeleteUser)
		r.POST("/users/register", controllers.Register)
		r.POST("/login", controllers.Login)
	}
}
