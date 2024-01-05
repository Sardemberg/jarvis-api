package main

import (
	"jarvisapi/routes"
	"jarvisapi/services"
	"jarvisapi/workers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	LoadEnvs()
	// database.InitializeDB()
	// workers.InitializeScheduler()
	log.Fatal(services.SendNewMessage("Olá, essa mensagem é de teste"))
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
