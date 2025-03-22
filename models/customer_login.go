package models

type CustomerLogin struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;size:100"`
	Email    string `json:"email" gorm:"unique;not null;size:100"`
	Password string `json:"password" gorm:"not null;size:255"`
}
