package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Period struct {
	gorm.Model
	ID          uuid.UUID `json:"ID" gorm:"unique"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	StartTime   time.Time `json:"start_time" gorm:"type:timestamptz;not null"`
	EndTime     time.Time `json:"end_time" gorm:"type:timestamptz;not null"`
	Watch       bool      `json:"watch" gorm:"default:false"`
	CreatedBy   uuid.UUID `gorm:"not null" json:"created_by"`
	User        User      `json:"user" gorm:"foreignKey:CreatedBy"`
}
