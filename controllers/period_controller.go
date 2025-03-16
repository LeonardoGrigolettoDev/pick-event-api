package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Listar todos os usuários
func GetPeriods(c *gin.Context) {
	periods, err := services.GetPeriods()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, periods)
}

// Buscar usuário por ID
func GetPeriodByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	period, err := services.GetPeriodByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Period not found."})
		return
	}
	c.JSON(http.StatusOK, period)
}

// Criar usuário
func CreatePeriod(c *gin.Context) {
	var period models.Period
	if err := c.ShouldBindJSON(period); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreatePeriod(&period); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, period)
}

// Atualizar usuário
func UpdatePeriod(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var period models.Period
	if err := c.ShouldBindJSON(&period); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	period.ID = id
	if err := services.UpdatePeriod(&period); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, period)
}

// Deletar usuário
func DeletePeriod(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := services.DeletePeriod(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
