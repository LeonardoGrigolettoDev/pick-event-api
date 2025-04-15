package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents_Success(t *testing.T) {
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
	defer func() { GetEventsFunc = services.GetEvents }()
	r := gin.Default()
	r.GET("/events", GetEvents)
	req, _ := http.NewRequest("GET", "/events", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetEventByID_Success(t *testing.T) {
	id := uuid.New()

	GetEventByIDFunc = func(reqID uuid.UUID) (models.Event, error) {
		assert.Equal(t, id, reqID)
		return models.Event{
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
		}, nil
	}
	defer func() { GetEventByIDFunc = services.GetEventByID }()
	router := gin.Default()
	router.GET("/events/:id", GetEventByID)

	req, _ := http.NewRequest("GET", "/events/"+id.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "manual")
}

func TestUpdateEvent_Sucess(t *testing.T) {
	id := uuid.New()
	event := models.Event{
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
	}
	UpdateEventFunc = func(event *models.Event) error {
		return nil
	}
	defer func() { UpdateEventFunc = services.UpdateEvent }()
	router := gin.Default()
	router.PUT("/events/:id", UpdateEvent)
	jsonValue, _ := json.Marshal(event)
	req, _ := http.NewRequest("PUT", "/events/"+id.String(), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "manual")
}

func TestDeleteEvent_Sucess(t *testing.T) {
	id := uuid.New()
	DeleteEventFunc = func(id uuid.UUID) error {
		return nil
	}
	defer func() { UpdateEventFunc = services.UpdateEvent }()
	router := gin.Default()
	router.PUT("/events/:id", DeleteEvent)
	req, _ := http.NewRequest("PUT", "/events/"+id.String(), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), id.String())
}
