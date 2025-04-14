package models

type Regency struct {
	ID         uint   `json:"id" gorm:"primaryKey;unique;not null"`
	Name       string `json:"name" gorm:"unique;not null;size:255"`
	ProvinceID uint   `json:"province_id" gorm:"not null"`

	Province  Province   `json:"province" gorm:"foreignKey:ProvinceID;references:ID"`
	Districts []District `json:"districts" gorm:"foreignKey:RegencyID;references:ID"`
}
