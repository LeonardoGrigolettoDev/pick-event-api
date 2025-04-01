package controllers

import (
	"io"
	"net/http"
	"os"
	"time"

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

func RegisterEvent(c *gin.Context) {
	// Verifica se há uma imagem na requisição
	file, err := c.FormFile("image")
	hasImage := err == nil // Se não houve erro, significa que tem uma imagem

	// Processa a imagem apenas se for enviada
	var imagePath string
	if hasImage {
		// Abrindo o arquivo recebido
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open image"})
			return
		}
		defer src.Close()

		// Lendo os bytes da imagem
		imageData, err := io.ReadAll(src)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read image"})
			return
		}

		// Salvando a imagem em um diretório local
		imagePath = "./uploads/events/" + file.Filename
		err = os.WriteFile(imagePath, imageData, 0644)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save image"})
			return
		}
	}

	// Pegando os outros dados da requisição
	var event models.Event
	event.DeviceID = uuid.MustParse(c.PostForm("device_id"))
	event.Type = c.PostForm("type")
	event.Action = c.PostForm("action")

	// Se veio uma imagem, forçamos um tipo específico
	if hasImage {
		if c.PostForm("event_time") == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "event_time is required"})
			return
		}
		event.Type = "image_event"
		event.EventTime, err = time.Parse(time.RFC3339, c.PostForm("event_time"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event_time"})
			return
		}
	}

	// Validações
	if event.DeviceID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device_id"})
		return
	}
	if event.Action == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Action is required"})
		return
	}

	// Verifica se o device existe
	device, err := services.GetDeviceByID(event.DeviceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
		return
	}
	switch event.Action {
	case "face":
		//TODO INCLUDE FACE RECOGNITION AND VERIFICATION
		break
	default:
		{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
			return
		}

	}
	// Criando o evento
	event.Device = device
	if err := services.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Criando histórico do evento (opcional)
	onPeriod, err := services.GetPeriodByTimestamp(time.Now().Unix())
	if err == nil {
		services.CreateHistory(&models.History{
			EventID:  event.ID,
			PeriodID: onPeriod.ID,
		})
	}

	// Retorno
	response := gin.H{
		"message": "Event registered successfully",
		"event":   event,
	}
	if hasImage {
		response["image"] = imagePath
	}
	c.JSON(http.StatusCreated, response)
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
