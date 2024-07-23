package handlers

import (
	"fastauth/auth"
	"fastauth/storage"
	"fastauth/structs"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTokenHandler(secretKey string, db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginPayload := new(structs.LoginRequest)

		if err := c.BodyParser(&loginPayload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		verifiedUser := storage.VerifyUser(*loginPayload, db)
		if !verifiedUser {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		token := auth.CreateToken(*loginPayload, secretKey)

		return c.JSON(fiber.Map{"token": token})
	}
}

func RegisterHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		registerPayload := new(structs.RegisterRequest)

		if err := c.BodyParser(&registerPayload); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if err := storage.CreateUser(*registerPayload, db); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendString(registerPayload.Username)
	}
}