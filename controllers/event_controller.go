package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-point.git/services"
	"github.com/LeonardoGrigolettoDev/pick-point.git/utils"
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

	typeEvent := c.PostForm("type")
	typeAction := c.PostForm("action")
	file, err := c.FormFile("image")

	if typeEvent == "" || typeAction == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Types are required"})
		return
	}

	if err != nil {
		if typeEvent != "manual" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
			return
		}
	}

	switch typeEvent {
	case "facial":
		{
			src, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot open image"})
				return
			}
			defer src.Close()

			buf := new(bytes.Buffer)
			_, err = io.Copy(buf, src)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on converting image"})
				return
			}
			switch typeAction {
			case "recognition":
				imageBase64, err := utils.EncodeImageToBase64(buf.Bytes())
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on image to convertion"})
					return
				}

				randomID := uuid.NewString()

				// Prepare message
				message := map[string]any{
					"id":    randomID,
					"type":  typeEvent,
					"image": imageBase64,
				}

				messageJSON, err := json.Marshal(message)
				if err != nil {
					println("Error on marshal message:", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on serializing message to JSON"})
					return
				}

				ctx := context.Background()
				err = redis.Redis.Publish(ctx, "compare", messageJSON).Err()
				if err != nil {
					println("Error on publish message to Redis:", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on publish message to Redis"})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"message": "Image sent for analysts.",
					"id":      randomID,
				})

			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
				return

			}
			break
		}

	case "manual":
		c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
		return
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event type"})
		return
	}

	// Pegando os outros dados da requisição
	// event.DeviceID = uuid.MustParse(c.PostForm("device_id"))

	// // Se veio uma imagem, forçamos um tipo específico
	// if hasImage {
	// 	if c.PostForm("event_time") == "" {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "event_time is required"})
	// 		return
	// 	}
	// 	event.Type = "image_event"
	// 	event.EventTime, err = time.Parse(time.RFC3339, c.PostForm("event_time"))
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event_time"})
	// 		return
	// 	}
	// }

	// // Validações
	// if event.DeviceID == uuid.Nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device_id"})
	// 	return
	// }
	// if event.Action == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Action is required"})
	// 	return
	// }

	// // Verifica se o device existe
	// device, err := services.GetDeviceByID(event.DeviceID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
	// 	return
	// }
	// switch event.Action {
	// case "face":
	// 	//TODO INCLUDE FACE RECOGNITION AND VERIFICATION
	// 	break
	// default:
	// 	{
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
	// 		return
	// 	}

	// }
	// // Criando o evento
	// event.Device = device
	// if err := services.CreateEvent(&event); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Criando histórico do evento (opcional)
	// onPeriod, err := services.GetPeriodByTimestamp(time.Now().Unix())
	// if err == nil {
	// 	services.CreateHistory(&models.History{
	// 		EventID:  event.ID,
	// 		PeriodID: onPeriod.ID,
	// 	})
	// }

	// // Retorno
	// response := gin.H{
	// 	"message": "Event registered successfully",
	// 	"event":   event,
	// }
	// if hasImage {
	// 	response["image"] = imagePath
	// }
	// c.JSON(http.StatusCreated, response)
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
