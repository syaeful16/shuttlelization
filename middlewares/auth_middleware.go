package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/syaeful16/shuttlelization/helpers"
	"github.com/syaeful16/shuttlelization/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header Authorization
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			return helpers.Response(c, "error", fiber.StatusUnauthorized, "Token is required", nil, nil)
		}

		// ? verifikasi token apakah masih valid
		claims := &utils.Claims{}
		if err := utils.VerifyToken(claims, accessToken, utils.AT_SECRET_KEY); err != nil {
			return helpers.Response(c, "error", fiber.StatusUnauthorized, err.Error(), nil, nil)
		}

		// Simpan user_id di context Fiber
		c.Locals("user_id", claims.UserID)

		// Lanjut ke handler berikutnya
		return c.Next()
	}
}
