package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventService interface {
	GetAll() ([]models.Event, error)
	GetByID(uuid.UUID) (models.Event, error)
	Update(*models.Event) error
	Delete(uuid.UUID) error
	Create(event *models.Event) error
}

type eventService struct {
	db *gorm.DB
}

func NewEventService(db *gorm.DB) EventService {
	return &eventService{db: db}
}

func (s *eventService) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := s.db.Find(&events).Error
	return events, err
}

func (s *eventService) GetByID(id uuid.UUID) (models.Event, error) {
	var event models.Event
	err := s.db.First(&event, id).Error
	return event, err
}

// Criar usuário
func (s *eventService) Create(event *models.Event) error {
	return s.db.Create(event).Error
}

// Atualizar usuário
func (s *eventService) Update(event *models.Event) error {
	return s.db.Save(event).Error
}

// Deletar usuário
func (s *eventService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.Event{}, id).Error
}
