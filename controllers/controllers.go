package controllers

import (
	"jarvisapi/commands"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	hello := commands.HelloCommand{}
	context := commands.Context{Command: &hello}
	result, err := context.ExecuteCommand()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
