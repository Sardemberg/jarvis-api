package commands

import (
	"fmt"
	"jarvisapi/database"
	"jarvisapi/models"
	"jarvisapi/services"
	"jarvisapi/workers"
	"strconv"

	"github.com/go-co-op/gocron/v2"
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

	cronString := getCronString(mapParams)

	mapParams["cronString"] = cronString

	notification := models.Notification{
		Status:  models.Pending,
		Every:   every,
		Cron:    cronString,
		Message: mapParams["message"],
		Number:  mapParams["number"],
	}

	result := database.DB.Create(&notification)

	if result.Error != nil {
		return "", result.Error
	}

	job, err := workers.Scheduler.NewJob(
		gocron.CronJob(cronString, false),
		gocron.NewTask(func(notificationId int) {
			var notification models.Notification

			database.DB.First(&notification, notificationId)

			err := services.SendNewMessage(
				notification.Message,
				notification.Number,
			)

			if err != nil {
				notification.Status = models.Failed
			}

			notification.Status = models.Processed

			database.DB.UpdateColumn("status", &notification)
		}, int(notification.ID)),
	)

	if err != nil {
		return "", err
	}

	notification.JobId = job.ID()

	result = database.DB.UpdateColumns(notification)

	if result.Error != nil {
		workers.Scheduler.RemoveJob(job.ID())
		return "", result.Error
	}

	fmt.Println(job.ID())

	return "Notificação criada com sucesso", nil
}

func getCronString(params map[string]string) string {
	var cronString string

	cronKeys := [5]string{"minute", "hour", "day", "month", "day_week"}

	for _, key := range cronKeys {
		if params[key] != "" {
			if params["every"] != "false" {
				cronString = cronString + "*/" + params[key] + " "
			} else {
				cronString = cronString + params[key] + " "
			}
		} else {
			cronString = cronString + "* "
		}
	}

	return cronString
}

func (n *NotificationCommand) GetDescription() string {
	return "Notification command add new scheduler to your bot"
}
