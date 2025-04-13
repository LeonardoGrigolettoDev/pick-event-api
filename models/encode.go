package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Encode struct {
	gorm.Model
	ID       string    `json:"ID" gorm:"not null"`
	Type     string    `json:"type" gorm:"not null"`
	EntityID uuid.UUID `json:"entity_id" gorm:"not null"`
}
