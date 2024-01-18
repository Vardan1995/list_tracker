package utils

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v4"
)

// func Protected() fiber.Handler {//todo
// 	return jwtware.New(jwtware.Config{
// 		SigningKey:     []byte(config.Config("SECRET")),
// 		ErrorHandler:   jwtError,
// 		SuccessHandler: jwtSuccess,
// 	})
// }

func jwtError(c fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func jwtSuccess(c fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	// in := int(id)

	c.Locals("id", uint(id))
	return c.Next()
}
