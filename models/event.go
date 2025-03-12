package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO DECIDIR O MODELO DE HISTORICO
type Event struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string    `json:"name" gorm:"unique"`
	Observation string    `json:"observation" gorm:"type:varchar(255)"`
}
