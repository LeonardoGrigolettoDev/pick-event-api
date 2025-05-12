package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

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

func CreateEvent(c *gin.Context) {
	typeEvent := c.PostForm("type")
	typeAction := c.PostForm("action")

	if typeEvent == "" || typeAction == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "Types are required")
		return
	}

	switch typeEvent {
	case "facial":
		if typeAction != "recognition" {
			utils.RespondWithError(c, http.StatusBadRequest, "Invalid action")
			return
		}
		handleFacialRecognition(c)
	case "manual":
		c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	default:
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid event type")
	}
}

func handleFacialRecognition(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Image is required")
		return
	}

	src, err := file.Open()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Cannot open image")
		return
	}
	defer src.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error converting image")
		return
	}

	imageBase64, err := EncodeImageToBase64(buf.Bytes())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error encoding image")
		return
	}

	randomID := uuid.NewString()
	message := map[string]any{
		"id":    randomID,
		"type":  "facial",
		"image": imageBase64,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error serializing message")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = redis.Redis.Publish(ctx, "compare", messageJSON).Err()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error publishing to Redis")
		return
	}

	face, err := redis.WaitForFaceComparisonResponse(ctx, randomID)
	if err != nil {
		if err == context.DeadlineExceeded {
			utils.RespondWithError(c, http.StatusRequestTimeout, "Timeout waiting for comparison")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if face.Status == "not_found" {
		utils.RespondWithError(c, http.StatusNotFound, "Face not found")
		return
	}

	entityID, err := uuid.Parse(face.MatchedID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid entity ID")
		return
	}

	entity, err := services.GetEntityByID(entityID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Entity not found")
		return
	}

	event := models.Event{
		EntityID: entity.ID,
		Entity:   entity,
		Type:     "facial",
		Action:   "recognize",
	}

	if err := services.CreateEvent(&event); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error creating event")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event processed successfully",
		"event":   event,
		"image":   imageBase64,
	})
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
