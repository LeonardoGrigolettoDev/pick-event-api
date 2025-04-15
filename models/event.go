package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uuid.UUID  `json:"ID" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Observation string     `json:"observation" gorm:"type:varchar(255)"`
	DeviceID    *uuid.UUID `json:"device_id" gorm:"type:uuid;default:null"`
	Device      *Device    `json:"device" gorm:"foreignKey:DeviceID"`
	EntityID    uuid.UUID  `json:"entity_id" gorm:"type:uuid;"`
	Entity      Entity     `json:"entity" gorm:"foreignKey:EntityID"`
	EventTime   time.Time  `gorm:"type:timestamptz not null;default:now()"`
	Type        string     `json:"type" gorm:"not null"`
	Action      string     `json:"action" gorm:"not null"`
}
