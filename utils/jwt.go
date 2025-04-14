package utils

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var AT_SECRET_KEY = []byte(os.Getenv("ACCESS_TOKEN_SECRET_KEY"))
var RT_SECRET_KEY = []byte(os.Getenv("REFRESH_TOKEN_SECRET_KEY"))

// * Struct untuk klaim token
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// * Generate Access Token
func GenerateToken(c *fiber.Ctx, userID uint, exp time.Time, secretKey []byte) (string, error) {
	tokenClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(claims *Claims, tokenString string, secretKey []byte) error {
	// Parsing token dengan claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Jika parsing gagal atau token tidak valid
	if err != nil || !token.Valid {
		return errors.New("invalid or expired token")
	}

	return nil
}
