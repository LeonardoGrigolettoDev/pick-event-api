package services

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
)

func GetEncodes() ([]models.Encode, error) {
	var encodes []models.Encode
	err := database.DB.Find(&encodes).Error
	return encodes, err
}

func GetEncodeByID(id string) (models.Encode, error) {
	var encode models.Encode
	err := database.DB.First(&encode, id).Error
	return encode, err
}

// Criar usuário
func CreateEncode(encode *models.Encode) error {
	return database.DB.Create(encode).Error
}

// Atualizar usuário
func UpdateEncode(encode *models.Encode) error {
	return database.DB.Save(encode).Error
}

// Deletar usuário
func DeleteEncode(id string) error {
	return database.DB.Delete(&models.Encode{}, id).Error
}
