package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/controllers"
	"github.com/syaeful16/shuttlelization/middlewares"
)

func RouteInit(r *fiber.App) {
	// Route for auth
	r.Post("/register", controllers.Register)
	r.Post("/login", controllers.Login)

	r.Get("/user", middlewares.AuthMiddleware(), controllers.CurrentUser)
	r.Get("/refresh-token", controllers.RefreshToken)
	r.Delete("/logout", controllers.Logout)
}
