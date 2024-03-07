package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ilhamsyaputra/bank-api-gorm/config"
	"github.com/ilhamsyaputra/bank-api-gorm/internal/entity"
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"github.com/spf13/viper"
)

func NewAuthorizationMiddleware() fiber.Handler {
	return Authorization
}

func Authorization(c *fiber.Ctx) error {
	JWT_SECRET := viper.GetString("JWT_SECRET")
	authorization := c.Get("Authorization")

	var tokenString string
	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	requestMap := make(map[string]interface{})

	if err := c.BodyParser(&requestMap); err != nil {
		return err
	}

	db := config.InitDatabase(config.InitViper(), logger.NewLogger("AUTH"))

	nasabah_ := entity.Nasabah{}
	db.Preload("Rekening").Joins("nasabah").First(&nasabah_, "no_nasabah = ?", claims["no_nasabah"])

	isAUthorized := false
	for _, value := range nasabah_.Rekening {
		if value.NoRekening == requestMap["no_rekening"] || value.NoRekening == requestMap["no_rekening_asal"] {
			isAUthorized = true
		}
	}

	if !isAUthorized {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":   fiber.StatusUnauthorized,
			"status": "error",
			"remark": "tidak memiliki izin untuk melanjutkan proses",
		})
	}

	return c.Next()
}
