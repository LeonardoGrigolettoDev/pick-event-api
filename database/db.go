package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var maxRetries = 5

func ConnectDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := connectWithRetry(dsn, 0)
	if err != nil {
		log.Fatal("Failed to connect to database after multiple attempts:", err)
		return err
	}

	DB = db
	log.Println("Successfully connected to DB.")

	// Start connection monitor in background
	go monitorDBConnection(dsn)

	return nil
}

func connectWithRetry(dsn string, attempt int) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Info),
		FullSaveAssociations: true,
	})

	if err != nil {
		if attempt >= maxRetries {
			return nil, err
		}

		// Exponential backoff: 1s, 2s, 4s, 8s, etc.
		waitTime := time.Duration(1<<attempt) * time.Second
		log.Printf("Connection attempt %d failed, retrying in %v: %v", attempt+1, waitTime, err)
		time.Sleep(waitTime)

		return connectWithRetry(dsn, attempt+1)
	}

	return db, nil
}

func checkDBConnection() bool {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error getting SQL DB:", err)
		return false
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Database ping failed:", err)
		return false
	}

	return true
}

func monitorDBConnection(dsn string) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if !checkDBConnection() {
			log.Println("Database connection lost, attempting to reconnect...")

			newDB, err := connectWithRetry(dsn, 0)
			if err != nil {
				log.Println("Failed to reconnect to database:", err)
				continue
			}

			DB = newDB
			log.Println("Successfully reconnected to the database")
		}
	}
}
