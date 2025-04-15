package redis

import (
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
)

func RedisMigrateAllEncodes() error {
	var encodes []models.Encode
	err := database.DB.Find(&encodes).Error
	if err != nil {
		return err
	}
	for _, encode := range encodes {
		err := Redis.Set(ctx, encode.ID, encode, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
