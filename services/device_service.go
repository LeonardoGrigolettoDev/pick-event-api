package services

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceService interface {
	GetAll() ([]models.Device, error)
	Create(device *models.Device) error
	GetByID(id uuid.UUID) (models.Device, error)
	Update(device *models.Device) error
	Delete(id uuid.UUID) error
}

type deviceService struct {
	db *gorm.DB
}

func NewDeviceService(db *gorm.DB) DeviceService {
	return &deviceService{db: db}
}

func (s *deviceService) GetAll() ([]models.Device, error) {
	var devices []models.Device
	err := s.db.Find(&devices).Error
	return devices, err
}

func (s *deviceService) GetByID(id uuid.UUID) (models.Device, error) {
	var device models.Device
	err := s.db.First(&device, id).Error
	return device, err
}

func (s *deviceService) Create(device *models.Device) error {
	return s.db.Create(device).Error
}

func (s *deviceService) Update(device *models.Device) error {
	return s.db.Save(device).Error
}

func (s *deviceService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.Device{}, id).Error
}
