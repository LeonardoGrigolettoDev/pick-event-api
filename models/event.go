package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID          uuid.UUID `json:"ID" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Observation string    `json:"observation" gorm:"type:varchar(255)"`
	Entity      Entity    `json:"entity" gorm:"foreignKey:EntityID"`
	EntityID    uuid.UUID `json:"entity_id" gorm:"not null"`
	EventTime   time.Time `gorm:"type:timestamptz not null"`
}

// parsedTime, err := time.Parse(layout, datetimeStr)
// if err != nil {
//     log.Fatalf("Erro ao converter data: %v", err)
// }
