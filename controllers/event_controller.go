package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Listar todos os usuários
func GetEvents(c *gin.Context) {
	events, err := services.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Buscar usuário por ID
func GetEventByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	events, err := services.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found."})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Criar usuário
func RegisterEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	switch event.Type {
	case "recognition":
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type."})
		return
	}
	deviceUUID, err := uuid.Parse(event.DeviceID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device_id format"})
		return
	}
	if _, err := services.GetDeviceByID(deviceUUID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found."})
		return
	}

	event.DeviceID = deviceUUID
	if event.Action == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action is required."})
		return
	}
	onPeriod, err := services.GetPeriodByTimestamp(event.EventTime.Unix())
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": "Could not find period to given timestamp, but event was created."})
		return
	}

	if err := services.CreateHistory(&models.History{
		EventID:  event.ID,
		PeriodID: onPeriod.ID,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, event)
}

// Atualizar usuário
func UpdateEvent(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = id
	if err := services.UpdateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Deletar usuário
func DeleteEvent(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := services.DeleteEvent(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
