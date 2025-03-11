package main

import (
	"log"
	"os"
	"strconv"

	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/routes"
	"github.com/LeonardoGrigolettoDev/pick-point.git/routines"
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

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on PORT " + port)
	r.Run(":" + port)
}
