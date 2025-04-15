package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEncodes(c *gin.Context) {
	encodes, err := services.GetEncodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, encodes)
}

func GetEncodeByID(c *gin.Context) {
	id := c.Param("id")
	encode, err := services.GetEncodeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encode not found."})
		return
	}
	c.JSON(http.StatusOK, encode)
}

func CreateEncode(c *gin.Context) {
	var encode models.Encode
	if err := c.ShouldBindJSON(&encode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateEncode(&encode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, encode)
}

func RegisterEncode(c *gin.Context) {
	strEntityID := c.PostForm("entity_id")
	typeEnconding := c.PostForm("type")
	if strEntityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entity ID haves to be provided"})
		return
	}

	switch typeEnconding {
	case "facial":
		file, _, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
			return
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on converting image"})
			return
		}

		imageBase64, err := utils.EncodeImageToBase64(buf.Bytes())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on image to convertion"})
			return
		}
		entityID, err := uuid.Parse(strEntityID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity uuid"})
			return
		}

		entityExists, err := services.GetEntityByID(entityID)
		log.Println(entityExists)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on serializing message to JSON"})
			return
		}

		ctx := context.Background()
		err = redis.Redis.Publish(ctx, "encode", messageJSON).Err()

		if err != nil {
			println("Error on publish message to Redis:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on publish message to Redis"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Image sent with success.",
			"result":  message,
		})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not encode this type.",
			"result":  nil,
		})
	}
}

func RecognizeEncode(c *gin.Context) {
	typeRecognition := c.PostForm("type")
	if typeRecognition == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing type param."})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
		return
	}

	switch typeRecognition {
	case "facial":
		file, err := image.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
			return
		}
		defer file.Close()

		// Convert to base64
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on converting image"})
			return
		}
		imageBase64, err := utils.EncodeImageToBase64(buf.Bytes())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on image to convertion"})
			return
		}

		randomID := uuid.NewString()

		// Prepare message
		message := map[string]any{
			"id":    randomID,
			"type":  typeRecognition,
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not encode this type.",
			"result":  nil,
		})
	}
}

func UpdateEncode(c *gin.Context) {
	id := c.Param("id")
	var encode models.Encode
	if err := c.ShouldBindJSON(&encode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	encode.ID = id
	if err := services.UpdateEncode(&encode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, encode)
}

func DeleteEncode(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteEncode(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": id})
}
