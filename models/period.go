package models

import (
	"time"

	"gorm.io/gorm"
)

type Period struct {
	gorm.Model
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	StartTime   time.Time `json:"start_time" gorm:"type:timestamptz;not null"`
	EndTime     time.Time `json:"end_time" gorm:"type:timestamptz;not null"`
}
