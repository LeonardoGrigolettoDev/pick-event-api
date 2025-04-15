package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	// Prepara o router
	GetEventsFunc = func() ([]models.Event, error) {
		return []models.Event{{
			ID:          uuid.New(),
			Observation: "Observation",
			EntityID:    uuid.New(),
			Entity: models.Entity{
				ID:   uuid.New(),
				Name: "Entity Name",
				Type: "Entity Type",
			},
			Type:   "facial",
			Action: "recognize",
		}, {
			ID:          uuid.New(),
			Observation: "Observation",
			EntityID:    uuid.New(),
			Entity: models.Entity{
				ID:   uuid.New(),
				Name: "Entity2 Name",
				Type: "Entity2 Type",
			},
			Type:   "manual",
			Action: "register",
		}}, nil
	}
	r := gin.Default()
	r.GET("/events", GetEvents)

	// Simula a requisição
	req, _ := http.NewRequest("GET", "/events", nil)
	resp := httptest.NewRecorder()

	// Executa a rota
	r.ServeHTTP(resp, req)

	// Verifica o status
	assert.Equal(t, http.StatusOK, resp.Code)

	// Aqui você pode testar o corpo, se quiser:
	// assert.JSONEq(t, `[{"id":1,"name":"Evento Teste"}]`, resp.Body.String())
}
