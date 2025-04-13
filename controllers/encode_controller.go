package controllers

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-point.git/services"
	"github.com/gin-gonic/gin"
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
	// Pega o ID enviado no formulário
	strEntityID := c.PostForm("entity_id")
	typeEnconding := c.PostForm("type")
	if strEntityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entity ID haves to be provided"})
		return
	}

	switch typeEnconding {
	case "face":
		// Pega o arquivo de imagem
		file, _, err := c.Request.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
			return
		}
		defer file.Close()

		// Converte a imagem para base64
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on converting image"})
			return
		}

		imageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		// Preparando a mensagem para ser enviada ao Redis
		entityID, err := uuid.Parse(strEntityID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity ID"})
			return
		}

		entityExists, err := services.GetEntityByID(entityID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			return
		}
		message := map[string]any{
			"id":    entityExists.ID,
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing typeRecognition"})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error on reading image"})
		return
	}

	switch typeRecognition {
	case "face":
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
		imageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		// 👉 Gerar ID único
		id := uuid.NewString()

		// Prepare message
		message := map[string]any{
			"id":    id,
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
			"id":      id,
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
