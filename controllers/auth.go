package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/database"
	"github.com/syaeful16/shuttlelization/helpers"
	"github.com/syaeful16/shuttlelization/models"
	"github.com/syaeful16/shuttlelization/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username    string `json:"username" validate:"required,min=3,max=20"`
		Password    string `json:"password" validate:"required,min=6"`
		Email       string `json:"email" validate:"required,email"`
		Fullname    string `json:"fullname" validate:"required,min=3"`
		PhoneNumber string `json:"phone_number" validate:"required,numeric,min=10"`
	}

	var input RegisterInput

	// Parse JSON input
	if err := c.BodyParser(&input); err != nil {
		return helpers.Response(c, "error", fiber.StatusBadRequest, "Invalid input", nil, nil)
	}

	// Validasi input
	// Validasi input menggunakan helper
	if errors := helpers.Validate(input); errors != nil {
		return helpers.Response(c, "error", fiber.StatusBadRequest, "Invalid input", nil, errors)
	}

	// Cek apakah username atau email sudah digunakan
	var existingUser models.CustomerLogin
	if err := database.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
		// Jika ada data ditemukan, berarti username atau email sudah terdaftar
		return helpers.Response(c, "error", fiber.StatusConflict, "Username or email is already registered", nil, nil)
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to hash password", nil, nil)
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
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to register user", nil, nil)
	}

	return helpers.Response(c, "success", fiber.StatusCreated, "User registered successfully", nil, nil)
}

func RegisterOAuth(c *fiber.Ctx) error {
	return c.SendString("Register OAuth")
}

func Login(c *fiber.Ctx) error {
	// TODO: Implement login with token

	// * 1️⃣ Struct untuk input login
	type loginInput struct {
		UsernameEmail string `json:"username_or_email" validate:"required"`
		Password      string `json:"password" validate:"required"`
	}

	// Parsing & validasi input
	var input loginInput
	if err := c.BodyParser(&input); err != nil || helpers.Validate(input) != nil {
		return helpers.Response(c, "error", fiber.StatusBadRequest, "Invalid input", nil, nil)
	}

	// * 4️⃣ - Cek apakah user ada di database
	var customer models.CustomerLogin
	if err := database.DB.Where("username = ? OR email = ?", input.UsernameEmail, input.UsernameEmail).First(&customer).Error; err != nil {
		return helpers.Response(c, "error", fiber.StatusNotFound, "User not found", nil, nil)
	}

	// * 5️⃣ - Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password)); err != nil {
		return helpers.Response(c, "error", fiber.StatusUnauthorized, "Invalid password", nil, nil)
	}

	// * 6️⃣ - Generate refresh & access token
	// generate refresh token
	expRefreshToken := time.Now().Add(120 * time.Minute) // ! Set expired time 120 menit
	refreshToken, err := utils.GenerateToken(c, customer.ID, expRefreshToken, utils.RT_SECRET_KEY)
	if err != nil {
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to generate refresh token", nil, nil)
	}

	// ? Cek apakah refresh token sudah ada di database
	// Simpan refresh token di database
	refreshTokenModel := models.RefreshToken{
		UserID: customer.ID,
		Token:  refreshToken,
	}
	database.DB.Where("user_id = ?", customer.ID).Assign(refreshTokenModel).FirstOrCreate(&refreshTokenModel)

	// * Simpan refresh token ke cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  expRefreshToken,
		Path:     "/",
		HTTPOnly: true,
		Secure:   false,
	})

	// * Generate access token
	expAccessToken := time.Now().Add(15 * time.Minute) // ! Set expired time 15 menit
	accessToken, err := utils.GenerateToken(c, customer.ID, expAccessToken, utils.AT_SECRET_KEY)
	if err != nil {
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to generate access token", nil, nil)
	}

	resAccessToken := map[string]string{
		"token": accessToken,
	}

	return helpers.Response(c, "success", fiber.StatusOK, "Login success", resAccessToken, nil)
}

func Logout(c *fiber.Ctx) error {
	// ambil refresh token dari cookie
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return helpers.Response(c, "error", fiber.StatusNoContent, "Missing refresh token", nil, nil)
	}

	// ambil data claims dari refresh token
	claims := &utils.Claims{}
	if err := utils.VerifyToken(claims, refreshToken, utils.RT_SECRET_KEY); err != nil {
		return helpers.Response(c, "error", fiber.StatusUnauthorized, err.Error(), nil, nil)
	}

	// hapus refresh token dari database
	if err := database.DB.Where("user_id = ? AND token = ?", claims.UserID, refreshToken).Delete(&models.RefreshToken{}).Error; err != nil {
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to delete refresh token", nil, nil)
	}
	// hapus cookie refresh token
	c.ClearCookie("refresh_token")

	return helpers.Response(c, "success", fiber.StatusOK, "Logout success", nil, nil)
}

func RefreshToken(c *fiber.Ctx) error {
	// ambil refresh token dari cookie
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return helpers.Response(c, "error", fiber.StatusUnauthorized, "Missing refresh token", nil, nil)
	}

	fmt.Println(refreshToken)

	// ? Verifikasi refresh token masih valid atau tidak
	claims := &utils.Claims{}
	if err := utils.VerifyToken(claims, refreshToken, utils.RT_SECRET_KEY); err != nil {
		return helpers.Response(c, "error", fiber.StatusUnauthorized, err.Error(), nil, nil)
	}

	// buat access token baru
	expAccessToken := time.Now().Add(1 * time.Minute) // ! Set expired time 1 menit
	accessToken, err := utils.GenerateToken(c, claims.UserID, expAccessToken, utils.AT_SECRET_KEY)
	if err != nil {
		return helpers.Response(c, "error", fiber.StatusInternalServerError, "Failed to generate access token", nil, nil)
	}

	resAccessToken := map[string]string{
		"token": accessToken,
	}

	return helpers.Response(c, "success", fiber.StatusOK, "Token refreshed", resAccessToken, nil)
}

func TestResponse(c *fiber.Ctx) error {
	return helpers.Response(c, "success", fiber.StatusOK, "Test response", nil, nil)
}
