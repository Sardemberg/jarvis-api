package models_test

import (
	"jarvisapi/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotificationValidation(t *testing.T) {
	notification := models.Notification{}
	assert.Equal(t, isNotValid(&notification), true)

	notification.Frequency = true
	assert.Equal(t, isNotValid(&notification), true)

	notification.StartedAt = time.Now()
	assert.Equal(t, isNotValid(&notification), true)

	notification.FinishedAt = time.Now()
	assert.Equal(t, isNotValid(&notification), true)

	notification.Metadata = make(map[string]interface{})
	notification.Metadata["teste"] = true

	assert.Equal(t, isNotValid(&notification), true)

	notification.Status = models.Processed

	assert.Equal(t, isNotValid(&notification), false)
}

func isNotValid(notification *models.Notification) bool {
	return notification.Validate() != nil
}
