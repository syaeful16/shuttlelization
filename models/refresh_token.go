package models

import "time"

type RefreshToken struct {
	UserID    uint          `json:"user_id" gorm:"primaryKey"`
	Token     string        `json:"token" gorm:"unique;not null"`
	Customer  CustomerLogin `json:"customer" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time     `json:"created_at"`
}
