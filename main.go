package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/LeonardoGrigolettoDev/pick-point.git/database"
	"github.com/LeonardoGrigolettoDev/pick-point.git/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar o .env, usando valores padrão")
	}
	database.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Servidor rodando na porta " + port)
	r.Run(":" + port)
}
