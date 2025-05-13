package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HistoryService interface {
	GetAll() ([]models.History, error)
	GetByID(id uuid.UUID) (models.Period, error)
	Create(history *models.History) error
	Update(history *models.History) error
	Delete(id uuid.UUID) error
}
type historyService struct {
	db *gorm.DB
}

func (s *historyService) GetAll() ([]models.History, error) {
	var histories []models.History
	err := s.db.Find(&histories).Error
	return histories, err
}

func (s *historyService) GetByID(id uuid.UUID) (models.Period, error) {
	var history models.Period
	err := s.db.First(&history, id).Error
	return history, err
}

func (s *historyService) Create(history *models.History) error {
	return s.db.Create(history).Error
}

func (s *historyService) Update(history *models.History) error {
	return s.db.Save(history).Error
}

func (s *historyService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.History{}, id).Error
}
