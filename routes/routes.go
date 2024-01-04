package routes

import (
	"jarvisapi/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(e *gin.Engine) {
	e.GET("/", controllers.HomeController)
}
