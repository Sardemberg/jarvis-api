package controllers

import (
	"encoding/json"
	"jarvisapi/commands"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	mappedCommands := loadCommands()
	var bodyParsed map[string]string
	err := json.NewDecoder(c.Request.Body).Decode(&bodyParsed)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	userCommand, params := formatRequest(bodyParsed["message"])
	context := commands.Context{Command: mappedCommands[userCommand]}
	result, err := context.ExecuteCommand(params...)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}

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

	return mapCommands
}
