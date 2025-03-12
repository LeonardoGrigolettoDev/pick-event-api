package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO DECIDIR O MODELO DE ENTIDADE

type Entity struct {
	gorm.Model
	ID   uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string    `json:"name"`
	Type string    `json:"type"`
}
