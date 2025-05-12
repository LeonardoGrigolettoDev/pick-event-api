package main

import (
	"log"
	"os"
	"strconv"

	"github.com/LeonardoGrigolettoDev/pick-event-api.git/database"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/redis"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/routes"
	"github.com/LeonardoGrigolettoDev/pick-event-api.git/routines"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file. Using default values.")
	}
	maxRetriesStr := os.Getenv("DB_CONNECTION_RETRIES")

	maxRetries, err := strconv.Atoi(maxRetriesStr)
	if err != nil {
		maxRetries = 5
	}

	db_connection_retries := 0
	err = database.ConnectDB()

	for err != nil && db_connection_retries < maxRetries {
		log.Println("Failed to connect to database. Retrying...")
		err = database.ConnectDB()
		db_connection_retries++
	}

	routines.VerifyDBTables()
	redis.SetupRedisClient()
	// go redis.ListenEncodedFaces()
	// go redis.ListenComparedFaces()

	gin.SetMode(gin.DebugMode) // Habilita o modo debug
	r := gin.Default()

	routes.SetupRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Suas rotas aqui...
	// r.Run()
	r.Run(":" + port)
	log.Println("Server running on PORT " + port)
}
