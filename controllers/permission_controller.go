package controllers

// import (
// 	"net/http"

// 	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
// 	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
// 	"github.com/gin-gonic/gin"
// )

// func GetPermissions(c *gin.Context) {
// 	var perm []models.Permission
// 	database.DB.Find(&perm)
// 	c.JSON(http.StatusOK, perm)
// }

// func CreatePermission(c *gin.Context) {
// 	var perm models.Permission
// 	if err := c.ShouldBindJSON(&perm); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	database.DB.Create(&perm)
// 	c.JSON(http.StatusCreated, perm)
// }
