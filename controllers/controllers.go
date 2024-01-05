package controllers

import (
	"jarvisapi/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	hello := commands.HelloCommand{}
	context := commands.Context{Command: &hello}
	result, err := context.ExecuteCommand("--test=1")

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
