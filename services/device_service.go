package services

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/google/uuid"
)

func GetDevices() ([]models.Device, error) {
	var devices []models.Device
	err := database.DB.Find(&devices).Error
	return devices, err
}

func GetDeviceByID(id uuid.UUID) (models.Device, error) {
	var device models.Device
	err := database.DB.First(&device, id).Error
	return device, err
}

func CreateDevice(device *models.Device) error {
	return database.DB.Create(device).Error
}

func UpdateDevice(device *models.Device) error {
	return database.DB.Save(device).Error
}

func DeleteDevice(id uuid.UUID) error {
	return database.DB.Delete(&models.Device{}, id).Error
}
