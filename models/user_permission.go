package models

import (
	"github.com/google/uuid"
)

type UserPermission struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name   string    `json:"name" gorm:"unique"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
}
