package services

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/google/uuid"
)

func GetPeriods() ([]models.Period, error) {
	var periods []models.Period
	err := database.DB.Find(&periods).Error
	return periods, err
}

func GetPeriodByID(id uuid.UUID) (models.Period, error) {
	var period models.Period
	err := database.DB.First(&period, id).Error
	return period, err
}

func GetPeriodByTimestamp(timestamp int64) (models.Period, error) {
	var period models.Period
	err := database.DB.Where("start <= ? AND end >= ?", timestamp, timestamp).First(&period).Error
	return period, err
}

// Criar usuário
func CreatePeriod(period *models.Period) error {
	return database.DB.Create(period).Error
}

// Atualizar usuário
func UpdatePeriod(period *models.Period) error {
	return database.DB.Save(period).Error
}

// Deletar usuário
func DeletePeriod(id uuid.UUID) error {
	return database.DB.Delete(&models.Period{}, id).Error
}
