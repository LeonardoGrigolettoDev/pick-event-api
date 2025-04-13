package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
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
	entityID := c.PostForm("entity_id")
	if entityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entity ID haves to be provided"})
		return
	}

	// Pega o arquivo de imagem
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler imagem"})
		return
	}
	defer file.Close()

	// Prepara para enviar ao Flask
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Campo "id"
	writer.WriteField("entity_id", entityID)

	// Campo "image"
	part, err := writer.CreateFormFile("image", header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create form file"})
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not copy file"})
		return
	}
	writer.Close()

	// Envia requisição para o backend Python
	req, err := http.NewRequest("POST", "http://localhost:5000/encode-face", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create request with pick-event-service"})
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send request to pick-event-service"})
		return
	}
	defer resp.Body.Close()

	// Lê a resposta
	respBody, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Encoding enviado para processamento",
		"result":  result,
	})
}
func RecognizeEncode(c *gin.Context) {
	var encode models.Encode
	if err := c.ShouldBindJSON(&encode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if err := services.CreateEncode(&encode); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusCreated, encode)
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
