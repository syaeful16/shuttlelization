package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/controllers"
)

func RouteInit(r *fiber.App) {
	// Route for auth
	r.Post("/register", controllers.Register)
}
