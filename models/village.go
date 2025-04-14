package models

type Village struct {
	ID         uint   `json:"id" gorm:"primaryKey;unique;not null"`
	Name       string `json:"name" gorm:"unique;not null;size:255"`
	DistrictID uint   `json:"district_id" gorm:"not null"`

	District District `json:"district" gorm:"foreignKey:DistrictID;references:ID"`
}
