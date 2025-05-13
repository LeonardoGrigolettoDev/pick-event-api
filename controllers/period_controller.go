package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PeriodController struct {
	Service services.PeriodService
}

func NewPeriodController(service services.PeriodService) *PeriodController {
	return &PeriodController{Service: service}
}

// Listar todos os usuários
func (c *PeriodController) GetPeriods(ctx *gin.Context) {
	periods, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, periods)
}

// Buscar usuário por ID
func (c *PeriodController) GetPeriodByID(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	period, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Period not found."})
		return
	}
	ctx.JSON(http.StatusOK, period)
}

// Criar usuário
func (c *PeriodController) CreatePeriod(ctx *gin.Context) {
	var period models.Period
	if err := ctx.ShouldBindJSON(period); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Service.Create(&period); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, period)
}

// Atualizar usuário
func (c *PeriodController) UpdatePeriod(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var period models.Period
	if err := ctx.ShouldBindJSON(&period); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	period.ID = id
	if err := c.Service.Update(&period); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, period)
}

// Deletar usuário
func (c *PeriodController) DeletePeriod(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}
