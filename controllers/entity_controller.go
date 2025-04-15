package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Listar todos os usuários
func GetEntities(c *gin.Context) {
	entities, err := services.GetEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities)
}

// Buscar usuário por ID
func GetEntityByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	entity, err := services.GetEntityByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Period not found."})
		return
	}
	c.JSON(http.StatusOK, entity)
}

// Criar usuário
func CreateEntity(c *gin.Context) {
	var entity models.Entity
	if err := c.ShouldBindJSON(entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateEntity(&entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, entity)
}

// Atualizar usuário
func UpdateEntity(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var entity models.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity.ID = id
	if err := services.UpdateEntity(&entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}

// Deletar usuário
func DeleteEntity(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := services.DeleteEntity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
