package controllers

import (
	"net/http"
	"strconv"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := c.Service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Service.Create(&user); err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists on database."})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)
	if err := c.Service.Update(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.Service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}
