package controllers_test

import (
	"jarvisapi/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var GinEngine *gin.Engine

func init() {
	GinEngine = gin.Default()
	routes.HandleRoutes(GinEngine)
}

func TestHomeMessage(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	GinEngine.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}
