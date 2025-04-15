package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var GetEventsFunc = services.GetEvents
var GetEventByIDFunc = services.GetEventByID
var UpdateEventFunc = services.UpdateEvent
var DeleteEventFunc = services.DeleteEvent
var CreateEventFunc = services.CreateEvent
var EncodeImageToBase64 = utils.EncodeImageToBase64

// Listar todos os usuários
func GetEvents(c *gin.Context) {
	events, err := GetEventsFunc()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

// Buscar usuário por ID
func GetEventByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	events, err := GetEventByIDFunc(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func RegisterEvent(c *gin.Context) {
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
				imageBase64, err := EncodeImageToBase64(buf.Bytes())
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
}

func UpdateEvent(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = id
	if err := UpdateEventFunc(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

// Deletar usuário
func DeleteEvent(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	if err := DeleteEventFunc(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
