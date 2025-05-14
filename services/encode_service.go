package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"gorm.io/gorm"
)

type EncodeService interface {
	GetAll() ([]models.Encode, error)
	Create(user *models.Encode) error
	GetByID(id string) (models.Encode, error)
	Update(user *models.Encode) error
	Delete(id string) error
}

type encodeService struct {
	db *gorm.DB
}

func NewEncodeService(db *gorm.DB) EncodeService {
	return &encodeService{db: db}
}

func (s *encodeService) GetAll() ([]models.Encode, error) {
	var encodes []models.Encode
	err := s.db.Find(&encodes).Error
	return encodes, err
}

func (s *encodeService) GetByID(id string) (models.Encode, error) {
	var encode models.Encode
	err := s.db.Where("id = ?", id).First(&encode)
	return encode, err.Error
}

// Criar usuário
func (s *encodeService) Create(encode *models.Encode) error {
	return s.db.Create(encode).Error
}

// Atualizar usuário
func (s *encodeService) Update(encode *models.Encode) error {
	return s.db.Save(encode).Error
}

// Deletar usuário
func (s *encodeService) Delete(id string) error {
	return s.db.Delete(&models.Encode{}, id).Error
}
