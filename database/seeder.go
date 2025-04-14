package database

import (
	"log"

	"github.com/syaeful16/shuttlelization/models"
	"gorm.io/gorm"
)

func SeedWilayah(db *gorm.DB) {
	provinces := []models.Province{
		{
			ID: 31, Name: "DKI Jakarta",
			Regencies: []models.Regency{
				{ID: 3171, Name: "Jakarta Selatan"},
				{ID: 3172, Name: "Jakarta Timur"},
				{ID: 3173, Name: "Jakarta Pusat"},
				{ID: 3174, Name: "Jakarta Barat"},
				{ID: 3175, Name: "Jakarta Utara"},
				{ID: 3176, Name: "Kepulauan Seribu"},
			},
		},
		{
			ID: 32, Name: "Jawa Barat",
			Regencies: []models.Regency{
				{ID: 3201, Name: "Bogor"},
				{ID: 3202, Name: "Sukabumi"},
				{ID: 3203, Name: "Cianjur"},
				{ID: 3204, Name: "Bandung"},
				{ID: 3273, Name: "Bekasi"},
				{ID: 3275, Name: "Depok"},
				{ID: 3277, Name: "Cimahi"},
			},
		},
		{
			ID: 33, Name: "Jawa Tengah",
			Regencies: []models.Regency{
				{ID: 3301, Name: "Cilacap"},
				{ID: 3302, Name: "Banyumas"},
				{ID: 3303, Name: "Purbalingga"},
				{ID: 3371, Name: "Semarang"},
				{ID: 3372, Name: "Surakarta"},
				{ID: 3373, Name: "Salatiga"},
			},
		},
		{
			ID: 35, Name: "Jawa Timur",
			Regencies: []models.Regency{
				{ID: 3501, Name: "Pacitan"},
				{ID: 3502, Name: "Ponorogo"},
				{ID: 3503, Name: "Trenggalek"},
				{ID: 3571, Name: "Surabaya"},
				{ID: 3572, Name: "Malang"},
				{ID: 3573, Name: "Kediri"},
			},
		},
		{
			ID: 51, Name: "Bali",
			Regencies: []models.Regency{
				{ID: 5101, Name: "Jembrana"},
				{ID: 5102, Name: "Tabanan"},
				{ID: 5103, Name: "Badung"},
				{ID: 5171, Name: "Denpasar"},
			},
		},
	}

	for _, province := range provinces {
		if err := db.Create(&province).Error; err != nil {
			log.Printf("Failed to seed province: %v", err)
		}
	}
}
