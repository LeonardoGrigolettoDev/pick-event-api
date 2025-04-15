package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
)

func GetHistories() ([]models.History, error) {
	var histories []models.History
	err := database.DB.Find(&histories).Error
	return histories, err
}

func GetHistoryByID(id uuid.UUID) (models.Period, error) {
	var history models.Period
	err := database.DB.First(&history, id).Error
	return history, err
}

// Criar usuário
func CreateHistory(history *models.History) error {
	return database.DB.Create(history).Error
}

// Atualizar usuário
func UpdateHistory(history *models.History) error {
	return database.DB.Save(history).Error
}

// Deletar usuário
func DeleteHistory(id uuid.UUID) error {
	return database.DB.Delete(&models.History{}, id).Error
}
