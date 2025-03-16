package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO DECIDIR O MODELO DE HISTORICO
type History struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID  uuid.UUID `json:"event_id" gorm:"not null"`
	Event    Event     `json:"event" gorm:"foreignKey:EventID"`
	PeriodID uuid.UUID `json:"period_id" gorm:"not null"`
	Period   Period    `json:"period" gorm:"foreignKey:PeriodID"`
}
