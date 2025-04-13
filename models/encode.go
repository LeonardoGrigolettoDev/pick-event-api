package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Encode struct {
	gorm.Model
	ID       string         `json:"id" gorm:"not null;unique"`
	Type     string         `json:"type" gorm:"not null"`
	EntityID uuid.UUID      `json:"entity_id" gorm:"not null"`
	Entity   Entity         `json:"entity" gorm:"foreignKey:EntityID"`
	Encoding datatypes.JSON `json:"encoding" gorm:"type:jsonb;not null"` // JSONB (ideal para PostgreSQL)
}
