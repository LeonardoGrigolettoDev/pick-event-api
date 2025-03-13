package services

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/google/uuid"
)

func GetEntity() ([]models.Entity, error) {
	var entities []models.Entity
	err := database.DB.Find(&entities).Error
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
