package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/database"
	"github.com/syaeful16/shuttlelization/helpers"
	"github.com/syaeful16/shuttlelization/models"
)

// untuk menampilkan semua user (Staff/Admin)
func AllUsers(c *fiber.Ctx) error {
	return c.SendString("All Users")
}

// mendapatkan data user yang sedang login
func CurrentUser(c *fiber.Ctx) error {
	// dapatkan user_id dari context
	user_id := c.Locals("user_id").(uint)

	var customer models.CustomerLogin
	// cek apakah user ada di database
	if err := database.DB.Where("id = ?", user_id).First(&customer).Error; err != nil {
		return helpers.Response(c, "error", fiber.StatusNotFound, "User not found", nil, nil)
	}

	return helpers.Response(c, "success", fiber.StatusOK, "User data", customer, nil)
}

// menampilkan data user berdasarkan id
func ShowUser(c *fiber.Ctx) error {
	return c.SendString("Show User")
}

// Update data user berdasarkan id
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

// Hapus data user berdasarkan id
func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
