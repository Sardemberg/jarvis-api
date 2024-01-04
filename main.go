package main

import (
	"jarvisapi/database"
	"jarvisapi/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	LoadEnvs()
	database.InitializeDB()
}

func main() {
	handler := gin.Default()
	routes.HandleRoutes(handler)
	handler.Run()
}

func LoadEnvs() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load enviroment variables: %s", err.Error())
		os.Exit(-1)
	}
}
