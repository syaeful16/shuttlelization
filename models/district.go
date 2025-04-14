package models

type District struct {
	ID        uint   `json:"id" gorm:"primaryKey;unique;not null"`
	Name      string `json:"name" gorm:"unique;not null;size:255"`
	RegencyID uint   `json:"regency_id" gorm:"not null"`

	Regency  Regency   `json:"regency" gorm:"foreignKey:RegencyID;references:ID"`
	Villages []Village `json:"villages" gorm:"foreignKey:DistrictID;references:ID"`
}
