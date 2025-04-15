package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
)

func GetEvents() ([]models.Event, error) {
	var events []models.Event
	err := database.DB.Find(&events).Error
	return events, err
}

func GetEventByID(id uuid.UUID) (models.Event, error) {
	var event models.Event
	err := database.DB.First(&event, id).Error
	return event, err
}

// Criar usuário
func CreateEvent(event *models.Event) error {
	return database.DB.Create(event).Error
}

// Atualizar usuário
func UpdateEvent(event *models.Event) error {
	return database.DB.Save(event).Error
}

// Deletar usuário
func DeleteEvent(id uuid.UUID) error {
	return database.DB.Delete(&models.Event{}, id).Error
}
