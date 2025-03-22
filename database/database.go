package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)

	// Koneksi ke PostgreSQL menggunakan GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	fmt.Println("✅ Koneksi ke PostgreSQL berhasil!")

	DB = db
}
