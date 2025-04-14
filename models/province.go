package models

type Province struct {
	ID   uint   `json:"id" gorm:"primaryKey;unique;not null"`
	Name string `json:"name" gorm:"unique;not null;size:255"`

	Regencies []Regency `json:"regencies" gorm:"foreignKey:ProvinceID;references:ID"`
}
