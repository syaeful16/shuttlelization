package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/database"
	"github.com/syaeful16/shuttlelization/helpers"
	"github.com/syaeful16/shuttlelization/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username    string `json:"username" validate:"required,min=3,max=20"`
		Password    string `json:"password" validate:"required,min=6"`
		Email       string `json:"email" validate:"required,email,unique_email"`
		Fullname    string `json:"fullname" validate:"required,min=3"`
		PhoneNumber string `json:"phone_number" validate:"required,numeric,min=10"`
	}

	var input RegisterInput

	// Parse JSON input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// Validasi input
	// Validasi input menggunakan helper
	if errors := helpers.ValidateStruct(input); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Simpan ke database
	customer := models.CustomerLogin{
		Username:    input.Username,
		Password:    string(hashedPassword),
		Email:       input.Email,
		Fullname:    input.Fullname,
		PhoneNumber: input.PhoneNumber,
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func RegisterOAuth(c *fiber.Ctx) error {
	return c.SendString("Register OAuth")
}

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}
