package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	gorm.Model
	ID   uuid.UUID `json:"ID" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `json:"name"`
	Type string    `json:"type"`
}
