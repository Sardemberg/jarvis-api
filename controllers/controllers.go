package controllers

import (
	"encoding/json"
	"fmt"
	"jarvisapi/commands"
	"jarvisapi/controllers/payloads"
	"jarvisapi/database"
	"jarvisapi/models"
	"jarvisapi/services"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CommandsController(c *gin.Context) {
	mappedCommands := loadCommands()

	var bodyPayload payloads.WhatsAppBusinessMessage

	err := c.Bind(&bodyPayload)

	if err != nil {
		fmt.Println(fmt.Sprintf("Erro ao processar payload: %s", err.Error()))
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	var message, from string

	if len(bodyPayload.Entry) > 0 &&
		len(bodyPayload.Entry[0].Changes) > 0 &&
		len(bodyPayload.Entry[0].Changes[0].Value.Messages) > 0 &&
		bodyPayload.Entry[0].Changes[0].Value.Messages[0].Text.Body != "" {
		message = bodyPayload.Entry[0].Changes[0].Value.Messages[0].Text.Body
		from = bodyPayload.Entry[0].Changes[0].Value.Messages[0].From
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "Invalid payload structure",
		})

		var payload map[string]interface{}

		json.NewDecoder(c.Request.Body).Decode(&payload)

		log := models.Log{
			Payload: payload,
			Module:  "Home Controller",
			Type:    "Error payload",
		}

		database.DB.Create(&log)

		return
	}

	responseNumber := fmt.Sprintf("number=%s", from)

	userCommand, params := formatRequest(message)

	fmt.Println(userCommand)

	params = append(params, responseNumber)
	context := commands.Context{Command: mappedCommands[userCommand]}
	result, err := context.ExecuteCommand(params...)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = services.SendNewMessage(result, responseNumber)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

func HomeController(c *gin.Context) {
	if c.Query("hub.mode") == "subscribe" && c.Query("hub.verify_token") == os.Getenv("WHATSAPP_TOKEN") {
		fmt.Println(c.Query("hub.challenge"))

		challenge, _ := strconv.Atoi(c.Query("hub.challenge"))

		c.JSON(http.StatusOK, challenge)
		return
	}

	c.JSON(http.StatusBadRequest, nil)
}

// Private

func formatRequest(message string) (string, []string) {
	messageSplitted := strings.Split(message, "--")
	userCommand := strings.TrimSpace(messageSplitted[0])
	params := messageSplitted[1:]

	return userCommand, params
}

func loadCommands() map[string]commands.Command {
	mapCommands := make(map[string]commands.Command)

	mapCommands["/notification"] = &commands.NotificationCommand{}
	mapCommands["/hello"] = &commands.HelloCommand{}
	mapCommands["/news"] = &commands.NewsCommand{}

	return mapCommands
}
