package redis

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func EnqueueFaceEncoding(rdb *redis.Client, id string, encoding []float64) error {
	encJson, _ := json.Marshal(encoding)
	return rdb.LPush(ctx, "face:"+id, encJson).Err()
}

func DequeueFaceEncoding(rdb *redis.Client, id string) ([]float64, error) {
	val, err := rdb.RPop(ctx, "face:"+id).Result()
	if err != nil {
		return nil, err
	}

	var encoding []float64
	json.Unmarshal([]byte(val), &encoding)
	return encoding, nil
}
