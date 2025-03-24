package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/database"
	"github.com/syaeful16/shuttlelization/routes"
)

func main() {
	// Import package database
	database.ConnectDB()

	// FIber app
	app := fiber.New()

	// Route
	routes.RouteInit(app)

	log.Fatal(app.Listen(":8000"))
}
