package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/syaeful16/shuttlelization/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ModelDefault struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index" `
}

func ConnectDB() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Gagal membaca file .env")
	}

	// Get environment variable
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	// Koneksi ke PostgreSQL menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	fmt.Println("✅ Koneksi ke PostgreSQL berhasil!")

	if err := db.AutoMigrate(&models.CustomerLogin{}); err != nil {
		log.Fatal("Gagal melakukan migrasi tabel customer_login:", err)
	}

	DB = db
}
