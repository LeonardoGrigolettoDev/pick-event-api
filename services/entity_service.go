package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityService interface {
	GetAll() ([]models.Entity, error)
	Create(entity *models.Entity) error
	GetByID(id uuid.UUID) (models.Entity, error)
	Update(entity *models.Entity) error
	Delete(id uuid.UUID) error
}
type entityService struct {
	db *gorm.DB
}

func NewEntityService(db *gorm.DB) EntityService {
	return &entityService{db: db}
}

func (s *entityService) GetAll() ([]models.Entity, error) {
	var entities []models.Entity
	err := s.db.Preload("Entity").Find(&entities).Error
	return entities, err
}

func (s *entityService) GetByID(id uuid.UUID) (models.Entity, error) {
	var entity models.Entity
	err := s.db.First(&entity, id).Error
	return entity, err
}

// Criar usuário
func (s *entityService) Create(entity *models.Entity) error {
	return s.db.Create(entity).Error
}

// Atualizar usuário
func (s *entityService) Update(entity *models.Entity) error {
	return s.db.Save(entity).Error
}

// Deletar usuário
func (s *entityService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.Entity{}, id).Error
}
