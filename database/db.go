package database

import (
	"fmt"
	"jarvisapi/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitializeDB() {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Fortaleza",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		"5432",
	)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error to connect database. Error: %s", err.Error())
		os.Exit(-1)
	}

	DB.AutoMigrate(&models.Notification{})

	fmt.Println("Succefuly connection to database")
}
