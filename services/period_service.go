package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PeriodService interface {
	GetAll() ([]models.Period, error)
	Create(user *models.Period) error
	GetByTimestamp(timestamp int64) (models.Period, error)
	GetByID(id uuid.UUID) (models.Period, error)
	Update(user *models.Period) error
	Delete(id uuid.UUID) error
}

type periodService struct {
	db *gorm.DB
}

func NewPeriodService(db *gorm.DB) PeriodService {
	return &periodService{db: db}
}

func (s *periodService) GetAll() ([]models.Period, error) {
	var periods []models.Period
	err := s.db.Find(&periods).Error
	return periods, err
}

func (s *periodService) GetByID(id uuid.UUID) (models.Period, error) {
	var period models.Period
	err := s.db.First(&period, id).Error
	return period, err
}

func (s *periodService) GetByTimestamp(timestamp int64) (models.Period, error) {
	var period models.Period
	err := s.db.Where("start <= ? AND end >= ?", timestamp, timestamp).First(&period).Error
	return period, err
}

// Criar usuário
func (s *periodService) Create(period *models.Period) error {
	return s.db.Create(period).Error
}

// Atualizar usuário
func (s *periodService) Update(period *models.Period) error {
	return s.db.Save(period).Error
}

// Deletar usuário
func (s *periodService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.Period{}, id).Error
}
