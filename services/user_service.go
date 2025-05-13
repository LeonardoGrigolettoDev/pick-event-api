package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"gorm.io/gorm"
)

type UserService interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) error
	GetByID(id uint) (models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

// Buscar todos os usuários
func (s *userService) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Preload("Entity").Find(&users).Error
	return users, err
}

// Criar usuário
func (s *userService) Create(user *models.User) error {
	return s.db.Create(user).Error
}

// Buscar usuário por ID
func (s *userService) GetByID(id uint) (models.User, error) {
	var user models.User
	err := s.db.Preload("Entity").First(&user, id).Error
	return user, err
}

// Atualizar usuário
func (s *userService) Update(user *models.User) error {
	return s.db.Save(user).Error
}

// Deletar usuário
func (s *userService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
