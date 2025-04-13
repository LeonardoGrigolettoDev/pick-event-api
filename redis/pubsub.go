package redis

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func StoreEncoding(rdb *redis.Client, id string, encoding []float64) error {
	encJson, _ := json.Marshal(encoding)
	return rdb.Set(ctx, "face:"+id, encJson, 0).Err() // 0 = sem expiração
}

func GetEncoding(rdb *redis.Client, id string) ([]float64, error) {
	val, err := rdb.Get(ctx, "face:"+id).Result()
	if err != nil {
		return nil, err
	}

	var encoding []float64
	json.Unmarshal([]byte(val), &encoding)
	return encoding, nil
}
