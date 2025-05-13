package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EntityController struct {
	Service services.EntityService
}

func NewEntityController(service services.EntityService) *EntityController {
	return &EntityController{Service: service}
}

// Listar todos os usuários
func (c *EntityController) GetEntities(ctx *gin.Context) {
	entities, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entities)
}

// Buscar usuário por ID
func (c *EntityController) GetEntityByID(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	entity, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Period not found."})
		return
	}
	ctx.JSON(http.StatusOK, entity)
}

// Criar usuário
func (c *EntityController) CreateEntity(ctx *gin.Context) {
	var entity models.Entity
	if err := ctx.ShouldBindJSON(entity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Service.Create(&entity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, entity)
}

// Atualizar usuário
func (c *EntityController) UpdateEntity(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var entity models.Entity
	if err := ctx.ShouldBindJSON(&entity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity.ID = id
	if err := c.Service.Update(&entity); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, entity)
}

// Deletar usuário
func (c *EntityController) DeleteEntity(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}
