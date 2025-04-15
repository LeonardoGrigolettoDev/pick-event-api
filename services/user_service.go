package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
)

// Buscar todos os usuários
func GetUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Preload("Entity").Find(&users).Error
	return users, err
}

// Buscar usuário por ID
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := database.DB.Preload("Entity").First(&user, id).Error
	return user, err
}

// Criar usuário
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// Atualizar usuário
func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

// Deletar usuário
func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
