package workers

import (
	"fmt"
	"log"

	"github.com/go-co-op/gocron/v2"
)

var (
	Scheduler gocron.Scheduler
	err       error
)

func InitializeScheduler() {
	Scheduler, err = gocron.NewScheduler()

	if err != nil {
		log.Fatalf("Erro ao inicializar scheduler, erro: %s", err.Error())
	}
}

func StopScheduler() {
	fmt.Println("Stopping scheduler")
	Scheduler.Shutdown()
}
