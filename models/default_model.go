package models

import (
	"time"

	"gorm.io/gorm"
)

type DefaultModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index" `
}
