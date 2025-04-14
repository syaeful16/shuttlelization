package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/syaeful16/shuttlelization/database"
	"github.com/syaeful16/shuttlelization/routes"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error load env")
	}
}

func main() {
	// Import package database
	database.ConnectDB()

	// FIber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000, https://your-frontend.com", // Ganti dengan domain frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Route
	routes.RouteInit(app)

	log.Fatal(app.Listen(":8000"))
}
