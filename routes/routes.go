package routes

import (
	"jarvisapi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(e *gin.Engine) {
	e.POST("/", controllers.CommandsController)
	e.GET("/", controllers.HomeController)
}
