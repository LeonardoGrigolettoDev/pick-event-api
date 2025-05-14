package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EncodeController struct {
	Service services.EncodeService
}

func NewEncodeController(service services.EncodeService) *EncodeController {
	return &EncodeController{Service: service}
}

func (c *EncodeController) GetEncodes(ctx *gin.Context) {
	encodes, err := c.Service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, encodes)
}

func (c *EncodeController) GetEncodeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID haves to be provided"})
		return
	}
	encode, err := c.Service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Encode not found."})
		return
	}
	ctx.JSON(http.StatusOK, encode)
}

func (c *EncodeController) RegisterEncode(ctx *gin.Context) {
	strEntityID := ctx.PostForm("entity_id")
	typeEnconding := ctx.PostForm("type")
	override := ctx.PostForm("override")
	if strEntityID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entity ID haves to be provided"})
		return
	}

	switch typeEnconding {
	case "facial":
		file, _, err := ctx.Request.FormFile("image")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
			return
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on converting image"})
			return
		}

		imageBase64, err := utils.EncodeImageToBase64(buf.Bytes())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on image to convertion"})
			return
		}
		entityID, err := uuid.Parse(strEntityID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity uuid"})
			return
		}
		entityService := services.NewEntityService(database.DB)
		entityExists, err := entityService.GetByID(entityID)
		log.Println(entityExists)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}
		message := map[string]any{
			"id":    entityID,
			"type":  typeEnconding,
			"image": imageBase64,
		}

		messageJSON, err := json.Marshal(message)
		if err != nil {
			println("Error on marshal message:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on serializing message to JSON"})
			return
		}

		redisContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = redis.Redis.Publish(redisContext, "encode", messageJSON).Err()

		if err != nil {
			println("Error on publish message to Redis:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error on publish message to Redis"})
			return
		}
		pubsub := redis.Redis.Subscribe(redisContext, "face_encoded")
		defer pubsub.Close()
		ch := pubsub.Channel()
		for msg := range ch {
			var face redis.FaceEncoded
			err := json.Unmarshal([]byte(msg.Payload), &face)
			if err != nil {
				log.Println("Could not decode message:", err)
				continue
			}

			if face.Status != "success" {
				log.Printf("[%s] Could not process: %s\n", face.ID, face.Message)
				continue
			}
			if face.ID != strEntityID {
				continue
			}
			entityID, err := uuid.Parse(face.ID)
			if err != nil {
				log.Printf("Invalid entity ID: %s\n", face.ID)
				ctx.JSON(http.StatusBadRequest, gin.H{
					"entity_id": face.ID,
					"message":   "Invalid entity ID",
					"result":    nil,
				})
				return
			}

			existingEncode, _ := c.Service.GetByID("facial:" + face.ID)
			if existingEncode.ID != "" {
				redis.SaveEncodeToRedis(existingEncode.ID, existingEncode)
				if override != "true" {
					log.Println("Encode already exists:", existingEncode.ID)
					ctx.JSON(http.StatusAccepted, gin.H{
						"entity_id": face.ID,
						"message":   "Encode already exists for this entity",
						"result":    nil,
					})
					return
				}
				existingEncode.Encoding = face.Encoding
				err = c.Service.Update(&existingEncode)
				if err != nil {
					log.Printf("Could not update encode: %s\n", err.Error())
					ctx.JSON(http.StatusBadRequest, gin.H{
						"entity_id": face.ID,
						"message":   "Could not update encode for this entity",
						"result":    nil,
					})
					return
				}
				log.Printf("Encode updated: %s\n", existingEncode.ID)
				redis.SaveEncodeToRedis(existingEncode.ID, existingEncode)
				ctx.JSON(http.StatusOK, gin.H{
					"message":   "Encode updated successfully",
					"entity_id": face.ID,
					"result":    message,
				})
				return
			}

			encode := models.Encode{
				ID:       "facial:" + face.ID,
				Type:     "facial",
				EntityID: entityID,
				Encoding: face.Encoding,
			}
			err = c.Service.Create(&encode)
			if err != nil {
				log.Printf("Could not create encode: %s\n", err.Error())
				ctx.JSON(http.StatusBadRequest, gin.H{
					"entity_id": face.ID,
					"message":   "Could not create encode for this entity",
					"result":    nil,
				})
				return
			}
			log.Printf("Encode created: %s\n", encode.ID)
			redis.SaveEncodeToRedis(encode.ID, encode)
			ctx.JSON(http.StatusOK, gin.H{
				"message":   "Encode created successfully",
				"entity_id": face.ID,
				"result":    message,
			})
			return
		}

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not encode this type.",
			"result":  nil,
		})
	}
}

func (c *EncodeController) UpdateEncode(ctx *gin.Context) {
	id := ctx.Param("id")
	var encode models.Encode
	if err := ctx.ShouldBindJSON(&encode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encode.ID = id
	if err := c.Service.Update(&encode); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, encode)
}

func (c *EncodeController) DeleteEncode(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": id})
}
