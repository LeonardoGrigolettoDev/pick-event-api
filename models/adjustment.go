package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Adjustment struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Observation string    `json:"observation" gorm:"type:varchar(255)"`
	HistoryID   uuid.UUID `json:"history_id" gorm:"foreignKey:HistoryID not null"`
	History     uint      `json:"history" gorm:"foreignKey:HistoryID not null"`
}
