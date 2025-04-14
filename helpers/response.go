package helpers

import (
	"github.com/gofiber/fiber/v2"
)

type Res struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func Response(c *fiber.Ctx, status string, statusCode int, message string, data interface{}, errors interface{}) error {
	return c.Status(statusCode).JSON(Res{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  errors,
	})
}
