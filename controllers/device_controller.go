package controllers

import (
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDevices(c *gin.Context) {
	devices, err := services.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func GetDeviceByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	device, err := services.GetDeviceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found."})
		return
	}
	c.JSON(http.StatusOK, device)
}

func CreateDevice(c *gin.Context) {
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateDevice(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	switch device.Type {
	case "esp32cam":
		if device.MAC == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "MAC address is required for this device type."})
			return
		}
	case "UI":
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device type."})
		return
	}
	c.JSON(http.StatusCreated, device)
}

func UpdateDevice(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device.ID = id
	if err := services.UpdateDevice(&device); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, device)
}

func DeleteDevice(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := services.DeleteDevice(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
