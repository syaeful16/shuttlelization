package models

type CustomerLogin struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username    string `json:"username" gorm:"unique;not null;size:100"`
	Password    string `json:"password" gorm:"not null;size:255"`
	Email       string `json:"email" gorm:"unique;not null;size:100"`
	Fullname    string `json:"fullname" gorm:"not null;size:255"`
	PhoneNumber string `json:"phone_number" gorm:"not null;size:20"`
	LoginBy     string `json:"login_by" gorm:"not null;size:10;default:'system'"`
	DefaultModel
}
