package redis

import (
	"encoding/json"
	"log"

	"github.com/LeonardoGrigolettoDev/pick-point.git/models"
	"github.com/LeonardoGrigolettoDev/pick-point.git/services"
	"github.com/google/uuid"
)

type FaceEncoded struct {
	ID       string `json:"id"`
	Encoding string `json:"encoding"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

func ListenEncodedFaces() {
	pubsub := Redis.Subscribe(ctx, "face_encoded")
	ch := pubsub.Channel()

	for msg := range ch {
		var face FaceEncoded
		err := json.Unmarshal([]byte(msg.Payload), &face)
		if err != nil {
			log.Println("Could not decode message:", err)
			continue
		}

		if face.Status != "success" {
			log.Printf("[%s] Could not process: %s\n", face.ID, face.Message)
			continue
		}

		entityID, err := uuid.Parse(face.ID)
		if err != nil {
			log.Printf("Invalid entity ID: %s\n", face.ID)
			continue
		}

		encode := models.Encode{
			ID:       "face-" + face.ID,
			Type:     "face",
			EntityID: entityID,
		}
		err = services.CreateEncode(&encode)
		if err != nil {
			log.Printf("Could not create encode: %s\n", err)
			continue
		}
		log.Printf("Encode created: %s\n", encode.ID)
		err = Redis.Set(ctx, "face:"+encode.ID, encode, 0).Err()
		if err != nil {
			log.Printf("Could not set encode in Redis: %s\n", err)
			continue
		}
	}
}
