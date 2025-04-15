package redis

import (
	"encoding/json"
	"log"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/models"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/services"
	"github.com/google/uuid"
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

func saveEncodeToRedis(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return Redis.Set(ctx, key, jsonData, 0).Err()
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

		existingEncode, _ := services.GetEncodeByID("facial-" + face.ID)
		if existingEncode.ID != "" {
			log.Println("Encode already exists:", existingEncode.ID)
			saveEncodeToRedis(existingEncode.ID, existingEncode)
			continue
		}

		encode := models.Encode{
			ID:       "facial:" + face.ID,
			Type:     "facial",
			EntityID: entityID,
			Encoding: face.Encoding,
		}
		err = services.CreateEncode(&encode)
		if err != nil {
			log.Printf("Could not create encode: %s\n", err)
			continue
		}
		log.Printf("Encode created: %s\n", encode.ID)
		saveEncodeToRedis(encode.ID, encode)
	}
}

func ListenComparedFaces() {
	pubsub := Redis.Subscribe(ctx, "face_compared")
	ch := pubsub.Channel()
	log.Println("Listening to compared faces...")
	for msg := range ch {
		var face FaceCompared
		err := json.Unmarshal([]byte(msg.Payload), &face)
		if err != nil {
			log.Println("Could not decode message:", err)
			continue
		}
		log.Println(face)
		if face.Status != "success" {
			log.Printf("[%s] Could not process: %s\n", face.ID, face.Status)
			continue
		}
		log.Println("face ID:", face.ID)

		entityID, err := uuid.Parse(face.MatchedID)
		if err != nil {
			log.Printf("Invalid entity ID: %s\n", face.MatchedID)
			continue
		}
		log.Println("Entity ID:", entityID)
		entity, err := services.GetEntityByID(entityID)
		if err != nil {
			log.Printf("Could not find entity: %s\n", err)
			continue
		}
		log.Println("Entity:", entity)
		event := models.Event{
			EntityID: entity.ID,
			Entity:   entity,
			Type:     "facial",
			Action:   "recognize",
		}
		err = services.CreateEvent(&event)
		if err != nil {
			log.Printf("Could not create event: %s\n", err)
			continue
		}
		log.Printf("Event created: %s\n", event.ID)
		//MAIS TARDE PUBLICAR EM ALGUMA PORTA DE RESPOSTA WEBSOCKET
	}

}
