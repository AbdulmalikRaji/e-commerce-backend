package authentication

import (
	"github.com/abdulmalikraji/e-commerce/dto/authDto"
	"github.com/abdulmalikraji/e-commerce/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	LoginByEmail(ctx *fiber.Ctx) error
}

type authHandler struct {
	service services.AuthService
}

func New(service services.AuthService) AuthHandler {
	return authHandler{
		service: service,
	}
}

func (c authHandler) LoginByEmail(ctx *fiber.Ctx) error {
	var loginRequest authDto.LoginByEmailRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	token, err := c.service.LoginByEmail(ctx, loginRequest)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"token": token})
}
