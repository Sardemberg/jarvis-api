package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	Payload map[string]interface{} `json:"payload"`
	Module  string                 `json:"module" validate:"nonzero"`
	Type    string                 `json:"type" validate:"nonzero"`
}

func (l *Log) Validate() error {
	err := validator.Validate(l)

	if err != nil {
		return err
	}

	return nil
}
