package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type status string

const (
	Processed  status = "processed"
	Failed     status = "failed"
	Processing status = "processing"
)

type Notification struct {
	gorm.Model
	Status     status                 `json:"status" validate:"nonzero"`
	Frequency  bool                   `json:"frequency" validate:"nonzero"`
	StartedAt  time.Time              `json:"started_at" validate:"nonzero"`
	FinishedAt time.Time              `json:"finished_at" validate:"nonzero"`
	Metadata   map[string]interface{} `json:"metadata" validate:"nonzero" gorm:"type:jsonb"`
}

func (n *Notification) Validate() error {
	if err := validator.Validate(n); err != nil {
		return err
	}

	return nil
}
