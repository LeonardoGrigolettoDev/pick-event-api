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

// func CreateEncode(c *gin.Context) {
// 	var encode models.Encode
// 	if err := c.ShouldBindJSON(&encode); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := services.CreateEncode(&encode); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, encode)
// }

func RegisterEncode(c *gin.Context) {
	strEntityID := c.PostForm("entity_id")
	typeEnconding := c.PostForm("type")
	override := c.PostForm("override")
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

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = redis.Redis.Publish(ctx, "encode", messageJSON).Err()

		if err != nil {
			println("Error on publish message to Redis:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on publish message to Redis"})
			return
		}
		pubsub := redis.Redis.Subscribe(ctx, "face_encoded")
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
				c.JSON(http.StatusBadRequest, gin.H{
					"entity_id": face.ID,
					"message":   "Invalid entity ID",
					"result":    nil,
				})
				return
			}

			existingEncode, _ := services.GetEncodeByID("facial:" + face.ID)
			if existingEncode.ID != "" {
				redis.SaveEncodeToRedis(existingEncode.ID, existingEncode)
				if override != "true" {
					log.Println("Encode already exists:", existingEncode.ID)
					c.JSON(http.StatusAccepted, gin.H{
						"entity_id": face.ID,
						"message":   "Encode already exists for this entity",
						"result":    nil,
					})
					return
				}
				existingEncode.Encoding = face.Encoding
				err = services.UpdateEncode(&existingEncode)
				if err != nil {
					log.Printf("Could not update encode: %s\n", err.Error())
					c.JSON(http.StatusBadRequest, gin.H{
						"entity_id": face.ID,
						"message":   "Could not update encode for this entity",
						"result":    nil,
					})
					return
				}
				log.Printf("Encode updated: %s\n", existingEncode.ID)
				redis.SaveEncodeToRedis(existingEncode.ID, existingEncode)
				c.JSON(http.StatusOK, gin.H{
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
			err = services.CreateEncode(&encode)
			if err != nil {
				log.Printf("Could not create encode: %s\n", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{
					"entity_id": face.ID,
					"message":   "Could not create encode for this entity",
					"result":    nil,
				})
				return
			}
			log.Printf("Encode created: %s\n", encode.ID)
			redis.SaveEncodeToRedis(encode.ID, encode)
			c.JSON(http.StatusOK, gin.H{
				"message":   "Encode created successfully",
				"entity_id": face.ID,
				"result":    message,
			})
			return
		}

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
