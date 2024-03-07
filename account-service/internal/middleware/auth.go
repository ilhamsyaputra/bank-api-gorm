package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtfiber "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
)

func NewAuthenticationMiddleware() fiber.Handler {
	JWT_SECRET := viper.GetString("JWT_SECRET")

	return jwtfiber.New(jwtfiber.Config{
		SigningKey:   []byte(JWT_SECRET),
		ErrorHandler: AuthErrorHandler,
	})
}

func AuthErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":   fiber.StatusBadRequest,
			"status": "error",
			"remark": "terjadi kesalahan dalam authentikasi",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":   fiber.StatusUnauthorized,
		"status": "error",
		"remark": "tidak memiliki izin untuk melanjutkan proses",
	})
}
