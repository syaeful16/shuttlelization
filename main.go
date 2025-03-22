package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/database"
)

func main() {
	// Import package database
	database.ConnectDB()

	// FIber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
