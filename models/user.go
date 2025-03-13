package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint             `json:"id" gorm:"primaryKey"`
	Name        string           `json:"name" gorm:"not null"`
	Email       string           `json:"email" gorm:"unique"`
	Password    string           `json:"password" gorm:"type: VARCHAR(255)"`
	Type        string           `json:"type"`
	Entity      Entity           `json:"entity" gorm:"foreignKey:EntityID"`
	EntityID    uuid.UUID        `json:"entity_id" gorm:"not null"`
	Permissions []UserPermission `json:"permissions" gorm:"many2many:permission;"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Comparação de senha para login
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
