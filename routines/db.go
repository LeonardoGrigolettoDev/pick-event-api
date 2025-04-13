package routines

import (
	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
)

func VerifyDBTables() {
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Period{})
	database.DB.AutoMigrate(&models.UserPermission{})
	database.DB.AutoMigrate(&models.Adjustment{})
	database.DB.AutoMigrate(&models.Entity{})
	database.DB.AutoMigrate(&models.Event{})
	database.DB.AutoMigrate(&models.History{})
	database.DB.AutoMigrate(&models.Encode{})
}
