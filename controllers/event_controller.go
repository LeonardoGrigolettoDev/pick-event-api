package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EventController struct {
	Service services.EventService
}

func NewEventController(service services.EventService) *EventController {
	return &EventController{Service: service}
}

// Listar todos os usuários
func (c *EventController) GetEvents(ctx *gin.Context) {
	events, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// Buscar usuário por ID
func (c *EventController) GetEventByID(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	events, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Event not found."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	typeEvent := ctx.PostForm("type")
	typeAction := ctx.PostForm("action")

	if typeEvent == "" || typeAction == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Types are required")
		return
	}

	switch typeEvent {
	case "facial":
		if typeAction != "recognition" {
			utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid action")
			return
		}
		event, imageBase64 := handleFacialRecognition(ctx)
		log.Println("Event: ", event)
		if event.EntityID == uuid.Nil {
			return
		}
		if err := c.Service.Create(&event); err != nil {
			if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_events_entity_id_type_action_key\" (SQLSTATE 23505)" {
				utils.RespondWithError(ctx, http.StatusBadRequest, "Event already exists on database.")
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Event processed successfully",
			"event":   event,
			"image":   imageBase64,
		})
	case "manual":
		ctx.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	default:
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid event type")
	}
}

func handleFacialRecognition(ctx *gin.Context) (models.Event, string) {
	file, err := ctx.FormFile("image")
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Image is required")
		return models.Event{}, ""
	}

	src, err := file.Open()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Cannot open image")
		return models.Event{}, ""
	}
	defer src.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Error converting image")
		return models.Event{}, ""
	}

	imageBase64, err := utils.EncodeImageToBase64(buf.Bytes())
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Error encoding image")
		return models.Event{}, ""
	}

	randomID := uuid.NewString()
	message := map[string]any{
		"id":    randomID,
		"type":  "facial",
		"image": imageBase64,
	}

	messageJSON, err := json.Marshal(message)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Error serializing message")
		return models.Event{}, ""
	}

	redisCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = redis.Redis.Publish(redisCtx, "compare", messageJSON).Err()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Error publishing to Redis")
		return models.Event{}, ""
	}

	face, err := redis.WaitForFaceComparisonResponse(ctx, randomID)
	if err != nil {
		if err == context.DeadlineExceeded {
			utils.RespondWithError(ctx, http.StatusRequestTimeout, "Timeout waiting for comparison")
			return models.Event{}, ""
		}
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return models.Event{}, ""
	}

	if face.Status == "not_found" {
		utils.RespondWithError(ctx, http.StatusNotFound, "Face not found")
		return models.Event{}, ""
	}

	entityID, err := uuid.Parse(face.MatchedID)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid entity ID")
		return models.Event{}, ""
	}

	entity, err := services.GetEntityByID(entityID)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, "Entity not found")
		return models.Event{}, ""
	}

	event := models.Event{
		EntityID: entity.ID,
		Entity:   entity,
		Type:     "facial",
		Action:   "recognize",
	}
	return event, imageBase64

}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = id
	if err := c.Service.Update(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

// Deletar usuário
func (c *EventController) DeleteEvent(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}
