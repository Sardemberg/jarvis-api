package commands

import (
	"jarvisapi/database"
	"jarvisapi/models"
	"jarvisapi/services"
	"jarvisapi/workers"
	"strconv"
)

type NotificationCommand struct{}

func (n *NotificationCommand) Execute(params ...string) (string, error) {
	mapParams, err := TransformParams(params...)

	if err != nil {
		return "", err
	}

	every, err := strconv.ParseBool(mapParams["every"])

	if err != nil {
		return "", err
	}

	notification := models.Notification{
		Status:   models.Pending,
		Every:    every,
		Metadata: mapParams,
	}

	result := database.DB.Create(&notification)

	if result.Error != nil {
		return "", result.Error
	}

	cronString := getCronString(mapParams)

	workers.Cron.AddFunc(cronString, func() {})

	return "Notificação criada com sucesso", nil
}

func getCronString(params map[string]string) string {
	var cronString string

	cronKeys := [5]string{"minute", "hour", "day", "month", "day_week"}

	for _, key := range cronKeys {
		if params[key] != "" {
			cronString = cronString + params[key]
		} else {
			cronString = cronString + "*"
		}
	}

	return cronString
}

func (n *NotificationCommand) GetDescription() string {
	return "Notification command add new scheduler to your bot"
}

func executeJob(id int) {
	var notification models.Notification

	database.DB.First(&notification, id)

	if !notification.Every && notification.Status == models.Processed {
		return
	}

	services.SendNewMessage(notification.Metadata["message"])
}
