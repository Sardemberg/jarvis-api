package workers

import (
	"fmt"

	"github.com/robfig/cron"
)

var Cron *cron.Cron

func InitializeScheduler() {
	Cron = cron.New()
	fmt.Println("Initializing scheduler")
	Cron.Start()
}

func StopScheduler() {
	fmt.Println("Stopping scheduler")
	Cron.Stop()
}
