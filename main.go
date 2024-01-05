package main

import (
	"jarvisapi/database"
	"jarvisapi/routes"
	"jarvisapi/workers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	LoadEnvs()
	database.InitializeDB()
	workers.InitializeScheduler()
}

func main() {
	handler := gin.Default()
	routes.HandleRoutes(handler)
	handler.Run()
	workers.StopScheduler()
}

func LoadEnvs() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error to load enviroment variables: %s", err.Error())
		os.Exit(-1)
	}
}
