package database

import (
	"fmt"
	"log"
	"os"

	"github.com/syaeful16/shuttlelization/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Get environment variable
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Koneksi ke PostgreSQL menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	fmt.Println("✅ Koneksi ke PostgreSQL berhasil!")

	if err = db.AutoMigrate(
		&models.CustomerLogin{},
		&models.RefreshToken{},
		&models.Province{},
		&models.Regency{},
		&models.District{},
		&models.Village{},
	); err != nil {
		log.Fatal("Gagal melakukan migrasi tabel:", err)
	}

	// seed wilayah
	// SeedWilayah(db)

	DB = db
}
