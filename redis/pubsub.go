package redis

import (
	"context"
	"encoding/json"
	"log"

	"gorm.io/datatypes"
)

type FaceEncoded struct {
	ID       string         `json:"id"`
	Encoding datatypes.JSON `json:"encoding"`
	Status   string         `json:"status"`
	Message  string         `json:"message,omitempty"`
}

type FaceCompared struct {
	ID        string `json:"id"`
	MatchedID string `json:"matched_id,omitempty"`
	Status    string `json:"status"`
}

func SaveEncodeToRedis(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return Redis.Set(ctx, key, jsonData, 0).Err()
}

func WaitForFaceComparisonResponse(ctx context.Context, expectedID string) (*FaceCompared, error) {
	pubsub := Redis.Subscribe(ctx, "face_compared")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			var face FaceCompared
			if err := json.Unmarshal([]byte(msg.Payload), &face); err != nil {
				log.Println("Error decoding Redis message:", err)
				continue
			}
			if face.ID != expectedID {
				continue
			}

			return &face, nil

		case <-ctx.Done():
			return nil, context.DeadlineExceeded
		}
	}
}
