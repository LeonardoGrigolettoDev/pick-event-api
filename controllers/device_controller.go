package controllers

import (
	"log"
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 🔓 🔓 🔓 libera qualquer origem
	},
}

type DeviceController struct {
	Service services.DeviceService
}

func NewDeviceController(service services.DeviceService) *DeviceController {
	return &DeviceController{Service: service}
}

func (c *DeviceController) GetDevices(ctx *gin.Context) {
	devices, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, devices)
}

func (c *DeviceController) GetDeviceByID(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	device, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Device not found."})
		return
	}
	ctx.JSON(http.StatusOK, device)
}

func (c *DeviceController) CreateDevice(ctx *gin.Context) {
	var device models.Device
	if err := ctx.ShouldBindJSON(&device); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.Service.Create(&device); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	switch device.Type {
	case "esp32cam":
		if device.MAC == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "MAC address is required for this device type."})
			return
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device type."})
		return
	}
	ctx.JSON(http.StatusCreated, device)
}

func (c *DeviceController) UpdateDevice(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var device models.Device
	if err := ctx.ShouldBindJSON(&device); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device.ID = id
	if err := c.Service.Update(&device); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, device)
}

func (c *DeviceController) DeleteDevice(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}

func (c *DeviceController) StreamDevice(ctx *gin.Context) {
	log.Println("Connection header:", ctx.Request.Header.Get("Connection"))
	log.Println("Upgrade header:", ctx.Request.Header.Get("Upgrade"))
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s", msg)

		err = conn.WriteMessage(websocket.TextMessage, []byte("ACK: "+string(msg)))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
