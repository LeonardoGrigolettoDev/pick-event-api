package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	ID          uuid.UUID `json:"ID" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Description string    `json:"description"`
	Type        string    `json:"type" gorm:"not null"`
	MAC         string    `json:"mac"`
}
