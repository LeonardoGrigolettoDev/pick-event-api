package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
)

func GetEntities() ([]models.Entity, error) {
	var entities []models.Entity
	err := database.DB.Preload("Entity").Find(&entities).Error
	return entities, err
}

func GetEntityByID(id uuid.UUID) (models.Entity, error) {
	var entity models.Entity
	err := database.DB.First(&entity, id).Error
	return entity, err
}

// Criar usuário
func CreateEntity(entity *models.Entity) error {
	return database.DB.Create(entity).Error
}

// Atualizar usuário
func UpdateEntity(entity *models.Entity) error {
	return database.DB.Save(entity).Error
}

// Deletar usuário
func DeleteEntity(id uuid.UUID) error {
	return database.DB.Delete(&models.Entity{}, id).Error
}
