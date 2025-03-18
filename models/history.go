package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ID       uuid.UUID `json:"ID" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID  uuid.UUID `json:"event_id" gorm:"not null"`
	Event    Event     `json:"event" gorm:"foreignKey:EventID"`
	PeriodID uuid.UUID `json:"period_id" gorm:"not null"`
	Period   Period    `json:"period" gorm:"foreignKey:PeriodID"`
}
